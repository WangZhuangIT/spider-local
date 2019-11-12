package scheduler

import "spider/engine"

type SimpleScheduler struct {
	WorkerChan engine.Worker
}

func (s *SimpleScheduler) WorkerChanType() engine.Worker {
	return s.WorkerChan
}

func (s *SimpleScheduler) WorkerReady(engine.Worker) {
}

func (s *SimpleScheduler) Run() {
	s.WorkerChan = make(engine.Worker)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkerChan <- r
	}()
}
