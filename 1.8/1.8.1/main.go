package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

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

func main() {
	fi, err := os.Open("./../sample")
	if err != nil {
		panic(err)
	}

	stack := NewStack(1000000)

	r := bufio.NewReader(fi)
	for {
		str, err := r.ReadString(byte('\n'))
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		stack.Push([]byte(str))
	}

	for {
		b := stack.Pop()
		if b == nil {
			break
		}
		fmt.Printf("%s", b)
	}
}
