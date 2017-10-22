package ops

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
	"testing"
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
	name   string
	config config
	want   opWant
}

type opTests struct {
	name  string
	op    Op
	tests []opTest
}

func stackWithNumbers(t *testing.T, ns ...int32) *stack {
	s := &stack{}

	for _, n := range ns {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, &n); err != nil {
			t.Errorf("Cannot parse number %d", n)
		}
		s.Push(buf.Bytes())
	}

	return s
}

type config struct {
	stack *stack
	alt   *stack
	buf   []byte
}

func runOpTests(t *testing.T, tests []opTests) {

	defaultStack := func(s *stack) *stack {
		if s == nil {
			s = &stack{}
		}
		return s
	}

	for _, opTest := range tests {
		for _, test := range opTest.tests {

			t.Run(opTest.name+" "+test.name, func(t *testing.T) {

				context := &context{
					defaultStack(test.config.stack),
					defaultStack(test.config.alt),
					bytes.NewBuffer(test.config.buf),
				}

				err := opTest.op(context)
				if err != test.want.err {
					t.Errorf("%s() error = %v, want err %v", opTest.name, err, test.want.err)
				}

				if !reflect.DeepEqual(test.want.stack, context.stack) {
					t.Errorf("want %v; got %v", test.want.stack, context.stack)
				}

				if !reflect.DeepEqual(test.want.alt, context.alt) {
					t.Errorf("want %v; got %v", test.want.alt, context.alt)
				}
			})
		}
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
