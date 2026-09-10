package camp

import (
	"slices"
	"testing"
)

func TestEveryItemIsFiled(t *testing.T) {
	for _, item := range Gear() {
		if _, ok := CategoryOf(item); !ok {
			t.Errorf("%q is on the list with no category", item)
		}
	}
}

func TestTheStoveIsInTheKitchen(t *testing.T) {
	if got := In(Kitchen); !slices.Contains(got, "gas stove") {
		t.Fatalf("In(Kitchen) = %q, want the gas stove among them", got)
	}
}
