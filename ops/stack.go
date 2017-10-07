package ops

import (
	"bytes"
	"encoding/binary"
)

// TODO(benjic): Document the operation symbols
const (
	OpToAltStack   uint8 = 0x6b
	OpFromAltStack uint8 = 0x6c
	OpIfDup        uint8 = 0x73
	OpDepth        uint8 = 0x74
	OpDrop         uint8 = 0x75
	OpDup          uint8 = 0x76
	OpNip          uint8 = 0x77
	OpOver         uint8 = 0x78
)

func opToAltStack(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}

	c.PushAlt(c.Pop())

	return nil
}

func opFromAltStack(c Context) error {
	if c.SizeAlt() < 1 {
		return ErrInvalidStackOperation
	}

	c.Push(c.PopAlt())
	return nil
}

func opIfDup(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}

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
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, int32(c.Size()))
	c.Push(buf.Bytes())
	return nil
}

func opDrop(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}
	c.Pop()
	return nil
}

func opDup(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}

	v := c.Pop()
	if v != nil {
		v2 := make([]byte, len(v))
		copy(v2, v)
		c.Push(v2)
	}

	c.Push(v)
	return nil
}

func opNip(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	v := c.Pop()
	c.Pop()
	c.Push(v)
	return nil
}

func opOver(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	v2 := c.Pop()
	v1 := c.Pop()
	c.Push(v1)
	c.Push(v2)

	if v1 != nil {
		v3 := make([]byte, len(v1))
		copy(v3, v1)
		c.Push(v3)
	}

	return nil
}
