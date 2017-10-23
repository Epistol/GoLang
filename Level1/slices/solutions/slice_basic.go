package main

import (
	"fmt"
)

func RemoveSpaces(in string) string {
	out := ""
	for _, c := range in {
		if c != ' ' {
			out = out + string(c)
		}
	}
	return out
}

func main() {

	quote := RemoveSpaces("Tracy no panic in a ponycart")

	// quote is a odd-length palindrom, we'd like to print the three characters located at mid-length

	// create a slice that when printed will allow us to print characters going from midlength-1 to midlength+1

	fmt.Println(q[len(q)/2-1 : len(q)/2+2])

	fmt.Println("This is a base code for an exercice, this file is not meant to be run \"as is\"\n")

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
