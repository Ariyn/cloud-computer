package main

import (
	cc "github.com/ariyn/cloud-computer"
	"log"
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

	log.Println("input", inputs)
	err := cc.RunRedis(func(inputs ...bool) (results []bool) {
		return inputs
	}, name, inputs, []cc.Element{e}, false, false, true)
	if err != nil {
		panic(err)
	}
}
