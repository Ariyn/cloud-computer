package main

import (
	cc "github.com/ariyn/cloud-computer"
	"log"
)

func main() {
	name := cc.Name
	if name == "" {
		name = "xor"
	}

	log.Println("start")

	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	err := cc.RunRedis(func(inputs ...bool) (result bool) {
		result = inputs[0]
		for _, b := range inputs[1:] {
			result = result != b
		}
		return
	}, name, inputs, outputs, cc.UseOptimization)
	if err != nil {
		panic(err)
	}
}
