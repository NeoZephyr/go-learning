package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server start")
}

type IndexHandler struct{}

func (this *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.RawQuery)
	fmt.Fprintln(w, r.URL.Host)
	fmt.Fprintln(w, r.URL.Path)

	rawQuery := r.URL.RawQuery
	params, _ := url.ParseQuery(rawQuery)

	fmt.Fprintln(w, params.Get("name"))
	fmt.Fprintln(w, "index handler")

	for key := range r.Header {
		fmt.Fprintf(w, "key: %s, value: %v\n", key, r.Header[key])
	}
}

func main() {
	http.HandleFunc("/", ServeHTTP)

	indexHandler := IndexHandler{}
	http.Handle("/index", &indexHandler)
	http.ListenAndServe(":8080", nil)
}
