package engine

import (
	"log"
	"spider-local/fetcher"
)

func Run(seeds ...Request) {
	queue := make([]Request, 0)
	queue = append(queue, seeds...)

	for len(queue) > 0 {
		curRequest := queue[0]
		queue = queue[1:]

		contents, err := fetcher.Fetch(curRequest.Url)

		log.Printf("fetching : %v", curRequest.Url)
		if err != nil {
			log.Printf("fetch err : %v", err)
			continue
		}

		result := curRequest.ParseFunc(contents)
		queue = append(queue, result.Requests...)

		for k, v := range result.Items {
			log.Printf("item%v: %v\n", k, v)
		}
	}
}
