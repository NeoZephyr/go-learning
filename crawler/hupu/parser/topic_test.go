package parser

import (
	"fmt"
	"pain.com/go-learning/crawler/fetcher"
	"testing"
)

func TestParseTopic(t *testing.T) {
	contents, err := fetcher.Fetch("https://bbs.hupu.com/bxj")

	//contents, err := ioutil.ReadFile("topic_test.html")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", contents)
	parsedResult := ParseTopic(contents)

	const totalTopic = 119
	if len(parsedResult.Requests) != totalTopic {
		t.Errorf("result should %d requests; but have %d requests", totalTopic, len(parsedResult.Requests))
	}
}
