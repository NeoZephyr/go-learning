package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	//testHttp1()
	testHttp2()
}

func testHttp1() {
	resp, err := http.Get("http://www.imooc.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	result, err := httputil.DumpResponse(resp, true)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("result: %s\n", result)
}

func testHttp2() {
	request, err := http.NewRequest(http.MethodGet, "https://bbs.hupu.com/bxj", nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Printf("request: %+v\n", *req)
			return nil
		},
	}

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	result, err := httputil.DumpResponse(response, true)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("result: %s\n", result)
}