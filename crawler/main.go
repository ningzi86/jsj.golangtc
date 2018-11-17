package main

import (
	"jsj.golangtc/crawler/zhenai/engine"
	"jsj.golangtc/crawler/zhenai/parser"
	"jsj.golangtc/crawler/scheduler"
)

func main() {

	ex := engine.ConcurrentEngine{
		WorkerCount: 1,
		Scheduler:   &scheduler.QueuedScheduler{},
	}

	ex.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
