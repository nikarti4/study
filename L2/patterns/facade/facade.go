package main

import (
	"fmt"
)

// фасад
type Aircraft struct {
	avionics Avionics
	engine Engine
	fcs FCS
}

func NewAircraft() Aircraft {
	return Aircraft {
		avionics: Avionics{name: "A"},
		engine: Engine{name: "B"},
		fcs: FCS{name: "C"},
	}
}

func (a *Aircraft) CheckAll() {
	fmt.Println("Check all systems...")
	a.engine.Check()
	a.avionics.Check()
	a.fcs.Check()
}

// подсистема 1
type Engine struct {
	name string
}

func (e *Engine) Check() {
	fmt.Println("Check engine...");
}

// подсистема 2
type Avionics struct {
	name string
}

func (a *Avionics) Check() {
	fmt.Println("Check avionics...");
}

// подсистема 3
type FCS struct {
	name string
}

func (f *FCS) Check() {
	fmt.Println("Check FCS...")
}


func main() {
	plane := NewAircraft()
	plane.CheckAll();
}