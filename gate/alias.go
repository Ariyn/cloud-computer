package gate

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"github.com/go-redis/redis/v9"
	"log"
)

var _ cc.Gater = (*Alias)(nil)

type Alias struct {
	cc.Gate
}

func NewAliasGate(name string, inputs []cc.Element, outputs []cc.Element) *Alias {
	return &Alias{
		Gate: cc.Gate{
			Name:    name,
			Inputs:  inputs,
			Outputs: outputs,
		},
	}
}

func (g *Alias) Init(ctx context.Context, client *redis.Client) {
	g.Gate.Init(ctx, client)

	for i, v := range g.PreviousInputs {
		g.PreviousOutputs, _ = g.Handler(i, v)

		for i, ch := range g.OutputChannels {
			ch <- g.PreviousOutputs[i]
		}
	}
}

func (g *Alias) Handler(index int, value bool) (results []bool, changed bool) {
	g.PreviousInputs[index] = value
	results = g.PreviousInputs

	if g.IsVerbose {
		log.Printf("name %s, inputs: %v, results: %v", g.Name, g.PreviousInputs, results)
	}

	changed = g.Changed(results, g.PreviousOutputs)
	g.PreviousOutputs = results

	return
}
