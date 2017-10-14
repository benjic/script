package ops

import (
	"reflect"
	"testing"
)

func TestFlowControlOps(t *testing.T) {
	tests := []opTests{
		{
			"opVerify",
			opVerify,
			[]opTest{
				{
					"empty stack",
					opArgs{emptyContext()},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"true",
					opArgs{contextWithStack(stackWithNumbers(t, 1))},
					opWant{
						&stack{},
						&stack{},
						nil,
					},
				},
				{
					"false",
					opArgs{contextWithStack(stackWithNumbers(t, 0))},
					opWant{
						&stack{},
						&stack{},
						ErrVerify,
					},
				},
				{
					"negative zero",
					opArgs{contextWithStack(stackWithNumbers(t, -0))},
					opWant{
						&stack{},
						&stack{},
						ErrVerify,
					},
				},
			},
		},
	}

	for _, opTest := range tests {
		for _, test := range opTest.tests {

			t.Run(opTest.name+" "+test.name, func(t *testing.T) {
				err := opTest.op(test.args.context)
				if err != test.want.err {
					t.Errorf("%s() error = %v, want err %v", opTest.name, err, test.want.err)
				}

				if !reflect.DeepEqual(test.want.stack, test.args.context.stack) {
					t.Errorf("want %v; got %v", test.want.stack, test.args.context.stack)
				}

				if !reflect.DeepEqual(test.want.alt, test.args.context.alt) {
					t.Errorf("want %v; got %v", test.want.alt, test.args.context.alt)
				}
			})

		}
	}
}
