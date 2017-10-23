package main

import (
	"fmt"
	"sync"
)

// What values will it output on the console ?

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {

		go func() {
			fmt.Println(i)

			wg.Done()
		}()
	}

	wg.Wait()
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
