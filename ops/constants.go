package ops

const (
	OP_FALSE   = 0x00
	OP_1NEGATE = 0x4f
	OP_TRUE    = 0x51
)

func opFalse(c Context) {
	c.Push([]byte{0x00, 0x00, 0x0, 0x00})
}

func op1Negate(c Context) {
	c.Push([]byte{0x40, 0x00, 0x00, 0x01})
}

func opTrue(c Context) {
	c.Push([]byte{0x00, 0x00, 0x00, 0x01})
}