package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	var flag float64
	for {
		z = z - (z*z-x)/(2*z)

		fmt.Println("seq", z)

		if flag == z || (flag-z) < 0.01 {
			break
		} else {
			flag = z
		}

	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
