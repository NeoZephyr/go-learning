package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//testWithoutCoroutine()
	//testWithCoroutine1()
	testWithCoroutine2()
}

func testWithoutCoroutine()  {
	for i := 0; i < 10; i++ {
		func(i int) {
			for {
				fmt.Printf("No coroutine, %d\n", i)
			}
		}(i)
	}
}

func testWithCoroutine1() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			// IO 操作，协程会交出控制权
			// select, channel, wait lock, call function, runtime.Gosched
			for {
				fmt.Printf("Coroutine, %d\n", i)
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
}

func testWithCoroutine2() {
	// 以单核运行
	//runtime.GOMAXPROCS(1)
	//runtime.NumCPU()

	var arr [10]int

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				arr[i]++

				// 手动交出控制权
				runtime.Gosched()

				// 结束当前协程
				//runtime.Goexit()
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Println(arr)
}