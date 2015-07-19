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

func NewCharacter(name string) Character {
	return Character{name: name, location: Location{State: "town"}}
}
