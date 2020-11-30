package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Sender: send elem: %v\n", i)
			ch <- i
		}

		fmt.Printf("Sender: close the channel\n")

		close(ch)
	}()

	for {
		elem, ok := <- ch

		if !ok {
			fmt.Printf("Receiver: channel close\n")
			break;
		}

		fmt.Printf("Receiver: receive elem: %v\n", elem)
	}

	fmt.Println("End...")
}
