package main

import (
	"fmt"
	"net/http"
)

func ServeHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server start")
}

func main() {
	http.HandleFunc("/", ServeHttp)
	http.ListenAndServe(":8080", nil)
}
