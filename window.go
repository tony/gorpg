package main

import "github.com/ryanfaerman/fsm"
import ui "github.com/gizak/termui"
import "time"

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

func showMainWindow() {

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	p := ui.NewPar(":PRESS q TO QUIT DEMO")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = ui.ColorWhite
	p.Border.Label = "Text Box"
	p.Border.FgColor = ui.ColorCyan

	draw := func(t int) {
		ui.Render(p)
	}

	evt := ui.EventCh()

	i := 0
	for {
		select {
		case e := <-evt:
			if e.Type == ui.EventKey && e.Ch == 'q' {
				return
			}
		default:
			draw(i)
			i++
			if i == 102 {
				return
			}
			time.Sleep(time.Second / 2)
		}
	}
}
