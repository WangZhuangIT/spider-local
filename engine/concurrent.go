package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler  Scheduler
	WorkNum int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	in := make(chan Request)
	out := make(chan ParseResult)

	for i:=0;i<e.WorkNum;i++ {
		createWork(in,out)
	}

	e.Scheduler.ConfigureMasterWorkChan(in)

	for _,r:=range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _,item := range result.Items{
			log.Printf("item :%v",item)
		}

		for _,request := range result.Requests{
			e.Scheduler.Submit(request)
		}

	}

}

func createWork(in chan Request, out chan ParseResult) {
	go func() {
		for  {
			r := <-in
			result,err:=worker(r)
			if err != nil{
				continue
			}
			out<-result 
		}
	}()
}
