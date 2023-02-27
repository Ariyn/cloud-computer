package main

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"github.com/ariyn/cloud-computer/gate"
)

func main() {
	name := cc.Name
	if name == "" {
		name = "flipflop"
	}

	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(2)

	flipFlop := gate.NewFlipFlopGate(name, inputs, outputs)
	flipFlop.UseOptimization = cc.UseOptimization

	err := cc.RunGateWithRedis(context.Background(), flipFlop)
	if err != nil {
		panic(err)
	}
}
