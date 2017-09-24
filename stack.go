package script

type stack [][]byte

func (s *stack) Push(v []byte) {
	*s = append(*s, v)
}

func (s *stack) Pop() []byte {
	if len(*s) == 0 {
		return nil
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}
