package ops

import (
	"bytes"
)

const (
	OP_TOALTSTACK   uint8 = 0x6b
	OP_FROMALTSTACK uint8 = 0x6c
	OP_IFDUP        uint8 = 0x73
)

func opToAltStack(c Context) error {
	c.PushAlt(c.Pop())
	return nil
}

func opFromAltStack(c Context) error {
	c.Push(c.PopAlt())
	return nil
}

func opIfDup(c Context) error {
	v := c.Pop()

	if !bytes.Equal(v, []byte{0x00, 0x00, 0x00, 0x00}) {
		v2 := make([]byte, len(v))
		copy(v2, v)
		c.Push(v2)
	}

	c.Push(v)

	return nil
}
