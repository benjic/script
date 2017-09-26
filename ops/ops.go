package ops

import (
	"io"
)

var (
	defaultProvider = &provider{
		map[uint8]Op{
			// Constants
			OP_FALSE:     opFalse,
			OP_PUSHDATA1: createOpPushNBytes(1),
			OP_PUSHDATA2: createOpPushNBytes(2),
			OP_PUSHDATA4: createOpPushNBytes(4),
			OP_1NEGATE:   op1Negate,
			OP_TRUE:      opTrue,

			// Logic
			OP_EQUAL: opEqual,

			// Stack
			OP_TOALTSTACK:   opToAltStack,
			OP_FROMALTSTACK: opFromAltStack,
			OP_IFDUP:        opIfDup,
			OP_DEPTH:        opDepth,
		},
		map[string]uint8{
			// Constants
			"OP_FALSE":     OP_FALSE,
			"OP_PUSHDATA1": OP_PUSHDATA1,
			"OP_PUSHDATA2": OP_PUSHDATA2,
			"OP_PUSHDATA4": OP_PUSHDATA4,
			"OP_1NEGATE":   OP_1NEGATE,
			"OP_TRUE":      OP_TRUE,

			// Logic
			"OP_EQUAL": OP_EQUAL,

			// Stack
			"OP_TOALTSTACK":   OP_TOALTSTACK,
			"OP_FROMALTSTACK": OP_FROMALTSTACK,
			"OP_IFDUP":        OP_IFDUP,
			"OP_DEPTH":        OP_DEPTH,
		},
	}

	// Default provides the general collection of operations.
	Default Provider = defaultProvider
)

func init() {
	for i := uint8(1); i <= 75; i++ {
		defaultProvider.ops[i] = createOpPushNBytes(i)
	}

	for i := OP_2; i <= OP_16; i++ {
		defaultProvider.ops[i] = createOpPushN(i)
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
