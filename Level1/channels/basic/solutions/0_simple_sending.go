package main

import (
	"fmt"
)

func main() {

	// Solution : démarrer la goroutine avant l'opération sur le canal,
	// de telle sorte que l'op d'écriture sur le canal ne bloquera pas indéfiniment

	ch := make(chan int)

	go func(ch chan int) {

		fmt.Println(<-ch)

	}(ch)

	ch <- 1

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
