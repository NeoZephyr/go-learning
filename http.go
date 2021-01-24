package main

import (
	"fmt"
	"io/ioutil"
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

	fmt.Fprintln(w, "query:")

	for key := range params {
		fmt.Fprintf(w, "key: %v, value: %v\n", key, params.Get(key))
	}
	fmt.Fprintln(w)

	fmt.Fprintln(w, "header:")
	for key := range r.Header {
		fmt.Fprintf(w, "key: %s, value: %v\n", key, r.Header[key])
	}
	fmt.Fprintln(w)

	r.ParseForm()
	fmt.Fprintln(w, "form:")
	fmt.Fprintf(w, "post form: %v\n", r.PostForm)
	fmt.Fprintf(w, "form: %v\n", r.Form)
	fmt.Fprintln(w)

	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["jd"][0]
	file, err := fileHeader.Open()

	fmt.Fprintln(w, "file:")
	if err != nil {
		fmt.Fprintln(w, "open file failed")
	} else {
		data, err := ioutil.ReadAll(file)

		if err != nil {
			fmt.Fprintln(w, "read from file failed")
		} else {
			fmt.Fprintln(w, string(data))
		}
	}
	fmt.Fprintln(w)

	size := r.ContentLength
	bodyBuf := make([]byte, size)
	r.Body.Read(bodyBuf)

	fmt.Fprintln(w, "request body:")
	fmt.Fprintf(w, "%v\n", string(bodyBuf))
	fmt.Fprintln(w)

	// type User struct {
	// 	Username string
	// 	Age int8
	// }

	// user := &User{
	// 	Username: "jack",
	// 	Age: 18
	// }

	// _json, _ := json.Marshal(user)
	// w.Write(_json)
}

func main() {
	http.HandleFunc("/", ServeHTTP)

	indexHandler := IndexHandler{}
	http.Handle("/index", &indexHandler)
	http.ListenAndServe(":8080", nil)
}
