package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	timeDemo()
	timerDemo()
	tickDemo()
	randDemo()
}

func timeDemo() {
	fmt.Println("=== timeDemo")

	t := time.Now()
	fmt.Printf("origin format: %v\n", t)
	fmt.Printf("dd/MM/yyyy HH:mm format: %v\n", t.Format("02/1/2006 15:04"))
	fmt.Printf("yyyy-MM-dd HH:mm format: %v\n", t.Format("2006-1-02 15:04"))
	fmt.Printf("yyyy/MM/dd format: %v\n", t.Format("2006/1/02"))
}

func timerDemo() {
	fmt.Println("=== timerDemo")

	t := time.NewTimer(1 * time.Second)

	fmt.Println("begin:", time.Now())
	fmt.Println("end1:", <- t.C)

	<- time.After(1 * time.Second)
	fmt.Println("end2")

	// 停止
	//t.Stop()
	// 重置
	//t.Reset(3 * time.Second)
}

func tickDemo()  {
	fmt.Println("=== tickDemo")

	t := time.NewTicker(1 * time.Second)
	<- t.C
	fmt.Println("tick 1")
	<- t.C
	fmt.Println("tick 2")
}

func randDemo() {
	fmt.Println("=== randDemo")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Printf("rand: %v\n", rand.Intn(100))
	}
}
