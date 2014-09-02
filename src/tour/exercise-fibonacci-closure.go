package main

import (
	"fmt"
)

//0 1 1 2 3 5
func fibonacci() func() int {
	s1 := 0
	s2 := 1

	return func() int {
		curval := s1 + s2
		s1 = s2
		s2 = curval

		return curval
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
