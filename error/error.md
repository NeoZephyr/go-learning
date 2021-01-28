# error

## panic

对于不可恢复的程序错误，例如索引越界、不可恢复的环境问题、栈溢出，才使用 panic。对于其他的错误情况，应该是期望使用 error 来进行判定

某个函数中的某行代码引发了一个 panic 时，初始的 panic 详情会被建立起来，并且该程序的控制权会从此行代码开始一级又一级地沿着调用栈的反方向传播最外层函数，最后被 Go 语言运行时系统收回。最后，程序崩溃并终止运行，承载程序这次运行的进程也会随之死亡并消失。在此过程中，panic 详情会被逐渐地积累和完善，并会在程序终止之前被打印出来


## error 分类

### error vs exception

1. 简单
2. 考虑失败，而不是成功
3. 没有隐藏的控制流
4. 由使用者完全控制 error
5. Error are values

### Sentinel Error

预定义的特定错误，称之为 sentinel error。使用 sentinel 值是最不灵活的错误处理策略，因为调用方必须使用 == 将结果与预先声明的值进行比较。当想要提供更多的上下文时，就会返回一个不同的错误而破坏相等性检查

甚至是一些有意义的 fmt.Errorf 携带一些上下文，也会破坏调用者的 == ，调用者将被迫查看 error.Error() 方法的输出，以查看它是否与特定的字符串匹配。我们不应该依赖检测 error.Error 的输出，因为 Error 方法的输出的字符串只用于记录日志、输出到 stdout 等

sentinel 成为 API 公共部分：如果公共函数或方法返回一个特定值的错误，那么该值必须是公共的，要有文档记录；如果 API 定义了一个返回特定错误的 interface，则该接口的所有实现都将被限制为仅返回该错误，即使它们可以提供更具描述性的错误

sentinel 最糟糕的问题是在两个包之间创建了源代码依赖关系。例如，检查错误是否等于 io.EOF，必须导入 io 包。当项目中的许多包导出错误值时，存在耦合，项目中的其他包必须导入这些错误值才能检查特定的错误条件

建议避免在编写的代码中使用 sentinel errors

### Error Type

Error type 是实现了 error 接口的自定义类型

```go
type MyError struct {
    Msg string
    File string
    Line int
}

func (e *MyError) Error() string {
    return fmt.Sprintf("%s:%d, %s", e.File, e.Line, e.Msg)
}

func test() error {
    return &MyError{"...", "server.go", 42}
}
```

调用者可以使用类型断言转换成这个类型，来获取更多的上下文信息。与错误值相比，错误类型的一大改进是它们能够包装底层错误以提供更多上下文

调用者要使用类型断言和类型 switch，就要让自定义的 error 变为 public，这种模型会导致和调用者产生强耦合

建议避免错误类型，或者至少避免将它们作为公共 API 的一部分

### Opaque Error

不透明错误处理，虽然知道发生了错误，但没有能力看到错误的内部。作为调用者，关于操作的结果，只知道的就是它成功还是失败了

```go
func fn() error {
    x, err := bar.Foo()

    if err != nil {
        return err
    }
}
```

在少数情况下，这种二分错误处理方法是不够的。例如，与进程外的世界进行交互（如网络活动），需要调用方调查错误的性质，以确定重试该操作是否合理。在这种情况下，可以断言错误实现了特定的行为，而不是断言错误是特定的类型或值

这个逻辑可以在不导入定义错误的包或者实际上不了解 err 的底层类型的情况下实现

```go
type temporary interface {
    Temporary() bool
}

func IsTemporary(err error) bool {
    te, ok := err.(temporary)
    return ok && te.Temporary()
}
```

## handle error

### eliminating error

```go
type errWriter struct {
    io.Writer
    err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
    if e.err != nil {
        return 0, e.err
    }

    var n int
    n, e.err = e.Writer.Write(buf)
    return n, nil
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
    ew := &errWriter{Writer: w}
    fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

    for _, h := range headers {
        fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
    }

    fmt.Fprint(ew, "\r\n");
    io.Copy(ew, body)

    return ew.err
}
```

### wrap error

```go
func AuthenticateRequest(r *Request) error {
    err := authenticate(r.User)

    if err != nil {
        return fmt.Errorf("authenticate failed: %v", err)
    }

    return nil
}
```

因为将错误值转换为字符串，将其与另一个字符串合并，然后将其转换回 fmt.Errorf 破坏了原始错误，导致等值判定失败

pkg errors

在应用代码中，使用 errors.New 或者  errors.Errorf 返回错误

```go
func parseArgs(args []string) error {
    if len(args) < 3 {
        return errors.Errorf("not enough args")
    }
    return nil
}
```

如果调用其他的函数，通常简单的直接返回

```go
if err != nil {
    return err
}
```

如果和其他库进行协作，考虑使用 errors.Wrap 或者 errors.Wrapf 保存堆栈信息。同样适用于和标准库协作的时候

```go
f, err := os.Open(path)

if err != nil {
    return errors.Wrapf(err, "failed to open %q", path)
}
```

直接返回错误，而不是每个错误产生的地方到处打日志

在程序的顶部或者是工作的 goroutine 顶部（请求入口），使用 %+v 把堆栈详情记录

```go
func main() {
    err := app.Run()

    if err != nil {
        fmt.Printf("FATAL: %+v\n", err)
        os.Exit(1)
    }
}
```

使用 errors.Cause 获取 root error，再进行和 sentinel error 判定

选择 wrap error 是只有 applications 可以选择应用的策略。具有最高可重用性的包只能返回根错误值

如果函数/方法不打算处理错误，那么用足够的上下文 wrap errors 并将其返回到调用堆栈中。例如，额外的上下文可以是使用的输入参数或失败的查询语句

一旦确定函数/方法将处理错误，错误就不再是错误。如果函数/方法仍然需要发出返回，则它不能返回错误值。它应该只返回零

### before 1.13

函数在调用栈中添加信息向上传递错误

```go
if err != nil {
    return fmt.Errof("decompress %v: %v", name, err)
}
```

使用 fmt.Errorf 创建新错误丢弃原始错误中除文本外的所有内容。有时可能需要定义一个包含底层错误的新错误类型，并将其保存以供代码检查：

```go
type QueryError struct [
    Query string
    Err error
]

if e, ok := err.(*QueryError); ok && e.Err == ErrPermission {}
```

包含另一个错误的 error 可以实现返回底层错误的 Unwrap 方法。如果 e1.Unwrap() 返回 e2，那说明 e1 包装 e2，可以展开 e1 以获得 e2

```go
func (e *QueryError) Unwrap() error {
    return e.Err
}
```

go1.13 errors 包中包含两个用于检查错误的新函数：Is 和 As

```go
// err or some error it wraps is a permission problem
if errors.Is(err, ErrPermission) {}
```

```go
var e *QueryError

// 类似于；
// if e, ok := err.(*QueryError); ok {}
if errors.As(err, &e) {}
```

fmt.Errorf 支持新的 %w 谓词，用 %w 包装错误可用于 errors.Is 以及 errors.As

```go
err := fmt.Errorf("access denied: %w", ErrPermission)

if errors.Is(err, ErrPermission)
```
