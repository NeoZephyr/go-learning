package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "Your name")
}

func main() {
	flag.Parse()

	gretting, err := hello(name)

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(gretting)
}

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	return fmt.Sprintf("Hello, %s!", name), nil
}

func getPrimes(max int) []int {
	if max <= 1 {
		return []int{}
	}
	marks := make([]bool, max)
	var count int
	squareRoot := int(math.Sqrt(float64(max)))
	for i := 2; i <= squareRoot; i++ {
		if marks[i] == false {
			for j := i * i; j < max; j += i {
				if marks[j] == false {
					marks[j] = true
					count++
				}
			}
		}
	}
	primes := make([]int, 0, max-count)
	for i := 2; i < max; i++ {
		if marks[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}
