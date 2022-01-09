package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

// curl -i -H 'name:jack' -H 'name:pain' 127.0.0.1

func index(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
			fmt.Printf("k: %s, v: %s\n", k, vv)
		}
	}

	w.Header().Set("VERSION", os.Getenv("VERSION"))

	host, _, err := net.SplitHostPort(req.RemoteAddr)

	if err == nil {
		log.Printf("remote ip: %s, http response code: %d\n", host, http.StatusOK)
	}
}

func health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	os.Setenv("VERSION", "1.0")
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(index))
	mux.Handle("/healthz", http.HandlerFunc(health))
	err := http.ListenAndServe(":80", mux)

	if err != nil {
		log.Fatalf("http server error: %s", err)
	}
}
