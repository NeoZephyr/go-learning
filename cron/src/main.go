package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	now := time.Now()

	for i := 20; i < 200; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("180.101.49.11:%d", port)
			conn, err := net.Dial("tcp", address)

			if err != nil {
				fmt.Printf("%s connect error: %s\n", address, err.Error())
				return
			}
			defer conn.Close()
			fmt.Printf("%s connect ok\n", address)
		}(i)
	}
	wg.Wait()
	duration := time.Since(now) / 1e9
	fmt.Printf("duration: %d\n", duration)
}
