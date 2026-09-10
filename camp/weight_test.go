package camp

import "testing"

func TestNoItemWeighsNothing(t *testing.T) {
	for _, item := range Gear() {
		if item.Grams <= 0 {
			t.Errorf("%q is on the list at %d g", item.Name, item.Grams)
		}
	}
}

func TestPackWeightTotalsTheList(t *testing.T) {
	if got, want := PackWeight(), 4920; got != want {
		t.Fatalf("PackWeight() = %d g, want %d g", got, want)
	}
}
