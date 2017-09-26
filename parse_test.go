package script

import (
	"reflect"
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
			"simple",
			args{"FF OP_FALSE"},
			[]byte{0xFF, ops.OP_FALSE},
			false,
		},
		{
			"bad data",
			args{"nothex"},
			nil,
			true,
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
