package main

import (
	"fmt"
)

func main() {

	mission := `Exercice : imagine a struct named Vehicle with methods Run & Stop.
We'd like to have another struct named Dragster with a Run method reusing Vehicle's Run method.
We'd also like to specialise Dragster's method so that it can execute its own code before running Vehicle's Run method.`

	fmt.Println(mission)
	fmt.Println("\nYour time to code know ! :)\n")

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
