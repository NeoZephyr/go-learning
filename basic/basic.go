// 1. 同一个目录，包名必需一样
// 2. 包可以与目录不同名
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	constDemo()
	bitDemo()
	boolDemo()
	intDemo()
	byteDemo()
	stringDemo()
	convertDemo()
	arrayDemo()
	sliceDemo()
	mapDemo()
	objectDemo()

	// 命令行参数
	fmt.Println(os.Args)

	// 返回值
	// 不会调用 defer
	os.Exit(-1)
}

func constDemo() {
	fmt.Println("=== constDemo")

	const uploadLimit = 5000
	fmt.Printf("uploadLimit is %d, the type of uploadLimit is %T\n", uploadLimit, uploadLimit)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Printf("b = %d, kb = %d, mb = %d, gb = %d, tb = %d, pd = %d\n", b, kb, mb, gb, tb, pb)
}

func bitDemo() {
	fmt.Println("=== bitDemo")

	const (
		read = 1 << iota
		write
		execute
	)

	permission := 7
	fmt.Printf("read: %v, write: %v, execute: %v\n",
		(permission&read == read), (permission&write == write), (permission&execute == execute))
	permission = permission &^ execute
	fmt.Printf("read: %v, write: %v, execute: %v\n",
		(permission&read == read), (permission&write == write), (permission&execute == execute))
}

func boolDemo() {
	fmt.Println("=== boolDemo")

	var b1 bool
	b2 := true
	b3 := false

	fmt.Printf("b1 = %v, the type of b1 is %T\n", b1, b1)
	fmt.Printf("b2 = %v, the type of b2 is %T\n", b2, b2)
	fmt.Printf("b3 = %v, the type of b3 is %T\n", b3, b3)
}

func intDemo() {
	fmt.Println("=== intDemo")

	var i0 int
	var i1 int64

	i2 := 100
	i3 := int64(110)
	i4 := int32(120)

	fmt.Printf("i0 = %v, the type of i0 is %T\n", i0, i0)
	fmt.Printf("i1 = %v, the type of i1 is %T\n", i1, i1)
	fmt.Printf("i2 = %v, the type of i2 is %T\n", i2, i2)
	fmt.Printf("i3 = %v, the type of i3 is %T\n", i3, i3)
	fmt.Printf("i4 = %v, the type of i4 is %T\n", i4, i4)

	// fmt.Println(10 / 0)
}

func byteDemo() {
	fmt.Println("=== byteDemo")
	var b1 byte = 'a'
	b2 := 'b'

	fmt.Printf("b1 = %v, b1 = %c, the type of b1 is %T\n", b1, b1, b1)
	fmt.Printf("b1 - 32 = %v\n", b1-32)

	fmt.Printf("b2 = %v, b2 = %c, the type of b2 is %T\n", b2, b2, b2)
}

func stringDemo() {
	fmt.Println("=== stringDemo")

	// string 是值类型，默认值为空字符串
	var s0 string
	s1 := "\"Hello Go\""
	s2 := `"Hello Go"`
	s3 := `
你好
`
	s4 := "Golang!"

	fmt.Printf("s0 = %v, the type of s0 is %T\n", s0, s0)
	fmt.Printf("s1 = %v, the type of s1 is %T\n", s1, s1)
	fmt.Printf("s2 = %v, the type of s2 is %T\n", s2, s2)
	fmt.Printf("s3 + s4 = %v\n", s3+s4)

	s5 := "Zephyr 喜欢 go 语言"

	// 获取字节长度
	fmt.Printf("s5 = %s, byte len(s5) = %d\n", s5, len(s5))
	// 获取字符数量
	fmt.Printf("s5 = %s, char len(s5) = %d\n", s5, utf8.RuneCountInString(s5))

	fmt.Println("byte range:")
	for i, b := range []byte(s5) {
		fmt.Printf("(%d, %X, %c) ", i, b, b)
	}
	fmt.Println()

	fmt.Println("range:")
	for i, c := range s5 {
		fmt.Printf("(%d, %X, %c) ", i, c, c)
	}
	fmt.Println()

	fmt.Println("rune range:")
	for i, c := range []rune(s5) {
		fmt.Printf("(%d, %X, %c) ", i, c, c)
	}
	fmt.Println()

	s6 := "\xE6\xB1\xBD"
	fmt.Printf("s6 = %s, len(s6) = %d, unicode: %X, utf: %X\n", s6, len(s6), []rune(s6)[0], s6)

	s7 := "汽"
	fmt.Printf("s7 = %s, len(s7) = %d, unicode: %X, utf: %X\n", s7, len(s7), []rune(s7)[0], s7)

	// 字符串分割
	s8 := "Golang,Java,C++,Scala"
	langs := strings.Split(s8, ",")
	for i, lang := range langs {
		fmt.Printf("(%d, %s) ", i, lang)
	}
	fmt.Println()

	// 字符串合并
	fmt.Println(strings.Join(langs, "/"))

	// 其他操作
	fmt.Println(strings.Contains(s8, "Go"))
	fmt.Println(strings.Index(s8, "Go"))
	fmt.Println(strings.Repeat("Go", 3))
	fmt.Println(strings.Trim("  hello golang  ", " "))
}

func convertDemo() {
	fmt.Println("=== convertDemo")

	si1 := strconv.Itoa(110)
	si2 := string(110)
	si3 := strconv.FormatInt(int64(110), 10)
	si4 := strconv.FormatInt(int64(110), 16)

	fmt.Printf("si1 = %v, si1 == '110', %v\n", si1, si1 == "110")
	fmt.Printf("si2 = %v, si2 == '110', %v\n", si2, si2 == "110")
	fmt.Printf("si3 = %v, si3 == '110', %v\n", si3, si3 == "110")
	fmt.Printf("si4 = %v, si4 == '6e', %v\n", si4, si4 == "6e")

	b := byte(1)
	sb := strconv.Itoa(int(b))
	fmt.Printf("sb = %v, ths type of sb is %T\n", sb, sb)

	i1, _ := strconv.Atoi("110")
	i2, _ := strconv.ParseInt("110", 10, 64)
	i3, _ := strconv.ParseInt("110", 10, 32)
	fmt.Printf("i1 = %v, the type of i1 is %T\n", i1, i1)
	fmt.Printf("i2 = %v, the type of i2 is %T\n", i2, i2)
	fmt.Printf("i3 = %v, the type of i3 is %T\n", i3, i3)

	// float64
	f1, _ := strconv.ParseFloat("3.14159265", 32)
	f2, _ := strconv.ParseFloat("3.14159265", 64)
	fmt.Printf("f1 = %v, the type of f1 is %T\n", f1, f1)
	fmt.Printf("f2 = %v, the type of f2 is %T\n", f2, f2)

	// complex128
	c1 := 2 + 3.14i
	fmt.Printf("c1 = %v, the type of c1 is %T\n", c1, c1)
	fmt.Printf("the real of c1 is: %v, the imag of c1 is %v\n", real(c1), imag(c1))
}

func arrayDemo() {
	fmt.Println("=== arrayDemo")

	var arr1 [3]int64
	arr1[0], arr1[1], arr1[2] = 1, 2, 3
	arr2 := [3]string{"red", "blue", "green"}
	arr3 := []string{"spark", "yarn", "linux"}

	// slice
	arr4 := [...]string{"hot", "cold"}
	arr5 := [2][2]int64{
		{100, 20},
		{40, 10},
	}

	fmt.Printf("arr1 = %v, the type of arr1 is %T\n", arr1, arr1)
	fmt.Printf("arr2 = %v, the type of arr2 is %T\n", arr2, arr2)
	fmt.Printf("arr3 = %v, the type of arr3 is %T\n", arr3, arr3)
	fmt.Printf("arr4 = %v, the type of arr4 is %T\n", arr4, arr4)
	fmt.Printf("arr5 = %v, the type of arr5 is %T\n", arr5, arr5)

	fmt.Printf("arr2 = %v, the len of arr2 is %T\n", arr2, len(arr2))
	// arr2[3] = "black"
	fmt.Printf("arr2 = %v, the type of arr2 is %T\n", arr2, arr2)
}

func sliceDemo() {
	fmt.Println("=== sliceDemo")

	// slice1, slice2 共用底层结构
	slice1 := []string{"lakers", "thunders", "knicks", "hawks", "heat"}
	// slice1 := [...]string{"lakers", "thunders", "knicks", "hawks", "heat"}
	slice2 := slice1[0:3]

	slice3 := make([]int, 0)

	for i := 0; i < 10; i++ {
		slice3 = append(slice3, i)
	}

	slice4 := append(slice1, slice2...)

	fmt.Printf("slice1 = %v, the type of slice1 is %T, the len of slice1 is %v, the cap of slice1 is %v\n", slice1, slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, the type of slice2 is %T, the len of slice2 is %v, the cap of slice2 is %v\n", slice2, slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3 = %v, the type of slice3 is %T, the len of slice3 is %v, the cap of slice3 is %v\n", slice3, slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4 = %v, the type of slice4 is %T, the len of slice4 is %v, the cap of slice4 is %v\n", slice4, slice4, len(slice4), cap(slice4))
}

func mapDemo() {
	fmt.Println("=== mapDemo")

	map1 := make(map[string]int)
	map1["lakers"] = 16
	map1["bulls"] = 6
	map1["spurs"] = 5

	fmt.Printf("map1 = %v, the type of map1 is %T\n", map1, map1)

	if v, ok := map1["bulls"]; ok {
		fmt.Printf("key exists, value is %v\n", v)
	}

	delete(map1, "bulls")
	fmt.Printf("map1 = %v, the type of map1 is %T\n", map1, map1)

	for k, v := range map1 {
		fmt.Printf("map[%s] = %v\n", k, v)
	}
}

type TreeNode struct {
	value int
	left, right *TreeNode
}

// 传值，不改变原变量的值，有复制的开销
func (node TreeNode) print() {
	fmt.Println(node.value)
}

// 传指针
func (node *TreeNode) setValue(value int) {
	node.value = value
}

func (node *TreeNode) traverse() {
	if node == nil {
		return
	}

	node.left.traverse()
	fmt.Print(node.value, " ")
	node.right.traverse()
}


// 通过包装实现扩展
type NewTreeNode struct {
	node *TreeNode
}


// 通过别名实现扩展
type Queue []int

func (queue *Queue) push(v int) {
	*queue = append(*queue, v)
}

func (queue *Queue) pop() int {
	head := (*queue)[0]
	*queue = (*queue)[1:]
	return head
}

func (queue *Queue) isEmpty() bool {
	return len(*queue) == 0
}

func objectDemo() {
	fmt.Println("=== objectDemo")

	root := TreeNode{value: 10}
	root.left = &TreeNode{value: 20}
	root.right = &TreeNode{30, nil, nil}
	root.right.left = new(TreeNode)
	// root.right.right = new(TreeNode){value: 40}

	root.traverse()
	fmt.Println()
}
