package camp

import "testing"

func TestEveryItemIsWeighed(t *testing.T) {
	for _, item := range Gear() {
		if _, ok := WeightGrams(item); !ok {
			t.Errorf("%q is on the list with no weight", item)
		}
	}
}

func TestPackWeightTotalsTheList(t *testing.T) {
	if got, want := PackWeight(), 4340; got != want {
		t.Fatalf("PackWeight() = %d g, want %d g", got, want)
	}
}
