package main

import (
	cc "github.com/ariyn/cloud-computer"
)

func main() {
	name := cc.Name
	inputs := cc.ParseInputs(cc.Inputs...)
	outputs := []cc.Element{{
		IsAlias: true,
	}}

	err := cc.RunRedis(func(i ...bool) []bool {
		return []bool{i[0]}
	}, name, inputs, outputs, false, true, false)
	if err != nil {
		panic(err)
	}
}
