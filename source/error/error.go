package main

import (
	"fmt"
	"os"
)

func main() {
	//handleError()
	handlePanic()
}

func handleError() {
	file, err := os.Open("abc.txt")

	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		} else {
			fmt.Println("unknown error: ", err)
		}
	}

	defer file.Close()
}

func handlePanic() {
	defer func() {
		i := recover()

		if err, ok := i.(error); ok {
			fmt.Println("error occurred", err)
		} else {
			panic(i)
		}
	}()
	//panic(errors.New("this is a error"))
	panic("this is a error")
}

