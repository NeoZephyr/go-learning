package main

import (
	"fmt"
	"time"
)

func main() {
	//testChannel()
	//testBufferedChannel()
	testChannelClose()
}

func worker(id int, c chan int) {
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker %d consume %c\n", id, n)
	//}

	for n := range c {
		fmt.Printf("worker %d consume %c\n", id, n)
	}
}

// 返回值，只能往里面发数据
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func testChannel() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func testBufferedChannel() {
	c := make(chan int, 3)

	go worker(0, c)
	c <- 'A'
	c <- 'B'
	c <- 'C'
	c <- 'D'

	time.Sleep(time.Millisecond)
}

func testChannelClose() {
	c := make(chan int, 3)

	go worker(0, c)
	c <- 'A'
	c <- 'B'
	c <- 'C'
	c <- 'D'

	// 关闭
	// 向关闭的 channel 发送数据，会发生 panic
	// 关闭 channel 可以用于通知多个协程
	close(c)
	time.Sleep(time.Millisecond)
}