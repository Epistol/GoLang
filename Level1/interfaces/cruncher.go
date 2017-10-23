package main

import (
	"fmt"
)

type NumberCruncher interface {
	JobDone() bool
	Process() bool
}

type BigNumberCruncher struct {
	jobDone      bool
	processCount int
}

func (o BigNumberCruncher) JobDone() bool { return o.jobDone }

func (o BigNumberCruncher) Process() bool {
	o.processCount++
	o.jobDone = true
	return o.jobDone
}

func ComputeStuff(cruncher NumberCruncher) bool {
	// ...
	return cruncher.Process()
}

func main() {

	mission := `Correct this program so that the call below to JobDone() returns a value equal to true
Your time to code now :)`
	fmt.Println(mission, "\n")

	bigcrunch := BigNumberCruncher{}

	ComputeStuff(bigcrunch)

	fmt.Println("BigNumberProcess : ", bigcrunch.JobDone())

}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
