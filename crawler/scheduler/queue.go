package scheduler

import (
	"jsj.golangtc/crawler/zhenai/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(r chan engine.Request) {
	s.workerChan <- r
}

func (s *QueuedScheduler) Run() {

	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workQ []chan engine.Request

		for {

			var activieRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 && len(workQ) > 0 {
				//fmt.Println("2.两个队列都存在")
				activieRequest = requestQ[0]
				activeWorker = workQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workQ = append(workQ, w)
			case activeWorker <- activieRequest:
				requestQ = requestQ[1:]
				workQ = workQ[1:]
			}

		}

	}()
}
