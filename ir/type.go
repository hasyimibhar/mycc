package ir

type Type interface {
	Name() string
}

type TypeVoid struct{}

func (t TypeVoid) Name() string {
	return "void"
}
