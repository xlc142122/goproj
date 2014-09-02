package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	b := Person{"Bob", 90}

	fmt.Println(a, b)
}
