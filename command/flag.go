package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	age := flag.Int("age", 18, "age")
	flag.Parse()
	fmt.Printf("Hello, %s! age: %d!\n", name, *age)
	fmt.Printf("flag args: %v\n", flag.Args())
	fmt.Printf("os args: %v\n", os.Args)
}
