package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

func funcWithVariableParamDemo(nums ...int) int {
	sum := 0

	for i := range nums {
		sum += nums[i]
	}

	return sum
}

func testFuncWithMultiReturn(a, b int, opt string) (int, error) {
	switch opt {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s\n", opt)
	}
}

func testFunctional() {
	result := apply(pow, 2, 3)
	fmt.Printf("result: %v\n", result)

	result = apply(func(a, b int) int {
		return a + b
	}, 100, 200)
	fmt.Printf("result: %v\n", result)

	powStat := performanceStat(pow)
	fmt.Printf("6^6 = %d\n", powStat(6, 6))
}

// 函数作为参数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling %s\n", opName)
	return op(a, b)
}

// 函数作为返回值
func performanceStat(inner func(int, int) int) func(int, int) int {
	wrapper := func(a, b int) int {
		start := time.Now()
		ret := inner(a, b)

		p := reflect.ValueOf(inner).Pointer()
		funcName := runtime.FuncForPC(p).Name()

		fmt.Printf("function %s spent %f seconds\n", funcName, time.Since(start).Seconds())

		return ret
	}

	return wrapper
}

func pow(a, b int) int {
	time.Sleep(time.Second * 2)
	return int(math.Pow(float64(a), float64(b)))
}

func testAnonymous() {
	var f1 = func(a, b int) int {
		return a + b
	}

	fmt.Println("f1(1, 3) =", f1(1, 3))

	var result = func(a, b int) int {
		return a * b
	}(4, 58)
	fmt.Println("result =", result)

	// 修改外部变量
	func() {
		result = 50
		fmt.Println("update result in anonymous, result = ", result)
	}()

	fmt.Println("result =", result)
}

func getClosure() func() int {
	// 自由变量
	x := 1

	// x 在闭包中被使用
	return func() int {
		x++
		return x * x
	}
}

func testClosure() {
	f := getClosure()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func testDefer()  {
	// 多个 defer 按照先进后出的顺序执行
	defer fmt.Println("defer1...")
	defer fmt.Println("defer2...")

	fmt.Println("test defer1")

	defer func() {
		fmt.Println("defer3...")
	}()

	// 即是有 panic，defer 也会执行
	panic("error happened")

	defer fmt.Println("defer3...")
}
