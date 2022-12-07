package main

import (
	"flag"
	cc "github.com/ariyn/cloud-computer"
)

func main() {
	flag.Parse()

	name := cc.Name
	if name == "" {
		name = "not"
	}

	err := cc.RunRedis(func(i ...bool) bool {
		return !i[0]
	}, name, []string{cc.InputName1}, 1)
	if err != nil {
		panic(err)
	}
}
