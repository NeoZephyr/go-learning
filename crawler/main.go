package main

import (
	"pain.com/go-learning/crawler/engine"
	"pain.com/go-learning/crawler/hupu/parser"
)

func main() {
	defaultEngine := engine.DefaultEngine{}
	defaultEngine.Run(engine.Request{
		Url: "https://bbs.hupu.com/bxj",
		ParseFunc: parser.ParseTopic,
	})
}