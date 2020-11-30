package main

import (
	"fmt"
)

func main() {
	testChannel()
}

type Worker struct {
	in chan int
	done chan bool
}

func worker(id int, c chan int, done chan bool) {
	for v := range c {
		fmt.Printf("worker %d receive %c\n", id, v)
		done <- true
	}
}

func createWorker(id int) Worker {
	w := Worker{
		in: make(chan int),
		done: make(chan bool),
	}
	go worker(id, w.in, w.done)
	return w
}

func testChannel() {
	var workers [10]Worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		<- worker.done
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for _, worker := range workers {
		<- worker.done
	}

	//time.Sleep(time.Millisecond)
}
