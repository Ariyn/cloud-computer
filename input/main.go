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

	log.Println("start")

	inputs := cc.ParseInputs(cc.Inputs...)

	err := cc.RunRedis(func(inputs ...bool) (results []bool) {
		return
	}, name, inputs, nil, cc.UseOptimization, false, true)
	if err != nil {
		panic(err)
	}
}
