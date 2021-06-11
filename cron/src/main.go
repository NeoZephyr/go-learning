package main

import (
	"fmt"
	"net"
	"sort"
	"time"
)

type ScanResult struct {
	port int
	opened bool
}

func worker(ports chan int, result chan ScanResult) {
	for port := range ports {
		address := fmt.Sprintf("cdh:%d", port)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			result <- ScanResult{port: port, opened: false}
			continue
		}

		result <- ScanResult{port: port, opened: true}
		conn.Close()
	}
}

func main() {
	now := time.Now()

	ports := make(chan int, 100)
	result := make(chan ScanResult)
	var openedPorts []int
	var closedPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, result)
	}

	go func() {
		for i := 20; i < 1024; i++ {
			ports <- i
		}
	}()

	for i := 20; i < 1024; i++ {
		scanResult := <- result
		if scanResult.opened {
			openedPorts = append(openedPorts, scanResult.port)
		} else {
			closedPorts = append(closedPorts, scanResult.port)
		}
	}

	close(ports)
	close(result)

	sort.Ints(openedPorts)
	sort.Ints(closedPorts)

	for _, port := range openedPorts {
		fmt.Printf("%d opened\n", port)
	}

	for _, port := range closedPorts {
		fmt.Printf("%d closed\n", port)
	}

	duration := time.Since(now) / 1e9
	fmt.Printf("duration: %d\n", duration)
}
