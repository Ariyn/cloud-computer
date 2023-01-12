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
	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	err := cc.RunRedis(func(inputs ...bool) (results []bool) {
		results = append(results, inputs[0])
		for _, b := range inputs[1:] {
			results[0] = results[0] && b
		}

		if cc.IsVerbose {
			log.Printf("name %s, inputs: %v, results: %v", name, inputs, results)
		}
		return
	}, name, inputs, outputs, cc.UseOptimization, false, false)
	if err != nil {
		panic(err)
	}
}
