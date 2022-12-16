package cloud_computer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
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
		}
	}
}

// type BoolHandler func(prev, curr bool) (o bool)
type BoolHandler func(inputs ...bool) (o bool)

// TODO: merge this and parse.Element
type Element struct {
	GateName    string
	Part        string
	IsAlias     bool
	IsParameter bool
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

	if name[0] == '$' && 7 <= len(name) && name[1:7] == "inputs" {
		name = fmt.Sprintf("${i%s}", strings.ReplaceAll(name, "$inputs.", ""))
		return name
	}

	return fmt.Sprintf("${name_variable}%s", name)
}

func parseElement(words ...string) Element {
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

	inputs := make([]<-chan bool, 0)
	for _, element := range inputElements {
		log.Println("input id", element.String())
		inputs = append(inputs, ReadAsyncRedis(ctx, client, element.String()))
	}

	outputs := make([]chan<- bool, 0)
	for _, element := range outputElements {
		element.GateName = name
		outputs = append(outputs, writeAsyncRedis(ctx, client, element.String()))
	}

	previousValues := make([]bool, len(inputs))
	previousOutput := false

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

		output := handler(previousValues...)
		if useShortcut && previousOutput == output {
			continue
		}

		err = writeRedis(ctx, client, name+".status", output)
		if err != nil {
			panic(err)
		}

		for _, ch := range outputs {
			ch <- output
		}

		previousOutput = output
	}

	return nil
}
