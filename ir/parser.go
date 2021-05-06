package ir

import "io"

type Parser struct{}

func (p Parser) Parse(rd io.Reader) (*Module, error) {
	return nil, nil
}
