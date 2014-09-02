package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68, -74.39}
	fmt.Println(m["Bell Labs"])

	var vm = map[string]Vertex{
		"a": {1.0, 2.1},
		"b": {3.1, 4.3},
	}

	fmt.Println(vm)
}
