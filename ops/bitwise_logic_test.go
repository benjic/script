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
					opArgs{contextWithStack(&stack{{0x00}, {0x00}})},
					opWant{
						&stack{{0x01, 0x00, 0x00, 0x00}},
						&stack{},
						nil,
					},
				},
				{
					"not equal",
					opArgs{contextWithStack(&stack{{0x00}, {0x01}})},
					opWant{
						&stack{{0x00, 0x00, 0x00, 0x00}},
						&stack{},
						nil,
					},
				},
				{
					"not enough arguments",
					opArgs{contextWithStack(&stack{{0x01}})},
					opWant{
						&stack{{0x01}},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
			},
		},
		{
			"opInvert",
			opInvert,
			[]opTest{
				{
					"empty context",
					opArgs{emptyContext()},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					opArgs{contextWithStack(&stack{{0xF0, 0xF0, 0xF0}})},
					opWant{
						&stack{{0x0F, 0x0F, 0x0F}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opAnd",
			opAnd,
			[]opTest{
				{
					"empty context",
					opArgs{emptyContext()},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					opArgs{contextWithStack(&stack{{0xF0, 0xF0, 0xF0}, {0x00, 0xFF, 0xF0}})},
					opWant{
						&stack{{0x00, 0xF0, 0xF0}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opOr",
			opOr,
			[]opTest{
				{
					"empty context",
					opArgs{emptyContext()},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					opArgs{contextWithStack(&stack{{0xF0, 0xF0, 0xF0}, {0x00, 0xFF, 0xF0}})},
					opWant{
						&stack{{0xF0, 0xFF, 0xF0}},
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opXor",
			opXor,
			[]opTest{
				{
					"empty context",
					opArgs{emptyContext()},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					opArgs{contextWithStack(&stack{{0xF0, 0xF0, 0xF0}, {0x00, 0xFF, 0xF0}})},
					opWant{
						&stack{{0xF0, 0x0F, 0x00}},
						&stack{},
						nil,
					},
				},
				{
					"small x1",
					opArgs{contextWithStack(&stack{{0xF0, 0xF0}, {0x00, 0xFF, 0xF0}})},
					opWant{
						&stack{{0xF0, 0x0F}},
						&stack{},
						nil,
					},
				},
				{
					"small x2",
					opArgs{contextWithStack(&stack{{0xF0, 0xF0, 0xF0}, {0x00, 0xFF}})},
					opWant{
						&stack{{0xF0, 0x0F}},
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
