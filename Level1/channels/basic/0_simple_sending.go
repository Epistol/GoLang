package main

import (
	"fmt"
	"time"
)

// The following program causes a deadlock, why ?

func main() {
	ch = make(chan int)

	ch <- 1

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	time.Sleep(10 * time.Millisecond)
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
