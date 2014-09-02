package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println("x =", v.X)
	fmt.Println("y =", v.Y)

	v.X = 10
	fmt.Println("x =", v.X)

	//通过指针访问
	p := &v
	fmt.Println(p.X)
}
