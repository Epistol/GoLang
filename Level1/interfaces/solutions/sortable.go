package main

import "fmt"

type Sortable interface {
	Length() int
	Lower(pos int) bool
	Swap(pos int)
}

type IntArray []int

func (o IntArray) Length() int {
	return len(o)
}

func (o IntArray) Lower(pos int) bool {
	if pos < len(o)-1 && o[pos] < o[pos+1] {
		return true
	}
	return false
}

func (p *IntArray) Swap(pos int) {
	if pos <= len(*p)-2 {
		(*p)[pos], (*p)[pos+1] = (*p)[pos+1], (*p)[pos]
	}
}

type FloatArray []float64

func (o FloatArray) Length() int {
	return len(o)
}

func (p *FloatArray) Swap(pos int) {
	if pos <= len(*p)-2 {
		(*p)[pos], (*p)[pos+1] = (*p)[pos+1], (*p)[pos]
	}
}

func (o FloatArray) Lower(pos int) bool {
	if pos < len(o)-1 && o[pos] < o[pos+1] {
		return true
	}
	return false
}

func Sort(col Sortable) {
	if col.Length() == 1 {
		return
	}

	for {
		not_swapped := true
		for i := 0; i <= col.Length()-2; i++ {
			if col.Lower(i) == false {
				col.Swap(i)
				not_swapped = false
			}
		}
		if not_swapped {
			return
		}
	}
}

func main() {
	i := IntArray{1, 3, 14, 2}
	f := FloatArray{1.0, 3.1, 2.5}

	Sort(&i)
	fmt.Println("sorted i:", i)

	Sort(&f)
	fmt.Println("sorted f:", f)
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
