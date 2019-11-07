package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkNum   int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkChanType() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkNum; i++ {
		createWork(e.Scheduler.WorkChanType(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("item :%v", item)
		}

		for _, request := range result.Requests {
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
