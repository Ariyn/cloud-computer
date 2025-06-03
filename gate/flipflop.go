package gate

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"github.com/go-redis/redis/v9"
	"log"
)

var _ cc.Gater = (*FlipFlop)(nil)

type FlipFlop struct {
	cc.Gate
	previousStatus bool
}

func NewFlipFlopGate(name string, inputs []cc.Element, outputs []cc.Element) *FlipFlop {
	return &FlipFlop{
		Gate: cc.Gate{
			Name:    name,
			Inputs:  inputs,
			Outputs: outputs,
		},
	}
}

func (g *FlipFlop) Init(ctx context.Context, client *redis.Client) {
	g.Gate.Init(ctx, client)

	for i, v := range g.PreviousInputs {
		g.PreviousOutputs, _ = g.Handler(i, v)

		for i, ch := range g.OutputChannels {
			ch <- g.PreviousOutputs[i]
		}
	}
}

func (g *FlipFlop) Handler(index int, value bool) (results []bool, changed bool) {
	g.PreviousInputs[index] = value
	if g.PreviousInputs[2] == false { // If Enable is false
		// Maintain previous state
		results = []bool{g.previousStatus, !g.previousStatus}
	} else { // If Enable is true
		s := g.PreviousInputs[0]
		r := g.PreviousInputs[1]
		currentStatus := g.previousStatus // Default to previous state

		if s && !r { // S=1, R=0
			currentStatus = true
		} else if !s && r { // S=0, R=1
			currentStatus = false
		} else if s && r { // S=1, R=1 (Invalid state)
			// This is an invalid state for an SR latch.
			// Maintaining existing behavior of setting Q=1, Q'=0.
			currentStatus = true
		}
		// If S=0, R=0, currentStatus remains g.previousStatus (previous state is maintained)

		g.previousStatus = currentStatus
		results = []bool{currentStatus, !currentStatus}
	}

	if g.IsVerbose {
		log.Printf("name %s, inputs: %v, results: %v", g.Name, g.PreviousInputs, results)
	}

	changed = g.Changed(results, g.PreviousOutputs)
	g.PreviousOutputs = results

	return
}

func (g *FlipFlop) GetType() string {
	return "flipflop"
}
