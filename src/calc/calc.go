package main

import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			Usage()
			return
		}

		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])

		if err1 != nil || err2 != nil {
			Usage()
			return
		}

		ret := simplemath.Add(v1, v2)
		fmt.Println("Result:", ret)
	case "sqrt":
		if len(args) != 3 {
			Usage()
			return
		}

		v, err := strconv.Atoi(args[2])

		if err != nil {
			Usage()
			return
		} else {
			fmt.Println("Result:", simplemath.Sqrt(v))
		}
	default:
		Usage()
	}
}
