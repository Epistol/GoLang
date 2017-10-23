package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var e error
	var fa, fb *os.File

	contents := make([]byte, 256)

	if fa, e = os.Open("fa"); e != nil {
		fmt.Println("error:", e)
		return
	}

	if fb, e = os.Create("fb"); e != nil {
		fmt.Println("error:", e)
		return
	}

	defer fa.Close()
	defer fb.Close()

	if contents, e = ioutil.ReadAll(fa); e != nil {
		fmt.Println("Error reading:", e)
	} else {
		fb.Write(contents)
	}

	fmt.Println("Done")
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
