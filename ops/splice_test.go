package ops

import (
	"testing"
)

func TestSpliceOps(t *testing.T) {
	runOpTests(t, []opTests{
		{
			"opCat",
			opCat,
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
					config{stack: &stack{[]byte("hello"), []byte("world")}},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: &stack{{0x0a, 0x0b, 0x0c, 0x0d}, {0x01, 0x00, 0x00, 0x00}, {0x02, 0x00, 0x00, 0x00}}},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: &stack{[]byte("helloworld"), {0x04, 0x00, 0x00, 0x00}}},
					opWant{
						&stack{[]byte("hell")},
						&stack{},
						nil,
					},
				},
				{
					"zero",
					config{stack: &stack{[]byte("helloworld"), {0x00, 0x00, 0x00, 0x00}}},
					opWant{
						&stack{[]byte("")},
						&stack{},
						nil,
					},
				},
				{
					"too large",
					config{stack: &stack{[]byte("helloworld"), {0xFF, 0x00, 0x00, 0x00}}},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: &stack{[]byte("helloworld"), {0x05, 0x00, 0x00, 0x00}}},
					opWant{
						&stack{[]byte("world")},
						&stack{},
						nil,
					},
				},
				{
					"zero",
					config{stack: &stack{[]byte("helloworld"), {0x00, 0x00, 0x00, 0x00}}},
					opWant{
						&stack{[]byte("helloworld")},
						&stack{},
						nil,
					},
				},
				{
					"too large",
					config{stack: &stack{[]byte("helloworld"), {0xFF, 0x00, 0x00, 0x00}}},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: &stack{[]byte("abc")}},
					opWant{
						&stack{[]byte("abc"), {0x03, 0x00, 0x00, 0x00}},
						&stack{},
						nil,
					},
				},
			},
		},
	})
}
