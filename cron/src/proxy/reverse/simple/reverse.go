package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	address = "127.0.0.1:2002"
	proxyAddress = "http://127.0.0.1:2003/proxy"
)

func main() {
	url, _ := url.Parse(proxyAddress)
	proxy := httputil.NewSingleHostReverseProxy(url)

	log.Printf("listen on %s\n", address)
	log.Fatalln(http.ListenAndServe(address, proxy))
}
