package main

import "testing"
import "fmt"

func TestGetLocationRules(t *testing.T) {
	rules := getLocationRules()
	location := Location{State: "town"}
	val := rules.Permitted(&location, "forest")
	if val == false {
		t.Error("Should be allowed to forest")
	}

	fmt.Println("HI")
}
