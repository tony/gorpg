package main

type Character struct {
	name       string
	location   Location
	hp         uint16
	mp         uint16
	energy     uint8
	level      uint8
	experience uint32
}

func (c Character) gainLevel() {
}

func (c Character) experienceForLevel() uint8 {
	if c.level > 30 {
		return (9 * c.level) - 158
	} else if c.level > 15 {
		return (5 * c.level) - 38
	}
	return (2 * c.level) + 7
}

func (c Character) nextLevel() {
	c.experience = 0
	c.level += 1
}

func NewCharacter(name string) Character {
	return Character{
		name:     name,
		location: Location{State: "town"},
		level:    1,
	}
}
