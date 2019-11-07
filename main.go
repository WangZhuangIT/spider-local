package main

import (
	"spider/engine"
	"spider/scheduler"
	"spider/zhenai/parser"
)

func main() {
	e:=engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkNum:   5,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
