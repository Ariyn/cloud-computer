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
		name = "input"
	}

	inputs := cc.ParseInputs(cc.Inputs...)
	e := cc.Element{
		GateName: name,
	}

	input := gate.NewInputGate(name, inputs, []cc.Element{e})
	input.UseOptimization = cc.UseOptimization

	if err = cc.RunGateWithRedis(context.Background(), config.Redis.Host, input); err != nil {
		log.Fatal(err)
	}
}
