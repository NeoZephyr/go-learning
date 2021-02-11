package main

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

func main() {
	// stringApp()
	// builderApp()
	readerApp()
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
		fmt.Printf("(%d, %X, %c, [% x]) ", i, item, item, []byte(string(item)))
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

func builderApp() {
	var builder strings.Builder
	builder.WriteString("hello go.")
	builder.WriteByte('\n')
	fmt.Printf("builder, len: %d, cap: %d, content: %q\n", builder.Len(), builder.Cap(), builder.String())

	builder.Grow(20)
	fmt.Printf("builder, len: %d, cap: %d, content: %q\n", builder.Len(), builder.Cap(), builder.String())

	// panic
	// func(builder strings.Builder) {
	// 	builder.Grow(20)
	// }(builder)

	builder.Reset()
	fmt.Printf("builder, len: %d, cap: %d, content: %q\n", builder.Len(), builder.Cap(), builder.String())

	func(builder strings.Builder) {
		builder.Grow(20)
		fmt.Printf("builder, len: %d, cap: %d, content: %q\n", builder.Len(), builder.Cap(), builder.String())
	}(builder)
}

func readerApp() {
	reader := strings.NewReader("This is the official reference guide for the HBase version it ships with")
	fmt.Printf("reader, size: %d, reading index: %d\n", reader.Size(), reader.Size()-int64(reader.Len()))

	buf := make([]byte, 12)
	n, _ := reader.Read(buf)
	fmt.Printf("reader, size: %d, reading index: %d, read bytes: %d, buf: %s\n", reader.Size(), reader.Size()-int64(reader.Len()), n, buf)

	n, _ = reader.Read(buf)
	fmt.Printf("reader, size: %d, reading index: %d, read bytes: %d, buf: %s\n", reader.Size(), reader.Size()-int64(reader.Len()), n, buf)

	offset := int64(12)
	n, _ = reader.ReadAt(buf, offset)
	fmt.Printf("reader, size: %d, reading index: %d, read bytes: %d, buf: %s\n", reader.Size(), reader.Size()-int64(reader.Len()), n, buf)

	readingIndex, _ := reader.Seek(offset, io.SeekCurrent)
	fmt.Printf("reader, size: %d, reading index: %d, read bytes: %d, readingIndex: %d\n", reader.Size(), reader.Size()-int64(reader.Len()), n, readingIndex)
}
