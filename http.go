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
}

func main() {
	http.HandleFunc("/", ServeHTTP)

	indexHandler := indexHandler{}
	http.Handle("/index", &indexHandler)
	http.ListenAndServe(":8080", nil)
}
