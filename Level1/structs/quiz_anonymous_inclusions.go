package main

import (
	"fmt"
)

func main() {

	type IngredientA struct {
		Formula string
	}

	type IngredientB struct {
		Formula string
	}

	type Sweet struct {
		IngredientA
		IngredientB
	}

	a := Sweet{IngredientA{"honey"}, IngredientB{"cinnamon"}}

	fmt.Println(a.IngredientA.Formula)

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
