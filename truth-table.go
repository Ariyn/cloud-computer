package cloud_computer

import (
	"errors"
	"strings"
)

var (
	ErrInvalidInputSize = errors.New("invalid input size")
	ErrNoSuchCase       = errors.New("no such case")
)

type TruthTableCase struct {
	Inputs  []bool
	Outputs []bool
}

type TruthTable struct {
	Cases      []TruthTableCase
	InputSize  int
	Inputs     []string
	OutputSize int
	Outputs    []string
}

func NewTruthTable(inputSize, outputSize int) TruthTable {
	return TruthTable{
		Cases:      make([]TruthTableCase, 1<<inputSize),
		InputSize:  inputSize,
		Inputs:     make([]string, inputSize),
		OutputSize: outputSize,
		Outputs:    make([]string, outputSize),
	}
}

func NewTruthTableWithCode(truthTableCode string) TruthTable {
	// TODO: 파싱하는 코드는 pg의 어딘가에 있을 것.
	lines := strings.Split(truthTableCode, "\n")
	if len(lines) < 1 {
		panic("no lines")
	}

	headerItems := strings.Split(lines[0], ",")
	if len(headerItems) < 2 {
		panic("minimum 1 input and 1 output required")
	}

	inputs := make([]string, 0)
	outputs := make([]string, 0)

	for _, item := range headerItems {
		item = strings.ToLower(item)
		if strings.HasPrefix(item, "i") {
			inputs = append(inputs, item)
		} else if strings.HasPrefix(item, "o") {
			outputs = append(outputs, item)
		}
	}

	possibleInputs := 1 << len(inputs)
	if len(lines) < possibleInputs+1 {
		panic("not enough rows")
	}

	cases := make([]TruthTableCase, 0)

	for _, line := range lines[1:] {
		truthItems := strings.Split(line, "")
		if len(truthItems) != len(headerItems) {
			panic("wrong items count")
		}

		c := TruthTableCase{
			Inputs:  make([]bool, 0),
			Outputs: make([]bool, 0),
		}

		for i, item := range truthItems {
			item = strings.ToLower(item)
			if item != "t" && item != "f" && item != "1" && item != "0" {
				panic("wrong item")
			}

			if i < len(inputs) {
				c.Inputs = append(c.Inputs, item == "t" || item == "1")
			} else {
				c.Outputs = append(c.Outputs, item == "t" || item == "1")
			}
		}

		cases = append(cases, c)
	}

	return TruthTable{
		Cases:      cases,
		InputSize:  len(inputs),
		Inputs:     inputs,
		OutputSize: len(outputs),
		Outputs:    outputs,
	}
}

func (t *TruthTable) GetCases(inputs []bool) ([]bool, error) {
	if len(inputs) != t.InputSize {
		return nil, ErrInvalidInputSize
	}

	truthTableMap := make(map[string][]bool)
	for _, c := range t.Cases {
		truthTableMap[boolArrayToString(c.Inputs)] = c.Outputs
	}

	if v, ok := truthTableMap[boolArrayToString(inputs)]; ok {
		return v, nil
	}
	return nil, ErrNoSuchCase
}

func boolArrayToString(b []bool) string {
	s := ""
	for _, v := range b {
		if v {
			s += "T"
		} else {
			s += "F"
		}
	}

	return s
}
