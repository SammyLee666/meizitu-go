package engine

import (
	"fmt"
	"Img/parser"
	"log"
	"Img/fetcher"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(parser.Request)
	ConfigureWorkerChan(chan parser.Request)
}

func (e *ConcurrentEngine) Run(seeds ...parser.Request) {
	in := make(chan parser.Request)
	out := make(chan parser.ParserResult)
	e.Scheduler.ConfigureWorkerChan(in)
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	//itemCount := 0
	for {
		result := <-out
		//for _, item := range result.Items {
		//	fmt.Printf("Got item #%d: %v\n",itemCount, item)
		//	itemCount++
		//}

		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}
	}
}
func createWorker(in chan parser.Request, out chan parser.ParserResult) {
	go func() {
		for {
			request := <-in
			result, err := Work(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func Work(r parser.Request) (parser.ParserResult, error) {

	fmt.Printf("Fetching %s \n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Println("fetcher.Fetch ERR: ", err)
		return parser.ParserResult{}, err
	}
	return r.ParserFunc(body), nil

}

