package ops

import "bytes"

const (
	OP_EQUAL uint8 = 0x87
)

func opEqual(c Context) {
	if bytes.Equal(c.Pop(), c.Pop()) {
		c.Push([]byte{0x00, 0x00, 0x00, 0x01})
	} else {
		c.Push([]byte{0x00, 0x00, 0x00, 0x00})
	}
}
