package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	mission := `Write a program that loads a text file and copies it into
	another file. 

	And make sure that calls to Close() will be done on opened file handles 
	when exiting the program`

	fmt.Println(mission)
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
