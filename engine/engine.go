package engine

import (
	"Img/parser"
	"log"
	"Img/fetcher"
	"fmt"
)

func Run(seeds parser.Request) {
	Q := []parser.Request{
		seeds,
	}

	for len(Q) > 0 {
		request := Q[0]
		Q = Q[1:]
		fmt.Println(request.Url)
		body, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Println("fetcher.Fetch ERR: ", err)
		}

		parserResult := request.ParserFunc(body)

		Q = append(Q, parserResult.Request...)
	}


}
