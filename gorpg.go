package main

import (
	"log"

	"github.com/ryanfaerman/fsm"

	//"github.com/spf13/viper"
)

func testBasics(rules *fsm.Ruleset, char *Character) {
	permission := rules.Permitted(&char.location, "forest")
	log.Printf("Do I have energy to go to forest? %t", permission)
	log.Print(char.name)
}

func main() {
	rules := getLocationRules()

	char := NewCharacter("tony")

	testBasics(&rules, &char)

	showMainWindow()
}
