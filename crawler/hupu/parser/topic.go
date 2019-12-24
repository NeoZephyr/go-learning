package parser

import (
	"pain.com/go-learning/crawler/engine"
	"regexp"
)

const topicReg = `<a  href="(/[0-9]+.html)"[^>]*>([^<]+)</a>`

func ParseTopic(contents []byte) engine.ParsedResult {

	// <a  href="/27480039.html" class="truetit" target="_blank">北京昨天下的冰雹，不躲会不会出人命？？</a>
	reg := regexp.MustCompile(topicReg)
	subMatchLists := reg.FindAllSubmatch(contents, -1)

	result := engine.ParsedResult{}

	for _, subMatchList := range subMatchLists {
		result.Items = append(result.Items, string(subMatchList[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(subMatchList[1]),
			ParseFunc: engine.DefaultParseFunc, // TODO change to topic detail parse func
		})
	}

	return result
}
