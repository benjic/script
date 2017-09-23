package script

type stacker interface {
	Push([]byte)
	Pop() []byte
}

type byteStack [][]byte

func (s *byteStack) Push(v []byte) {
	*s = append(*s, v)
}

func (s *byteStack) Pop() []byte {
	if len(*s) == 0 {
		return nil
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}
