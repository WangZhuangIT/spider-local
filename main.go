package main

import (
	"spider/engine"
	"spider/persist"
	"spider/scheduler"
	"spider/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkNum:   10,
		ItemSaver: itemChan,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
