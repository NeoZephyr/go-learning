GOROOT: GO 语言的安装路径
GOPATH: 工作空间
GOBIN: GO 程序生成的可执行文件路径

代码包的名称一般会与源码文件所在的目录同名

```sh
go build hello.go
go run hello.go
```

如果一个源码文件声明属于main包，并且包含一个无参数声明且无结果声明的main函数，那么它就是命令源码文件。
```go
package main

import (
    "fmt"
    // 用于接收和解析命令参数
    "flag"
)

var name string

func init() {
    flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
    // flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
    flag.CommandLine.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
        flag.PrintDefaults()
    }

    // param addr, param name, param default value, prompt
    flag.StringVar(&name, "name", "everyone", "The greeting object.")
    // var name = flag.String("name", "everyone", "The greeting object.")
}

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
        flag.PrintDefaults()
    }

    // 真正解析命令参数，并把它们的值赋给相应的变量
    flag.Parse()
    fmt.Printf("Hello, %s!\n", name)
    hello(name)
}
```
```sh
go run demo.go -name="Jack"
go build demo.go
```

```go
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)
cmdLine.Parse(os.Args[1:])
```

库源码文件是不能被直接运行的源码文件，它仅用于存放程序实体，这些程序实体可以被其他代码使用。

在同一个目录下的源码文件都需要被声明为属于同一个代码包
```go
package main

import "fmt"

func hello(name string) {
    fmt.Printf("Hello, %s!\n", name)
}
```

源码文件声明的代码包的名称可以与其所在的目录的名称不同。在针对代码包进行构建时，生成的结果文件的主名称与其父目录的名称一致。

```sh
go run demo.go hello.go
```

```go
var name = *flag.String("name", "everyone", "The greeting object")
```
```go
// 短变量声明只能在函数体内部使用
name := *flag.String("name", "everyone", "The greeting object.")
```



