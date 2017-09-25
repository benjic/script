package ops

import (
	"io"
)

var (
	// OpRegistry provides a lookup for an operation.
	OpRegistry = map[uint8]Op{
		// Constants
		OP_FALSE:     opFalse,
		OP_PUSHDATA1: createOpPushNBytes(1),
		OP_PUSHDATA2: createOpPushNBytes(2),
		OP_PUSHDATA4: createOpPushNBytes(4),
		OP_1NEGATE:   op1Negate,
		OP_TRUE:      opTrue,

		// Logic
		OP_EQUAL: opEqual,
	}
)

func init() {
	for i := uint8(1); i <= 75; i++ {
		OpRegistry[i] = createOpPushNBytes(i)
	}

	for i := OP_2; i <= OP_16; i++ {
		OpRegistry[i] = createOpPushN(i)
	}
}

// An Op is a single instruction to be executed.
type Op func(Context) error

// Context allows an operation to manipulate the execution context.
type Context interface {
	io.Reader

	Pop() []byte
	PopAlt() []byte

	Push([]byte)
	PushAlt([]byte)
}
