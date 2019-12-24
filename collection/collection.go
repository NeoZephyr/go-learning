package main

import "fmt"

func main() {
	//testArray()
	//testSlice1()
	//testSlice2()
	testSlice3()

	//testMap()
}

func testArray()  {
	var arr1 [3]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8}
	var grid [2][3]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 调用会拷贝数组，是值传递，即函数内改变值之后，不会影响原数组
	printArray(arr1)

	updateArray1(arr2)
	fmt.Println("after first time update:")
	printArray(arr2)
	updateArray2(&arr2)
	fmt.Println("after second time update:")
	printArray(arr2)
}

func printArray(arr [3]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func updateArray1(arr [3]int)  {
	for i := range arr {
		arr[i] *= 100
	}
}

func updateArray2(arr *[3]int) {
	for i := range arr {
		arr[i] *= 100
	}
}

func testSlice1() {
	arr := [...]int{100, 200, 300, 400, 500, 600, 700, 800, 900}

	fmt.Println("arr[2:5] = ", arr[2:5])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:5] = ", arr[:5])
	fmt.Println("arr[:] = ", arr[:])

	s := arr[2:5]
	fmt.Println("arr = ", arr)
	fmt.Println("s = ", s)

	fmt.Println("after update slice:")
	updateSlice(s)
	fmt.Println("arr = ", arr)
	fmt.Println("s = ", s)

	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
}

func updateSlice(slice []int) {
	for i := range slice {
		slice[i] = slice[i] * 3
	}
}

func testSlice2() {
	arr := [...]int{100, 200, 300, 400, 500, 600, 700, 800}
	s1 := arr[2:7]
	fmt.Printf("arr = %v, s1 = %v\n", arr, s1)

	s2 := append(s1, -100)

	// s3, s4, s5 是其它数组的 view
	// 系统重新分配更大的底层数组
	s3 := append(s2, -200)
	s4 := append(s3, -300)
	s5 := append(s4, -400)

	fmt.Printf("s1 = %v\n", s1)
	fmt.Printf("s2 = %v\n", s2)
	fmt.Printf("s3 = %v\n", s3)
	fmt.Printf("s4 = %v\n", s4)
	fmt.Printf("s5 = %v\n", s5)
	fmt.Printf("arr = %v\n", arr)
}

func testSlice3() {
	var s1 []int

	for i := 1; i < 50; i += 2 {
		printSlice(s1)
		s1 = append(s1, i)
	}

	printSlice(s1)

	s2 := make([]int, 12)
	s3 := make([]int, 23, 33)

	printSlice(s2)
	printSlice(s3)

	copy(s2, s1)
	printSlice(s2)

	// delete
	s1 = append(s1[:4], s1[5:]...)
	printSlice(s1)

	front := s1[0]
	s1 = s1[1:]

	tail := s1[len(s1) - 1]
	s1 = s1[:len(s1) - 1]

	fmt.Printf("front = %v, tail = %d\n", front, tail)
	printSlice(s1)
}

func printSlice(s []int) {
	fmt.Printf("slice = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
}

func testMap() {
	m1 := map[string]string {
		"name": "jack",
		"email": "a@qq.com",
		"mobile": "110",
	}

	fmt.Println(m1)

	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println(m2, m3)

	for k, v := range m1 {
		fmt.Printf("%v: %v\n", k, v)
	}

	name, exist := m1["name"]
	fmt.Printf("name: %v, exist: %v\n", name, exist)

	delete(m1, "name")
	name, exist = m1["name"]
	fmt.Printf("name: %v, exist: %v\n", name, exist)
}