package main

import (
	"fmt"
)

// The goal of this program is to add extra functionality to this simple function
// The change to user code should be minimal
func Connected() bool { return true }

// The signature of this function is simply:   func() bool

// This function take as param a function like Connected, and returns a closure with extra operations
// that be performed before or after calling the user provided function
func MakeConnectedWrapper(userFunc func() bool) func() bool {
	NumQueries := 0
	return func() bool {
		fmt.Println("NumQueries before:", NumQueries)
		result := userFunc()
		NumQueries++
		fmt.Println("NumQueries after:", NumQueries)
		return result
	}
}

func main() {
	connected := MakeConnectedWrapper(Connected)

	// Instead of directly calling Connected, user code can call the wrapper:
	connected()
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
