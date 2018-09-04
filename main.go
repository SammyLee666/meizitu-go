package main

import (
	"Img/scheduler"
	"Img/engine"
	parser2 "Img/parser"
)

func main() {
	//engine.Run(parser.Request{
	//	Url:        "http://www.meizitu.com/a/more_1.html",
	//	ParserFunc: parser.List,
	//})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 20,
	}
	e.Run(parser2.Request{
		Url:        "http://www.meizitu.com/a/more_1.html",
		ParserFunc: parser2.List,
	})
}
