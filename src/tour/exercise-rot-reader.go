package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(p byte) byte {
	switch {
	case p >= 'A' && p <= 'M':
		p = p - 'A' + 'N'
	case p >= 'N' && p <= 'Z':
		p = p - 'N' + 'A'
	case p >= 'a' && p <= 'm':
		p = p - 'a' + 'n'
	case p >= 'n' && p <= 'z':
		p = p - 'n' + 'a'
	}

	return p
}

func (v rot13Reader) Read(p []byte) (n int, err error) {
	original := make([]byte, 50)
	i, err := v.r.Read(original)
	for index, value := range original[:i] {
		p[index] = rot13(value)
	}

	return i, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
