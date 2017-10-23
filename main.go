package main

import (
	"fmt"
	/*"math"
	"strings"*/

)
func Concat(integers []int){
	integers = append(integers,1)
	fmt.Println("In concact : ", integers)
}

func ConcatByAddr(integers *[]int){
	*integers = append(*integers,1)
}



func main() {

	slice0 := []int{-1,0}
	Concat(slice0)
	ConcatByAddr(&slice0)
	fmt.Println("After : ", slice0)

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
