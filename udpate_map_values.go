package main

import (
	"fmt"
)

func main() {

	type Cluster struct {
		NumVideos int
	}

	// If you want to easily modify a map item, store pointers to values instead
	// of just values

	Clusters := map[string]*Cluster{
		"pluto": &Cluster{},
	}

	Clusters["pluto"].NumVideos++

	fmt.Println("Clusters:", Clusters["pluto"])
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
