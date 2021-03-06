package main

import "github.com/ryanfaerman/fsm"

type Location struct {
	State fsm.State

	machine *fsm.Machine
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

func (t *Location) CurrentState() fsm.State { return t.State }
func (t *Location) SetState(s fsm.State)    { t.State = s }

func (t *Location) Apply(r *fsm.Ruleset) *fsm.Machine {
	if t.machine == nil {
		t.machine = &fsm.Machine{Subject: t}
	}

	t.machine.Rules = r
	return t.machine
}
