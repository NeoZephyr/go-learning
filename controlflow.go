package main

import (
    "fmt"
	"time"
)

func main() {
	forApp()
	switchApp()
}

func forApp() {
	fmt.Println()
	fmt.Println("=== for App")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Printf("sum = %d\n", sum)

	n := 0
	for {
		if n >= 3 {
			break
		}

		n++
		time.Sleep(time.Second)
		fmt.Println("for sleep...")
	}

	numArr := [...]int{1, 2, 3, 4, 5}
	numSlice := []int{1, 2, 3, 4, 5}
	maxIdx := len(numArr) - 1

	// range 表达式只会在 for 语句开始执行时被求值一次
	// range 表达式的求值结果会被复制，被迭代的对象是 range 表达式结果值的副本
	for idx, elem := range(numArr) {
		if (idx == maxIdx) {
			numArr[0] += elem
		} else {
			numArr[idx + 1] += elem
		}
	}

	for idx, elem := range(numSlice) {
		if (idx == maxIdx) {
			numSlice[0] += elem
		} else {
			numSlice[idx + 1] += elem
		}
	}

	fmt.Printf("numArr: %v\n", numArr)
	fmt.Printf("numSlice: %v\n", numSlice)
}

func switchApp() {
	fmt.Println()
	fmt.Println("=== switch App")

	score := 'E'

	// 没有条件的 switch 同 switch true 一样
	switch score {
	case 'A':
		fmt.Println("90 - 100")
	case 'B':
		fmt.Println("80 - 89")
	case 'C':
		fmt.Println("70 - 79")
	case 'D', 'E', 'F':
		fallthrough
	default:
		fmt.Println("not passed")
	}
}

