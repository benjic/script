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
