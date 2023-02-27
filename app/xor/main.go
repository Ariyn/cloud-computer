package main

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"github.com/ariyn/cloud-computer/gate"
)

func main() {
	name := cc.Name
	if name == "" {
		name = "xor"
	}
	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	xor := gate.NewXorGate(name, inputs, outputs)
	xor.UseOptimization = cc.UseOptimization

	err := cc.RunGateWithRedis(context.Background(), xor)
	if err != nil {
		panic(err)
	}
}
