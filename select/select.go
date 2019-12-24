package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	testSelect()
}

func createProducer() chan int {
	out := make(chan int)

	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func createConsumer(id int) chan int {
	in := make(chan int)

	go func(id int) {
		for n := range in {
			time.Sleep(time.Second)
			fmt.Printf("consumer %d consume %d\n", id, n)
		}
	}(id)

	return in
}

func testSelect() {
	p1, p2 := createProducer(), createProducer()
	c1 := createConsumer(9527)

	var bufValues []int
	t := time.After(10 * time.Second)
	tick := time.Tick(2 * time.Second)

	for {
		var activeWorker chan int
		var activeValue int

		if len(bufValues) > 0 {
			activeWorker = c1
			activeValue = bufValues[0]
		}

		select {
		case n := <- p1:
			// 产生数据数据可能比消耗速度快
			bufValues = append(bufValues, n)
		case n := <- p2:
			bufValues = append(bufValues, n)
		case activeWorker <- activeValue: // nil channel 无法被 select 到
			bufValues = bufValues[1:]
		case <- time.After(800 * time.Millisecond):
			fmt.Println("time out")
		case <- t: // 10 秒钟结束
			fmt.Println("end")
			return
		case <- tick:
			fmt.Println("buf size:", len(bufValues))
		}
	}
}