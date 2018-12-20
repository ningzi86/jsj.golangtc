package engine

type ConcurrentEngine struct {
	WorkerCount int
	Scheduler   Scheduler
	Item        chan Item
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	Run()
	ReadyNotifier
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for i := 0; i <= e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() { e.Item <- item }()
		}

		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {

	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
