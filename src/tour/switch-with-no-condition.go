package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	//无条件的switch用于替代较长的if-then-else
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
