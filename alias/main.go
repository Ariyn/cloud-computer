package main

import (
	"flag"
	cc "github.com/ariyn/cloud-computer"
	"log"
)

func main() {
	flag.Parse()

	name := cc.Name

	log.Println("start")

	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := []cc.Element{{
		IsAlias: true,
	}}

	log.Println("aliasing", name, "input", inputs[0].String())
	err := cc.RunRedis(func(i ...bool) bool {
		return i[0]
	}, name, inputs, outputs, false)
	if err != nil {
		panic(err)
	}
}
