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
		name    string
		args    args
		want    *stack
		wantErr bool
	}{
		{
			"equal",
			args{contextWithStack(&stack{[]byte{0x00}, []byte{0x00}})},
			&stack{[]byte{0x00, 0x00, 0x00, 0x01}},
			false,
		},
		{
			"not equal",
			args{contextWithStack(&stack{[]byte{0x00}, []byte{0x01}})},
			&stack{[]byte{0x00, 0x00, 0x00, 0x00}},
			false,
		},
		{
			"not enough arguments",
			args{contextWithStack(&stack{[]byte{0x01}})},
			&stack{[]byte{0x00, 0x00, 0x00, 0x00}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := opEqual(tt.args.c)

			if (err != nil) != tt.wantErr {
				t.Errorf("opEqual() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.args.c.stack, tt.want) {
				t.Errorf("Want: %+v; Got: %+v", tt.want, tt.args.c)
			}
		})
	}
}
