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
	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := []cc.Element{{
		IsAlias: true,
	}}

	alias := gate.NewAliasGate(name, inputs, outputs)
	alias.UseOptimization = cc.UseOptimization

	if err = cc.RunGateWithRedis(context.Background(), config.Redis.Host, alias); err != nil {
		log.Fatal(err)
	}
}
