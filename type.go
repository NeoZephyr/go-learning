// 1. 同一个目录，包名必需一样
// 2. 包可以与目录不同名

package main

import "fmt"
import "unicode/utf8"
import "strings"

func main() {
	constApp()
	bitApp()
	boolApp()
	byteApp()
	intApp()
	stringApp()
	optStringApp()
	typeAliasApp()
	typeMatchApp()
}

func constApp() {
	fmt.Println()
	fmt.Println("=== const app")

	const rateLimit = 5000

	fmt.Printf("rateLimit is %d, type is %T\n", rateLimit, rateLimit)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)

	fmt.Printf("b = %d, kb = %d, mb = %d, gb = %d\n", b, kb, mb, gb)
}

func bitApp() {
	fmt.Println()
	fmt.Println("=== big app")

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

func boolApp() {
	fmt.Println()
	fmt.Println("=== bool App")

	// 默认 false
	var b1 bool
	b2 := true
	b3 := false

	fmt.Printf("b1: %v, b2: %v, b3: %v\n", b1, b2, b3)
}

func byteApp() {
	fmt.Println()
	fmt.Println("=== byte App")

	// uint8
	var b1 byte = 'a'

	// int32
	b2 := 'b'

	fmt.Printf("b1 = %v, b1 = %c, the type of b1 is %T\n", b1, b1, b1)
	fmt.Printf("b2 = %v, b2 = %c, the type of b2 is %T\n", b2, b2, b2)
	fmt.Printf("b1 - 32 = %v\n", b1-32)
}

func intApp() {
	fmt.Println()
	fmt.Println("=== int App")

	var i0 int
	var i1 int64

	// int
	i2 := 100
	i3 := int32(200)
	i4 := int64(300)

	fmt.Printf("i0 = %v, the type of i0 is %T\n", i0, i0)
	fmt.Printf("i1 = %v, the type of i1 is %T\n", i1, i1)
	fmt.Printf("i2 = %v, the type of i2 is %T\n", i2, i2)
	fmt.Printf("i3 = %v, the type of i3 is %T\n", i3, i3)
	fmt.Printf("i4 = %v, the type of i4 is %T\n", i4, i4)
}

func stringApp() {
	fmt.Println()
	fmt.Println("=== string App")

	// 默认空字符串
	var s0 string
	s1 := "\"Hello Go\""
	s2 := `"Hello Go"`
	s3 := `
	Hello Go
	`
	s4 := "go 语言"

	fmt.Printf("s0 = %v, the type of s0 is %T\n", s0, s0)
	fmt.Printf("s1 = %v, the type of s1 is %T\n", s1, s1)
	fmt.Printf("s2 = %v, the type of s2 is %T\n", s2, s2)
	fmt.Printf("s3 = %v, the type of s3 is %T\n", s3, s3)
	fmt.Printf("s4 = %v, the type of s4 is %T\n", s4, s4)

	// 获取字节、字符长度
	fmt.Printf("byte len of s4 is %d, char len of s4 is %d\n",
		len(s4),
		utf8.RuneCountInString(s4))

	// 打印 数值、unicode、字符
	fmt.Println("range:")
	for i, item := range s4 {
		fmt.Printf("(%d, %X, %c) ", i, item, item)
	}
	fmt.Println()

	fmt.Println("byte range:")
	for i, item := range []byte(s4) {
		fmt.Printf("(%d, %X, %c) ", i, item, item)
	}
	fmt.Println()

	// int32
	fmt.Println("rune range:")
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

func optStringApp() {
	fmt.Println("=== opt string App")
	fmt.Println()

	s0 := "Golang,Java,C++,Scala"

	// 字符串分割
	langs := strings.Split(s0, ",")

	for i, lang := range langs {
		fmt.Printf("(%d, %s) ", i, lang)
	}
	fmt.Println()

	// 字符串合并
	fmt.Printf("langs: %s\n", strings.Join(langs, "/"))

	// 其他操作
	fmt.Printf("s0 contains Go: %v\n", strings.Contains(s0, "Go"))
	fmt.Printf("repeat three times Go: %v\n", strings.Repeat("Go", 3))
	fmt.Printf("trim string: %v\n", strings.Trim("  Hello Golang  ", " "))
}

func typeAliasApp() {
	fmt.Println("=== type alias App")
	fmt.Println()

	type AString = string
	str := "a string"
	aStr := AString(str)
	fmt.Printf("%T(%q) == %T(%q): %v\n", str, str, aStr, aStr, str == aStr)

	strs := []string{"E", "F", "G"}
	aStrs := []AString(strs)
	fmt.Printf("%T(%q) == %T(%q)\n", strs, strs, aStrs, aStrs)

	type BString string
	str = "b string"
	bStr := BString(str)
	fmt.Printf("%T(%q) != %T(%q)\n", str, str, bStr, bStr)
}

func typeMatchApp() {
	fmt.Println("=== type match App")
	fmt.Println()

	m := map[string]int{"lakers": 17, "heat": 3}

	switch n := interface{}(m).(type) {
	case []string:
		fmt.Printf("match []string, %v\n", n)
	case map[string]int:
		fmt.Printf("match map[string]int, %v\n", n)
	default:
		fmt.Printf("unsupported type, %v\n", n)
		return
	}
}

