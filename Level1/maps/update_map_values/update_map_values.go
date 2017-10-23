package main

import (
	"fmt"
)

func main() {
	mission := `We want to store information about servers hosting videos.

	1) Write a simple program that gathers some fields, like a video counter, into a ServerInfo struct type
	2) and then, using a map[string]ServerInfo, assign a few values of this struct type to server names
	3) finally, pick a server stored in the map and attemp to modify one of its members.`

	fmt.Println(mission)
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
