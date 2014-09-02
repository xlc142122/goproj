package main

import (
	"fmt"
)

// func add(x int, y int) int {
// 	return x + y
// }

//当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型外，其他的都可以省略
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(2, 3))
}
