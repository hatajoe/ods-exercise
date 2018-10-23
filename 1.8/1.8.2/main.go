package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/hatajoe/ods-excercise/stack"
)

func main() {
	fi, err := os.Open("./../sample")
	if err != nil {
		panic(err)
	}

	st := stack.NewStack(50)

	buf := make(chan string, 50)
	r := bufio.NewReader(fi)
	for {
		str, err := r.ReadString(byte('\n'))
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		select {
		case buf <- str:
		default:
			print(st, buf)
			buf <- str
		}
	}
	print(st, buf)
}

func print(st *stack.Stack, buf chan string) {
	for i := 0; i < cap(buf); i++ {
		st.Push([]byte(<-buf))
	}
	for {
		b := st.Pop()
		if b == nil {
			break
		}
		fmt.Printf("%s", b)
	}
}
