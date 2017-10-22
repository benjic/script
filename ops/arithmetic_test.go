package ops

import (
	"testing"
)

func TestArithmeticOps(t *testing.T) {
	runOpTests(t, []opTests{
		{
			"op1Add",
			op1Add,
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
					"adds 1 to number",
					config{stack: stackWithNumbers(t, 7)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"adds 1 to number",
					config{stack: stackWithNumbers(t, 7)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"multiplies number by 2",
					config{
						stack: stackWithNumbers(t, 7),
					},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"divides number by 2",
					config{
						stack: stackWithNumbers(t, 7),
					},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"makes negatives positive",
					config{stack: stackWithNumbers(t, -1)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"makes positives negative",
					config{stack: stackWithNumbers(t, 1)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"negatives",
					config{stack: stackWithNumbers(t, -4)},
					opWant{
						stackWithNumbers(t, 4),
						&stack{},
						nil,
					},
				},
				{
					"positives",
					config{stack: stackWithNumbers(t, 4)},
					opWant{
						stackWithNumbers(t, 4),
						&stack{},
						nil,
					},
				},
				{
					"zero",
					config{stack: stackWithNumbers(t, 0)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"true when less than zero",
					config{stack: stackWithNumbers(t, -2)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"false when zero",
					config{stack: stackWithNumbers(t, 0)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"false when greater than zero",
					config{stack: stackWithNumbers(t, 1)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"true when less than zero",
					config{stack: stackWithNumbers(t, -2)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"false when zero",
					config{stack: stackWithNumbers(t, 0)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"true when greater than zero",
					config{stack: stackWithNumbers(t, 1)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: stackWithNumbers(t, 2, 2)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: stackWithNumbers(t, 4, 2)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: stackWithNumbers(t, 4, 2)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: stackWithNumbers(t, 12, 3)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: stackWithNumbers(t, 12, 5)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"negative",
					config{stack: stackWithNumbers(t, -1, 2)},
					opWant{
						stackWithNumbers(t, -4),
						&stack{},
						nil,
					},
				},
				{
					"positive",
					config{stack: stackWithNumbers(t, 1, 2)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"negative",
					config{stack: stackWithNumbers(t, -4, 2)},
					opWant{
						stackWithNumbers(t, -1),
						&stack{},
						nil,
					},
				},
				{
					"positive",
					config{stack: stackWithNumbers(t, 4, 2)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"both zero",
					config{stack: stackWithNumbers(t, 0, 0)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"both not zero",
					config{stack: stackWithNumbers(t, -1, 1)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"a zero",
					config{stack: stackWithNumbers(t, 0, 1)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"b zero",
					config{stack: stackWithNumbers(t, -1, 0)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"both zero",
					config{stack: stackWithNumbers(t, 0, 0)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"both not zero",
					config{stack: stackWithNumbers(t, -1, 1)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"a zero",
					config{stack: stackWithNumbers(t, 0, 1)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"b zero",
					config{stack: stackWithNumbers(t, -1, 0)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"equal",
					config{stack: stackWithNumbers(t, 7, 7)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"not equal",
					config{stack: stackWithNumbers(t, 7, 8)},
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
					"equal",
					config{stack: stackWithNumbers(t, 7, 7)},
					opWant{
						&stack{},
						&stack{},
						nil,
					},
				},
				{
					"not equal",
					config{stack: stackWithNumbers(t, 7, 8)},
					opWant{
						&stack{},
						&stack{},
						ErrNumEqualVerify,
					},
				},
			},
		},
		{
			"opNumNotEqual",
			opNumNotEqual,
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
					"equal",
					config{stack: stackWithNumbers(t, 7, 7)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"not equal",
					config{stack: stackWithNumbers(t, 7, 8)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"equal",
					config{stack: stackWithNumbers(t, 7, 7)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"less than",
					config{stack: stackWithNumbers(t, 7, 8)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"greater than",
					config{stack: stackWithNumbers(t, 8, 7)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"equal",
					config{stack: stackWithNumbers(t, 7, 7)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"less than",
					config{stack: stackWithNumbers(t, 7, 8)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"greater than",
					config{stack: stackWithNumbers(t, 8, 7)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"equal",
					config{stack: stackWithNumbers(t, 7, 7)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"less than",
					config{stack: stackWithNumbers(t, 7, 8)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"greater than",
					config{stack: stackWithNumbers(t, 8, 7)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"equal",
					config{stack: stackWithNumbers(t, 7, 7)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"less than",
					config{stack: stackWithNumbers(t, 7, 8)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"greater than",
					config{stack: stackWithNumbers(t, 8, 7)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"left",
					config{stack: stackWithNumbers(t, 7, 8)},
					opWant{
						stackWithNumbers(t, 7),
						&stack{},
						nil,
					},
				},
				{
					"right",
					config{stack: stackWithNumbers(t, 8, 7)},
					opWant{
						stackWithNumbers(t, 7),
						&stack{},
						nil,
					},
				},
				{
					"same",
					config{stack: stackWithNumbers(t, 7, 7)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"left",
					config{stack: stackWithNumbers(t, 8, 7)},
					opWant{
						stackWithNumbers(t, 8),
						&stack{},
						nil,
					},
				},
				{
					"right",
					config{stack: stackWithNumbers(t, 7, 8)},
					opWant{
						stackWithNumbers(t, 8),
						&stack{},
						nil,
					},
				},
				{
					"same",
					config{stack: stackWithNumbers(t, 8, 8)},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"less than",
					config{stack: stackWithNumbers(t, 6, 7, 9)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"min inclusive",
					config{stack: stackWithNumbers(t, 7, 7, 9)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"within",
					config{stack: stackWithNumbers(t, 8, 7, 9)},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
				{
					"max exclusive",
					config{stack: stackWithNumbers(t, 9, 7, 9)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"greater than",
					config{stack: stackWithNumbers(t, 10, 7, 9)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
			},
		},
	},
	)
}
