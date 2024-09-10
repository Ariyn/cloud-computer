package main

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"github.com/ariyn/cloud-computer/gate"
	"log"
)

func main() {
	config, err := cc.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

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

	if err = cc.RunGateWithRedis(context.Background(), config.Redis.Host, and); err != nil {
		log.Fatal(err)
	}
}
