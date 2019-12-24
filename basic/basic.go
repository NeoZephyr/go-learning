// 同一个目录，包名必需一样
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 包内部变量
var (
	count = 100
	topic = "rule_process"
)

func main() {
	//testVar()
	//testInput()
	//testConst()
	testType()
	//testSwitch('F')
	//testIf("README.md")
	//testFor("README.md")
	//testPointer()
	//testClearBit()
	testString()

	// 命令行参数
	fmt.Println(os.Args)

	// 返回值
	// 不会调用 defer
	os.Exit(-1)
}

func testVar()  {
	var ia int
	fmt.Printf("ia = %d, the type of ia is %T\n", ia, ia)

	// string 是值类型，默认值为空字符串
	var sa string
	fmt.Printf("sa = %q, the type of sa is %T\n", sa, sa)

	var ib = 3
	fmt.Printf("ib = %d, the type of ib is %T\n", ib, ib)

	var sb = "hello world"
	fmt.Printf("sa = %s, the type of sa is %T\n", sb, sb)

	var ic, id, sc = 11, 23, "hello go"
	fmt.Println(ic, id, sc)

	ie, ig, sd := 100, 200, "hello go world"
	fmt.Println(ie, ig, sd)

	ii, _, ij := 300, 400, 500
	fmt.Println(ii, ij)

	var (
		im = 600
		in = 700
		se = "hello go var"
	)
	fmt.Println(im, in, se)

	var ba bool
	fmt.Printf("ba = %v, the type of ba is %T\n", ba, ba)

	bb := true
	fmt.Printf("bb = %v, the type of bb is %T\n", bb, bb)

	var fa float64 = 3.14
	fmt.Printf("fa = %f, the type of fa is %T\n", fa, fa)
	fmt.Printf("fa = %v, the type of fa is %T\n", fa, fa)
	fb := 3.1415926
	fmt.Printf("fb = %.2f, the type of fb is %T\n", fb, fb)

	var za complex128 = 2 + 5.12i
	fmt.Printf("za = %v, the type of za is %T\n", za, za)
	fmt.Printf("the real of za is: %v, the imag of za is %v\n", real(za), imag(za))
	zb := 5 - 3.14i
	fmt.Printf("zb = %v, the type of zb is %T\n", zb, zb)
	fmt.Printf("the real of zb is: %v, the imag of zb is %v\n", real(zb), imag(zb))

	var ca byte = 'a'
	fmt.Printf("ca = %v, ca = %c, the type of ca is %T\n", ca, ca, ca)
	fmt.Printf("ca - 32 = %v\n", ca - 32)
	cb := 'b'
	fmt.Printf("cb = %v, the type of cb is %T\n", cb, cb)
}

func testConst()  {
	const UploadLimit = 5000
	fmt.Printf("UploadLimit is %d, the type of UploadLimit is %T\n", UploadLimit, UploadLimit)

	const (
		ListProcessingLimit = 20
		TagProcessingLimit = 10
	)

	fmt.Printf("ListProcessingLimit is %d, the type of ListProcessingLimit is %T\n", ListProcessingLimit, ListProcessingLimit)
	fmt.Printf("TagProcessingLimit is %d, the type of TagProcessingLimit is %T\n", TagProcessingLimit, TagProcessingLimit)

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

func testInput()  {
	var num int
	fmt.Println("please input num:")
	//fmt.Scanf("%d", &num)
	fmt.Scan(&num)
	fmt.Printf("the input num is %d\n", num)
}

func testType() {
	type bigint int64
	var ba bigint = 100
	fmt.Printf("ba = %d, the type of ba is %T\n", ba, ba)

	type (
		long int64
		char byte
	)

	var la long = 200
	fmt.Printf("la = %v, the type of la is %T\n", la, la)
	var ca char = 'A'
	fmt.Printf("ca = %v, the type of ca is %T\n", ca, ca)

	// 函数类型
	type compute func(a, b int) int
	var c compute = pow
	fmt.Println(c(2, 7))
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func testIf(filename string)  {
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("content:\n%s\n", content)
	}
}

func testSwitch(score byte) {
	switch score {
	case 'A':
		fmt.Printf("90 - 100")
	case 'B':
		fmt.Printf("80 - 89")
	case 'C':
		fmt.Printf("70 - 79")
	case 'D', 'E', 'F':
		fallthrough
	default:
		fmt.Printf("not passed")
	}
}

func testFor(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func swap(a, b *float64)  {
	*b, *a = *a, *b
}

func testPointer()  {
	a, b := 2.78, 3.14
	fmt.Printf("before swap, a = %v, b = %v\n", a, b)

	swap(&a, &b)
	fmt.Printf("after swap, a = %v, b = %v\n", a, b)
}

func testString()  {
	s := "Pain 喜欢 go 语言"

	// 获取字节长度
	fmt.Printf("s = %s, len(s) = %d\n", s, len(s))

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s {
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println()

	// 获取字符数量
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}
	fmt.Println()

	s = "\xE6\xB1\xBD"
	fmt.Printf("s = %s, len(s) = %d\n", s, len(s))

	fmt.Printf("unicode %x\n",[]rune(s)[0])
	fmt.Printf("utf8 %x\n", s)

	s = "汽"
	fmt.Printf("s = %s, len(s) = %d\n", s, len(s))

	fmt.Printf("unicode %x\n",[]rune(s)[0])
	fmt.Printf("utf8 %x\n", s)

	// 存储任意二进制数据
	s = "\xE6\xC1\xBE\xCF"
	fmt.Printf("s = %s, len(s) = %d\n", s, len(s))

	// 字符串分割
	s = "Golan,Java,C++"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		fmt.Printf("%s ", part)
	}
	fmt.Println()

	// 字符串合并
	fmt.Println(strings.Join(parts, "/"))

	count := 10
	fmt.Printf("int convert to string: %s\n", strconv.Itoa(count) + " apples")
	countStr := "109"
	if value, ok := strconv.Atoi(countStr); ok == nil {
		fmt.Printf("string convert to int: %d\n", value)
	} else {
		fmt.Printf("string convert to int failed, %s\n", ok)
	}

	fmt.Println(strings.Contains("hello go", "go"))
	fmt.Println(strings.Index("hello go", "Go"))
	fmt.Println(strings.Repeat("Go", 3))
	fmt.Println(strings.Trim("  hello go  ", " "))
}

func testClearBit() {
	const (
		read = 1 << iota
		write
		execute
	)

	permission := 7
	fmt.Printf("read: %v, write: %v, execute: %v\n",
		(permission & read == read), (permission & write == write), (permission & execute == execute))
	permission = permission &^ execute
	fmt.Printf("read: %v, write: %v, execute: %v\n",
		(permission & read == read), (permission & write == write), (permission & execute == execute))
}