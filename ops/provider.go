package ops

// A Provider enables the query of op information. This allows for custom sets
// of Ops to be defined and used by a parser or evaluator.
type Provider interface {
	// GetOp returns a function for a given op code.
	GetOp(uint8) (Op, bool)

	// GetCode returns an op code for a given word.
	GetCode(string) (uint8, bool)
}

type provider struct {
	ops   map[uint8]Op
	codes map[string]uint8
}

func (p *provider) GetOp(code uint8) (Op, bool) {
	op, ok := p.ops[code]
	return op, ok
}

func (p *provider) GetCode(word string) (uint8, bool) {
	code, ok := p.codes[word]
	return code, ok
}
