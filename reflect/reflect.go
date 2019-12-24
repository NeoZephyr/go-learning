package main

import (
	"fmt"
	"reflect"
)

func main() {
	testInterface(3.14)
	testInterface(Student{"jack", "female"})
}

type Student struct {
	name string
	gender string
}

func testInterface(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.ValueOf(i))
	fmt.Println(reflect.ValueOf(i).Kind())
}
