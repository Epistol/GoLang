package main

import (
	"fmt"
)

func adder() func(int) int {
	var x int

	return func(increment int) int {
		x += increment
		return x
	}
}

func main() {
	f := adder()

	fmt.Println(f(0), f(1), f(2))

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
