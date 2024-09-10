package gate

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"log"
)

var _ cc.Gater = (*Input)(nil)

type Input struct {
	cc.Gate
}

func NewInputGate(name string, inputs []cc.Element, outputs []cc.Element) *Input {
	return &Input{
		Gate: cc.Gate{
			Name:    name,
			Inputs:  inputs,
			Outputs: outputs,
		},
	}
}

func (g *Input) Init(ctx context.Context, client cc.DataSourcer) {
	g.Gate.Init(ctx, client)

	for i, v := range g.PreviousInputs {
		initialOutputs, _ := g.Handler(i, v)
		for i := range initialOutputs {
			g.PreviousOutputs[i] = initialOutputs[i]
		}

		for i, ch := range g.OutputChannels {
			ch <- g.PreviousOutputs[i]
		}
	}
}

func (g *Input) Handler(index int, value bool) (results []bool, changed bool) {
	g.PreviousInputs[index] = value
	for i := range g.PreviousInputs {
		results = append(results, g.PreviousInputs[i])
	}

	if g.IsVerbose {
		log.Printf("name %s, inputs: %v, results: %v", g.Name, g.PreviousInputs, results)
	}

	changed = g.Changed(g.PreviousInputs, g.PreviousOutputs)
	for i := range g.PreviousOutputs {
		g.PreviousOutputs[i] = results[i]
	}

	return
}

func (g *Input) GetType() string {
	return "input"
}
