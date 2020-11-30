package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//testClient1()
	testClient2()
}

func testClient1() {
	conn, err := net.Dial("tcp", "127.0.0.1:8088")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	conn.Write([]byte("hello go server"))
}

func testClient2() {
	conn, err := net.Dial("tcp", "127.0.0.1:8088")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	go func() {
		output := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(output)

			if err != nil {
				fmt.Println(err)
				continue
			}

			conn.Write(output[:n])
		}
	}()

	input := make([]byte, 1024)
	addr := conn.LocalAddr().String()

	for {
		n, err := conn.Read(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s receive %s\n", addr, string(input[:n]))
	}
}
