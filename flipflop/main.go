package main

import (
	cc "github.com/ariyn/cloud-computer"
)

func main() {
	name := cc.Name
	if name == "" {
		name = "flipflop"
	}

	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(2)

	isUsingClock := len(inputs) == 3
	previousStatus := false
	err := cc.RunRedis(func(i ...bool) []bool {
		if isUsingClock && i[2] == false {
			return []bool{previousStatus, !previousStatus}
		}

		currentStatus := previousStatus
		if i[0] {
			currentStatus = true
		} else if i[1] {
			currentStatus = false
		}

		previousStatus = currentStatus
		return []bool{currentStatus, !currentStatus}
	}, name, inputs, outputs, cc.UseOptimization)
	if err != nil {
		panic(err)
	}
}
