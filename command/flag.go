// 同一个目录下的源码文件都需要被声明为属于同一个代码包
// 包名与其父目录的名称一致

// 在 Go 1.5 及后续版本中，可以通过创建 internal 代码包让一些程序实体仅仅能被当前模块中的其他代码引用
// internal 代码包中声明的公开程序实体仅能被该代码包的直接父包及其子包中的代码引用

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
