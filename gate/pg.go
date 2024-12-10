package gate

import (
	"context"
	cc "github.com/ariyn/cloud-computer"
	"log"
)

var _ cc.Gater = (*ProgrammableGate)(nil)

type ProgrammableGate struct {
	cc.Gate
	table cc.TruthTable
}

func NewPGGate(name string, inputs []cc.Element, outputs []cc.Element, truthTable cc.TruthTable) *ProgrammableGate {
	return &ProgrammableGate{
		Gate: cc.Gate{
			Name:    name,
			Inputs:  inputs,
			Outputs: outputs,
		},
		table: truthTable,
	}
}

func (g *ProgrammableGate) Init(ctx context.Context, client cc.DataSourcer) {
	g.Gate.Init(ctx, client)

	for i, v := range g.PreviousInputs {
		g.PreviousOutputs, _ = g.Handler(i, v)

		for i, ch := range g.OutputChannels {
			ch <- g.PreviousOutputs[i]
		}
	}
}

func (g *ProgrammableGate) Handler(index int, value bool) (results []bool, changed bool) {
	inputs := append([]bool{}, g.PreviousInputs...)
	inputs[index] = value

	var err error
	results, err = g.table.GetCases(inputs)
	if err != nil {
		panic(err)
	}

	if g.IsVerbose {
		log.Printf("name %s, inputs: %v, results: %v", g.Name, g.PreviousInputs, results)
	}

	changed = g.Changed(results, g.PreviousOutputs)

	g.PreviousInputs[index] = value
	g.PreviousOutputs = results

	return
}

func (g *ProgrammableGate) GetType() string {
	return "programmable-gate"
}
