package main

import (
	cc "github.com/ariyn/cloud-computer"
	"log"
)

func main() {
	name := cc.Name
	if name == "" {
		name = "and"
	}

	log.Println("start")

	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	err := cc.RunRedis(func(inputs ...bool) (results []bool) {
		results[0] = inputs[0]
		for _, b := range inputs[1:] {
			results[0] = results[0] && b
		}
		return
	}, name, inputs, outputs, cc.UseOptimization)
	if err != nil {
		panic(err)
	}
}
