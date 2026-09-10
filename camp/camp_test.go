package camp

import "testing"

func TestGearHasATent(t *testing.T) {
	if got := Gear()[0]; got != "tent" {
		t.Fatalf("first item = %q, want tent", got)
	}
}

// The list is packing order, so the mat has to be listed before the bag it
// goes under.
func TestTheMatIsPackedUnderTheBag(t *testing.T) {
	mat, bag := -1, -1
	for i, item := range Gear() {
		switch item {
		case "sleeping mat":
			mat = i
		case "sleeping bag":
			bag = i
		}
	}
	if mat < 0 || bag < 0 {
		t.Fatalf("mat at %d, bag at %d: an item left the list", mat, bag)
	}
	if mat > bag {
		t.Fatalf("mat at %d, bag at %d: the bag goes in under the mat", mat, bag)
	}
}
