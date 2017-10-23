package main

import (
	"fmt"
)

type Vehicule struct{}

type Dragster struct {
	Vehicule
}

func (o Vehicule) Run()  { fmt.Println("Vehicle started") }
func (o Vehicule) Stop() { fmt.Println("Vehicle stopped") }

// ====================================================================

func (o Dragster) PrepareParachute() { fmt.Println("Parachute ready") }

func (o Dragster) Run() {
	o.PrepareParachute()
	o.Vehicule.Run()
}

func main() {

	d := &Dragster{}
	d.Run()
	d.Stop()
}

// ____________________________________________________________________________
// Training content part of "Go introductory level"
// Copyright Fred Ménez, March 2017, fred.menez@scalera.io

// Permission granted to reproduce for personal and educational use only.
// Commercial copying, hiring, lending is prohibited.
// ____________________________________________________________________________
