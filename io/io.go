package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//testStandard()
	//testWriteFile()
	//testReadFile()
	//testReadFileByLine()

	testFileStat()
}

func testStandard() {
	var num int
	fmt.Println("please input num:")
	fmt.Scan(&num)
	fmt.Println("input num:", num)

	os.Stdout.WriteString("hello world")
	os.Stdout.Close()
	os.Stdout.WriteString("hello world")
}

func testWriteFile() {
	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	file.WriteString("hello world\n")
	file.WriteString("hello pain\n")
}

func testReadFile() {
	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	buf := make([]byte, 2 * 1024)
	n, err := file.Read(buf)

	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf[:n]))
}

func testReadFileByLine() {
	file, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		// \n 也会读取
		bytes, err := reader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		fmt.Println("line: ", string(bytes))
	}
}

func testFileStat() {
	info, err := os.Stat("io/io.go")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(info.Name())
	fmt.Println(info.Size())
}