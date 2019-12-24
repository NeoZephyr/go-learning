package main

import (
	"fmt"
	"time"
)

func main() {
	testTime()
	//testTimer()
	//testTick()
}

func testTime() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format("02/1/2006 15:04"))
	fmt.Println(t.Format("2006-1-02 15:04"))
	fmt.Println(t.Format("2006/1/02"))
}

func testTimer() {
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

func testTick()  {
	t := time.NewTicker(1 * time.Second)
	<- t.C
	fmt.Println("tick 1")
	<- t.C
	fmt.Println("tick 2")
}