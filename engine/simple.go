package engine

import (
	"log"
	"spider/fetcher"
)

type SimpleEngine struct {
}

func (SimpleEngine) Run(seeds ...Request) {
	queue := make([]Request, 0)
	queue = append(queue, seeds...)

	for len(queue) > 0 {
		curRequest := queue[0]
		queue = queue[1:]
		result, err := worker(curRequest)
		if err != nil {
			continue
		}
		queue = append(queue, result.Requests...)

		for k, v := range result.Items {
			log.Printf("item%v: %v\n", k, v)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	contents, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch err : %v", err)
		return ParseResult{}, err
	}

	return r.ParseFunc(contents), nil
}
