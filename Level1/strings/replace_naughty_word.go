package main

import (
	"fmt"
)

func main() {
	sentence := "James is such an uvavu. He should stay on Earth."

	mission := `

We all know that uvavu is the more unbearable insult in Klingon.
Below you'll find a message intercepted by our nicebot on alt.mars101.
Write code that will replace uvavu by something more appropriate like fruitcake.
And also remove the n so that it reads properly : such a fruitcake.

Get going, your time to code now ! :)`

	fmt.Println(mission, "\n")
	sentence = sentence
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
