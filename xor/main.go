package main

import (
	"flag"
	cc "github.com/ariyn/cloud-computer"
	"log"
)

func main() {
	flag.Parse()

	name := cc.Name
	if name == "" {
		name = "xor"
	}

	log.Println("start")

	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := cc.CreateOutputs(1)

	err := cc.RunRedis(func(i ...bool) bool {
		return i[0] != i[1]
	}, name, inputs, outputs, true)
	if err != nil {
		panic(err)
	}
}
