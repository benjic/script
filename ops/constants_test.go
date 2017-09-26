package ops

import (
	"reflect"
	"testing"
)

func Test_opFalse(t *testing.T) {
	type args struct {
		c *context
	}
	tests := []struct {
		name    string
		args    args
		want    *stack
		wantErr bool
	}{
		{
			"pushes false",
			args{emptyContext()},
			&stack{[]byte{0x00, 0x00, 0x0, 0x00}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opFalse(tt.args.c)

			if (err != nil) != tt.wantErr {
				t.Errorf("opFalse() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.args.c.stack, tt.want) {
				t.Errorf("Want %+v; Got %+v", tt.args.c, tt.want)
			}
		})
	}
}

func Test_op1Negate(t *testing.T) {
	type args struct {
		c *context
	}
	tests := []struct {
		name    string
		args    args
		want    *stack
		wantErr bool
	}{
		{
			"pushes negative 1",
			args{emptyContext()},
			&stack{[]byte{0x40, 0x00, 0x00, 0x01}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := op1Negate(tt.args.c)

			if (err != nil) != tt.wantErr {
				t.Errorf("op1Negate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.args.c.stack, tt.want) {
				t.Errorf("Want %+v; Got %+v", tt.args.c, tt.want)
			}
		})
	}
}

func Test_opTrue(t *testing.T) {
	type args struct {
		c *context
	}
	tests := []struct {
		name    string
		args    args
		want    *stack
		wantErr bool
	}{
		{
			"pushes 1",
			args{emptyContext()},
			&stack{[]byte{0x00, 0x00, 0x00, 0x01}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opTrue(tt.args.c)

			if (err != nil) != tt.wantErr {
				t.Errorf("opTrue() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.args.c.stack, tt.want) {
				t.Errorf("Want %+v; Got %+v", tt.args.c, tt.want)
			}
		})
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
