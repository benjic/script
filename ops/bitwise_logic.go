package ops

import "bytes"

// TODO(benjic): Document the operation symbols
const (
	// OpInvert flips all the bits in the input.
	//
	// [in] => [out]
	OpInvert uint8 = 0x83

	// OpAnd applies a boolean *and* between each bit in the inputs.
	//
	// [x1, x2] => [x1 & x2]
	OpAnd uint8 = 0x84

	// OpOr applies a boolean *or* between each bit in the inputs.
	//
	// [x1, x2] => [x1 | x2]
	OpOr uint8 = 0x85

	// OpXor applies a boolean *xor* between each bit in the inputs.
	//
	// [x1, x2] => [x1 ^ x2]
	OpXor uint8 = 0x86

	// OpEqual return 1 if the inputs are exactly equal, otherwise 0.
	//
	// [x1, x2] => [x1 == x2]
	OpEqual uint8 = 0x87

	// OpEqualVerify uint8 = 0x88
	// TODO(benjic): Implement when VERIFY exists
)

func opInvert(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}

	in := duplicate(c.Pop())

	for i := range in {
		in[i] ^= 0xFF
	}

	c.Push(in)

	return nil
}

func opAnd(c Context) error {
	return applyOp(c, func(a, b byte) byte { return a & b })
}

func opOr(c Context) error {
	return applyOp(c, func(a, b byte) byte { return a | b })
}

func opXor(c Context) error {
	return applyOp(c, func(a, b byte) byte { return a ^ b })
}

func opEqual(c Context) error {
	if bytes.Equal(c.Pop(), c.Pop()) {
		c.Push([]byte{0x00, 0x00, 0x00, 0x01})
	} else {
		c.Push([]byte{0x00, 0x00, 0x00, 0x00})
	}

	return nil
}

func applyOp(c Context, binaryOp func(a, b byte) byte) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}

	x2 := c.Pop()
	x1 := c.Pop()

	// Pick the smaller of the two
	x := x2
	if len(x1) < len(x) {
		x = x1
	}

	out := make([]byte, len(x))
	for i := range x {
		out[i] = binaryOp(x1[i], x2[i])
	}

	c.Push(out)

	return nil
}
