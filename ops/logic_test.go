package ops

import (
	"reflect"
	"testing"

	"github.com/benjic/script"
)

func Test_equal(t *testing.T) {
	type args struct {
		c *script.Stack
	}
	tests := []struct {
		name string
		args args
		want *script.Stack
	}{
		{
			"equal",
			args{&script.Stack{[]byte{0x00}, []byte{0x00}}},
			&script.Stack{[]byte{0x00, 0x00, 0x00, 0x01}},
		},
		{
			"not equal",
			args{&script.Stack{[]byte{0x00}, []byte{0x01}}},
			&script.Stack{[]byte{0x00, 0x00, 0x00, 0x00}},
		},
		{
			"not enough arguments",
			args{&script.Stack{[]byte{0x01}}},
			&script.Stack{[]byte{0x00, 0x00, 0x00, 0x00}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equal(tt.args.c)

			if !reflect.DeepEqual(tt.args.c, tt.want) {
				t.Errorf("Want: %+v; Got: %+v", tt.want, tt.args.c)
			}
		})
	}
}
