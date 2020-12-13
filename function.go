package main

import (
	"fmt"
	"errors"
)

type operate func(x int, y int) int
type calcFunc func(x int, y int) (int, error)

func main() {
	variableParamApp(100, 200, 300)

	result, err := multiReturnApp(100, 5, "*")

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("result: %v\n", result)
	}

	deferApp()

	result = functionParamApp(func(a int, b int) int {
		return a + b
	}, 100, 6)

	fmt.Printf("function param result: %v\n", result)

	add := genCalculator(func(x int, y int) int {
		return x + y
	})
	result, err = add(3, 5)

	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
    	fmt.Printf("add(3, 5) = %v\n", result)
	}
}

func deferApp()  {
	fmt.Println()
	fmt.Println("=== defer App")

	// 多个 defer 按照先进后出的顺序执行
	defer fmt.Println("defer1...")
	defer fmt.Println("defer2...")

	fmt.Println("do something")

	defer func() {
		fmt.Println("defer3...")
	}()

	// 即是有 panic，defer 也会执行
	// panic("error happened")

	defer fmt.Println("defer4...")
}

func variableParamApp(nums ...int) {
	fmt.Println()
	fmt.Println("=== variable param App")

	sum := 0

	for i := range nums {
		sum += nums[i]
	}

	fmt.Printf("sum: %v\n", sum)
}

func multiReturnApp(a, b int, opt string) (int, error) {
	fmt.Println()
	fmt.Println("=== multi return App")

	switch opt {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s\n", opt)
	}
}

func functionParamApp(fn func(int, int) int, a int, b int) int {
	fmt.Println()
	fmt.Println("=== function param App")

	return fn(a, b)
}

func genCalculator(op operate) calcFunc {
	return func(x int, y int) (int, error) {
		if (op == nil) {
			return 0, errors.New("invalid operate")
		}

		return op(x, y), nil
	}
}
