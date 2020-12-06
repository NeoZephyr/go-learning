package main

import (
	"fmt"
)

func main() {
	intPointerApp()
    structPointerApp()
}

func intPointerApp() {
	fmt.Println()
	fmt.Println("=== int pointer app")

	num := 100
	p := &num

	fmt.Printf("num addr: %p, num value: %v, p value: %v\n", p, num, *p)

	*p = *p * 6

	fmt.Printf("num addr: %p, num value: %v, p value: %v\n", p, num, *p)

}

type Vertex struct {
	X int
	Y int
}

func structPointerApp() {
	fmt.Println()
	fmt.Println("=== struct pointer app")

	v := Vertex{Y: 300}
	fmt.Printf("v vlaue: %v\n", v)

	p := &v
	p.X = 300
	fmt.Printf("v vlaue: %v\n", v)

}
