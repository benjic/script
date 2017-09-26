package script

import (
	"bytes"
	"encoding/hex"
	"strings"

	"github.com/benjic/script/ops"
)

// Parse creates a binary representation of a script that can be evaluated.
func Parse(in string) (*bytes.Buffer, error) {
	buf := []byte{}
	for _, token := range strings.Fields(in) {
		op, ok := ops.OpCodeLookup[token]
		if ok {
			buf = append(buf, byte(op))
		} else {
			str, err := hex.DecodeString(token)
			if err != nil {
				return new(bytes.Buffer), err
			}
			buf = append(buf, str...)
		}
	}
	return bytes.NewBuffer(buf), nil
}
