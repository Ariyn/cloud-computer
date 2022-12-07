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
		name = "and"
	}

	log.Println("start")
	err := cc.RunRedis(func(i ...bool) bool {
		return i[0] && i[1]
	}, name, []string{cc.InputName1, cc.InputName2}, 1)
	if err != nil {
		panic(err)
	}
}
