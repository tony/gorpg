package main

import (
	"log"

	"github.com/ryanfaerman/fsm"
)

type Location struct {
	State fsm.State

	machine *fsm.Machine
}

type Character struct {
	name     string
	location Location
}

func (t *Location) CurrentState() fsm.State { return t.State }
func (t *Location) SetState(s fsm.State)    { t.State = s }

func (t *Location) Apply(r *fsm.Ruleset) *fsm.Machine {
	if t.machine == nil {
		t.machine = &fsm.Machine{Subject: t}
	}

	t.machine.Rules = r
	return t.machine
}

func main() {
	println("la")
	var err error
	loc := Location{State: "town"}
	rules := fsm.Ruleset{}
	rules.AddTransition(fsm.T{"town", "forest"})
	rules.AddTransition(fsm.T{"forest", "battle"})

	println("Hi")
	log.Print("Lol")
	if err != nil {
		log.Fatal(err)
	}
	println("Hi")

	rules.AddRule(fsm.T{"town", "forest"}, func(subject fsm.Stater, goal fsm.State) bool {
		println("You must have enough energy to go to forest.")
		return true
	})
	err = loc.Apply(&rules).Transition("town")
	permission := rules.Permitted(&Location{State: "town"}, "forest")
	println(permission)

	char := getCharacter("tony")
	permission = rules.Permitted(&char.location, "forest")
	println(permission)
}

func getCharacter(name string) *Character {
	return &Character{name: name, location: Location{State: "town"}}
}
