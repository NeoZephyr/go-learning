## go 语言学习
### 变量
### 条件控制
### 循环

```sh
go get github.com/gpmgo/gopm
gopm get -g -v golang.org/x/tools/cmd/goimports
gopm get -g -v golang.org/x/text
gopm get -g -v golang.org/x/net/html

go install pain.com/basic/

# 安装目录下的所有文件
go install ./...

# 运行当前目录下测试
go test .

# 查看测试代码覆盖率 
go test -coverprofile=c.out
go tool cover -html=c.out

# 性能测试
go test -bench .

go test -bench . -cpuprofile cpu.out
go tool pprof cpu.out

# http 性能
# 导入 net/http/pprof 包
go tool pprof http://localhost:8888/debug/pprof
go tool pprof http://localhost:8888/debug/pprof/heap

# 安装 graphviz
# web
```
```
go doc
godoc -http :6060
```
```
go run -race a.go
```