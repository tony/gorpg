package main

import (
	"github.com/ryanfaerman/fsm"
)

type Battle struct {
	State fsm.State

	machine *fsm.Machine
}

func NewBattleRules() fsm.Ruleset {
	rules := fsm.Ruleset{}
	rules.AddTransition(fsm.T{"wait", "attack"})
	rules.AddTransition(fsm.T{"wait", "run"})
	rules.AddTransition(fsm.T{"wait", "use item"})
	rules.AddTransition(fsm.T{"wait", "end"})
	rules.AddRule(fsm.T{"wait", "attack"}, func(subject fsm.Stater, goal fsm.State) bool {
		return true
	})

	return rules
}

func NewBattle() Battle {
	return Battle{
		State: "wait",
	}
}
