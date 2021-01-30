package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	// block()
	// blockWait()
	// appAndDebug()

	var a App
	a.track.Shutdown()
}

func block() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello go")
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func blockWait() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello go")
	})

	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}

func appAndDebug() {
	done := make(chan error, 2)
	stop := make(chan struct{})

	go func() {
		done <- serveDebug(stop)
	}()

	go func() {
		done <- serveApp(stop)
	}()

	var stopped bool

	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			fmt.Println("error: %v", err)
		}

		if !stopped {
			stopped = true
			close(stop)
		}
	}
}

func serveApp(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello go")
	})

	return serve("0.0.0.0:8000", mux, stop)
}

func serveDebug(stop <-chan struct{}) error {
	// /debug/pprof
	return serve("127.0.0.1:8001", http.DefaultServeMux, stop)
}

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

type Tracker struct {
	wg sync.WaitGroup
}

type App struct {
	track Tracker
}

func (t *Tracker) Event(data string) {
	t.wg.Add(1)

	go func() {
		defer t.wg.Done()
		time.Sleep(time.Millisecond)
		log.Println(data)
	}()
}

func (t *Tracker) Shutdown(ctx context.Context) error {
	ch := make(chan struct{})

	go func() {
		t.wg.Wait()
		close(ch)
	}()

	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		return errors.New("timeout")
	}
}

func (a *App) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	a.tract.Event("event...")
}
