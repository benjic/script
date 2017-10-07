package ops

import (
	"bytes"
	"io"
)

// A Stack is a FIFO data structure which orders data.
type stack [][]byte

type context struct {
	*stack
	alt *stack
	io.Reader
}

func (c *context) Pop() []byte                 { return c.stack.Pop() }
func (c *context) PopAlt() []byte              { return c.alt.Pop() }
func (c *context) Push(val []byte)             { c.stack.Push(val) }
func (c *context) PushAlt(val []byte)          { c.alt.Push(val) }
func (c *context) Read(bs []byte) (int, error) { return c.Reader.Read(bs) }
func (c *context) Size() int                   { return c.stack.Size() }
func (c *context) SizeAlt() int                { return c.alt.Size() }

type opArgs struct {
	context *context
}

type opWant struct {
	stack *stack
	alt   *stack
	err   error
}

type opTest struct {
	name string
	args opArgs
	want opWant
}

type opTests struct {
	name  string
	op    Op
	tests []opTest
}

func emptyContext() *context {
	return &context{&stack{}, &stack{}, new(bytes.Buffer)}
}

func contextWithStack(s *stack) *context {
	return &context{s, &stack{}, new(bytes.Buffer)}
}

func contextWithStackAndAlt(s *stack, alt *stack) *context {
	return &context{s, alt, new(bytes.Buffer)}
}

func contextWithData(buf []byte) *context {
	return &context{
		&stack{},
		&stack{},
		bytes.NewBuffer(buf),
	}
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

func (s *stack) Size() int {
	return len(*s)
}
