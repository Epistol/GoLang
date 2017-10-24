package main

import (
	"fmt"
	"math/rand"
)
type User struct {
	/*NumVideos int*/
}
func main() {

/*

	serverMars := VideoServer{
		100,
	}*/

	users := make(map[int]*User)

	for i:=0; i < 10; i++{
		users[i] = &User{rand.Intn(20)}
	}


	for idx,_ := range users{
		fmt.Println("%d\n", idx)
	}
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
