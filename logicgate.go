package cloud_computer

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"
)

var ErrEmptySignals = errors.New("empty signals")

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

func getSelectCaseSignals(signals ...os.Signal) (sc reflect.SelectCase, err error) {
	if len(signals) == 0 {
		err = ErrEmptySignals
		return
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, signals...)

	sc = reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(sigs),
	}

	return
}

func RunGateWithRedis(ctx context.Context, hostUrl string, gate Gater) (err error) {
	client := ConnectRedis(hostUrl)

	gate.Init(ctx, client)

	if gate.GetType() == "input" {
		addInput(ctx, client, gate.GetName(), gate.GetName())
		addChildren(ctx, client, gate.GetName())
	}
	if gate.GetType() == "alias" {
		addOutput(ctx, client, gate.GetName(), gate.GetName())
	}

	// TODO: inputs, outputs가 redis에 붙는 동작은 여기서 진행해야 한다.
	// Gate type은 좀 더 게이트 동작 그 자체에 집중할 수 있도록 수정 필요

	cases := gate.SelectCases()

	sc, err := getSelectCaseSignals(syscall.SIGINT, syscall.SIGTERM)
	if err != nil {
		panic(err)
	}
	cases = append(cases, sc)

	defer func() {
		deleteRedis(ctx, client, gate.GetName()+".status")
		for _, element := range gate.GetOutputs() {
			element.GateName = gate.GetName()
			deleteRedis(ctx, client, element.String()+".status")
		}

		if gate.GetType() == "input" {
			parents := strings.Split(gate.GetName(), ".")
			grandParent := strings.Join(parents[:len(parents)-1], ".")
			deleteRedis(ctx, client, grandParent+".inputs")
			deleteRedis(ctx, client, grandParent+".children")
		}
		if gate.GetType() == "alias" {
			parents := strings.Split(gate.GetName(), ".")
			deleteRedis(ctx, client, strings.Join(parents[:len(parents)-1], ".")+".outputs")
		}
	}()

	for {
		index, value, ok := reflect.Select(cases)
		if !ok {
			return nil
		}

		if index == len(cases)-1 {
			return nil
		}

		outputs, changed := gate.Handler(index, value.Bool())
		if !changed {
			continue
		}

		err = writeRedis(ctx, client, gate.GetName()+".status", outputs[0])
		if err != nil {
			panic(err)
		}

		for i, ch := range gate.GetOutputChannels() {
			ch <- outputs[i]
		}
	}

	return nil
}
