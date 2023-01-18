package main

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"github.com/ariyn/cloud-computer/gate"
)

func main() {
	name := cc.Name
	if name == "" {
		name = "input"
	}

	inputs := cc.ParseInputs(cc.Inputs...)
	e := cc.Element{
		GateName: name,
	}

	input := gate.NewInputGate(name, inputs, []cc.Element{e})
	input.UseOptimization = cc.UseOptimization

	err := cc.RunGateWithRedis(context.Background(), input)
	if err != nil {
		panic(err)
	}
}
