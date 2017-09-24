package ops

const (
	OP_0         uint8 = 0x00
	OP_FALSE     uint8 = 0x00
	OP_PUSHDATA1 uint8 = 0x4c
	OP_PUSHDATA2 uint8 = 0x4d
	OP_PUSHDATA4 uint8 = 0x4e
	OP_1NEGATE   uint8 = 0x4f
	OP_TRUE      uint8 = 0x51
	OP_1         uint8 = 0x51
	OP_2         uint8 = 0x52
	OP_3         uint8 = 0x53
	OP_4         uint8 = 0x54
	OP_5         uint8 = 0x55
	OP_6         uint8 = 0x56
	OP_7         uint8 = 0x57
	OP_8         uint8 = 0x58
	OP_9         uint8 = 0x59
	OP_10        uint8 = 0x5a
	OP_11        uint8 = 0x5b
	OP_12        uint8 = 0x5c
	OP_13        uint8 = 0x5d
	OP_14        uint8 = 0x5e
	OP_15        uint8 = 0x5f
	OP_16        uint8 = 0x60
)

func createOpPushNBytes(n uint8) Op {
	return func(c Context) error {
		bs := make([]byte, n)
		_, err := c.Read(bs)
		if err != nil {
			return err
		}

		c.Push(bs)
		return nil
	}
}

func createOpPushN(n uint8) Op {
	return func(c Context) error {
		c.Push([]byte{0x00, 0x00, 0x00, n})
		return nil
	}
}

func opFalse(c Context) error {
	c.Push([]byte{0x00, 0x00, 0x0, 0x00})
	return nil
}

func op1Negate(c Context) error {
	c.Push([]byte{0x40, 0x00, 0x00, 0x01})
	return nil
}

func opTrue(c Context) error {
	c.Push([]byte{0x00, 0x00, 0x00, 0x01})
	return nil
}
