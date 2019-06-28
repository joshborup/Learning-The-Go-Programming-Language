package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	c := 30

	// go doesnt need parenthesis for if statements
	if a < b {
		fmt.Println(a, "is greater than", b)
	}
	// go allows if else statememt for logic flow control
	if b < a {
		fmt.Println(b, "is greater than", a)
	} else if c > b {
		fmt.Println(c, "is greater than", b)
	}
	// go allows you to assign variables inline with the if statment that is testing the variables value
	if fraction := b / a; fraction < 3 {
		fmt.Println(fraction, "is less than", 3)
	}
}
