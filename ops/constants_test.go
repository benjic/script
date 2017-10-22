package ops

import (
	"io"
	"testing"
)

func TestConstantsOps(t *testing.T) {

	runOpTests(t, []opTests{
		{
			"opFalse",
			opFalse,
			[]opTest{
				{
					"pushes false",
					config{},
					opWant{
						stackWithNumbers(t, 0),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"op1Negate",
			op1Negate,
			[]opTest{
				{
					"pushes negative 1",
					config{},
					opWant{
						stackWithNumbers(t, -1),
						&stack{},
						nil,
					},
				},
			},
		},
		{
			"opTrue",
			opTrue,
			[]opTest{
				//
				{
					"pushes 1",
					config{},
					opWant{
						stackWithNumbers(t, 1),
						&stack{},
						nil,
					},
				},
			},
		},
	},
	)

}

func Test_createOpPushNBytes(t *testing.T) {
	runOpTests(t, []opTests{
		{
			"createOpPushNBytes(0)",
			createOpPushNBytes(0),
			[]opTest{
				{
					"zero",
					config{buf: []byte{0x00}},
					opWant{
						&stack{},
						&stack{},
						nil,
					},
				},
			},
		},
		{

			"createOpPushNBytes(2)",
			createOpPushNBytes(2),
			[]opTest{
				{
					"correct number of bytes",
					config{buf: []byte{0x00, 0x00}},
					opWant{
						&stack{{0x00, 0x00}},
						&stack{},
						nil,
					},
				},
				{
					"insufficient  number of bytes",
					config{buf: []byte{0x00}},
					opWant{
						&stack{},
						&stack{},
						ErrInsufficientNumberOfBytes,
					},
				},
				{
					"empty reader",
					config{buf: []byte{}},
					opWant{
						&stack{},
						&stack{},
						io.EOF,
					},
				},
			},
		},
	})
}

func Test_createOpPushN(t *testing.T) {
	runOpTests(t, []opTests{
		{
			"createOpPushN",
			createOpPushN(0xff),
			[]opTest{
				{
					"simple",
					config{},
					opWant{
						stackWithNumbers(t, 255),
						&stack{},
						nil,
					},
				},
			},
		},
	})
}
