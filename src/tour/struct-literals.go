package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} //类型为vertex

	v2 = Vertex{X: 1} //省略了Y：0

	v3 = Vertex{} //省略了X:0和Y:0

	p = &Vertex{1, 2} //类型为*Vertex
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
