package ops

const (
	//OpVerify Marks transaction as invalid if top stack value is not true.
	OpVerify uint8 = 0x69
)

func opVerify(c Context) error {
	if c.Size() < 1 {
		return ErrInvalidStackOperation
	}

	b, err := readBool(c)
	if err != nil {
		return err
	}

	if !b {
		return ErrVerify
	}

	return nil

}
