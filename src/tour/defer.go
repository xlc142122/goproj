package main

import (
	"fmt"
)

func main() {
	//defer语句会延迟函数的执行到上层函数返回
	defer fmt.Println("world")
	fmt.Println("Hello")
}
