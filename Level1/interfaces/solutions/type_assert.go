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

	switch s.(type) {

	case Trickerer:
		fmt.Println("Trickerer")

	default:
		fmt.Println("not a trickerer")
	}
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
