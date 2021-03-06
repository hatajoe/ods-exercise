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

	stack := stack.NewStack(1000000)

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
