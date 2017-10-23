package main

import (
	"fmt"
	"time"
)

var (
	N_ELTS = 5
)

func SendValues(ch chan int) {
	for i := 0; i < N_ELTS; i++ {

		ch <- i

		time.Sleep(500 * time.Millisecond)
	}

	close(ch)
}

func main() {
	numIterDuringFreeTime := 0

	intChan := make(chan int)

	go SendValues(intChan)

	done := false

	for !done {
		select {

		case val, ok := <-intChan:

			if !ok {
				done = true
			} else {
				fmt.Println(val)
			}

		default:
			numIterDuringFreeTime++

		}
	}

	fmt.Println("Num iterations during \"free time\":", numIterDuringFreeTime)
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
