package ops

import "io"
import "bytes"

// A Stack is a FIFO data structure which orders data.
type stack [][]byte

type context struct {
	stack  *stack
	reader io.Reader
}

func (c *context) Push(val []byte)             { c.stack.Push(val) }
func (c *context) Pop() []byte                 { return c.stack.Pop() }
func (c *context) Read(bs []byte) (int, error) { return c.reader.Read(bs) }

func emptyContext() *context {
	return &context{&stack{}, new(bytes.Buffer)}
}

func contextWithStack(s *stack) *context {
	return &context{s, new(bytes.Buffer)}
}

// Push appends a given value to the top of the stack.
//
// The nil value is ignored.
func (s *stack) Push(v []byte) {
	if v == nil {
		return
	}

	*s = append(*s, v)
}

// Pop yields the top of the stack.
//
// If the stack is empty nil is returned.
func (s *stack) Pop() []byte {
	if len(*s) == 0 {
		return nil
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return v
}
