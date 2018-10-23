package stack

type Stack struct {
	buffers [][]byte
}

func NewStack(s int) *Stack {
	return &Stack{
		buffers: make([][]byte, 0, s),
	}
}

func (s *Stack) Push(b []byte) {
	s.buffers = append(s.buffers, b)
}

func (s *Stack) Pop() []byte {
	if len(s.buffers) <= 0 {
		return nil
	}
	defer func() {
		s.buffers = s.buffers[:len(s.buffers)-1]
	}()

	return s.buffers[len(s.buffers)-1]
}
