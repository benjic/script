package ops

import (
	"reflect"
	"testing"
)

func Test_opToAltStack(t *testing.T) {
	type args struct {
		c *context
	}
	type want struct {
		stack *stack
		alt   *stack
		err   bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"simple",
			args{contextWithStackAndAlt(&stack{[]byte{0x00}, []byte{0xff}}, &stack{[]byte{0x00}})},
			want{
				&stack{[]byte{0x00}},
				&stack{[]byte{0x00}, []byte{0xff}},
				false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opToAltStack(tt.args.c)
			if (err != nil) != tt.want.err {
				t.Errorf("opToAltStack() error = %v, want err %v", err, tt.want.err)
			}

			if !reflect.DeepEqual(tt.want.stack, tt.args.c.stack) {
				t.Errorf("want %v; got %v", tt.want.stack, tt.args.c.stack)
			}

			if !reflect.DeepEqual(tt.want.alt, tt.args.c.alt) {
				t.Errorf("want %v; got %v", tt.want.alt, tt.args.c.alt)
			}
		})
	}
}

func Test_opFromAltStack(t *testing.T) {
	type args struct {
		c *context
	}
	type want struct {
		stack *stack
		alt   *stack
		err   bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"simple",
			args{contextWithStackAndAlt(&stack{[]byte{0x00}}, &stack{[]byte{0x00}, []byte{0xff}})},
			want{
				&stack{[]byte{0x00}, []byte{0xff}},
				&stack{[]byte{0x00}},
				false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opFromAltStack(tt.args.c)
			if (err != nil) != tt.want.err {
				t.Errorf("opFromAltStack() error = %v, want err %v", err, tt.want.err)
			}

			if !reflect.DeepEqual(tt.want.stack, tt.args.c.stack) {
				t.Errorf("want %v; got %v", tt.want.stack, tt.args.c.stack)
			}

			if !reflect.DeepEqual(tt.want.alt, tt.args.c.alt) {
				t.Errorf("want %v; got %v", tt.want.alt, tt.args.c.alt)
			}
		})
	}
}

func Test_opIfDup(t *testing.T) {
	type args struct {
		c *context
	}
	type want struct {
		err   bool
		stack *stack
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"non zero top stack",
			args{contextWithStack(&stack{[]byte{0x00, 0x00, 0x00, 0x01}})},
			want{false, &stack{[]byte{0x00, 0x00, 0x00, 0x01}, []byte{0x00, 0x00, 0x00, 0x01}}},
		},
		{
			"zero top stack",
			args{contextWithStack(&stack{[]byte{0x00, 0x00, 0x00, 0x00}})},
			want{false, &stack{[]byte{0x00, 0x00, 0x00, 0x00}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opIfDup(tt.args.c)
			if (err != nil) != tt.want.err {
				t.Errorf("opFromAltStack() error = %v, want err %v", err, tt.want.err)
			}

			if !reflect.DeepEqual(tt.want.stack, tt.args.c.stack) {
				t.Errorf("want %v; got %v", tt.want.stack, tt.args.c.stack)
			}
		})
	}
}

func Test_opDepth(t *testing.T) {
	type args struct {
		c *context
	}
	type want struct {
		err   bool
		stack *stack
		alt   *stack
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"empty stack",
			args{contextWithStack(&stack{})},
			want{false, &stack{[]byte{0x00, 0x00, 0x00, 0x00}}, &stack{}},
		},
		{
			"depth 1",
			args{contextWithStack(&stack{[]byte{0x00}})},
			want{
				false,
				&stack{[]byte{0x00}, []byte{0x00, 0x00, 0x00, 0x01}},
				&stack{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opDepth(tt.args.c)
			if (err != nil) != tt.want.err {
				t.Errorf("opFromAltStack() error = %v, want err %v", err, tt.want.err)
			}

			if !reflect.DeepEqual(tt.want.stack, tt.args.c.stack) {
				t.Errorf("want %v; got %v", tt.want.stack, tt.args.c.stack)
			}

			if !reflect.DeepEqual(tt.want.alt, tt.args.c.alt) {
				t.Errorf("want %v; got %v", tt.want.alt, tt.args.c.alt)
			}
		})
	}
}

func Test_opDrop(t *testing.T) {
	type args struct {
		c *context
	}
	type want struct {
		err   bool
		stack *stack
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"empty stack",
			args{contextWithStack(&stack{})},
			want{false, &stack{}},
		},
		{
			"single value",
			args{contextWithStack(&stack{[]byte{0x00}})},
			want{false, &stack{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opDrop(tt.args.c)
			if (err != nil) != tt.want.err {
				t.Errorf("opFromAltStack() error = %v, want err %v", err, tt.want.err)
			}

			if !reflect.DeepEqual(tt.want.stack, tt.args.c.stack) {
				t.Errorf("want %v; got %v", tt.want.stack, tt.args.c.stack)
			}
		})
	}
}
