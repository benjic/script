package ops

import (
	"reflect"
	"testing"
)

func TestArithmeticOps(t *testing.T) {
	tests := []opTests{}

	for _, opTests := range tests {
		for _, test := range opTests.tests {

			t.Run(opTests.name+" "+test.name, func(t *testing.T) {
				err := opTests.op(test.args.context)
				if err != test.want.err {
					t.Errorf("%s() error = %v, want err %v", opTests.name, err, test.want.err)
				}

				if !reflect.DeepEqual(test.want.stack, test.args.context.stack) {
					t.Errorf("want %v; got %v", test.want.stack, test.args.context.stack)
				}

				if !reflect.DeepEqual(test.want.alt, test.args.context.alt) {
					t.Errorf("want %v; got %v", test.want.alt, test.args.context.alt)
				}
			})
		}
	}
}
