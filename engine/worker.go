package engine

import (
	"log"
	"spider/fetcher"
)

func worker(r Request) (ParseResult, error) {
	contents, err := fetcher.Fetch(r.Url, "local")
	if err != nil {
		log.Printf("fetch err : %v", err)
		return ParseResult{}, err
	}

	return r.ParseFunc(contents, r.Url), nil
}
