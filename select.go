package main

import (
	"fmt"
)

func main() {
	chs := [3]chan int {
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	chs[0] <- 100
	chs[1] <- 300

	// 如果前面的表达式都阻塞了，那么默认分支就会被选中并执行
	// 如果没有加入默认分支，一旦所有 case 表达式都不满足条件，select 语句就会被阻塞。直到至少有一个e表达式满足条件为止
	select {
	case elem, ok := <- chs[0]:
		fmt.Printf("first case, elem: %v, ok: %v\n", elem, ok)
	case elem, ok := <- chs[1]:
		fmt.Printf("second case, elem: %v, ok: %v\n", elem, ok)
	case elem, ok := <- chs[2]:
		fmt.Printf("third case, elem: %v, ok: %v\n", elem, ok)
	default:
		fmt.Println("default case")
	}
}
