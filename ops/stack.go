package ops

import (
	"bytes"
	"encoding/binary"
)

const (
	OP_TOALTSTACK   uint8 = 0x6b
	OP_FROMALTSTACK uint8 = 0x6c
	OP_IFDUP        uint8 = 0x73
	OP_DEPTH        uint8 = 0x74
	OP_DROP         uint8 = 0x75
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

func opDepth(c Context) error {
	buf := []byte{0x00, 0x00, 0x00, 0x00}
	var d uint32
	for {
		v := c.Pop()
		if v == nil {
			break
		}
		c.PushAlt(v)
		d++
	}

	for i := uint32(0); i < d; i++ {
		c.Push(c.PopAlt())
	}

	binary.BigEndian.PutUint32(buf, d)
	c.Push(buf)
	return nil
}

func opDrop(c Context) error {
	c.Pop()
	return nil
}
