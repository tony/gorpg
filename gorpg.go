package main

import (
	"log"

	"github.com/ryanfaerman/fsm"
	//"github.com/spf13/viper"
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

func getLocationRules() fsm.Ruleset {
	rules := fsm.Ruleset{}
	rules.AddTransition(fsm.T{"town", "forest"})
	rules.AddTransition(fsm.T{"forest", "battle"})
	rules.AddRule(fsm.T{"town", "forest"}, func(subject fsm.Stater, goal fsm.State) bool {
		println("You must have enough energy to go to forest.")
		return true
	})

	return rules
}

func testBasics(rules *fsm.Ruleset, char *Character) {
	permission := rules.Permitted(&char.location, "forest")
	log.Printf("Do I have energy to go to forest? %t", permission)
	log.Print(char.name)
}
func main() {
	rules := getLocationRules()

	char := getCharacter("tony")

	testBasics(&rules, &char)
}

func getCharacter(name string) Character {
	return Character{name: name, location: Location{State: "town"}}
}
