package ops

import (
	"bytes"
	"encoding/binary"
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
			OpOver:         opOver,
			OpPick:         opPick,
			OpRoll:         opRoll,
			OpRot:          opRot,
			OpSwap:         opSwap,
			OpTuck:         opTuck,
			Op2Drop:        op2Drop,
			Op2Dup:         op2Dup,
			Op3Dup:         op3Dup,
			Op2Over:        op2Over,
			Op2Rot:         op2Rot,
			Op2Swap:        op2Swap,

			// Arithmetic
			Op1Add:               op1Add,
			Op1Sub:               op1Sub,
			Op2Mul:               op2Mul,
			Op2Div:               op2Div,
			OpNegate:             opNegate,
			OpAbs:                opAbs,
			OpNot:                opNot,
			Op0NotEqual:          op0NotEqual,
			OpAdd:                opAdd,
			OpSub:                opSub,
			OpMul:                opMul,
			OpDiv:                opDiv,
			OpMod:                opMod,
			OpLShift:             opLShift,
			OpRShift:             opRShift,
			OpBoolAnd:            opBoolAnd,
			OpBoolOr:             opBoolOr,
			OpNumEqual:           opNumEqual,
			OpNumEqualVerify:     opNumEqualVerify,
			OpNumNotEqual:        opNumNotEqual,
			OpLessThan:           opLessThan,
			OpGreaterThan:        opGreaterThan,
			OpLessThanOrEqual:    opLessThanOrEqual,
			OpGreaterThanOrEqual: opGreaterThanOrEqual,
			OpMin:                opMin,
			OpMax:                opMax,
			OpWithin:             opWithin,
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
			"OP_OVER":         OpOver,
			"OP_PICK":         OpPick,
			"OP_ROLL":         OpRoll,
			"OP_ROT":          OpRot,
			"OP_SWAP":         OpSwap,
			"OP_TUCK":         OpTuck,
			"OP_2DROP":        Op2Drop,
			"OP_2DUP":         Op2Dup,
			"OP_3DUP":         Op3Dup,
			"OP_2OVER":        Op2Over,
			"OP_2ROT":         Op2Rot,
			"OP_2SWAP":        Op2Swap,

			// Arithmetic
			"OP_1ADD":               Op1Add,
			"OP_1SUB":               Op1Sub,
			"OP_2MUL":               Op2Mul,
			"OP_2DIV":               Op2Div,
			"OP_NEGATE":             OpNegate,
			"OP_ABS":                OpAbs,
			"OP_NOT":                OpNot,
			"OP_0NOTEQUAL":          Op0NotEqual,
			"OP_ADD":                OpAdd,
			"OP_SUB":                OpSub,
			"OP_MUL":                OpMul,
			"OP_DIV":                OpDiv,
			"OP_MOD":                OpMod,
			"OP_LSHIFT":             OpLShift,
			"OP_RSHIFT":             OpRShift,
			"OP_BOOLAND":            OpBoolAnd,
			"OP_BOOLOR":             OpBoolOr,
			"OP_NUMEQUAL":           OpNumEqual,
			"OP_NUMEQUALVERIFY":     OpNumEqualVerify,
			"OP_NUMNOTEQUAL":        OpNumNotEqual,
			"OP_LESSTHAN":           OpLessThan,
			"OP_GREATERTHAN":        OpGreaterThan,
			"OP_LESSTHANOREQUAL":    OpLessThanOrEqual,
			"OP_GREATERTHANOREQUAL": OpGreaterThanOrEqual,
			"OP_MIN":                OpMin,
			"OP_MAX":                OpMax,
			"OP_WITHIN":             OpWithin,
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

// duplicate creates a copy of a given byte slice.
func duplicate(v []byte) []byte {
	c := make([]byte, len(v))
	copy(c, v)

	return c
}

func writeNum(c Context, num int32) error {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, &num); err != nil {
		return err
	}

	c.Push(buf.Bytes())

	return nil
}

func readInt(c Context) (d int32, err error) {
	if c.Size() < 1 {
		return d, ErrInvalidStackOperation
	}

	buf := bytes.NewBuffer(c.Pop())
	err = binary.Read(buf, binary.LittleEndian, &d)

	return d, err
}
