package ops

import (
	"reflect"
	"testing"
)

func TestSpliceOps(t *testing.T) {
	tests := []opTests{
		{
			"opCat",
			opCat,
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
					opArgs{contextWithStack(&stack{[]byte("hello"), []byte("world")})},
					opWant{
						&stack{[]byte("helloworld")},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opSubstr",
			opSubstr,
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
					opArgs{contextWithStack(&stack{{0x0a, 0x0b, 0x0c, 0x0d}, {0x01, 0x00, 0x00, 0x00}, {0x02, 0x00, 0x00, 0x00}})},
					opWant{
						&stack{{0xb, 0xc}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opLeft",
			opLeft,
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
					opArgs{contextWithStack(&stack{[]byte("helloworld"), {0x04, 0x00, 0x00, 0x00}})},
					opWant{
						&stack{[]byte("hell")},
						&stack{},
						nil,
					},
				},
				{
					"zero",
					opArgs{contextWithStack(&stack{[]byte("helloworld"), {0x00, 0x00, 0x00, 0x00}})},
					opWant{
						&stack{[]byte("")},
						&stack{},
						nil,
					},
				},
				{
					"too large",
					opArgs{contextWithStack(&stack{[]byte("helloworld"), {0xFF, 0x00, 0x00, 0x00}})},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
			},
		},
		{
			"opRight",
			opRight,
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
					opArgs{contextWithStack(&stack{[]byte("helloworld"), {0x05, 0x00, 0x00, 0x00}})},
					opWant{
						&stack{[]byte("world")},
						&stack{},
						nil,
					},
				},
				{
					"zero",
					opArgs{contextWithStack(&stack{[]byte("helloworld"), {0x00, 0x00, 0x00, 0x00}})},
					opWant{
						&stack{[]byte("helloworld")},
						&stack{},
						nil,
					},
				},
				{
					"too large",
					opArgs{contextWithStack(&stack{[]byte("helloworld"), {0xFF, 0x00, 0x00, 0x00}})},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
			},
		},
		{
			"opSize",
			opSize,
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
					opArgs{contextWithStack(&stack{[]byte("abc")})},
					opWant{
						&stack{[]byte("abc"), {0x03, 0x00, 0x00, 0x00}},
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
