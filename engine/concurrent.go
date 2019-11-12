package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkNum   int
	ItemSaver chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChanType() Worker
	Run()
}

type ReadyNotifier interface {
	WorkerReady(Worker)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkNum; i++ {
		createWork(e.Scheduler.WorkerChanType(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			log.Printf("duplicate url : %v", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemSaver <- item
			}()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}

	}

}

func createWork(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//tell scheduler i am ready
			ready.WorkerReady(in)
			r := <-in
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
