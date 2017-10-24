package Gofiles

import (
	"fmt"
)
type VideoServer struct {
	NumVideos int
}
func main() {



	serverMars := VideoServer{
		100,
	}

	serverRepo := make(map[string]VideoServer)
	serverRepo["mars"] = serverMars

	fmt.Println( serverRepo["mars"])
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
