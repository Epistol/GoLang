package main

import (
	"errors"
	"fmt"
	"time"
)

func GetRouterInfo(routerId string) string {

	if routerId == "juniper" {
		time.Sleep(300 * time.Millisecond)
		return "Juniper is go"
	} else {
		// fake failure
		time.Sleep(1300 * time.Millisecond)
		return "Cisco is going down"
	}
}

func main() {
	mission := `
In this program GetRouterInfo has a varying executation duration 
depending on the type of router being inspected.

We'd like GetRouterInfo to also return an error to indicate whether
timeout has occurred after a certain duration passed as argument.

Hint : the select statement can monitor more than one case clause...`

	fmt.Println(mission)

	fmt.Println(GetRouterInfo("juniper"))
	fmt.Println(GetRouterInfo("cisco"))
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
