package main

import (
	"flag"
	"fmt"
	"os"
)

var fooCmdLine = flag.NewFlagSet("foo", flag.ExitOnError)
var barCmdLine = flag.NewFlagSet("bar", flag.ExitOnError)

func init() {
	fooCmdLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage of %s:\n", os.Args[1])
		fooCmdLine.PrintDefaults()
	}
	barCmdLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[1])
		barCmdLine.PrintDefaults()
	}
}

func main() {
	fooEnable := fooCmdLine.Bool("enable", false, "foo enable")
	fooName := fooCmdLine.String("name", "foo", "foo name")
	barLevel := barCmdLine.Int("level", -1, "bar level")

	if (len(os.Args) < 2) {
		fmt.Printf("expect subcommand\n")
		os.Exit(1)
	}

	switch (os.Args[1]) {
	case "foo":
		fooCmdLine.Parse(os.Args[2:])
		fmt.Printf("subcommand foo, enable: %v, name: %s\n", *fooEnable, *fooName)
		fmt.Printf("foo args: %v\n", fooCmdLine.Args())
		fmt.Printf("os args: %v\n", os.Args)
	case "bar":
		barCmdLine.Parse(os.Args[2:])
		fmt.Printf("subcommand bar, level: %d\n", *barLevel)
		fmt.Printf("bar args: %v\n", barCmdLine.Args())
		fmt.Printf("os args: %v\n", os.Args)
	default:
		fmt.Printf("expect subcommand foo or bar\n")
	}
}
