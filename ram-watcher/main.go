package main

import (
	"context"
	"flag"
	"fmt"
	cc "github.com/ariyn/cloud-computer"
	"log"
	"reflect"
)

var ramName string
var wordSize int

func init() {
	flag.StringVar(&ramName, "name", "", "name of ram watch")
	flag.IntVar(&wordSize, "words", 16, "size of ram")
}

// TODO: 이부분 RunRedis와 거의 동일함. 추상화 할 방법 찾아보기
func main() {
	flag.Parse()

	log.Println("running")
	client := cc.ConnectRedis()
	log.Println("connected")

	inputs := make(map[int][]<-chan bool)
	for i := 0; i < wordSize; i++ {
		inputs[i] = make([]<-chan bool, 0)
		for j := 0; j < 4; j++ {
			name := fmt.Sprintf("%s.reg%d.o%d", ramName, i, j)
			inputs[i] = append(inputs[i], cc.ReadAsyncRedis(context.TODO(), client, name))
		}
	}

	cases := make([]reflect.SelectCase, 0)
	for _, channels := range inputs {
		for _, ch := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			})
		}
	}

	previousValues := make([][]bool, wordSize)
	for i := 0; i < wordSize; i++ {
		previousValues[i] = make([]bool, 4)
	}

	log.Printf("%#v", previousValues)

	for {
		index, value, ok := reflect.Select(cases)
		if !ok {
			break
		}

		previousValues[index/4][index%4] = value.Bool()

		fmt.Println("")
		for index, pv := range previousValues {
			fmt.Printf("index %d:\t", index)
			printBits(pv)
		}
		//log.Printf("%d: %v", index, previousValues[index])
	}
}

func printBits(bits []bool) {
	l := ""
	for i := range bits {
		b := bits[len(bits)-i-1]
		if b {
			l += "1"
		} else {
			l += "0"
		}
	}

	fmt.Println("  ", l)
}
