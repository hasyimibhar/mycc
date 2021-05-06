package ir

type BasicBlock struct {
	label        string
	instructions []Instruction
}

func (b *BasicBlock) Label() string {
	return b.label
}

func (b *BasicBlock) Instructions() []Instruction {
	return b.instructions
}

func (b *BasicBlock) TermInstruction() Instruction {
	return b.instructions[len(b.instructions)-1]
}
