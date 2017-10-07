package ops

import (
	"reflect"
	"testing"
)

func TestLogicOps(t *testing.T) {
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
			"opEqual",
			opEqual,
			[]test{
				{
					"equal",
					args{contextWithStack(&stack{[]byte{0x00}, []byte{0x00}})},
					want{
						&stack{[]byte{0x00, 0x00, 0x00, 0x01}},
						&stack{},
						nil,
					},
				},
				{
					"not equal",
					args{contextWithStack(&stack{[]byte{0x00}, []byte{0x01}})},
					want{
						&stack{[]byte{0x00, 0x00, 0x00, 0x00}},
						&stack{},
						nil,
					},
				},
				{
					"not enough arguments",
					args{contextWithStack(&stack{[]byte{0x01}})},
					want{
						&stack{[]byte{0x00, 0x00, 0x00, 0x00}},
						&stack{},
						nil,
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
