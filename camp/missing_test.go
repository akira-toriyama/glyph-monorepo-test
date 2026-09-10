package camp

import (
	"slices"
	"testing"
)

func TestMissingIsTheWholeListWhenNothingIsPacked(t *testing.T) {
	if got, want := len(Missing(nil)), len(Gear()); got != want {
		t.Fatalf("Missing(nil) has %d items, want %d", got, want)
	}
}

func TestMissingIgnoresWhatThePackHoldsBeyondTheList(t *testing.T) {
	packed := append(names(Gear()), "harmonica")
	if got := Missing(packed); len(got) != 0 {
		t.Fatalf("still missing %q with the whole list packed", names(got))
	}
}

func TestMissingAnswersInListOrder(t *testing.T) {
	packed := []string{"water filter", "sleeping bag", "tent"}
	want := []string{"sleeping mat", "gas stove"}
	if got := names(Missing(packed)); !slices.Equal(got, want) {
		t.Fatalf("Missing() = %q, want %q", got, want)
	}
}
