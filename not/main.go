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

	err := cc.RunRedis(func(i ...bool) (results []bool) {
		results = append(results, !i[0])
		return
	}, name, inputs, outputs, cc.UseOptimization)
	if err != nil {
		panic(err)
	}
}
