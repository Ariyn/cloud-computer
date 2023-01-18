package gate

import (
	cc "github.com/ariyn/cloud-computer"
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

func (gate Or) Handler(inputs ...bool) (results []bool) {
	results = append(results, inputs[0])

	for _, b := range inputs[1:] {
		results[0] = results[0] || b
	}

	if gate.IsVerbose {
		log.Printf("name %s, inputs: %v, results: %v", gate.Name, inputs, results)
	}
	return
}
