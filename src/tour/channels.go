package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum //将sum送入c
}

func main() {
	a := []int{7, 2, 8, 19, 4, 0}

	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c //从c中获取

	fmt.Println(x, y, x+y)
}
