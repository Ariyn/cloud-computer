package main

import (
	"flag"
	cc "github.com/ariyn/cloud-computer"
	"log"
)

// clk means clock. It should be integer Hz
var clk int

func init() {
	flag.IntVar(&clk, "clk", 0, "clock for purse")
}

func main() {
	flag.Parse()

	name := cc.Name
	if name == "" {
		name = "clock"
	}

	log.Println("start")

	//inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	cc.Clock(clk, outputs)
	//err := cc.RunRedis(func(inputs ...bool) (result bool) {
	//	result = inputs[0]
	//	for _, b := range inputs[1:] {
	//		result = result && b
	//	}
	//	return
	//}, name, nil, outputs, true)
	//if err != nil {
	//	panic(err)
	//}
}
