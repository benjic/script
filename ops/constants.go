package ops

import (
	"errors"
)

// TODO(benjic): Document the operation symbols
const (
	Op0         uint8 = 0x00
	OpFalse     uint8 = 0x00
	OpPushData1 uint8 = 0x4c
	OpPushData2 uint8 = 0x4d
	OpPushData4 uint8 = 0x4e
	Op1Negate   uint8 = 0x4f
	OpTrue      uint8 = 0x51
	Op1         uint8 = 0x51
	Op2         uint8 = 0x52
	Op3         uint8 = 0x53
	Op4         uint8 = 0x54
	Op5         uint8 = 0x55
	Op6         uint8 = 0x56
	Op7         uint8 = 0x57
	Op8         uint8 = 0x58
	Op9         uint8 = 0x59
	Op10        uint8 = 0x5a
	Op11        uint8 = 0x5b
	Op12        uint8 = 0x5c
	Op13        uint8 = 0x5d
	Op14        uint8 = 0x5e
	Op15        uint8 = 0x5f
	Op16        uint8 = 0x60
)

var (
	// ErrInvalidStackOperation indicates an operation was attempted against a
	// stack with an invalid size.
	ErrInvalidStackOperation = errors.New("operation not valid with the current stack size")

	// ErrVerify indicates an OpVerify operation failed.
	ErrVerify = errors.New("script failed an OP_VERIFY operation")

	// ErrEqualVerify indicates an OpEqualVerify operation has failed.
	ErrEqualVerify = errors.New("script failed an OP_EQUALVERIFY operation")

	// ErrNumEqualVerify indicates an OpNumEqualVerify has failed.
	ErrNumEqualVerify = errors.New("script failed an OP_NUMEQUALVERIFY operation")

	// ErrInsufficientNumberOfBytes indicates a push the number of bytes was unsuccessful
	// due to the reader failing to supply the number of bytes.
	ErrInsufficientNumberOfBytes = errors.New("insufficient number of bytes available")
)

func createOpPushNBytes(n uint8) Op {
	return func(c Context) error {
		if n == 0 {
			// TODO(benjic): Verify this is the correct behavior
			return nil
		}

		bs := make([]byte, n)
		cnt, err := c.Read(bs)
		if err != nil {
			return err
		}

		if cnt != int(n) {
			return ErrInsufficientNumberOfBytes
		}

		c.Push(bs)

		return nil
	}
}

func createOpPushN(n uint8) Op {
	return func(c Context) error {
		return writeInt(c, int32(n))
	}
}

func opFalse(c Context) error {
	return writeBool(c, false)
}

func op1Negate(c Context) error {
	return writeInt(c, -1)
}

func opTrue(c Context) error {
	return writeBool(c, true)
}
