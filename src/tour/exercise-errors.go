package main

import (
	"fmt"
)

type ErrNegativeSqrt struct {
	Num float64
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e.Num)
}

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

func run(x float64) error {
	if x < 0 {
		return &ErrNegativeSqrt{
			x,
		}
	}

	return nil
}
func main() {

	x := -100.0
	if err := run(x); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(Sqrt(x))
	}
}
