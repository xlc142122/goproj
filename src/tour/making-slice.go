package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	d = append(d, 2, 3, 4)
	printSlice("new d", d)

	for i, v := range d {
		fmt.Printf("%d = %d\n", i, v)
	}

	//忽略序号
	for _, value := range d {
		fmt.Println(value)
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
