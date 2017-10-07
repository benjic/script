package ops

import (
	"reflect"
	"testing"
)

func TestStackOps(t *testing.T) {
	type args struct {
		context *context
	}
	type want struct {
		stack *stack
		alt   *stack
		err   error
	}
	type test struct {
		name string
		args args
		want want
	}

	tests := []struct {
		name  string
		op    Op
		tests []test
	}{
		{
			"opToAltStack",
			opToAltStack,
			[]test{
				{
					"empty stack",
					args{contextWithStackAndAlt(&stack{}, &stack{[]byte{0x00}})},
					want{
						&stack{},
						&stack{[]byte{0x00}},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					args{contextWithStackAndAlt(&stack{[]byte{0x00}, []byte{0xff}}, &stack{[]byte{0x00}})},
					want{
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
			[]test{
				{
					"simple",
					args{contextWithStackAndAlt(&stack{[]byte{0x00}}, &stack{[]byte{0x00}, []byte{0xff}})},
					want{
						&stack{[]byte{0x00}, []byte{0xff}},
						&stack{[]byte{0x00}},
						nil,
					},
				},
				{
					"empty stack",
					args{contextWithStackAndAlt(&stack{[]byte{0x00}}, &stack{})},
					want{
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
			[]test{
				{
					"non zero top stack",
					args{contextWithStack(&stack{[]byte{0x00, 0x00, 0x00, 0x01}})},
					want{
						&stack{[]byte{0x00, 0x00, 0x00, 0x01}, []byte{0x00, 0x00, 0x00, 0x01}},
						&stack{},
						nil,
					},
				},
				{
					"zero top stack",
					args{contextWithStack(&stack{[]byte{0x00, 0x00, 0x00, 0x00}})},
					want{
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
			[]test{
				{
					"empty stack",
					args{contextWithStack(&stack{})},
					want{
						&stack{[]byte{0x00, 0x00, 0x00, 0x00}},
						&stack{},
						nil,
					},
				},
				{
					"depth 1",
					args{contextWithStack(&stack{[]byte{0x00}})},
					want{
						&stack{[]byte{0x00}, []byte{0x01, 0x00, 0x00, 0x00}},
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
			[]test{
				{
					"empty stack",
					args{contextWithStack(&stack{})},
					want{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"single value",
					args{contextWithStack(&stack{[]byte{0x00}})},
					want{
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
			[]test{
				{
					"simple",
					args{contextWithStack(&stack{{0x1}})},
					want{
						&stack{{0x1}, {0x1}},
						&stack{},
						nil,
					},
				},
				{
					"empty stack",
					args{contextWithStack(&stack{})},
					want{
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
			[]test{
				{
					"simple",
					args{contextWithStack(&stack{{0x1}, {0x2}})},
					want{
						&stack{{0x1}, {0x2}, {0x1}},
						&stack{},
						nil,
					},
				},
				{
					"too small stack",
					args{contextWithStack(&stack{{0x1}})},
					want{
						&stack{{0x1}},
						&stack{},
						ErrInvalidStackOperation,
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
