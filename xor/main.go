package main

import (
	"flag"
	cc "github.com/ariyn/cloud-computer"
	"log"
)

func main() {
	flag.Parse()

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
	}, name, inputs, outputs, true)
	if err != nil {
		panic(err)
	}
}
