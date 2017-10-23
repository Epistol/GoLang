package main

import (
	"fmt"
	"time"
)

var (
	N_ELTS     = 5
	count  int = 1
)

func SendValues(c chan int) {
	for i := 0; i < N_ELTS; i++ {
		c <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(c)
}

func main() {
	numCallsDuringfreeTime := 0

	intChan := make(chan int)

	go SendValues(intChan)

	fmt.Println(`
Complete this program with a for loop with a select statement inside that

* prints receives values from intChan until done

`)

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
