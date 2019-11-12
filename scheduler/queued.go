package scheduler

import "spider/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workerChan  chan engine.Worker
}

func (s *QueueScheduler) WorkerChanType() engine.Worker {
	return make(chan engine.Request)
}

func (s *QueueScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueueScheduler) WorkerReady(worker engine.Worker) {
	s.workerChan <- worker
}

func (s *QueueScheduler) Run() {
	s.workerChan = make(chan engine.Worker)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workQ []engine.Worker
		for {
			var activeRequest engine.Request
			var activeWorker engine.Worker
			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workQ = append(workQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workQ = workQ[1:]
			}
		}
	}()
}
