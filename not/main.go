package main

import (
	cc "github.com/ariyn/cloud-computer"
)

func main() {
	name := cc.Name
	if name == "" {
		name = "not"
	}

	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	err := cc.RunRedis(func(i ...bool) bool {
		return !i[0]
	}, name, inputs, outputs, true)
	if err != nil {
		panic(err)
	}
}
