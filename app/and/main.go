package main

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"github.com/ariyn/cloud-computer/gate"
	"log"
)

func main() {
	name := cc.Name
	if name == "" {
		name = "and"
	}
	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	if cc.IsVerbose {
		log.Println(inputs)
	}

	and := gate.NewAndGate(name, inputs, outputs)
	and.UseOptimization = cc.UseOptimization

	err := cc.RunGateWithRedis(context.Background(), and)
	if err != nil {
		panic(err)
	}
}
