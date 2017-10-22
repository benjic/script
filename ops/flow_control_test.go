package ops

import (
	"testing"
)

func TestFlowControlOps(t *testing.T) {
	runOpTests(t, []opTests{
		{
			"opVerify",
			opVerify,
			[]opTest{
				{
					"empty stack",
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"true",
					config{stack: stackWithNumbers(t, 1)},
					opWant{
						&stack{},
						&stack{},
						nil,
					},
				},
				{
					"false",
					config{stack: stackWithNumbers(t, 0)},
					opWant{
						&stack{},
						&stack{},
						ErrVerify,
					},
				},
				{
					"negative zero",
					config{stack: stackWithNumbers(t, -0)},
					opWant{
						&stack{},
						&stack{},
						ErrVerify,
					},
				},
			},
		},
	})
}
