package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/persist"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url: "http://www.7799520.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
	//e.Run(engine.Request{
	//	Url: "http://www.7799520.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}