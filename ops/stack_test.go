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
		err   error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"empty stack",
			args{contextWithStackAndAlt(&stack{}, &stack{[]byte{0x00}})},
			want{
				&stack{},
				&stack{[]byte{0x00}},
				ErrInvalidStackOperation,
			},
		},
		{
			"simple",
			args{contextWithStackAndAlt(&stack{[]byte{0x00}, []byte{0xff}}, &stack{[]byte{0x00}})},
			want{
				&stack{[]byte{0x00}},
				&stack{[]byte{0x00}, []byte{0xff}},
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opToAltStack(tt.args.c)
			if err != tt.want.err {
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
		err   error
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
				nil,
			},
		},
		{
			"empty stack",
			args{contextWithStackAndAlt(&stack{[]byte{0x00}}, &stack{})},
			want{
				&stack{[]byte{0x00}},
				&stack{},
				ErrInvalidStackOperation,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opFromAltStack(tt.args.c)
			if err != tt.want.err {
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
				&stack{[]byte{0x00}, []byte{0x01, 0x00, 0x00, 0x00}},
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
		err   error
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
			want{ErrInvalidStackOperation, &stack{}},
		},
		{
			"single value",
			args{contextWithStack(&stack{[]byte{0x00}})},
			want{nil, &stack{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opDrop(tt.args.c)
			if err != tt.want.err {
				t.Errorf("opFromAltStack() error = %v, want err %v", err, tt.want.err)
			}

			if !reflect.DeepEqual(tt.want.stack, tt.args.c.stack) {
				t.Errorf("want %v; got %v", tt.want.stack, tt.args.c.stack)
			}
		})
	}
}

func Test_opDup(t *testing.T) {
	type args struct {
		c *context
	}
	type want struct {
		err   error
		stack *stack
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"simple",
			args{contextWithStack(&stack{{0x1}})},
			want{nil, &stack{{0x1}, {0x1}}},
		},
		{
			"empty stack",
			args{contextWithStack(&stack{})},
			want{ErrInvalidStackOperation, &stack{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := opDup(tt.args.c); err != tt.want.err {
				t.Errorf("opDup() error = %v, wantErr %v", err, tt.want.err)
			}

			if !reflect.DeepEqual(tt.want.stack, tt.args.c.stack) {
				t.Errorf("want %v; got %v", tt.want.stack, tt.args.c.stack)
			}
		})
	}
}

func Test_opOver(t *testing.T) {
	type args struct {
		c *context
	}
	type want struct {
		err   error
		stack *stack
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"simple",
			args{contextWithStack(&stack{{0x1}, {0x2}})},
			want{nil, &stack{{0x1}, {0x2}, {0x1}}},
		},
		{
			"too small stack",
			args{contextWithStack(&stack{{0x1}})},
			want{ErrInvalidStackOperation, &stack{{0x1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := opOver(tt.args.c); err != tt.want.err {
				t.Errorf("opOver() error = %v, wantErr %v", err, tt.want.err)
			}
			if !reflect.DeepEqual(tt.want.stack, tt.args.c.stack) {
				t.Errorf("want %v; got %v", tt.want.stack, tt.args.c.stack)
			}

		})
	}
}
