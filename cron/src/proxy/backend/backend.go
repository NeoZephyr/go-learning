package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type BackendServer struct {
	address string
}

func main() {
	server1 := &BackendServer{
		address: "127.0.0.1:2003",
	}
	server2 := &BackendServer{
		address: "127.0.0.1:2004",
	}
	server1.run()
	server2.run()

	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<- signals
}

func (server *BackendServer) run() {
	log.Printf("start server: %s\n", server.address)
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", server.debugHandler)
	serveMux.HandleFunc("/ping", server.pingHandler)
	serveMux.HandleFunc("/error", server.errorHandler)
	serveMux.HandleFunc("/timeout", server.timeoutHandler)

	httpServer := &http.Server{
		Addr:         server.address,
		WriteTimeout: time.Second * 1,
		Handler:      serveMux,
	}
	go func() {
		log.Fatalln(httpServer.ListenAndServe())
	}()
}

func (server *BackendServer) debugHandler(w http.ResponseWriter, r *http.Request)  {
	pathInfo := fmt.Sprintf("http://%s%s\n", server.address, r.URL.Path)
	ipInfo := fmt.Sprintf("RemoteAddr=%s,X-Forwarded-For=%s,X-Real-Ip=%s\n",
		r.RemoteAddr, r.Header.Get("X-Forwarded-For"), r.Header.Get("X-Real-Ip"))
	headerInfo := fmt.Sprintf("Headers=%v\n", r.Header)

	io.WriteString(w, pathInfo)
	io.WriteString(w, ipInfo)
	io.WriteString(w, headerInfo)
}

func (server *BackendServer) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "pong")
}

func (server *BackendServer) errorHandler(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, http.StatusText(http.StatusInternalServerError))
}

func (server *BackendServer) timeoutHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, http.StatusText(http.StatusOK))
}