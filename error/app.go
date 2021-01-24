package main

import (
	"errors"
	"fmt"

	xerrors "github.com/pkg/errors"
)

func Positive(n int) (bool, error) {
	if n == 0 {
		return false, errors.New("undefined")
	}

	return n > 0, nil
}

func check(n int) {
	pos, err := Positive(n)

	if err != nil {
		fmt.Println(n, err)
		return
	}

	if pos {
		fmt.Println(n, "is positive")
	} else {
		fmt.Println(n, "is negative")
	}
}

func defaultErrorApp() {
	check(-1)
	check(0)
	check(1)
}

type stringError string

func (e stringError) Error() string {
	return string(e)
}

func NewStringError(text string) error {
	return stringError(text)
}

type structError struct {
	s string
}

func (e structError) Error() string {
	return e.s
}

func NewStructError(text string) error {
	return structError{text}
}

type TypeError struct {
	Msg  string
	File string
	Line int
}

func (e *TypeError) Error() string {
	return fmt.Sprintf("%s:%d, %s", e.File, e.Line, e.Msg)
}

func typeErrorApp() {
	var err1 = NewStringError("EOF")
	var err2 = NewStructError("EOF")
	var err3 = errors.New("EOF")

	if err1 == NewStringError("EOF") {
		fmt.Println("String type error:", err1)
	}

	if err2 == NewStructError("EOF") {
		fmt.Println("Struct type error:", err2)
	}

	if err3 == errors.New("EOF") {
		fmt.Println("Default type error:", err3)
	}

	err4 := &TypeError{"... error message ...", "server.go", 100}

	fmt.Println("Type error:", err4)

	err5 := fmt.Errorf("outer error: %v", err4)

	fmt.Println("Type error:", err5)

	switch err := interface{}(err4).(type) {
	case nil:
	case *TypeError:
		fmt.Println("error occurred on line: ", err.Line)
	default:
	}
}

func pkgWrapErrorApp() {
	err := xerrors.Errorf("no such file")
	err1 := xerrors.Wrap(err, "wrap error")

	fmt.Println("Wrap error:", err1)

	err2 := xerrors.WithMessage(err, "with some message")

	fmt.Printf("oring error, type: %T, value: %v\n", xerrors.Cause(err2), xerrors.Cause(err2))
	fmt.Printf("With message error: \n%+v\n", err2)
}

func main() {
	// defaultErrorApp()
	// typeErrorApp()
	// pkgWrapErrorApp()
}
