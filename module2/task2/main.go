package main

import (
	"fmt"
	"sync"
	"time"
)

func produce(no int, wg *sync.WaitGroup, c chan <- string) {
	counter := 0
	wg.Add(1)
	for {
		time.Sleep(time.Second)
		c <- fmt.Sprintf("Producer %d ---> %d", no, counter)
		counter++
	}
	wg.Done()
}

func consume(no int, wg *sync.WaitGroup, c <- chan string) {
	wg.Add(1)
	for v := range c {
		fmt.Printf("Consumer %d consume message: %s", no, v)
	}
	wg.Done()
}

func main() {
	var c chan string
	pwg := &sync.WaitGroup{}
	cwg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		go produce(i, pwg, c)
	}

	go consume(0, cwg, c)

	pwg.Wait()
	cwg.Wait()
}
