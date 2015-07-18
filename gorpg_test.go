package main

import "testing"

func TestGetLocationRules(t *testing.T) {
	rules := getLocationRules()
	location := Location{State: "town"}
	val := rules.Permitted(&location, "forest")
	if val == false {
		t.Error("Should be allowed to forest")
	}

}

func TestGetCharacter(t *testing.T) {
	name := "tony"
	char := getCharacter(name)
	if char.name != name {
		t.Errorf("should return the right name (%s)", name)
	}
}

func TestGetBattle(t *testing.T) {
	battle := getBattle()
	if battle.State != "wait" {
		t.Error("Battles should always start wait.")
	}
}
