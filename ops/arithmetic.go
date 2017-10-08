package ops

// TODO(benjic): Document operation descriptions
const (
	Op1Add               uint8 = 0x8b
	Op1Sub               uint8 = 0x8c
	Op2Mul               uint8 = 0x8d
	Op2Div               uint8 = 0x8e
	OpNegate             uint8 = 0x8f
	OpAbs                uint8 = 0x90
	OpNot                uint8 = 0x91
	Op0NotEqual          uint8 = 0x92
	OpAdd                uint8 = 0x93
	OpSub                uint8 = 0x94
	OpMul                uint8 = 0x95
	OpDiv                uint8 = 0x96
	OpMod                uint8 = 0x97
	OpLShift             uint8 = 0x98
	OpRShift             uint8 = 0x99
	OpBoolAnd            uint8 = 0x9a
	OpBoolOr             uint8 = 0x9b
	OpNumEqual           uint8 = 0x9c
	OpNumEqualVerify     uint8 = 0x9d
	OpNumNotEqual        uint8 = 0x9e
	OpLessThan           uint8 = 0x9f
	OpGreaterThan        uint8 = 0xa0
	OpLessThanOrEqual    uint8 = 0xa1
	OpGreaterThanOrEqual uint8 = 0xa2
	OpMin                uint8 = 0xa3
	OpMax                uint8 = 0xa4
	OpWithin             uint8 = 0xa5
)

func op1Add(c Context) error {
	d, err := readInt(c)
	if err != nil {
		return err
	}
	writeInt(c, d+1)
	return nil
}

func op1Sub(c Context) error {
	// TODO: implement me
	return nil
}

func op2Mul(c Context) error {
	// TODO: implement me
	return nil
}

func op2Div(c Context) error {
	// TODO: implement me
	return nil
}

func opNegate(c Context) error {
	// TODO: implement me
	return nil
}

func opAbs(c Context) error {
	// TODO: implement me
	return nil
}

func opNot(c Context) error {
	// TODO: implement me
	return nil
}

func op0NotEqual(c Context) error {
	// TODO: implement me
	return nil
}

func opAdd(c Context) error {
	// TODO: implement me
	return nil
}

func opSub(c Context) error {
	// TODO: implement me
	return nil
}

func opMul(c Context) error {
	// TODO: implement me
	return nil
}

func opDiv(c Context) error {
	// TODO: implement me
	return nil
}

func opMod(c Context) error {
	// TODO: implement me
	return nil
}

func opLShift(c Context) error {
	// TODO: implement me
	return nil
}

func opRShift(c Context) error {
	// TODO: implement me
	return nil
}

func opBoolAnd(c Context) error {
	// TODO: implement me
	return nil
}

func opBoolOr(c Context) error {
	// TODO: implement me
	return nil
}

func opNumEqual(c Context) error {
	// TODO: implement me
	return nil
}

func opNumEqualVerify(c Context) error {
	// TODO: implement me
	return nil
}

func opNumNotEqual(c Context) error {
	// TODO: implement me
	return nil
}

func opLessThan(c Context) error {
	// TODO: implement me
	return nil
}

func opGreaterThan(c Context) error {
	// TODO: implement me
	return nil
}

func opLessThanOrEqual(c Context) error {
	// TODO: implement me
	return nil
}

func opGreaterThanOrEqual(c Context) error {
	// TODO: implement me
	return nil
}

func opMin(c Context) error {
	// TODO: implement me
	return nil
}

func opMax(c Context) error {
	// TODO: implement me
	return nil
}

func opWithin(c Context) error {
	// TODO: implement me
	return nil
}
