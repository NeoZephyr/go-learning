package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//testServer1()
	testServer2()
}

func testServer1() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")

	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := listener.Accept()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("buf:", string(buf[:n]))
}

func testServer2() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println("accept connection from", addr)

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println(err)
			break
		}

		msg := strings.Trim(string(buf[:n]), "\n")

		if msg == "quit" {
			fmt.Printf("%s disnonnected from server\n", addr)
			break
		}

		fmt.Printf("receive %s from %s\n", msg, addr)
		conn.Write([]byte(strings.ToUpper(msg)))
	}
}