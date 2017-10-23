package main

import "fmt"

func main() {

	mission := `Sort 2 collection types using the methods defined in an interface.
* code 2 collection types that can store ints and floats respectively
* Add a function Length() to these types
* Add a function LowerThan(pos int) bool :compares values at pos, pos+1 and return true if value[pos] <= value[pos+1]
* Add a function Swap(pos int) to swap value[pos] with value[pos+1]
* Create a interface named Sortable that expose these functions
* Create a Sort function that takes as parameter a Sortable and uses its methods to sort the elements

Get going, it's your time to code now :-)`

	fmt.Println(mission)
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
