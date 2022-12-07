package cloud_computer

import (
	"context"
	"flag"
	"fmt"
	"log"
	"reflect"
)

var InputName1 string
var InputName2 string
var Name string

func init() {
	flag.StringVar(&Name, "name", "", "and, or, not, etc...")
	flag.StringVar(&InputName1, "i1", "input_1", "input_1")
	flag.StringVar(&InputName2, "i2", "input_2", "input_2")
}

type Handler func(i ...int) int
type BoolHandler func(i ...bool) (o bool)

func Run(handler Handler, inputPaths []string, outputPaths []string) (err error) {
	log.Println(handler, inputPaths, outputPaths)

	inputs := make([]chan int, 0)
	for _, path := range inputPaths {
		input, err := openInput(path)
		if err != nil {
			log.Println(path, err)
			return err
		}

		defer input.Close()

		ch := readAsync(input)
		inputs = append(inputs, ch)
	}

	outputs := make([]chan int, 0)
	for _, path := range outputPaths {
		output, err := openOutput(path)
		if err != nil {
			log.Println(path, err)
			return err
		}
		defer output.Close()

		ch := writeAsync(output)
		outputs = append(outputs, ch)
	}

	previousValues := make([]int, len(inputPaths))

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

		previousValues[index] = int(value.Int())

		output := handler(previousValues...)
		log.Println(previousValues, output)

		for _, ch := range outputs {
			ch <- output
		}
	}

	return nil
}

// TODO: combine this and run
func RunRedis(handler BoolHandler, name string, inputIds []string, outputSize int) (err error) {
	log.Println(handler, name, inputIds, outputSize)

	ctx := context.TODO()
	client := ConnectRedis()

	err = writeRedis(ctx, client, name+"_status", false)
	if err != nil {
		panic(err)
	}
	defer deleteRedis(ctx, client, name+"_status")

	inputs := make([]<-chan bool, 0)
	for _, inputId := range inputIds {
		log.Println("input id", inputId)
		inputs = append(inputs, ReadAsyncRedis(ctx, client, inputId))
	}

	outputs := make([]chan<- bool, 0)
	for i := 0; i < outputSize; i++ {
		outputName := fmt.Sprintf("%s_output_%d", name, i+1)
		outputs = append(outputs, writeAsyncRedis(ctx, client, outputName))
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
		if previousOutput == output {
			continue
		}

		err = writeRedis(ctx, client, name+"_status", output)
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
