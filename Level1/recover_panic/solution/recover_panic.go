package main

import (
	"fmt"
)

func letUsPanic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Panic detected: ", e)
		}
	}()
	panic("boom!")
	fmt.Println("This will never be called")
}

func main() {
	letUsPanic()

	fmt.Println("Continuing execution after panic was recovered...")
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
