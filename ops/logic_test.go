package ops

import (
	"reflect"
	"testing"
)

func Test_equal(t *testing.T) {
	type args struct {
		c *context
	}
	tests := []struct {
		name string
		args args
		want *stack
	}{
		{
			"equal",
			args{contextWithStack(&stack{[]byte{0x00}, []byte{0x00}})},
			&stack{[]byte{0x00, 0x00, 0x00, 0x01}},
		},
		{
			"not equal",
			args{contextWithStack(&stack{[]byte{0x00}, []byte{0x01}})},
			&stack{[]byte{0x00, 0x00, 0x00, 0x00}},
		},
		{
			"not enough arguments",
			args{contextWithStack(&stack{[]byte{0x01}})},
			&stack{[]byte{0x00, 0x00, 0x00, 0x00}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opEqual(tt.args.c)

			if !reflect.DeepEqual(tt.args.c.stack, tt.want) {
				t.Errorf("Want: %+v; Got: %+v", tt.want, tt.args.c)
			}
		})
	}
}
