package main

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"github.com/ariyn/cloud-computer/gate"
)

func main() {
	name := cc.Name
	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := []cc.Element{{
		IsAlias: true,
	}}

	alias := gate.NewAliasGate(name, inputs, outputs)
	alias.UseOptimization = cc.UseOptimization

	err := cc.RunGateWithRedis(context.Background(), alias)
	if err != nil {
		panic(err)
	}
}
