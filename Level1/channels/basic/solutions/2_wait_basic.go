package main

import (
	"fmt"
	"time"
)

var (
	N_ELTS = 5
)

func SendValues(c chan int) {
	for i := 0; i < N_ELTS; i++ {

		c <- i

		time.Sleep(500 * time.Millisecond)
	}

	close(c)
}

func main() {
	intChan := make(chan int)

	go SendValues(intChan)

	for {
		val, ok := <-intChan

		if !ok {
			return
		}

		fmt.Println(val)
	}
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
