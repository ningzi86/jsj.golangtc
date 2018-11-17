package engine

import (
	"fmt"
)

type ConcurrentEngine struct {
	WorkerCount int
	Scheduler   Scheduler
}

type Scheduler interface {
	Submit(Request)
	ConfigerMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for _, r := range seeds {
		fmt.Println("0.发现一个请求")
		e.Scheduler.Submit(r)
	}

	for i := 0; i <= e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	itemCount := 0
	for {
		itemCount++
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item #%d: %v\n", itemCount, item)
		}

		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {

			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
