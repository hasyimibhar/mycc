package ir

type Opcode string

const (
	OpcodeAdd = "add"
)

type Instruction interface {
	Serializable

	Opcode() Opcode
}
