package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
)

const (
	proxyAddress = "http://127.0.0.1:2003"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatalln(http.ListenAndServe("127.0.0.1:2002", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse(proxyAddress)
	r.URL.Scheme = url.Scheme
	r.URL.Host = url.Host

	transport := http.DefaultTransport
	response, err := transport.RoundTrip(r)

	if err != nil {
		log.Println(err.Error())
		return
	}

	for k, vv := range response.Header {
		for _, v := range vv {
			w.Header().Add(k, v)
		}
	}

	defer response.Body.Close()
	bufio.NewReader(response.Body).WriteTo(w)
}