package main

import (
	"goExProject/crawler/engine"
	"goExProject/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parser.GetCityList,
	})
}
