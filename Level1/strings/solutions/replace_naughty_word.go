package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "James is such an uvavu. He should stay on Earth."

	// Value 1 means perform one replace operation. Be careful not to mistake this value for a start index
	sentence = strings.Replace(sentence, "uvavu", "fruitcake", 1)

	words := strings.Split(sentence, " ")

	// remove "n" in "an" by slicing from pos 0 and cutting at pos 1
	words[3] = words[3][:1]

	sentence = strings.Join(words, " ")

	fmt.Println(sentence)

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
