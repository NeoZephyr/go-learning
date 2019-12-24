package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	testContext()
}

func testContext() {
	// 根 context.Background()
	// 产生子 Context
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go func(i int, context context.Context) {
			for {
				if isCanceled(ctx) {
					break
				}

				time.Sleep(time.Millisecond * 5)
			}

			fmt.Println(i, "canceled")
		}(i, ctx)
	}

	cancel()

	time.Sleep(time.Millisecond * 100)
}

func isCanceled(context context.Context) bool {
	select {
	case <- context.Done():
		return true
	default:
		return false
	}
}