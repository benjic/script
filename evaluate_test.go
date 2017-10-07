package script

import (
	"bytes"
	"io"
	"testing"

	"github.com/benjic/script/ops"
)

func TestEvaluate(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"top stack value is true",
			args{bytes.NewReader([]byte{
				ops.OpTrue,
			})},
			false,
		},
		{
			"top value is false",
			args{bytes.NewReader([]byte{
				ops.OpFalse,
			})},
			true,
		},

		{
			"unknown op code",
			args{bytes.NewReader([]byte{
				0xFF,
			})},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Evaluate(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
