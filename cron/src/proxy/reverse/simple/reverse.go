package main

import (
	"cron/src/proxy/core/lb"
	"cron/src/proxy/core/proxy"
	"log"
	"net/http"
)

const (
	address = "127.0.0.1:2002"
	proxyAddress = "http://127.0.0.1:2003/proxy"
)

func main() {
	rb := lb.GetLoadBalancer(lb.WeightRound)

	if err := rb.Add("http://127.0.0.1:2003/base", "10"); err != nil {
		log.Fatalln(err)
	}

	if err := rb.Add("http://127.0.0.1:2003/base", "10"); err != nil {
		log.Fatalln(err)
	}

	proxyHandler := proxy.NewSingleHostReverseProxy(rb)
	log.Printf("listen on %s\n", address)
	log.Fatalln(http.ListenAndServe(address, proxyHandler))
}