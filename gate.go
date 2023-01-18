package cloud_computer

var _ Gater = (*Gate)(nil)

type Gate struct {
	Name            string
	Inputs          []Element
	Outputs         []Element
	UseOptimization bool
	IsVerbose       bool
}

func (g *Gate) Handler(inputs ...bool) (o []bool) {
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

func (g *Gate) SetOptimization(v bool) {
	g.UseOptimization = v
}

func (g *Gate) SetVerbose(v bool) {
	g.IsVerbose = v
}

type Gater interface {
	Handler(inputs ...bool) (o []bool)
	GetName() string
	GetInputs() (inputs []Element)
	GetInputSize() int
	GetOutputs() (outputs []Element)
	GetOutputSize() int
}
