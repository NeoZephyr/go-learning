package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// bitApp()
	// byteApp()
	// intApp()
	// stringApp()
	// typeAliasApp()
	typeMatchApp()
}

func bitApp() {
	const (
		read = 1 << iota
		write
		execute
	)

	permission := 7

	// 注意多行格式
	fmt.Printf("read: %v, write: %v, execute: %v\n",
		(permission&read == read),
		(permission&write == write),
		(permission&execute == execute))

	permission = permission &^ execute

	fmt.Printf("read: %v, write: %v, execute: %v\n",
		(permission&read == read),
		(permission&write == write),
		(permission&execute == execute))
}

func byteApp() {
	// uint8
	var b1 byte = 'a'

	// int32
	b2 := 'b'

	fmt.Printf("b1 = %v, b1 = %c, the type of b1 is %T\n", b1, b1, b1)
	fmt.Printf("b2 = %v, b2 = %c, the type of b2 is %T\n", b2, b2, b2)
	fmt.Printf("b1 - 32 = %v\n", b1-32)
}

func intApp() {
	// int
	i0 := 100
	i1 := int32(200)
	i2 := int64(300)

	fmt.Printf("i0 = %v, the type of i0 is %T, size: %d\n", i0, i0, unsafe.Sizeof(i0))
	fmt.Printf("i1 = %v, the type of i1 is %T, size: %d\n", i1, i1, unsafe.Sizeof(i1))
	fmt.Printf("i2 = %v, the type of i2 is %T, size: %d\n", i2, i2, unsafe.Sizeof(i2))
}

func stringApp() {
	// 默认空字符串
	var s0 string
	s1 := "\"Hello Go\""
	s2 := `"Hello Go"`
	s3 := `
	Hello Go
	`
	s4 := "go 语言"

	fmt.Printf("s0 = %v\ns1 = %v\ns2 = %v\ns3 = %v\ns4 = %v\n", s0, s1, s2, s3, s4)

	// 获取字节、字符长度
	fmt.Printf("byte len of s4 is %d, char len of s4 is %d\n",
		len(s4),
		utf8.RuneCountInString(s4))

	// 打印数值、unicode、字符
	fmt.Printf("range: ")
	for i, item := range s4 {
		fmt.Printf("(%d, %X, %c) ", i, item, item)
	}
	fmt.Println()

	fmt.Printf("byte range: ")
	for i, item := range []byte(s4) {
		fmt.Printf("(%d, %X, %c) ", i, item, item)
	}
	fmt.Println()

	// int32
	fmt.Printf("rune range: ")
	for i, item := range []rune(s4) {
		fmt.Printf("(%d, %X, %c) ", i, item, item)
	}
	fmt.Println()

	s5 := "\xE6\xB1\xBD"
	fmt.Printf("s5: %v, len of s5: %d, unicode of s5: %X, utf of s5: %X\n",
		s5, len(s5), []rune(s5)[0], s5)

	s6 := "汽"
	fmt.Printf("s6: %v, len of s6: %d, unicode of s6: %X, utf of s6: %X\n",
		s6, len(s6), []rune(s6)[0], s6)
}

func typeAliasApp() {
	// 别名类型
	type AString = string
	str := "a string"
	aStr := AString(str)
	fmt.Printf("str: %v, type of str: %T\naStr: %v, type of aStr: %T\n(str == aStr): %v\n", str, str, aStr, aStr, (str == aStr))

	strs := []string{"E", "F", "G"}
	aStrs := []AString(strs)
	fmt.Printf("strs: %v, type of strs: %T\naStrs: %v, type of aStrs: %T\n(strs == aStrs): %v\n", strs, strs, aStrs, aStrs, reflect.DeepEqual(strs, aStrs))

	// 类型再定义
	type BString string
	str = "b string"
	bStr := BString(str)
	fmt.Printf("str: %v, type of str: %T\nbStr: %v, type of bStr: %T\n", str, str, bStr, bStr)
	// fmt.Printf("str: %v, type of str: %T\nbStr: %v, type of bStr: %T\n(str == bStr): %v\n", str, str, bStr, bStr, (str == bStr))
}

func typeMatchApp() {
	m := map[string]int{"lakers": 17, "heat": 3}

	switch n := interface{}(m).(type) {
	case []string:
		fmt.Printf("match type: []string, value: %v\n", n)
	case map[string]int:
		fmt.Printf("match type: map[string]int, value: %v\n", n)
	default:
		fmt.Printf("unsupported type, value: %v\n", n)
		return
	}

	n, ok := interface{}(m).(map[string]int)
	fmt.Printf("match ok: %v, value: %v\n", ok, n)
}
