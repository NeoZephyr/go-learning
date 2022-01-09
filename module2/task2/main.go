package main

import (
	"fmt"
	"sync"
	"time"
)

var stop = false

func produce(no int, wg *sync.WaitGroup, c chan <- string) {
	counter := 0
	for !stop {
		time.Sleep(time.Second)
		c <- fmt.Sprintf("Producer %d ---> %d", no, counter)
		counter++
	}
	wg.Done()
}

func consume(no int, wg *sync.WaitGroup, c <- chan string) {
	for v := range c {
		time.Sleep(time.Second)
		fmt.Printf("Consumer %d consume message: %s\n", no, v)
	}
	wg.Done()
}

func main() {
	c := make(chan string, 10)
	pwg := &sync.WaitGroup{}
	cwg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		pwg.Add(1)
		go produce(i, pwg, c)
	}

	for i := 0; i < 2; i++ {
		cwg.Add(1)
		go consume(i, cwg, c)
	}

	go func() {
		time.AfterFunc(time.Second * 10, func() {
			stop = true
			fmt.Println("stop produce!!!")
		})
	}()

	pwg.Wait()
	close(c)
	cwg.Wait()
}
