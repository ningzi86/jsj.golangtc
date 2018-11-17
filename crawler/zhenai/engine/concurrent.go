package engine

import "fmt"

type ConcurrentEngine struct {
	WorkerCount int
	Scheduler   Scheduler
}

type Scheduler interface {
	Submit(Request)
	ConfigerMasterWorkerChan(chan Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.ConfigerMasterWorkerChan(in)

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for i := 0; i <= e.WorkerCount; i++ {
		createWorker(in, out)
	}

	itemCount := 0
	for {
		itemCount++

		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item #%d: %v", itemCount, item)
		}

		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
