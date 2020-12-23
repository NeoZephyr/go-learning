package main

import "fmt"

// 对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的
// 发送操作和接收操作中对元素值的处理都是不可分割的
// 发送操作在完全完成之前会被阻塞。接收操作也是如此

// 对于值为 nil 的通道，对它的发送操作和接收操作都会永久地处于阻塞状态
// 通道一旦关闭，再对它进行发送操作，就会引发 panic
// 关闭一个已经关闭了的通道，会引发 panic
// 接收操作可以感知到通道的关闭，并能够安全退出

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

	sch := getIntChan()

	for elem := range sch {
		fmt.Printf("Receiver: receive elem: %v\n", elem)
	}

	fmt.Println("End...")
}

func getIntChan() <- chan int {
	ch := make(chan int, 5)

	for i := 1; i < 5; i++ {
		ch <- i * 100
	}

	close(ch)

	return ch
}
