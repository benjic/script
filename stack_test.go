package script

import (
	"reflect"
	"testing"
)

func Test_stack_Push(t *testing.T) {
	type args struct {
		v []byte
	}
	tests := []struct {
		name string
		s    *stack
		args args
		want *stack
	}{
		{
			"to empty stack",
			&stack{},
			args{[]byte{0x0}},
			&stack{[]byte{0x0}},
		},
		{
			"add as last element",
			&stack{[]byte{0x0}},
			args{[]byte{0x1}},
			&stack{[]byte{0x0}, []byte{0x1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.v)

			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("stack.Push() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func Test_stack_Pop(t *testing.T) {
	tests := []struct {
		name string
		s    *stack
		want []byte
	}{
		{"from empty stack", &stack{}, nil},
		{"returns last element", &stack{[]byte{0x0}, []byte{0x1}}, []byte{0x1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
