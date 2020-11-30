package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//testNoSafe()
	testSafe()
}

type noSafeInt int

func (i *noSafeInt) increment() {
	*i++
}

func (i *noSafeInt) get() noSafeInt {
	return *i
}

func testNoSafe() {
	var i noSafeInt = 0

	for n := 0; n < 1000; n++ {
		go func() {
			i.increment()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("i = ", i.get())
}

type safeInt struct {
	value int
	lock sync.Mutex
}

func (s *safeInt) increment() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.value++
}

func (s *safeInt) get() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.value
}

func testSafe() {
	var s safeInt

	for n := 0; n < 1000; n++ {
		go func() {
			s.increment()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("s = ", s.get())
}