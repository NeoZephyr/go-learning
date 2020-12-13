package main

import (
	"fmt"
	"time"
	"sync/atomic"
)

func main() {
	// withoutParamTimeApp()
	// withoutParamChanApp()
	// withParamApp()
	withParamOrderApp()
}

func withoutParamTimeApp() {
	fmt.Println()
	fmt.Println("=== without param time app")

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(time.Millisecond * 50)
}

func withoutParamChanApp() {
	fmt.Println()
	fmt.Println("=== without param chan app")

	sign := make(chan struct{}, 5)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	for i := 0; i < 5; i++ {
		<- sign
	}
}

func withParamApp() {
	fmt.Println()
	fmt.Println("=== with param app")

	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Millisecond * 50)
}

func withParamOrderApp() {
	fmt.Println()
	fmt.Println("=== with param order app")

	var count uint32

	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break;
			}

			time.Sleep(time.Nanosecond)
		}
	}

	for i := uint32(0); i < 5; i++ {
		go func(i uint32) {
			trigger(i, func() {
				fmt.Println(i)
			})
		}(i)
	}

	trigger(5, func() {})
}
