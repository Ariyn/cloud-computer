package cloud_computer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var InvalidElement error

type ArrayStringFlag []string

func (asf *ArrayStringFlag) String() string {
	return strings.Join(*asf, " ")
}

func (asf *ArrayStringFlag) Set(v string) error {
	*asf = append(*asf, v)
	return nil
}

var Inputs ArrayStringFlag
var Name string
var UseOptimization bool = true

func init() {
	//flag.StringVar(&Name, "name", "", "and, or, not, etc...")
	//flag.StringVar(&InputName1, "i1", "input_1", "input_1")
	//flag.StringVar(&InputName2, "i2", "input_2", "input_2")
	//flag.Var(&Inputs, "inputs", "inputs. first input will be set to input 1\neg) -inputs a.o1 -inputs b.o1")

	parseArguments()

	InvalidElement = errors.New("invalid element")
}

func parseArguments() {
	for index := 1; index < len(os.Args); index++ {
		arg := os.Args[index]
		if !strings.HasPrefix(arg, "-") {
			continue
		}

		if strings.HasPrefix(arg, "-i") {
			log.Println(arg, os.Args[index+1])
			Inputs = append(Inputs, os.Args[index+1])
		} else if arg == "-name" {
			Name = os.Args[index+1]
		} else if arg == "-no-optimization" {
			UseOptimization = false
		}
	}
}

// type BoolHandler func(prev, curr bool) (o bool)
type BoolHandler func(inputs ...bool) (o []bool)

// TODO: merge this and parse.Element
type Element struct {
	GateName      string
	Part          string
	IsAlias       bool
	IsParameter   bool
	StaticValue   bool
	IsStaticValue bool
}

func (e Element) String() string {
	es := []string{e.GateName}
	if !e.IsAlias && e.Part != "" {
		es = append(es, e.Part)
	}

	return strings.Join(es, ".")
}

func (e Element) Bash() string {
	name := e.String()

	if name == "" && e.IsStaticValue {
		name = ""
		if e.StaticValue == true {
			name = "1"
		} else {
			name = "0"
		}
		return name
	}

	if name[0] == '$' && 7 <= len(name) && name[1:7] == "inputs" {
		name = fmt.Sprintf("${i%s}", strings.ReplaceAll(name, "$inputs.", ""))
		return name
	}

	return fmt.Sprintf("${name_variable}%s", name)
}

func parseElement(words ...string) Element {
	if n, err := strconv.Atoi(words[0]); err == nil {
		e := Element{
			IsStaticValue: true,
		}

		if n == 0 {
			e.StaticValue = false
		} else {
			e.StaticValue = true
		}

		return e
	}

	e := Element{
		GateName: strings.Join(words, "."),
	}

	//if 2 <= len(words) {
	//	// Parse Element recursive
	//	e.Part = strings.Join(words[1:], ".")
	//}
	//
	//log.Println(e)

	return e
}

func ParseInputs(inputs ...string) (elements []Element) {
	for _, i := range inputs {
		words := strings.Split(i, ".")
		e := parseElement(words...)
		elements = append(elements, e)
	}

	return
}

func CreateOutputs(size int) (elements []Element) {
	for i := 0; i < size; i++ {
		e := Element{
			Part:    fmt.Sprintf("o%d", i+1),
			IsAlias: false,
		}

		elements = append(elements, e)
	}
	return
}

func RunRedis(handler BoolHandler, name string, inputElements []Element, outputElements []Element, useShortcut bool) (err error) {
	log.Println(handler, name, inputElements, outputElements)

	ctx := context.TODO()
	client := ConnectRedis()

	err = writeRedis(ctx, client, name+".status", false)
	if err != nil {
		panic(err)
	}
	defer deleteRedis(ctx, client, name+".status")

	previousValues := make([]bool, len(inputElements))
	previousOutputs := make([]bool, len(inputElements))

	inputs := make([]<-chan bool, 0)
	for i, element := range inputElements {
		if element.IsStaticValue {
			previousValues[i] = element.StaticValue
			continue
		}
		inputs = append(inputs, ReadAsyncRedis(ctx, client, element.String()))
	}

	outputChannels := make([]chan<- bool, 0)
	for _, element := range outputElements {
		element.GateName = name
		outputChannels = append(outputChannels, writeAsyncRedis(ctx, client, element.String()))
	}

	outputs := handler(previousValues...)
	for i, ch := range outputChannels {
		ch <- outputs[i]
	}

	cases := make([]reflect.SelectCase, 0)
	for _, ch := range inputs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	for {
		index, value, ok := reflect.Select(cases)
		if !ok {
			break
		}

		previousValues[index] = value.Bool()

		outputs := handler(previousValues...)
		if useShortcut && equalOutputs(previousOutputs, outputs) {
			continue
		}

		for i, ch := range outputChannels {
			ch <- outputs[i]
		}

		previousOutputs = outputs
	}

	return nil
}

func equalOutputs(v1, v2 []bool) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if v1[i] != v2[i] {
			return false
		}
	}

	return true
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
		outputs = append(outputs, writeAsyncRedis(ctx, client, element.String()))
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
