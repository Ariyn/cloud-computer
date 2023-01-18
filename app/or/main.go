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
		name = "or"
	}
	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	if cc.IsVerbose {
		log.Println(inputs)
	}

	or := gate.NewOrGate(name, inputs, outputs)
	or.UseOptimization = cc.UseOptimization

	err := cc.RunGateWithRedis(context.Background(), or)
	if err != nil {
		panic(err)
	}
}
