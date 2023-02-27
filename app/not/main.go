package main

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"github.com/ariyn/cloud-computer/gate"
)

func main() {
	name := cc.Name
	if name == "" {
		name = "not"
	}

	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	not := gate.NewNotGate(name, inputs, outputs)
	not.UseOptimization = cc.UseOptimization

	err := cc.RunGateWithRedis(context.Background(), not)
	if err != nil {
		panic(err)
	}
}
