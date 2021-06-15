package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const (
	address = "127.0.0.1:2002"
	proxyAddress = "http://127.0.0.1:2003/proxy"
)

func main() {
	url, _ := url.Parse(proxyAddress)

	proxy := newSingleHostReverseProxy(url)
	log.Printf("listen on %s\n", address)
	log.Fatalln(http.ListenAndServe(address, proxy))
}

func newSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
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
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
}

func modifyResponse (response *http.Response) error {
	if response.StatusCode != http.StatusOK {
		return nil
	}
	sourceBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	destBytes := []byte("=====" + string(sourceBytes))
	response.Body = ioutil.NopCloser(bytes.NewBuffer(destBytes))
	response.ContentLength = int64(len(destBytes))
	response.Header.Set("Content-Length", fmt.Sprint(len(destBytes)))
	return nil
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