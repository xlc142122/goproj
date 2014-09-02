package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Answer"] = 43
	fmt.Println(m)

	m["Answer"] = 48
	fmt.Println(m)

	delete(m, "Answer")
	fmt.Println(m)

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
