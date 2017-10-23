package main

import (
	"fmt"
)

type NumberCruncher interface {
	JobDone() bool
	Process() bool
}

type BigNumberCruncher struct {
	jobDone bool
}

func (o BigNumberCruncher) JobDone() bool { return o.jobDone }

func (o *BigNumberCruncher) Process() bool {
	o.jobDone = true
	return o.jobDone
}

func ComputeStuff(cruncher NumberCruncher) bool {
	// ...
	res := cruncher.Process()
	// ...
	return res
}

func main() {

	var crunch NumberCruncher = &BigNumberCruncher{}

	fmt.Println(ComputeStuff(crunch))

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
