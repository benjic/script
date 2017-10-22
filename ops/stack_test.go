package ops

import (
	"testing"
)

func TestStackOps(t *testing.T) {
	runOpTests(t, []opTests{
		{
			"opToAltStack",
			opToAltStack,
			[]opTest{
				{
					"empty stack",
					config{alt: &stack{[]byte{0x00}}},
					opWant{
						&stack{},
						&stack{[]byte{0x00}},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{
						alt:   &stack{{0x00}},
						stack: &stack{{0x00}, {0xff}},
					},
					opWant{
						&stack{[]byte{0x00}},
						&stack{[]byte{0x00}, []byte{0xff}},
						nil,
					},
				},
			},
		},
		{
			"opFromAltStack",
			opFromAltStack,
			[]opTest{
				{
					"simple",
					config{
						alt:   &stack{{0xff}},
						stack: &stack{{0x00}},
					},
					opWant{
						&stack{{0x00}, {0xff}},
						&stack{},
						nil,
					},
				},
				{
					"empty stack",
					config{
						stack: &stack{{0x00}},
					},
					opWant{
						&stack{{0x00}},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
			},
		},
		{
			"opIfDup",
			opIfDup,
			[]opTest{
				{
					"non zero top stack",
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"non zero top stack",
					config{stack: &stack{[]byte{0x00, 0x00, 0x00, 0x01}}},
					opWant{
						&stack{[]byte{0x00, 0x00, 0x00, 0x01}, []byte{0x00, 0x00, 0x00, 0x01}},
						&stack{},
						nil,
					},
				},
				{
					"zero top stack",
					config{stack: &stack{[]byte{0x00, 0x00, 0x00, 0x00}}},
					opWant{
						&stack{[]byte{0x00, 0x00, 0x00, 0x00}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opDepth",
			opDepth,
			[]opTest{
				{
					"empty stack",
					config{stack: &stack{}},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
				{
					"depth 1",
					config{stack: stackWithNumbers(t, 0)},
					opWant{
						stackWithNumbers(t, 0, 1),
						&stack{},
						nil,
					},
				},
				//
			},
		},
		{
			"opDrop",
			opDrop,
			[]opTest{
				{
					"empty stack",
					config{stack: &stack{}},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"single value",
					config{stack: &stack{[]byte{0x00}}},
					opWant{
						&stack{},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opDup",
			opDup,
			[]opTest{
				{
					"simple",
					config{stack: &stack{{0x1}}},
					opWant{
						&stack{{0x1}, {0x1}},
						&stack{},
						nil,
					},
				},
				{
					"empty stack",
					config{stack: &stack{}},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
			},
		},
		{
			"opOver",
			opOver,
			[]opTest{
				{
					"simple",
					config{stack: &stack{{0x1}, {0x2}}},
					opWant{
						&stack{{0x1}, {0x2}, {0x1}},
						&stack{},
						nil,
					},
				},
				{
					"too small stack",
					config{stack: &stack{{0x1}}},
					opWant{
						&stack{{0x1}},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
			},
		},
		{
			"opNip",
			opNip,
			[]opTest{
				{
					"invalid size",
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: &stack{{0x1}, {0x2}}},
					opWant{
						&stack{{0x2}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opPick",
			opPick,
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
					"not enough elements",
					config{stack: stackWithNumbers(t, 0, 2)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"negative n",
					config{stack: stackWithNumbers(t, 0, -1)},
					opWant{
						&stack{{0x00, 0x00, 0x00, 0x00}},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: stackWithNumbers(t, 1, 2, 1)},
					opWant{
						stackWithNumbers(t, 1, 2, 1),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opRoll",
			opRoll,
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
					"not enough elements",
					config{stack: stackWithNumbers(t, 0, 2)},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"negative n",
					config{stack: stackWithNumbers(t, 0, -1)},
					opWant{
						&stack{{0x00, 0x00, 0x00, 0x00}},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: stackWithNumbers(t, 1, 2, 1)},
					opWant{
						stackWithNumbers(t, 2, 1),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opRot",
			opRot,
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
					config{stack: &stack{{0x3}, {0x2}, {0x1}}},
					opWant{
						&stack{{0x2}, {0x3}, {0x1}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opSwap",
			opSwap,
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
					config{stack: &stack{{0x2}, {0x1}, {0x0}}},
					opWant{
						&stack{{0x2}, {0x0}, {0x1}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opTuck",
			opTuck,
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
					config{stack: &stack{{0x1}, {0x2}}},
					opWant{
						&stack{{0x2}, {0x1}, {0x2}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"op2Drop",
			op2Drop,
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
					config{stack: &stack{{0x3}, {0x2}, {0x1}}},
					opWant{
						&stack{{0x3}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"op2Dup",
			op2Dup,
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
					config{stack: &stack{{0x1}, {0x2}}},
					opWant{
						&stack{{0x1}, {0x2}, {0x1}, {0x2}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"op3Dup",
			op3Dup,
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
					config{stack: &stack{{0x1}, {0x2}, {0x3}}},
					opWant{
						&stack{{0x1}, {0x2}, {0x3}, {0x1}, {0x2}, {0x3}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"op2Over",
			op2Over,
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
					config{stack: &stack{{0x1}, {0x2}, {0x3}, {0x4}}},
					opWant{
						&stack{{0x1}, {0x2}, {0x3}, {0x4}, {0x1}, {0x2}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"op2Rot",
			op2Rot,
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
					config{stack: &stack{{0x6}, {0x5}, {0x4}, {0x3}, {0x2}, {0x1}}},
					opWant{
						&stack{{0x2}, {0x1}, {0x6}, {0x5}, {0x4}, {0x3}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"op2Swap",
			op2Swap,
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
					config{stack: &stack{{0x1}, {0x2}, {0x3}, {0x4}}},
					opWant{
						&stack{{0x2}, {0x1}, {0x4}, {0x3}},
						&stack{},
						nil,
					},
				},
			},
		},
	})
}
