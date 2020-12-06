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
	fmt.Println("=== for App")
	fmt.Println()

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
}

func switchApp() {
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

