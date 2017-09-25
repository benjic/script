package ops

const (
	OP_TOALTSTACK   uint8 = 0x6b
	OP_FROMALTSTACK uint8 = 0x6c
)

func opToAltStack(c Context) error {
	c.PushAlt(c.Pop())
	return nil
}

func opFromAltStack(c Context) error {
	c.Push(c.PopAlt())
	return nil
}
