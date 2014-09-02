package main

import (
	"fmt"
	"io"
	"strings"
)

type MyReader struct {
	Source string
}

func (mr *MyReader) Read() (int, error) {
	r := strings.NewReader(mr.Source)
	b := make([]byte, 1)

	n := 0
	var err error

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v b = %v\n", n, b)

		if err == io.EOF {
			break
		}
	}

	return n, err
}

func main() {
	mr := &MyReader{
		"A",
	}

	mr.Read()
}
