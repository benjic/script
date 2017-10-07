package ops

import (
	"io"
)

var (
	defaultProvider = &provider{
		map[uint8]Op{
			// Constants
			OpFalse:     opFalse,
			OpPushData1: createOpPushNBytes(1),
			OpPushData2: createOpPushNBytes(2),
			OpPushData4: createOpPushNBytes(4),
			Op1Negate:   op1Negate,
			OpTrue:      opTrue,

			// Logic
			OpEqual: opEqual,

			// Stack
			OpToAltStack:   opToAltStack,
			OpFromAltStack: opFromAltStack,
			OpIfDup:        opIfDup,
			OpDepth:        opDepth,
			OpDup:          opDup,
			OpNip:          opNip,
		},
		map[string]uint8{
			// Constants
			"OP_FALSE":     OpFalse,
			"OP_PUSHDATA1": OpPushData1,
			"OP_PUSHDATA2": OpPushData2,
			"OP_PUSHDATA4": OpPushData4,
			"OP_1NEGATE":   Op1Negate,
			"OP_TRUE":      OpTrue,

			// Logic
			"OP_EQUAL": OpEqual,

			// Stack
			"OP_TOALTSTACK":   OpToAltStack,
			"OP_FROMALTSTACK": OpFromAltStack,
			"OP_IFDUP":        OpIfDup,
			"OP_DEPTH":        OpDepth,
			"OP_DUP":          OpDup,
			"OP_NIP":          OpNip,
		},
	}

	// Default provides the general collection of operations.
	Default Provider = defaultProvider
)

func init() {
	for i := uint8(1); i <= 75; i++ {
		defaultProvider.ops[i] = createOpPushNBytes(i)
	}

	for i := Op2; i <= Op16; i++ {
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

	Size() int
	SizeAlt() int
}
