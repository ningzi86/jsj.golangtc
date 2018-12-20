package main

import (
	"jsj.golangtc/crawler/zhenai/engine"
	"jsj.golangtc/crawler/zhenai/parser"
	"jsj.golangtc/crawler/scheduler"
	"jsj.golangtc/crawler_distributed/persist/client"
)

func main() {

	ex := engine.ConcurrentEngine{
		WorkerCount: 100,
		Scheduler:   &scheduler.QueuedScheduler{},
		Item:       client.ItemSaver(),
	}

	ex.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
