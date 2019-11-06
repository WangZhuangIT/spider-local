package main

import (
	"spider-local/engine"
	"spider-local/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
