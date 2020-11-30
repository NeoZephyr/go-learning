package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	testPool()
}

type Connection struct {}

type ConnectionPool struct {
	connectionBuf chan *Connection
}

func NewConnectionPool(count int) *ConnectionPool {
	connectionPool := ConnectionPool{}
	connectionPool.connectionBuf = make(chan *Connection, count)

	for i := 0; i < count; i++ {
		connectionPool.connectionBuf <- &Connection{}
	}

	return &connectionPool
}

func (pool *ConnectionPool) GetConnection(timeout time.Duration) (*Connection, error) {
	select {
	case ret := <- pool.connectionBuf:
		return ret, nil
	case <- time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (pool *ConnectionPool) ReleaseConnection(connection *Connection) error {
	select {
	case pool.connectionBuf <- connection:
		return nil
	default:
		return errors.New("overflow")
	}
}

func testPool() {
	pool := NewConnectionPool(5)

	for i := 0; i < 6; i++ {
		connection, err := pool.GetConnection(time.Second * 1)

		if err != nil {
			pool.ReleaseConnection(connection)
			fmt.Println(err)
			break
		}
	}
}