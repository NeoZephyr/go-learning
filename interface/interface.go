package main

import (
	"fmt"
	"pain.com/go-learning/interface/mock"
	"pain.com/go-learning/interface/text"
)

type Downloader interface {
	Download(url string) string
}

type Uploader interface {
	Upload(url string, form map[string]string) string
}

// 接口组合
type NetDisk interface {
	Downloader
	Uploader
}

const url = "http://www.baidu.com"

func main() {
	// 接口变量接收值或者指针类型都可以
	var downloader Downloader

	downloader = &mock.Downloader{Content: "hello mock downloader"}
	download(downloader)

	downloader = &text.Downloader{UserAgent: "Mozilla/5.0"}
	download(downloader)

	testEmptyInterface1(10)
	testEmptyInterface1("hello")
	testEmptyInterface1(3.14)

	testEmptyInterface2(10)
	testEmptyInterface2("hello")
	testEmptyInterface2(3.14)

	var netdisk NetDisk
	netdisk = &mock.NetDisk{Content: "hello mock netdisk"}
	download(netdisk)
	testNetDisk(netdisk)
}

func download(downloader Downloader) {
	fmt.Printf("downloader type: %T, %v\n", downloader, downloader)
	switch downloader.(type) {
	case *mock.Downloader:
		fmt.Printf("mock download result: %s\n", downloader.Download(url))
	case *text.Downloader:
		fmt.Printf("text download result: %s\n", downloader.Download(url))
	default:
		fmt.Printf("unknow downloader type\n")
	}
}

func testNetDisk(netDisk NetDisk) {
	fmt.Printf("netdisk type: %T, %v\n", netDisk, netDisk)
	netDisk.Upload(url, map[string]string {
		"content": "netdisk upload",
	})
	fmt.Printf("mock netdisk result: %s\n", netDisk.Download(url))
	fmt.Println("NetDisk", netDisk)
}

func testEmptyInterface1(p interface{}) {
	if value, ok := p.(int); ok {
		fmt.Printf("the type of p is int, p = %v\n", value)
		return
	}

	if value, ok := p.(string); ok {
		fmt.Printf("the type of p is string, p = %v\n", value)
		return
	}

	fmt.Printf("the type of p is unknown\n")
}

func testEmptyInterface2(p interface{})  {
	switch value := p.(type) {
	case int:
		fmt.Printf("the type of p is int, p = %v\n", value)
	case string:
		fmt.Printf("the type of p is string, p = %v\n", value)
	default:
		fmt.Printf("the type of p is unknown\n")
	}
}

// err.Error
// errors.New