package ops

import (
	"io"
)

var (
	// OpRegistry provides a lookup for an operation.
	OpRegistry = map[uint8]Op{
		OP_1NEGATE: op1Negate,
		OP_EQUAL:   opEqual,
		OP_FALSE:   opFalse,
		OP_TRUE:    opTrue,
	}
)

// An Op is a single instruction to be executed.
type Op func(Context)

// Context allows an operation to manipulate the execution context.
type Context interface {
	io.Reader
	Pop() []byte
	Push([]byte)
}
