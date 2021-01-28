package main

import (
	"fmt"
)

func main() {
	variableParamApp(100, 200, 300)
	closureApp()
}

func variableParamApp(nums ...int) {
	sum := 0

	for i := range nums {
		sum += nums[i]
	}

	fmt.Printf("sum: %v\n", sum)
}

func closureApp() {
	addFunc, subFunc, echoFunc := calc(100)

	fmt.Println("add:", addFunc(10))
	fmt.Println("echo: ", echoFunc())
	fmt.Println("sub:", subFunc(20))
	fmt.Println("echo: ", echoFunc())
}

func calc(base int) (func(int) int, func(int) int, func() int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	echo := func() int {
		return base
	}

	return add, sub, echo
}
