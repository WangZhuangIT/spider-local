package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkNum   int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkNum; i++ {
		createWork(out, e.Scheduler)
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

func createWork(out chan ParseResult, s Scheduler) {
	job := make(chan Request)
	go func() {
		for {
			//tell scheduler i am ready
			s.WorkerReady(job)
			r := <-job
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
