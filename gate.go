package cloud_computer

import (
	"context"
	"github.com/go-redis/redis/v9"
	"reflect"
)

var _ Gater = (*Gate)(nil)

type Gate struct {
	Name            string
	Inputs          []Element
	Outputs         []Element
	UseOptimization bool
	IsVerbose       bool
	PreviousInputs  []bool
	InputChannels   []<-chan bool
	PreviousOutputs []bool
	OutputChannels  []chan<- bool
}

func (g *Gate) Init(ctx context.Context, client *redis.Client) {
	g.PreviousInputs = make([]bool, len(g.Inputs))
	g.PreviousOutputs = make([]bool, len(g.Outputs))

	g.InputChannels = make([]<-chan bool, len(g.Inputs))
	for i, element := range g.GetInputs() {
		if element.IsStaticValue {
			g.PreviousInputs[i] = element.StaticValue
			continue
		}

		g.InputChannels[i] = ReadAsyncRedis(ctx, client, element.String())

		v, err := ReadRedis(ctx, client, element.String()+".status")
		if err != nil {
			panic(err)
		}
		g.PreviousInputs[i] = v
	}

	for i, element := range g.Outputs {
		element.GateName = g.GetName()
		g.Outputs[i] = element
	}

	g.OutputChannels = make([]chan<- bool, len(g.Outputs))
	for i, element := range g.GetOutputs() {
		g.OutputChannels[i] = WriteAsyncRedis(ctx, client, element.String())
	}
}

func (g *Gate) SelectCases() (cases []reflect.SelectCase) {
	cases = make([]reflect.SelectCase, len(g.InputChannels))

	for i, ch := range g.InputChannels {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		}
	}

	return
}

func (g *Gate) Handler(_ int, _ bool) ([]bool, bool) {
	panic("implement me")
}

func (g *Gate) GetName() string {
	return g.Name
}

func (g *Gate) GetInputs() (inputs []Element) {
	return g.Inputs
}

func (g *Gate) GetInputSize() int {
	return len(g.Inputs)
}

func (g *Gate) GetOutputs() (outputs []Element) {
	return g.Outputs
}

func (g *Gate) GetOutputSize() int {
	return len(g.Outputs)
}

func (g *Gate) GetOutputChannels() []chan<- bool {
	return g.OutputChannels
}

func (g *Gate) SetOptimization(v bool) {
	g.UseOptimization = v
}

func (g *Gate) SetVerbose(v bool) {
	g.IsVerbose = v
}

func (g *Gate) Changed(v1, v2 []bool) bool {
	for i := range v1 {
		if v1[i] != v2[i] {
			return true
		}
	}

	return false
}

func (g *Gate) GetType() string {
	panic("not implements")
}

type Gater interface {
	Init(ctx context.Context, client *redis.Client)
	Handler(index int, input bool) (o []bool, changed bool)
	GetName() string
	GetInputs() (inputs []Element)
	GetInputSize() int
	SelectCases() (cases []reflect.SelectCase)
	GetOutputs() (outputs []Element)
	GetOutputSize() int
	GetOutputChannels() []chan<- bool
	GetType() string
}
