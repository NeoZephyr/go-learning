# goroutine

如果你的 goroutine 在从另一个 goroutine 获得结果之前无法取得进展，那么通常情况下，你自己去做这项工作比委托它( go func() )更简单。
这通常消除了将结果从 goroutine 返回到其启动器所需的大量状态跟踪和 chan 操作

// G: goroutine
// P: 可以承载若干个 G，且能够使这些 G 适时地与 M 进行对接
// M: 系统级线程


Leave concurrency to the caller
```go
func ListDirectory(dir string) ([]string, error)
```

将目录读取到一个切片中，然后返回。如果出现错误，则返回错误。这是同步调用的，调用方会阻塞，直到读取所有目录条目。根据目录的大小，这可能需要很长时间，并且可能会分配大量内存来存储目录条目

```go
func ListDirectory(dir string) chan string
```

返回一个通道，通过通道传递目录。当通道关闭时，这表示不再有目录。由于在 ListDirectory 返回后发生通道的填充，ListDirectory 可能内部启动 goroutine 来填充通道

调用方无法区分空目录与错误之间的区别。这两种方法都会导致从 ListDirectory 返回的通道会立即关闭。调用者必须持续从通道读取，直到它关闭，即使已经收到了想要的答案

```go
func ListDirectory(dir string, fn func(string))
```
如果函数启动 goroutine，则必须向调用方提供显式停止该 goroutine 的方法。通常，将异步执行函数的决定权交给该函数的调用方通常更容易

Never start a goroutine without knowning when it will stop

```go
func leak() {
    ch := make(chan int)

    go func() {
        val := <-ch
        fmt.Println("Received:", val)
    }()
}
```


