package main

import (
	cc "github.com/ariyn/cloud-computer"
)

func main() {
	name := cc.Name
	if name == "" {
		name = "or"
	}
	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	err := cc.RunRedis(func(inputs ...bool) (results []bool) {
		results = append(results, inputs[0])

		for _, b := range inputs[1:] {
			results[0] = results[0] || b
		}
		return
	}, name, inputs, outputs, cc.UseOptimization, false, false)
	if err != nil {
		panic(err)
	}
}
