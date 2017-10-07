package ops

import "bytes"

// TODO(benjic): Document the operation symbols
const (
	OpEqual uint8 = 0x87
)

func opEqual(c Context) error {
	if bytes.Equal(c.Pop(), c.Pop()) {
		c.Push([]byte{0x00, 0x00, 0x00, 0x01})
	} else {
		c.Push([]byte{0x00, 0x00, 0x00, 0x00})
	}

	return nil
}
