package camp

import (
	"slices"
	"testing"
)

// names is the shorthand every list assertion in this package reads through.
func names(items []Item) []string {
	out := make([]string, len(items))
	for i, item := range items {
		out[i] = item.Name
	}
	return out
}

func TestEveryItemIsFiled(t *testing.T) {
	for _, item := range Gear() {
		switch item.Category {
		case Shelter, Sleep, Kitchen:
		default:
			t.Errorf("%q is filed under %q", item.Name, item.Category)
		}
	}
}

func TestInReadsTheKitchenInListOrder(t *testing.T) {
	want := []string{"gas stove", "pot", "mug", "water filter"}
	if got := names(In(Kitchen)); !slices.Equal(got, want) {
		t.Fatalf("In(Kitchen) = %q, want %q", got, want)
	}
}

func TestTheMatIsFiledWithTheBag(t *testing.T) {
	want := []string{"sleeping mat", "sleeping bag"}
	if got := names(In(Sleep)); !slices.Equal(got, want) {
		t.Fatalf("In(Sleep) = %q, want %q", got, want)
	}
}
