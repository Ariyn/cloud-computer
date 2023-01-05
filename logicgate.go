package cloud_computer

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"strings"
	"syscall"
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

const DebugChannelName = "CLOUD_COMPUTER_DEBUG"

var Inputs ArrayStringFlag
var Outputs ArrayStringFlag
var Name string
var UseOptimization bool = true
var IsDebugging bool
var Client redis.Conn

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
			// TODO: -i 뒤의 숫자를 파싱해서 해당 번호에 넣기
			//log.Println(arg, os.Args[index+1])
			Inputs = append(Inputs, os.Args[index+1])
			index += 1
		} else if arg == "-name" {
			Name = os.Args[index+1]
		} else if arg == "-no-optimization" {
			UseOptimization = false
		} else if strings.HasPrefix(arg, "-o") {
			// TODO: -o 뒤의 숫자를 파싱해서 해당 번호에 넣기
			Outputs = append(Outputs, os.Args[index+1])
			index += 1
		} else if arg == "-debug" {
			if os.Args[index+1] == "1" {
				IsDebugging = true
				log.Println("RUNNING DEBUGGING MODE")
			}
			index += 1
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

func RunRedis(handler BoolHandler, name string, inputElements []Element, outputElements []Element, useShortcut bool, isAlias, isInput bool) (err error) {
	ctx := context.TODO()
	client := ConnectRedis()

	if isInput {
		addInput(ctx, client, name, name)
		addChildren(ctx, client, name)
	}

	previousValues := make([]bool, len(inputElements))
	previousOutputs := make([]bool, len(inputElements))

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
			for _, element := range outputElements {
				element.GateName = name
				deleteRedis(ctx, client, element.String()+".status")
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
		if useShortcut && equalOutputs(previousOutputs, outputs) {
			continue
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
