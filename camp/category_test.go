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

func TestInReadsTheKitchenInListOrder(t *testing.T) {
	want := []string{"gas stove", "water filter"}
	if got := In(Kitchen); !slices.Equal(got, want) {
		t.Fatalf("In(Kitchen) = %q, want %q", got, want)
	}
}
