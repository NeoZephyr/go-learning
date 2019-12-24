package engine

import (
	"log"
	"pain.com/go-learning/crawler/fetcher"
)

type DefaultEngine struct {}

func (e *DefaultEngine) Run(seeds ...Request) {
	var requests []Request

	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parsedResult, err := e.worker(request)

		if err != nil {
			log.Printf("Engine: process url %s, %s\n", request.Url, err)
		}

		requests = append(requests, parsedResult.Requests...)

		for _, item := range parsedResult.Items {
			log.Printf("item: %s ", item)
		}
		log.Println()
	}
}

func (e *DefaultEngine) worker(request Request) (ParsedResult, error) {
	log.Printf("Fetcher: Begin fetch url %s\n", request.Url)
	content, err := fetcher.Fetch(request.Url)

	if err != nil {
		log.Printf("Fetcher: fetch url %s, %s\n", request.Url, err)
		return ParsedResult{}, err
	}

	parsedResult := request.ParseFunc(content)
	return parsedResult, nil
}
