package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	testServer()
}

func testServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	go sendAll()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

type Client struct {
	Buf chan string
	Name string
	Address string
}

var clients map[string]Client

var msgBuf = make(chan string)

func handleConn(conn net.Conn) {
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	client := Client{make(chan string), clientAddr, clientAddr}
	clients[clientAddr] = client

	msgBuf <- fmt.Sprintf("[%s] login at %v\n", clientAddr, time.Now().Format("2006-1-02 15:04"))

	quit := make(chan int)
	alive := make(chan int)

	go receive(alive, quit, client, conn)
	go send(client, conn)

	for {
		select {
		case <-alive:
		case <- quit:
			delete(clients, clientAddr)
			msgBuf <- fmt.Sprintf("[%s] logout at %v\n", clientAddr, time.Now().Format("2006-1-02 15:04"))
			return
		case <- time.After(30 * time.Second):
			delete(clients, clientAddr)
			msgBuf <- fmt.Sprintf("[%s] timeout at %v, please reconnect\n", clientAddr, time.Now().Format("2006-1-02 15:04"))
			return
		}
	}
}

func sendAll() {
	clients = make(map[string]Client, 1000)

	for {
		message := <- msgBuf

		for _, client := range clients {
			client.Buf <- message
		}
	}
}

func send(client Client, conn net.Conn) {
	for msg := range client.Buf {
		conn.Write([]byte(msg))
	}
}

func receive(alive chan int, quit chan int, client Client, conn net.Conn) {
	for {
		buf := make([]byte, 1024 * 32)
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println(err)
			quit <- 0
			break
		}

		if n != 0 {
			buf = buf[:n]
			msgBuf <- fmt.Sprintf("[%s] send %s at %v\n", client.Address, strings.Trim(string(buf), "\n"), time.Now().Format("2006-1-02 15:04"))
		}

		alive <- 0
	}
}