package ops

import (
	"reflect"
	"testing"
)

func TestConstantsOps(t *testing.T) {

	tests := []opTests{
		{
			"opFalse",
			opFalse,
			[]opTest{
				{
					"pushes false",
					opArgs{emptyContext()},
					opWant{
						&stack{[]byte{0x00, 0x00, 0x0, 0x00}},
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
					opArgs{emptyContext()},
					opWant{
						&stack{[]byte{0x40, 0x00, 0x00, 0x01}},
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
					opArgs{emptyContext()},
					opWant{
						&stack{[]byte{0x00, 0x00, 0x00, 0x01}},
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

func Test_createOpPushNBytes(t *testing.T) {
	type args struct {
		context *context
		n       uint8
	}
	tests := []struct {
		name    string
		args    args
		want    *stack
		wantErr bool
	}{
		{
			"zero",
			args{contextWithData([]byte{0x00}), 0},
			&stack{},
			false,
		},
		{
			"correct number of bytes available",
			args{contextWithData([]byte{0x00}), 1},
			&stack{[]byte{0x00}},
			false,
		},
		{
			"incorrect number of bytes available",
			args{contextWithData([]byte{0x01}), 2},
			&stack{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := createOpPushNBytes(tt.args.n)

			err := op(tt.args.context)

			if (err != nil) != tt.wantErr {
				t.Errorf("op() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.args.context.stack, tt.want) {
				t.Errorf("Want %+v; Got %+v", tt.want, tt.args.context.stack)
			}
		})
	}
}

func Test_createOpPushN(t *testing.T) {
	type args struct {
		n uint8
	}
	tests := []struct {
		name    string
		args    args
		want    *stack
		wantErr bool
	}{
		{
			"simple",
			args{0xff},
			&stack{[]byte{0x00, 0x00, 0x00, 0xff}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := createOpPushN(tt.args.n)
			context := emptyContext()

			err := op(context)

			if (err != nil) != tt.wantErr {
				t.Errorf("op() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(context.stack, tt.want) {
				t.Errorf("Want %+v; Got %+v", tt.want, context.stack)
			}
		})
	}
}
