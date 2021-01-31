package main

import (
	"log"
	"sync"
	"time"
)

type counter struct {
	number uint
	mutex  sync.RWMutex
}

func (c *counter) num() uint {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.number
}

func (c *counter) incr(i uint) uint {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.number += i
	return c.number
}

func main() {
	// countApp()
	sendAndRecvApp()
}

func send(lock *sync.Mutex, sendCond *sync.Cond, recvCond *sync.Cond, mailbox *uint8, id int, index int) {
	lock.Lock()

	for *mailbox == 1 {
		sendCond.Wait()
	}

	log.Printf("sender%d [%d]: empty", id, index)

	*mailbox = 1

	log.Printf("sender%d [%d]: the letter has been sent", id, index)

	lock.Unlock()
	recvCond.Broadcast()
}

func recv(lock *sync.Mutex, sendCond *sync.Cond, recvCond *sync.Cond, mailbox *uint8, id int, index int) {
	lock.Lock()

	for *mailbox == 0 {
		recvCond.Wait()
	}

	log.Printf("receiver%d [%d]: full", id, index)

	*mailbox = 0

	log.Printf("receiver%d [%d]: the letter has been received", id, index)

	lock.Unlock()
	sendCond.Broadcast()
}

func sendAndRecvApp() {
	var mailbox uint8
	var lock sync.Mutex

	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(&lock)

	sign := make(chan struct{}, 3)
	max := 6

	go func(id int, max int) {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 0; i < max; i++ {
			time.Sleep(time.Millisecond * 500)
			send(&lock, sendCond, recvCond, &mailbox, id, i)
		}
	}(0, max)

	go func(id int, max int) {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 0; i < max; i++ {
			time.Sleep(time.Millisecond * 200)
			recv(&lock, sendCond, recvCond, &mailbox, id, i)
		}
	}(1, max/2)

	go func(id int, max int) {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 0; i < max; i++ {
			time.Sleep(time.Millisecond * 200)
			recv(&lock, sendCond, recvCond, &mailbox, id, i)
		}
	}(2, max/2)

	<-sign
	<-sign
	<-sign
}

func countApp() {
	c := counter{}
	count(&c)
}

func count(c *counter) {
	sign := make(chan struct{}, 3)

	go func() {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 500)
			c.incr(1)
		}
	}()

	go func() {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 200)
			log.Printf("Go 1 read number in counter: [%d - %d]", i, c.num())
		}
	}()

	go func() {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 300)
			log.Printf("Go 2 read number in counter: [%d - %d]", i, c.num())
		}
	}()

	<-sign
	<-sign
	<-sign
}
