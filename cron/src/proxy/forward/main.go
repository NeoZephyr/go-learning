package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Proxy struct {}

func (proxy *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("received request %s %s %s\n", r.Method, r.Host, r.RemoteAddr)
	transport := http.DefaultTransport
	request := new(http.Request)
	*request = *r
	if clientIp, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		if pre, ok := request.Header["X-Forwarded-For"]; ok {
			clientIp = strings.Join(pre, ", ") + ", " + clientIp
		}
		request.Header.Set("X-Forwarded-For", clientIp)
	}

	// 请求下游
	response, err := transport.RoundTrip(request)

	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	for key, value := range response.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
	response.Body.Close()
}

//func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

//	// step 3, 把下游请求内容返回给上游
//	rw.WriteHeader(res.StatusCode)
//	io.Copy(rw, res.Body)
//	res.Body.Close()
//}
//


func main() {
	http.Handle("/", &Proxy{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}