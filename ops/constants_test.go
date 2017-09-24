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
		name string
		args args
		want *stack
	}{
		{
			"pushes false",
			args{emptyContext()},
			&stack{[]byte{0x00, 0x00, 0x0, 0x00}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opFalse(tt.args.c)

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
		name string
		args args
		want *stack
	}{
		{
			"pushes negative 1",
			args{emptyContext()},
			&stack{[]byte{0x40, 0x00, 0x00, 0x01}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op1Negate(tt.args.c)

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
		name string
		args args
		want *stack
	}{
		{
			"pushes 1",
			args{emptyContext()},
			&stack{[]byte{0x00, 0x00, 0x00, 0x01}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opTrue(tt.args.c)

			if !reflect.DeepEqual(tt.args.c.stack, tt.want) {
				t.Errorf("Want %+v; Got %+v", tt.args.c, tt.want)
			}
		})
	}
}
