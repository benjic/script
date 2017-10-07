package ops

import (
	"reflect"
	"testing"
)

func TestLogicOps(t *testing.T) {
	tests := []opTests{
		{
			"opEqual",
			opEqual,
			[]opTest{
				{
					"equal",
					opArgs{contextWithStack(&stack{[]byte{0x00}, []byte{0x00}})},
					opWant{
						&stack{[]byte{0x00, 0x00, 0x00, 0x01}},
						&stack{},
						nil,
					},
				},
				{
					"not equal",
					opArgs{contextWithStack(&stack{[]byte{0x00}, []byte{0x01}})},
					opWant{
						&stack{[]byte{0x00, 0x00, 0x00, 0x00}},
						&stack{},
						nil,
					},
				},
				{
					"not enough arguments",
					opArgs{contextWithStack(&stack{[]byte{0x01}})},
					opWant{
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
