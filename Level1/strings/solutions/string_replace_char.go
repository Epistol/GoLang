package main

import (
	"fmt"
)

func ReplaceCharUsingRuneArray(str string, idx int) string {

	stmod := []rune(str)
	stmod[idx] = 'a'

	return string(stmod)
}

func ReplaceCharUsingWriteToSlice(str string, idx int) string {
	// even though it's legit to slice a string : the slice will refer to a newly created buffer,
	// thus preventing code from modifying the string => strings being immutable

	s := []rune(str)

	s[idx] = 'a'

	return str
}

func ReplaceCharUsingSliceCopy(str string, idx int) string {
	// var sl []rune
	sl := make([]rune, len(str))

	// will succeed if sl has enough capacity
	copy(sl, []rune(str[:idx]))

	sl = append(sl, 'a')

	// ... operator allow expansion of slice/tab into a list of arguments
	sl = append(sl, []rune(str[idx+1:])...)

	return string(sl)
}

func main() {

	str := "Wizz"

	fmt.Println(ReplaceCharUsingRuneArray(str, 1))

	fmt.Println(ReplaceCharUsingWriteToSlice(str, 1))

	fmt.Println(ReplaceCharUsingSliceCopy(str, 1))
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
