package main

import (
	"spider/engine"
	"spider/persist"
	"spider/scheduler"
	"spider/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkNum:   10,
		ItemSaver: persist.ItemSaver(),
	}
	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/shanghai",
		ParseFunc: parser.ParseCity,
	})
}
