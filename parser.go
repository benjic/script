package script

import (
	"bytes"
	"fmt"
	"io"

	"github.com/benjic/script/ops"
)

type evaluationContext struct {
	*stack
	reader io.Reader
}

func (c *evaluationContext) Push(value []byte)           { c.stack.Push(value) }
func (c *evaluationContext) Pop() []byte                 { return c.stack.Pop() }
func (c *evaluationContext) Read(bs []byte) (int, error) { return c.reader.Read(bs) }

// Evaluate executes the binary representation of a script.
func Evaluate(input io.Reader) error {
	op := make([]byte, 1)
	context := &evaluationContext{&stack{}, input}

	for {
		_, err := input.Read(op)
		switch err {
		case nil:
			opCode := uint8(op[0])
			op, ok := ops.OpRegistry[opCode]
			if !ok {
				return fmt.Errorf("unknown op for code %d", opCode)
			}

			err := op(context)
			if err != nil {
				return fmt.Errorf("op (%d) failed: %s", opCode, err)
			}

		case io.EOF:
			if bytes.Equal(context.stack.Pop(), []byte{0x00, 0x00, 0x00, 0x00}) {
				return fmt.Errorf("top value of stack is false")
			}
			return nil
		default:
			return err
		}
	}
}
