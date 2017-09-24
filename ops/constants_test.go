package ops

import (
	"reflect"
	"testing"

	"github.com/benjic/script"
)

func Test_opFalse(t *testing.T) {
	type args struct {
		c context
	}
	tests := []struct {
		name string
		args args
		want *script.Stack
	}{
		{
			"pushes false",
			args{&script.Stack{}},
			&script.Stack{[]byte{0x00, 0x00, 0x0, 0x00}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opFalse(tt.args.c)

			if !reflect.DeepEqual(tt.args.c, tt.want) {
				t.Errorf("Want %+v; Got %+v", tt.args.c, tt.want)
			}
		})
	}
}

func Test_op1Negate(t *testing.T) {
	type args struct {
		c context
	}
	tests := []struct {
		name string
		args args
		want *script.Stack
	}{
		{
			"pushes negative 1",
			args{&script.Stack{}},
			&script.Stack{[]byte{0x40, 0x00, 0x00, 0x01}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op1Negate(tt.args.c)
		})
	}
}

func Test_opTrue(t *testing.T) {
	type args struct {
		c context
	}
	tests := []struct {
		name string
		args args
		want *script.Stack
	}{
		{
			"pushes 1",
			args{&script.Stack{}},
			&script.Stack{[]byte{0x00, 0x00, 0x00, 0x01}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opTrue(tt.args.c)
		})
	}
}
