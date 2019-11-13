package engine

import (
	"log"
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
