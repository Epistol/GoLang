package main

import (
	"errors"
	"fmt"
	"time"
)

var Routers []string = []string{"Juniper00", "Cisco01"}

func main() {

	for routerIndex, rname := range Routers {

		info, e := GetRouterInfoWithTimeout(routerIndex, 1*time.Second)

		fmt.Printf("%s status: %s\n", rname, info)
	}

}

func GetRouterInfoWithTimeout(routerID int, d time.Duration) string {

	// Local func that immediately returns a result channel to be used below in select
	chRouterStatus := func(routerID int) chan string {
		chRes := make(chan string)

		// Simulate async call to remote service with artificial latency
		go func(chan string) {
			defer close(chRes)

			if routerID == 0 {
				time.Sleep(300 * time.Millisecond)
			} else {
				// fake delay
				time.Sleep(1300 * time.Millisecond)
			}

			chRes <- "live"

		}(chRes)

		return chRes
	}(routerID)

	select {

	case res := <-chRouterStatus:
		// all good, we received an answer and then we just exit
		return res

	case <-time.After(d):
		// we didn't receive in a timely manner and time.After gets a chance
		// to execute and send over its channel
		return "timeout"
	}

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
