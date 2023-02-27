package gate

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"github.com/go-redis/redis/v9"
	"log"
)

var _ cc.Gater = (*Or)(nil)

type Or struct {
	cc.Gate
}

func NewOrGate(name string, inputs []cc.Element, outputs []cc.Element) *Or {
	return &Or{
		Gate: cc.Gate{
			Name:    name,
			Inputs:  inputs,
			Outputs: outputs,
		},
	}
}

func (g *Or) Init(ctx context.Context, client *redis.Client) {
	g.Gate.Init(ctx, client)

	for i, v := range g.PreviousInputs {
		g.PreviousOutputs, _ = g.Handler(i, v)

		for i, ch := range g.OutputChannels {
			ch <- g.PreviousOutputs[i]
		}
	}
}

func (g *Or) Handler(index int, value bool) (results []bool, changed bool) {
	g.PreviousInputs[index] = value
	results = append(results, g.PreviousInputs[0])

	for _, b := range g.PreviousInputs[1:] {
		results[0] = results[0] || b
	}

	if g.IsVerbose {
		log.Printf("name %s, inputs: %v, results: %v", g.Name, g.PreviousInputs, results)
	}

	changed = g.Changed(results, g.PreviousOutputs)
	g.PreviousOutputs = results

	return
}

func (g *Or) GetType() string {
	return "or"
}
