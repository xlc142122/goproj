package main

import (
	"fmt"
	"strings"
)

var f string = "I am learning Go!"

func WordCount(s string) map[string]int {
	splitret := strings.Fields(s)

	m := make(map[string]int)

	for i := 0; i < len(splitret); i++ {
		v, ok := m[splitret[i]]
		if ok == true {
			v = v + 1
			m[splitret[i]] = v
		} else {
			m[splitret[i]] = 1
		}
	}

	return m
}

func main() {
	fmt.Println(WordCount(f))
}
