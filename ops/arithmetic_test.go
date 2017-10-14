package ops

import (
	"reflect"
	"testing"
)

func TestArithmeticOps(t *testing.T) {
	tests := []opTests{
		{
			"op1Add",
			op1Add,
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
					"adds 1 to number",
					opArgs{contextWithStack(stackWithNumbers(t, 7))},
					opWant{
						stackWithNumbers(t, 8),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"op1Sub",
			op1Sub,
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
					"adds 1 to number",
					opArgs{contextWithStack(stackWithNumbers(t, 7))},
					opWant{
						stackWithNumbers(t, 6),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"op2Mul",
			op2Mul,
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
					"multiplies number by 2",
					opArgs{contextWithStack(stackWithNumbers(t, 7))},
					opWant{
						stackWithNumbers(t, 14),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"op2Div",
			op2Div,
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
					"divides number by 2",
					opArgs{contextWithStack(stackWithNumbers(t, 7))},
					opWant{
						stackWithNumbers(t, 3),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opNegate",
			opNegate,
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
					"makes negatives positive",
					opArgs{contextWithStack(stackWithNumbers(t, -1))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"makes positives negative",
					opArgs{contextWithStack(stackWithNumbers(t, 1))},
					opWant{
						stackWithNumbers(t, -1),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opAbs",
			opAbs,
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
					"negatives",
					opArgs{contextWithStack(stackWithNumbers(t, -4))},
					opWant{
						stackWithNumbers(t, 4),
						&stack{},
						nil,
					},
				},
				{
					"positives",
					opArgs{contextWithStack(stackWithNumbers(t, 4))},
					opWant{
						stackWithNumbers(t, 4),
						&stack{},
						nil,
					},
				},
				{
					"zero",
					opArgs{contextWithStack(stackWithNumbers(t, 0))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opNot",
			opNot,
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
					"true when less than zero",
					opArgs{contextWithStack(stackWithNumbers(t, -2))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"false when zero",
					opArgs{contextWithStack(stackWithNumbers(t, 0))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"false when greater than zero",
					opArgs{contextWithStack(stackWithNumbers(t, 1))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"op0NotEqual",
			op0NotEqual,
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
					"true when less than zero",
					opArgs{contextWithStack(stackWithNumbers(t, -2))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"false when zero",
					opArgs{contextWithStack(stackWithNumbers(t, 0))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"true when greater than zero",
					opArgs{contextWithStack(stackWithNumbers(t, 1))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opAdd",
			opAdd,
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
					"simple",
					opArgs{contextWithStack(stackWithNumbers(t, 2, 2))},
					opWant{
						stackWithNumbers(t, 4),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opSub",
			opSub,
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
					"simple",
					opArgs{contextWithStack(stackWithNumbers(t, 4, 2))},
					opWant{
						stackWithNumbers(t, 2),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opMul",
			opMul,
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
					"simple",
					opArgs{contextWithStack(stackWithNumbers(t, 4, 2))},
					opWant{
						stackWithNumbers(t, 8),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opDiv",
			opDiv,
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
					"simple",
					opArgs{contextWithStack(stackWithNumbers(t, 12, 3))},
					opWant{
						stackWithNumbers(t, 4),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opMod",
			opMod,
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
					"simple",
					opArgs{contextWithStack(stackWithNumbers(t, 12, 5))},
					opWant{
						stackWithNumbers(t, 2),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opLShift",
			opLShift,
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
					"negative",
					opArgs{contextWithStack(stackWithNumbers(t, -1, 2))},
					opWant{
						stackWithNumbers(t, -4),
						&stack{},
						nil,
					},
				},
				{
					"positive",
					opArgs{contextWithStack(stackWithNumbers(t, 1, 2))},
					opWant{
						stackWithNumbers(t, 4),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opRShift",
			opRShift,
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
					"negative",
					opArgs{contextWithStack(stackWithNumbers(t, -4, 2))},
					opWant{
						stackWithNumbers(t, -1),
						&stack{},
						nil,
					},
				},
				{
					"positive",
					opArgs{contextWithStack(stackWithNumbers(t, 4, 2))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opBoolAnd",
			opBoolAnd,
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
					"both zero",
					opArgs{contextWithStack(stackWithNumbers(t, 0, 0))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"both not zero",
					opArgs{contextWithStack(stackWithNumbers(t, -1, 1))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"a zero",
					opArgs{contextWithStack(stackWithNumbers(t, 0, 1))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"b zero",
					opArgs{contextWithStack(stackWithNumbers(t, -1, 0))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opBoolOr",
			opBoolOr,
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
					"both zero",
					opArgs{contextWithStack(stackWithNumbers(t, 0, 0))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"both not zero",
					opArgs{contextWithStack(stackWithNumbers(t, -1, 1))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"a zero",
					opArgs{contextWithStack(stackWithNumbers(t, 0, 1))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"b zero",
					opArgs{contextWithStack(stackWithNumbers(t, -1, 0))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opNumEqual",
			opNumEqual,
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
					"equal",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 7))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"not equal",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 8))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opNumEqualVerify",
			opNumEqualVerify,
			[]opTest{
			// TODO(benjic): Implement when OP_VERIfY is available
			},
		},
		{
			"opNumNotEqual",
			opNumNotEqual,
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
					"equal",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 7))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"not equal",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 8))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opLessThan",
			opLessThan,
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
					"equal",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 7))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"less than",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 8))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"greater than",
					opArgs{contextWithStack(stackWithNumbers(t, 8, 7))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opGreaterThan",
			opGreaterThan,
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
					"equal",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 7))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"less than",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 8))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"greater than",
					opArgs{contextWithStack(stackWithNumbers(t, 8, 7))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opLessThanOrEqual",
			opLessThanOrEqual,
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
					"equal",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 7))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"less than",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 8))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"greater than",
					opArgs{contextWithStack(stackWithNumbers(t, 8, 7))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opGreaterThanOrEqual",
			opGreaterThanOrEqual,
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
					"equal",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 7))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"less than",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 8))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"greater than",
					opArgs{contextWithStack(stackWithNumbers(t, 8, 7))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opMin",
			opMin,
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
					"left",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 8))},
					opWant{
						stackWithNumbers(t, 7),
						&stack{},
						nil,
					},
				},
				{
					"right",
					opArgs{contextWithStack(stackWithNumbers(t, 8, 7))},
					opWant{
						stackWithNumbers(t, 7),
						&stack{},
						nil,
					},
				},
				{
					"same",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 7))},
					opWant{
						stackWithNumbers(t, 7),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opMax",
			opMax,
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
					"left",
					opArgs{contextWithStack(stackWithNumbers(t, 8, 7))},
					opWant{
						stackWithNumbers(t, 8),
						&stack{},
						nil,
					},
				},
				{
					"right",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 8))},
					opWant{
						stackWithNumbers(t, 8),
						&stack{},
						nil,
					},
				},
				{
					"same",
					opArgs{contextWithStack(stackWithNumbers(t, 8, 8))},
					opWant{
						stackWithNumbers(t, 8),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opWithin",
			opWithin,
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
					"less than",
					opArgs{contextWithStack(stackWithNumbers(t, 6, 7, 9))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"min inclusive",
					opArgs{contextWithStack(stackWithNumbers(t, 7, 7, 9))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"within",
					opArgs{contextWithStack(stackWithNumbers(t, 8, 7, 9))},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"max exclusive",
					opArgs{contextWithStack(stackWithNumbers(t, 9, 7, 9))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"greater than",
					opArgs{contextWithStack(stackWithNumbers(t, 10, 7, 9))},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
			},
		},
	}

	for _, opTests := range tests {
		for _, test := range opTests.tests {

			t.Run(opTests.name+" "+test.name, func(t *testing.T) {
				err := opTests.op(test.args.context)
				if err != test.want.err {
					t.Errorf("%s() error = %v, want err %v", opTests.name, err, test.want.err)
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
