package main

import (
	"fmt"
	"time"
)

// Usage example of time.After to schedule an event in the future
func StartTicker() {
	ticker := time.NewTicker(100 * time.Millisecond)
	ticks := 0
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick:", t)
			ticks++
		}
		if ticks >= 10 {
			return
		}
	}
}

func main() {
	go StartTicker()

	// Using time.After to schedule an event in the future
	select {
	case t := <-time.After(1 * time.Second):
		fmt.Println("Time after 1 sec:", t)
	}
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
