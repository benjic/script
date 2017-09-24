package script

import (
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
			// TODO: Evaluate operation
			opCode := uint8(op[0])
			op, ok := ops.OpRegistry[opCode]
			if !ok {
				return fmt.Errorf("unknown op for code %d", opCode)
			}

			op(context)

		case io.EOF:
			// TODO: Evaluate success after reading program input
			return nil
		default:
			return err
		}
	}
}
