package scheduler

import "spider/engine"

type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (s *SimpleScheduler) WorkChanType() chan engine.Request {
	return s.WorkChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.WorkChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkChan <- r
	}()
}
