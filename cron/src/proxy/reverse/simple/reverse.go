package main

import (
	"cron/src/proxy/core/lb"
	"cron/src/proxy/core/proxy"
	"cron/src/proxy/router"
	"fmt"
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

	seqRouter := router.NewSeqRouter()
	seqRouter.Group("/hello", func(routerContext *router.SeqRouterContext) {
		routerContext.Rw.Write([]byte("hello world"))
	})
	seqRouter.Group("/", func(routerContext *router.SeqRouterContext) {
		fmt.Println("proxy....")
		reverseProxy := proxy.NewSingleHostReverseProxy(rb, routerContext)
		reverseProxy.ServeHTTP(routerContext.Rw, routerContext.Req)
	})

	log.Printf("listen on %s\n", address)
	log.Fatalln(http.ListenAndServe(address, seqRouter))
}