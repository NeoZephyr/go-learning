package main

import (
	"fmt"
	"sync"
)

func main() {
	testChannel()
}

type Worker struct {
	in chan int
	//wg *sync.WaitGroup
	done func()
}

func worker(id int, c chan int, done func()) {
	for v := range c {
		fmt.Printf("worker %d receive %c\n", id, v)
		done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) Worker {
	w := Worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go worker(id, w.in, w.done)
	return w
}

func testChannel() {
	var workers [10]Worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	wg.Wait()

	//time.Sleep(time.Millisecond)
}