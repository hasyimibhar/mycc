package ir

type Variable struct {
	tp Type
}

func (v *Variable) Type() Type {
	return v.tp
}
