package main

import (
    "fmt"
	"time"
)

func main() {
	ifDemo()
	switchDemo()
	forDemo()
}

func ifDemo() {
	fmt.Println("=== ifDemo")

	lang := "golang"

	if lang == "python" {
		fmt.Println("need tabs")
	} else if lang == "java" {
		fmt.Println("need spaces")
	} else {
		fmt.Println("spaces or tabs")
	}
}

func switchDemo() {
	fmt.Println("=== switchDemo")

	score := 'E'

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

func forDemo() {
	fmt.Println("=== forDemo")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for {
		time.Sleep(time.Second)
		fmt.Println("sleep")
	}
}
