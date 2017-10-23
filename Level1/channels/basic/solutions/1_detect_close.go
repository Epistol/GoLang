package main

import (
	"fmt"
)

func ConsumeChannel1(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			return
		}

		fmt.Println(v)
	}
}

func ConsumeChannel2(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {

	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)

	ConsumeChannel1(ch)

	// ou aussi :

	// ConsumeChannel2(ch)
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
