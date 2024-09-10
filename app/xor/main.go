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
		name = "xor"
	}
	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	xor := gate.NewXorGate(name, inputs, outputs)
	xor.UseOptimization = cc.UseOptimization

	if err = cc.RunGateWithRedis(context.Background(), config.Redis.Host, xor); err != nil {
		log.Fatal(err)
	}
}
