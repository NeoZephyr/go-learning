package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

func main() {
	//testOnce()
	//testFirstResponse()
	//testAllResponse()
	testPool()
}

type Singleton struct {}

var singleton *Singleton
var once sync.Once

func testOnce() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			singleton := createSingleton()
			fmt.Printf("address: %x\n", unsafe.Pointer(singleton))
			wg.Done()
		}()
	}

	wg.Wait()
}

func createSingleton() *Singleton {

	once.Do(func() {
		fmt.Println("create singleton")
		singleton = new(Singleton)
	})

	return singleton
}

func firstResponse() string {

	// 防止其它线程阻塞
	c := make(chan string, 5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			time.Sleep(50 * time.Millisecond)
			res := fmt.Sprintf("The result from thread %d", i)
			c <- res
		}(i)
	}

	return <- c
}

func testFirstResponse() {
	fmt.Println("Before:", runtime.NumGoroutine())
	res := firstResponse()
	fmt.Println("res:", res)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("After:", runtime.NumGoroutine())
}

func allResponse() string {
	c := make(chan string)

	for i := 0; i < 5; i++ {
		go func(i int) {
			time.Sleep(50 * time.Millisecond)
			res := fmt.Sprintf("The result from thread %d", i)
			c <- res
		}(i)
	}

	allRes := ""

	for i := 0; i < 5; i++ {
		allRes += <- c + "\n"
	}

	return allRes
}

func testAllResponse() {
	fmt.Println("Before:", runtime.NumGoroutine())
	res := allResponse()
	fmt.Println("res:", res)
	fmt.Println("After:", runtime.NumGoroutine())
}

func testPool() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new object")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)

	pool.Put(3)
	runtime.GC()
	v = pool.Get().(int)
	fmt.Println(v)
}