package main

import (
	"bytes"
	"fmt"
)

func main() {
	bytesApp()
}

func bytesApp() {
	var buffer bytes.Buffer

	fmt.Printf("buffer, len: %d, cap: %d\n", buffer.Len(), buffer.Cap())

	buffer.WriteString("Hello bytes")

	fmt.Printf("buffer, len: %d, cap: %d\n", buffer.Len(), buffer.Cap())

	buf := make([]byte, 5)
	n, _ := buffer.Read(buf)

	fmt.Printf("buffer, len: %d, cap: %d, read: %d\n", buffer.Len(), buffer.Cap(), n)

	var pBuffer = bytes.NewBufferString("Hello buffer")

	fmt.Printf("buffer, len: %d, cap: %d\n", pBuffer.Len(), pBuffer.Cap())

	// 10, 20, 30
	pBuffer.Grow(21)

	fmt.Printf("buffer, len: %d, cap: %d\n", pBuffer.Len(), pBuffer.Cap())
}
