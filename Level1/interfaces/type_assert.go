package main

import (
	"fmt"
)

type Trickerer interface {
	Trick()
}

func main() {
	var s interface{}

	s = 2

	mission := `Continue this program by using a type assertion to test whether 
	s is a value of type Trickerer or not, and display a message indicating the result`

	s = s
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
