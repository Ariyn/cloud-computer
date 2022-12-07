package main

import (
	"context"
	"flag"
	cc "github.com/ariyn/cloud-computer"
	"log"
	"reflect"
	"strings"
)

type arrayFlags []string

func (af *arrayFlags) String() string {
	return strings.Join(*af, "\n")
}

func (af *arrayFlags) Set(value string) error {
	*af = append(*af, value)
	return nil
}

var watches arrayFlags

func init() {
	flag.Var(&watches, "names", "names for watch")
}

// TODO: 이부분 RunRedis와 거의 동일함. 추상화 할 방법 찾아보기
func main() {
	flag.Parse()

	log.Println("running")
	client := cc.ConnectRedis()
	log.Println("connected")

	inputs := make([]<-chan bool, 0)
	for _, name := range watches {
		inputs = append(inputs, cc.ReadAsyncRedis(context.TODO(), client, name))
	}

	cases := make([]reflect.SelectCase, 0)
	for _, ch := range inputs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	previousValues := make([]bool, len(inputs))
	for {
		index, value, ok := reflect.Select(cases)
		if !ok {
			break
		}

		previousValues[index] = value.Bool()
		log.Printf("%d: %v", index, previousValues[index])
	}
}
