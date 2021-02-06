package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
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

type IntToStrMap struct {
	m sync.Map
}

func (iMap *IntToStrMap) Delete(key int) {
	iMap.m.Delete(key)
}

func (iMap *IntToStrMap) Load(key int) (value string, ok bool) {
	v, ok := iMap.m.Load(key)

	if v != nil {
		value = v.(string)
	}

	return
}

func (iMap *IntToStrMap) LoadOrStore(key int, value string) (actual string, loaded bool) {
	v, loaded := iMap.m.LoadOrStore(key, value)
	actual = v.(string)
	return
}

func (iMap *IntToStrMap) Range(f func(key int, value string) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(int), value.(string))
	}
	iMap.m.Range(f1)
}

func (iMap *IntToStrMap) Store(key int, value string) {
	iMap.m.Store(key, value)
}

type ConcurrentMap struct {
	m         sync.Map
	keyType   reflect.Type
	valueType reflect.Type
}

func NewConcurrentMap(keyType, valueType reflect.Type) (*ConcurrentMap, error) {
	if keyType == nil {
		return nil, errors.New("nil key type")
	}

	if valueType == nil {
		return nil, errors.New("nil value type")
	}

	if !keyType.Comparable() {
		return nil, fmt.Errorf("incomparable key type: %s", keyType)
	}

	cMap := &ConcurrentMap{
		keyType:   keyType,
		valueType: valueType,
	}

	return cMap, nil
}

func (cMap *ConcurrentMap) Delete(key interface{}) {
	if reflect.TypeOf(key) != cMap.keyType {
		return
	}

	cMap.m.Delete(key)
}

func (cMap *ConcurrentMap) Load(key interface{}) (value interface{}, ok bool) {
	if reflect.TypeOf(key) != cMap.keyType {
		return
	}

	return cMap.m.Load(key)
}

func (cMap *ConcurrentMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	if reflect.TypeOf(key) != cMap.keyType {
		panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(key)))
	}

	if reflect.TypeOf(value) != cMap.valueType {
		panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(value)))
	}

	actual, loaded = cMap.m.LoadOrStore(key, value)
	return
}

func (cMap *ConcurrentMap) Range(f func(key, value interface{}) bool) {
	cMap.m.Range(f)
}

func (cMap *ConcurrentMap) Store(key, value interface{}) {
	if reflect.TypeOf(key) != cMap.keyType {
		panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(key)))
	}

	if reflect.TypeOf(value) != cMap.valueType {
		panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(value)))
	}

	cMap.m.Store(key, value)
}

var pairs = []struct {
	k int
	v string
}{
	{k: 1, v: "a"},
	{k: 2, v: "b"},
	{k: 3, v: "c"},
	{k: 4, v: "d"},
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

	// intToStrMapApp()
	concurrentMapApp()
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

func intToStrMapApp() {
	var iMap IntToStrMap

	for _, pair := range pairs {
		iMap.Store(pair.k, pair.v)
	}

	iMap.Range(func(key int, value string) bool {
		fmt.Printf("iMap iteration, k: %d, v: %s\n", key, value)
		return true
	})

	iMap.Delete(3)
	fmt.Println()

	iMap.Range(func(key int, value string) bool {
		fmt.Printf("iMap iteration, k: %d, v: %s\n", key, value)
		return true
	})
}

func concurrentMapApp() {
	cMap, err := NewConcurrentMap(reflect.TypeOf(pairs[0].k), reflect.TypeOf(pairs[0].v))

	if err != nil {
		fmt.Printf("fatal error: %s", err)
		return
	}

	for _, pair := range pairs {
		cMap.Store(pair.k, pair.v)
	}

	cMap.Range(func(key, value interface{}) bool {
		fmt.Printf("cMap iteration, k: %d, v: %s\n", key, value)
		return true
	})

	cMap.Delete(3)
	fmt.Println()

	cMap.Range(func(key, value interface{}) bool {
		fmt.Printf("cMap iteration, k: %d, v: %s\n", key, value)
		return true
	})
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
