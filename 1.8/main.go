package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fo, err := os.Create("./sample")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(fo)
	for i := 0; i < 1000000; i++ {
		r := rand.Int31n(10000)
		if _, err := w.Write([]byte(strconv.Itoa(int(r)) + "\n")); err != nil {
			panic(err)
		}
	}
	if err := w.Flush(); err != nil {
		panic(err)
	}
}
