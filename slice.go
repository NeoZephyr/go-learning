package main

import "fmt"

func main() {
	sliceApp()
}

func sliceApp() {
	fmt.Println()
	fmt.Println("=== slice App")

	s0 := [7]int{1, 2, 3, 4, 5, 6, 7}
	s1 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("s0, length: %d, capacity: %d, value: %v, type: %T\n", len(s0), cap(s0), s0, s0)
	fmt.Printf("s1, length: %d, capacity: %d, value: %v, type: %T\n", len(s1), cap(s1), s1, s1)

	// s1, s2 共用底层结构
	s2 := s1[4:6]
	fmt.Printf("s2, length: %d, capacity: %d, value: %v\n", len(s2), cap(s2), s2)

	for i := 1; i <= 3; i++ {
		s2 = append(s2, i * 100)
		fmt.Printf("s2, length: %d, capacity: %d, value: %v\n", len(s2), cap(s2), s2)
	}

	fmt.Printf("s1, length: %d, capacity: %d, value: %v\n", len(s1), cap(s1), s1)


	s3 := make([]int, 5, 10)
	fmt.Printf("s3, length: %d, capacity: %d, value: %v\n", len(s3), cap(s3), s3)

	s4 := append(s3, make([]int, 8)...)
	fmt.Printf("s4, length: %d, capacity: %d, value: %v\n", len(s4), cap(s4), s4)

	s5 := make([]int, 5)
	copy(s5, s4)
	fmt.Printf("s5, length: %d, capacity: %d, value: %v\n", len(s5), cap(s5), s5)

	s6 := make([][]int, 3)

	for i := 0; i < len(s6); i++ {
		s6[i] = make([]int, i + 1)

		for j := 0; j < i + 1; j++ {
			s6[i][j] = j
		}
	}

	fmt.Printf("s6, length: %d, capacity: %d, value: %v\n", len(s6), cap(s6), s6)
}
