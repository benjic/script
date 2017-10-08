package ops

import (
	"reflect"
	"testing"
)

func TestStackOps(t *testing.T) {
	tests := []opTests{
		{
			"opToAltStack",
			opToAltStack,
			[]opTest{
				{
					"empty stack",
					opArgs{contextWithStackAndAlt(&stack{}, &stack{[]byte{0x00}})},
					opWant{
						&stack{},
						&stack{[]byte{0x00}},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					opArgs{contextWithStackAndAlt(&stack{[]byte{0x00}, []byte{0xff}}, &stack{[]byte{0x00}})},
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
					opArgs{contextWithStackAndAlt(&stack{[]byte{0x00}}, &stack{[]byte{0x00}, []byte{0xff}})},
					opWant{
						&stack{[]byte{0x00}, []byte{0xff}},
						&stack{[]byte{0x00}},
						nil,
					},
				},
				{
					"empty stack",
					opArgs{contextWithStackAndAlt(&stack{[]byte{0x00}}, &stack{})},
					opWant{
						&stack{[]byte{0x00}},
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
					opArgs{emptyContext()},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"non zero top stack",
					opArgs{contextWithStack(&stack{[]byte{0x00, 0x00, 0x00, 0x01}})},
					opWant{
						&stack{[]byte{0x00, 0x00, 0x00, 0x01}, []byte{0x00, 0x00, 0x00, 0x01}},
						&stack{},
						nil,
					},
				},
				{
					"zero top stack",
					opArgs{contextWithStack(&stack{[]byte{0x00, 0x00, 0x00, 0x00}})},
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
					opArgs{contextWithStack(&stack{})},
					opWant{
						stackWithNumbers(t, []int32{0}),
						&stack{},
						nil,
					},
				},
				{
					"depth 1",
					opArgs{contextWithStack(stackWithNumbers(t, []int32{0}))},
					opWant{
						stackWithNumbers(t, []int32{0, 1}),
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
					opArgs{contextWithStack(&stack{})},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"single value",
					opArgs{contextWithStack(&stack{[]byte{0x00}})},
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
					opArgs{contextWithStack(&stack{{0x1}})},
					opWant{
						&stack{{0x1}, {0x1}},
						&stack{},
						nil,
					},
				},
				{
					"empty stack",
					opArgs{contextWithStack(&stack{})},
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
					opArgs{contextWithStack(&stack{{0x1}, {0x2}})},
					opWant{
						&stack{{0x1}, {0x2}, {0x1}},
						&stack{},
						nil,
					},
				},
				{
					"too small stack",
					opArgs{contextWithStack(&stack{{0x1}})},
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
					opArgs{emptyContext()},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					opArgs{contextWithStack(&stack{{0x1}, {0x2}})},
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
					opArgs{emptyContext()},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"not enough elements",
					opArgs{contextWithStack(stackWithNumbers(t, []int32{0, 2}))},
					opWant{
						stackWithNumbers(t, []int32{0}),
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"negative n",
					opArgs{contextWithStack(stackWithNumbers(t, []int32{0, -1}))},
					opWant{
						&stack{{0x00, 0x00, 0x00, 0x00}},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					opArgs{contextWithStack(stackWithNumbers(t, []int32{1, 2, 1}))},
					opWant{
						stackWithNumbers(t, []int32{1, 2, 1}),
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
