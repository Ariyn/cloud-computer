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
	if g.PreviousInputs[2] == false {
		results = []bool{g.previousStatus, !g.previousStatus}
	} else {
		currentStatus := g.previousStatus
		if g.PreviousInputs[0] {
			currentStatus = true
		} else if !g.PreviousInputs[0] || g.PreviousInputs[1] { // TODO: 원래 의도인 SR래치와 동작이 다름. 이건 D flip flop이 구현되어 있는 상태임
			currentStatus = false
		}

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

//var previousStatus = false
//
//func FlipFlopWithClock(inputs ...bool) []bool {
//	if inputs[2] == false {
//		return []bool{previousStatus, !previousStatus}
//	}
//
//	currentStatus := previousStatus
//	if inputs[0] {
//		currentStatus = true
//	} else if !inputs[0] || inputs[1] {
//		currentStatus = false
//	}
//
//	previousStatus = currentStatus
//	return []bool{currentStatus, !currentStatus}
//}
