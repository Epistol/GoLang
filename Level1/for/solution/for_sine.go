package main

import (
	"fmt"
	"math"
	"strings"
)

func makeLines() [100][]rune {
	var lines [100][]rune
	for i := 0; i < 100; i++ {
		l := int(math.Abs(math.Sin(0.05*float64(i)*3.14) * 100.0))
		str := strings.Repeat("*", int(l))
		for c := 0; c < l; c++ {
			lines[i] = append(lines[i], rune(str[c]))
		}
	}
	return lines
}

func main() {

	lines := makeLines()

	for _, l := range lines {
		for _, c := range l {
			fmt.Printf("%c", c)
		}
		fmt.Println("")
	}

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
