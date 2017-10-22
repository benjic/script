package ops

import (
	"bytes"
)

const (
	// OpCat concatenates two strings.
	//
	// [x1, x2] => [ x1x2 ]
	OpCat uint8 = 0x7e

	// OpSubstr returns a section of a string
	//
	// [in, begin, size] => [ in[begin:begin+size] ]
	OpSubstr uint8 = 0x7f

	// OpLeft returns characters to the left of the specified point in a string.
	//
	// [in, size] => [ out ]
	OpLeft uint8 = 0x80

	// OpRight returns characters to the right of the specified point in a
	// string.
	//
	// [in, size] => [ out ]
	OpRight uint8 = 0x81

	// OpSize returns the string length of the top element in the stack without
	// poping it from the stack.
	//
	// [in] => [in, len(in)]
	OpSize uint8 = 0x82
)

func opCat(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}
	x2 := c.Pop()
	x1 := c.Pop()

	var buf bytes.Buffer
	_, err := buf.Write(x1)
	if err != nil {
		return err
	}
	_, err = buf.Write(x2)
	if err != nil {
		return err
	}

	c.Push(buf.Bytes())

	return nil
}

func opSubstr(c Context) error {
	if c.Size() < 3 {
		return ErrInvalidStackOperation
	}

	size, err := readInt(c)
	if err != nil {
		return err
	}

	begin, err := readInt(c)
	if err != nil {
		return err
	}

	in := string(duplicate(c.Pop()))
	c.Push([]byte(in[begin : begin+size]))

	return nil
}

func opLeft(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}
	size, err := readInt(c)
	if err != nil {
		return err
	}
	in := string(duplicate(c.Pop()))

	if size >= int32(len(in)) {
		return ErrInvalidStackOperation
	}

	c.Push([]byte(in[0:size]))
	return nil
}

func opRight(c Context) error {
	if c.Size() < 2 {
		return ErrInvalidStackOperation
	}
	size, err := readInt(c)
	if err != nil {
		return err
	}

	in := string(duplicate(c.Pop()))

	if size >= int32(len(in)) {
		return ErrInvalidStackOperation
	}

	c.Push([]byte(in[size:]))
	return nil
}

func opSize(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}

	in := c.Pop()
	c.Push(in)

	return writeInt(c, int32(len(string(in))))
}
