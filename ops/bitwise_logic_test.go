package ops

import (
	"testing"
)

func TestLogicOps(t *testing.T) {
	runOpTests(t, []opTests{
		{
			"opEqual",
			opEqual,
			[]opTest{
				{
					"equal",
					config{stack: &stack{{0x00}, {0x00}}},
					opWant{
						&stack{{0x01, 0x00, 0x00, 0x00}},
						&stack{},
						nil,
					},
				},
				{
					"not equal",
					config{stack: &stack{{0x00}, {0x01}}},
					opWant{
						&stack{{0x00, 0x00, 0x00, 0x00}},
						&stack{},
						nil,
					},
				},
				{
					"not enough arguments",
					config{stack: &stack{{0x01}}},
					opWant{
						&stack{{0x01}},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
			},
		},
		{
			"opEqualVerify",
			opEqualVerify,
			[]opTest{
				{
					"equal",
					config{stack: &stack{{0x00}, {0x00}}},
					opWant{
						&stack{{0x01, 0x00, 0x00, 0x00}},
						&stack{},
						nil,
					},
				},
				{
					"not equal",
					config{stack: &stack{{0x00}, {0x01}}},
					opWant{
						&stack{{0x00, 0x00, 0x00, 0x00}},
						&stack{},
						ErrEqualVerify,
					},
				},
				{
					"not enough arguments",
					config{stack: &stack{{0x01}}},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: &stack{{0xF0, 0xF0, 0xF0}}},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: &stack{{0xF0, 0xF0, 0xF0}, {0x00, 0xFF, 0xF0}}},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: &stack{{0xF0, 0xF0, 0xF0}, {0x00, 0xFF, 0xF0}}},
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
					config{},
					opWant{
						&stack{},
						&stack{},
						ErrInvalidStackOperation,
					},
				},
				{
					"simple",
					config{stack: &stack{{0xF0, 0xF0, 0xF0}, {0x00, 0xFF, 0xF0}}},
					opWant{
						&stack{{0xF0, 0x0F, 0x00}},
						&stack{},
						nil,
					},
				},
				{
					"small x1",
					config{stack: &stack{{0xF0, 0xF0}, {0x00, 0xFF, 0xF0}}},
					opWant{
						&stack{{0xF0, 0x0F}},
						&stack{},
						nil,
					},
				},
				{
					"small x2",
					config{stack: &stack{{0xF0, 0xF0, 0xF0}, {0x00, 0xFF}}},
					opWant{
						&stack{{0xF0, 0x0F}},
						&stack{},
						nil,
					},
				},
			},
		},
	},
	)
}
