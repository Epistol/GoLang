package main

import (
	"fmt"
)

func main() {

	type IntType int

	var i int = 9

	// The following does not compile because IntType and int are different types
	// Types are said to be equal when they have the same name

	// var it IntType = IntType(i)

	// Explicit conversion is required :
	var it IntType = IntType(i)

	fmt.Printf("it = %d\n", it)
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
