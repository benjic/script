package script

import (
	"reflect"
	"strings"
	"testing"

	"github.com/benjic/script/ops"
)

func TestParse(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"good data",
			args{"FF "},
			[]byte{0x01, 0xFF},
			false,
		},
		{
			"bad data",
			args{"nothex"},
			nil,
			true,
		},
		{
			"too large data",
			args{strings.Repeat("A", 1000)},
			nil,
			true,
		},
		{
			"op",
			args{"OP_FALSE"},
			[]byte{ops.OP_FALSE},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Bytes(), tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
