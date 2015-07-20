package main

import "testing"

func TestExpTilLevel(t *testing.T) {
	c := NewCharacter("TestChar")

	if c.experienceForLevel() != 9 {
		t.Fatal("Experience until level 2 should be 9")
	}
	c.level = 16

	level16Expected := uint8((16 * 5) - 38)
	if c.experienceForLevel() != level16Expected {
		t.Error("Experience should jump 5* after level 15")
	}

	c.level = 31

	level31Expected := uint8((31 * 9) - 158)
	if c.experienceForLevel() != level31Expected {
		t.Error("Experience should jump 9* after level 30")
	}

}
