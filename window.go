package main

import "github.com/ryanfaerman/fsm"

type Window struct {
	name string
}

func NewWindowRules() fsm.Ruleset {
	rules := fsm.Ruleset{}
	rules.AddTransition(fsm.T{"title", "character selection"})
	rules.AddTransition(fsm.T{"character selection", "game"})
	rules.AddTransition(fsm.T{"character selection", "create character"})
	rules.AddRule(fsm.T{"title", "character selection"}, func(subject fsm.Stater, goal fsm.State) bool {
		return true
	})

	return rules
}
