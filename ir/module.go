package ir

type Module struct {
	globals *Globals
}

type Globals struct {
	variables []*Variable
	functions []*Function
}

func (m *Module) Globals() *Globals {
	return m.globals
}

func (g *Globals) Variables() []*Variable {
	return g.variables
}

func (g *Globals) Functions() []*Function {
	return g.functions
}
