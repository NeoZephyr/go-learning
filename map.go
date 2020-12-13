package main

import "fmt"

func main() {
	mapDemo()
}

func mapDemo() {
	fmt.Println()
	fmt.Println("=== map App")

	m := make(map[string]int)
	m["lakers"] = 17
	m["bulls"] = 6
	m["spurs"] = 5

	fmt.Printf("m: %v, len: %d, the type of m is %T\n", m, len(m), m)

	elem, ok := m["bulls"]
	fmt.Printf("elem: %d, exist: %v\n", elem, ok)

	delete(m, "bulls")
	fmt.Printf("m = %v, len: %d, the type of m is %T\n", m, len(m), m)

	fmt.Println("map range:")
	for k, v := range m {
		fmt.Printf("m[%s] = %v\n", k, v)
	}

	n := map[string]int{"lebron": 4, "jordan": 6}
	fmt.Printf("n: %v, len: %d\n", n, len(n))

	k, success := interface{}(n).(map[string]int)
	fmt.Printf("n convert to map[string]int, value: %v , success: %v\n", k, success)
}
