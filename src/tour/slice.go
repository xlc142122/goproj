package main

import (
	"fmt"
)

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	fmt.Println("p[1:3] ==", p[1:4]) //从下标1取到4
	fmt.Println("p[0:3] ==", p[:3])  //从0-2
	fmt.Println("p[4:] ==", p[4:])   //从4-结束
}
