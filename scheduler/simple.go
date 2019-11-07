package scheduler

import "spider/engine"

type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan engine.Request) {
	s.WorkChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkChan<-r
	}()
}




