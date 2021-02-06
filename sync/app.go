package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"sync/atomic"
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

type Buffer interface {
	Delimiter() byte

	Write(content string) (err error)

	Read() (content string, err error)

	Free()
}

type delimiterBuffer struct {
	buf       bytes.Buffer
	delimiter byte
}

var bufPool sync.Pool
var delimiter = byte('\n')

func (b *delimiterBuffer) Delimiter() byte {
	return b.delimiter
}

func (b *delimiterBuffer) Write(content string) (err error) {
	if _, err := b.buf.WriteString(content); err != nil {
		return err
	}

	return b.buf.WriteByte(b.delimiter)
}

func (b *delimiterBuffer) Read() (content string, err error) {
	return b.buf.ReadString(b.delimiter)
}

func (b *delimiterBuffer) Free() {
	bufPool.Put(b)
}

func getBuffer() Buffer {
	return bufPool.Get().(Buffer)
}

func init() {
	bufPool = sync.Pool{
		New: func() interface{} {
			return &delimiterBuffer{delimiter: delimiter}
		},
	}
}

func main() {
	// mutexCountApp()
	// sendAndRecvApp()
	// chanCountApp()
	// wgCountApp()
	// contextCountApp()

	// syncPoolApp()
}

func mutexCountApp() {
	c := counter{}
	count(&c)
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

func chanCountApp() {
	sign := make(chan struct{}, 2)
	count := int32(0)
	max := int32(10)

	go incr(&count, 0, max, func() {
		sign <- struct{}{}
	})

	go incr(&count, 1, max, func() {
		sign <- struct{}{}
	})

	<-sign
	<-sign

	log.Printf("count: %d\n", count)
}

func wgCountApp() {
	var wg sync.WaitGroup
	wg.Add(2)
	count := int32(0)
	max := int32(10)

	go incr(&count, 0, max, wg.Done)
	go incr(&count, 1, max, wg.Done)

	wg.Wait()

	log.Printf("count: %d\n", count)
}

func contextCountApp() {
	count := int32(0)
	max := int32(10)

	ctx, cancelFunc := context.WithCancel(context.Background())

	go incr(&count, 0, max, cancelFunc)
	go incr(&count, 1, max, cancelFunc)

	<-ctx.Done()

	log.Printf("count: %d\n", count)
}

func syncPoolApp() {
	buf := getBuffer()
	defer buf.Free()

	buf.Write("A Pool is a set of temporary objects")
	buf.Write("A Pool is safe for use by multiple goroutines simultaneously")
	buf.Write("A Pool must not be copied after first use")

	for {
		block, err := buf.Read()

		if err != nil {
			if err == io.EOF {
				break
			}

			panic(fmt.Errorf("unexpected error: %+v", err))
		}

		fmt.Print(block)
	}
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

func incr(pCount *int32, id int32, max int32, deferFunc func()) {
	defer func() {
		deferFunc()
	}()

	for i := 0; ; i++ {
		curNum := atomic.LoadInt32(pCount)

		if curNum >= max {
			break
		}

		newNum := curNum + 2
		time.Sleep(time.Millisecond * 200)

		if atomic.CompareAndSwapInt32(pCount, curNum, newNum) {
			log.Printf("[OK] operator: %d, iterator count: %d, number: %d\n", id, i, newNum)
		} else {
			log.Printf("[FAILED] operator: %d, iterator count: %d, number: %d\n", id, i, curNum)
		}
	}
}
