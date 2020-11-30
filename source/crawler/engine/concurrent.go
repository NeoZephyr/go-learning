package engine

import (
	"log"
	"pain.com/go-learning/crawler/fetcher"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)
	ConfigureRequestPool(chan Request)
}

func (e *ConcurrentEngine) run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParsedResult)

	e.Scheduler.ConfigureRequestPool(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	for {
		parsedResult := <- out

		for _, item := range parsedResult.Items {
			log.Printf("item: %s ", item)
		}
		log.Println()

		for _, request := range parsedResult.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParsedResult) {
	go func() {
		for {
			request := <- in
			parsedResult, err := worker(request)

			if err != nil {
				log.Printf("Engine: process url %s, %s\n", request.Url, err)
				continue
			}

			out <- parsedResult
		}
	}()
}

func worker(request Request) (ParsedResult, error) {
	log.Printf("Fetcher: Begin fetch url %s\n", request.Url)
	content, err := fetcher.Fetch(request.Url)

	if err != nil {
		log.Printf("Fetcher: fetch url %s, %s\n", request.Url, err)
		return ParsedResult{}, err
	}

	parsedResult := request.ParseFunc(content)
	return parsedResult, nil
}