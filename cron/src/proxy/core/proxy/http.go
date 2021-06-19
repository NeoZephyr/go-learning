package proxy

import (
	"bytes"
	"compress/gzip"
	"cron/src/proxy/core/lb"
	"cron/src/proxy/router"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second, // 连接超时
		KeepAlive: 30 * time.Second, // 长连接超时时间
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

func NewSingleHostReverseProxy(balancer lb.LoadBalancer, ctx *router.SeqRouterContext) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		addr, err := balancer.Get(req.RemoteAddr)

		if err != nil {
			log.Fatalln("get server addr failed")
		}

		target, err := url.Parse(addr)

		if err != nil {
			log.Fatalln(err)
		}

		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path, req.URL.RawPath = joinURLPath(target, req.URL)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
		req.Header.Set("X-Real-Ip", req.RemoteAddr)
	}

	modifyResponse := func(response *http.Response) error {
		if strings.Contains(response.Header.Get("Connection"), "Upgrade") {
			return nil
		}

		var payload []byte
		var readErr error

		if strings.Contains(response.Header.Get("Content-Encoding"), "gzip") {
			reader, err := gzip.NewReader(response.Body)

			if err != nil {
				return err
			}

			payload, readErr = ioutil.ReadAll(reader)
			response.Header.Del("Content-Encoding")
		} else {
			payload, readErr = ioutil.ReadAll(response.Body)
		}

		if readErr != nil {
			return readErr
		}

		if response.StatusCode != http.StatusOK {
			payload = []byte("StatusCode error: " + string(payload))
		}

		ctx.Set("status_code", response.StatusCode)
		ctx.Set("payload", payload)

		response.Body = ioutil.NopCloser(bytes.NewBuffer(payload))
		response.ContentLength = int64(len(payload))
		response.Header.Set("Content-Length", strconv.FormatInt(int64(len(payload)), 10))
		return nil
	}

	fmt.Println("=======create")
	return &httputil.ReverseProxy{Director: director, Transport: transport, ModifyResponse: modifyResponse, ErrorHandler: errorHandler}
}

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, "error: "+err.Error(), 500)
}

func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
