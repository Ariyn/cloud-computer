package cloud_computer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
	"time"
)

const DebugChannelName = "CLOUD_COMPUTER_DEBUG"

var InvalidElement = errors.New("invalid element")
var Inputs = make([]string, 0)
var Outputs = make([]string, 0)
var Name string
var UseOptimization bool = true
var IsDebugging bool
var IsVerbose bool

type BoolHandler func(inputs ...bool) (o []bool)

func init() {
	parseArguments()
}

// TODO: make redis as interface
func RunRedis(handler BoolHandler, name string, inputElements []Element, outputElements []Element, useShortcut bool, isAlias, isInput bool) (err error) {
	ctx := context.TODO()
	client := ConnectRedis()

	if isInput {
		addInput(ctx, client, name, name)
		addChildren(ctx, client, name)
	}

	previousValues := make([]bool, len(inputElements))
	previousOutputs := make([]bool, len(outputElements))

	inputs := make([]<-chan bool, 0)
	for i, element := range inputElements {
		if element.IsStaticValue {
			previousValues[i] = element.StaticValue
			continue
		}

		inputs = append(inputs, ReadAsyncRedis(ctx, client, element.String()))

		v, err := ReadRedis(ctx, client, element.String()+".status")
		if err != nil {
			panic(err)
		}
		if IsVerbose {
			log.Println(i, v)
		}
		previousValues[i] = v
	}

	if isAlias {
		addOutput(ctx, client, name, name)
	}

	// TODO: inputs와 이름의 차이가 큼. 수정할 것
	outputChannels := make([]chan<- bool, 0)
	for _, element := range outputElements {
		element.GateName = name
		outputChannels = append(outputChannels, WriteAsyncRedis(ctx, client, element.String()))
	}

	// TODO: 여기서 문제가 없을지 확인 필요하다. 과연...
	previousOutputs = handler(previousValues...)
	for i, ch := range outputChannels {
		ch <- previousOutputs[i]
	}

	cases := make([]reflect.SelectCase, 0)
	if IsDebugging {
		debugInput := ReadAsyncRedis(ctx, client, DebugChannelName)

		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(debugInput),
		})
	}

	for _, ch := range inputs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	cases = append(cases, reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(sigs),
	})

	for {
		index, value, ok := reflect.Select(cases)
		if !ok {
			break
		}

		if index == len(cases)-1 {
			deleteRedis(ctx, client, name+".status")
			for _, element := range outputElements {
				element.GateName = name
				deleteRedis(ctx, client, element.String()+".status")
				//log.Println("removing", element.String()+".status")
			}

			if isInput {
				parents := strings.Split(name, ".")
				grandParent := strings.Join(parents[:len(parents)-1], ".")
				deleteRedis(ctx, client, grandParent+".inputs")
				deleteRedis(ctx, client, grandParent+".children")
			}
			if isAlias {
				parents := strings.Split(name, ".")
				deleteRedis(ctx, client, strings.Join(parents[:len(parents)-1], ".")+".outputs")
			}
			break
		}

		if IsDebugging {
			if index == 0 {
				outputs := handler(previousValues...)
				if useShortcut && equalOutputs(previousOutputs, outputs) {
					continue
				}

				for i, ch := range outputChannels {
					ch <- outputs[i]
				}

				previousOutputs = outputs
				continue
			}

			index -= 1
		}
		previousValues[index] = value.Bool()

		outputs := handler(previousValues...)
		if IsVerbose {
			log.Println(outputs)
		}

		if useShortcut && equalOutputs(previousOutputs, outputs) {
			continue
		}

		if IsVerbose {
			log.Println("write", outputs)
			log.Println(name+".status", outputs[0])
		}

		// WARN: 임시 코드. 좀 더 우아하게 수정할 것
		err = writeRedis(ctx, client, name+".status", outputs[0])
		if err != nil {
			panic(err)
		}

		if !IsDebugging {
			for i, ch := range outputChannels {
				ch <- outputs[i]
			}

			previousOutputs = outputs
		}
	}

	return nil
}

func Clock(clk int, outputElements []Element) (err error) {
	log.Println(clk, outputElements)

	ctx := context.TODO()
	client := ConnectRedis()

	name := fmt.Sprintf("clock.%dHz", clk)

	err = writeRedis(ctx, client, name+".status", false)
	if err != nil {
		panic(err)
	}
	defer deleteRedis(ctx, client, name+".status")

	outputs := make([]chan<- bool, 0)
	for _, element := range outputElements {
		element.GateName = name
		outputs = append(outputs, WriteAsyncRedis(ctx, client, element.String()))
	}

	previousValues := false

	for {
		start := time.Now()
		curr := !previousValues

		err = writeRedis(ctx, client, name+".status", curr)
		if err != nil {
			panic(err)
		}

		for _, ch := range outputs {
			ch <- curr
		}

		previousValues = curr

		time.Sleep(time.Second/time.Duration(clk) - time.Now().Sub(start))
	}
}
