package gate

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"github.com/go-redis/redis/v9"
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

func (g *Input) Init(ctx context.Context, client *redis.Client) {
	g.Gate.Init(ctx, client)

	for i, v := range g.PreviousInputs {
		g.PreviousOutputs, _ = g.Handler(i, v)

		for i, ch := range g.OutputChannels {
			ch <- g.PreviousOutputs[i]
		}
	}
}

func (g *Input) Handler(index int, value bool) (results []bool, changed bool) {
	g.PreviousInputs[index] = value
	results = g.PreviousInputs

	if g.IsVerbose {
		log.Printf("name %s, inputs: %v, results: %v", g.Name, g.PreviousInputs, results)
	}

	changed = g.Changed(results, g.PreviousOutputs)
	g.PreviousOutputs = results

	return
}

func (g *Input) GetType() string {
	return "input"
}
