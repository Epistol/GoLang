package main

import (
	"fmt"
)

func ConcatByReturningACopy(integers []int) []int {
	integers = append(integers, 1)
	return integers
}

func ConcatByAddress(integers *[]int) {
	*integers = append(*integers, 1)
}

func main() {
	tab := []int{-1, 0}

	fmt.Println(ConcatByReturningACopy(tab))

	ConcatByAddress(&tab)
	fmt.Println(tab)
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
