package ops

type context interface {
	Pop() []byte
	Push([]byte)
}
