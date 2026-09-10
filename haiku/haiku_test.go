package haiku

import (
	"strings"
	"testing"
)

func TestSeasonsHoldsEverySeason(t *testing.T) {
	if got := len(Seasons()); got != 4 {
		t.Fatalf("poems = %d, want 4", got)
	}
}

func TestSeasonsOpensOnSpring(t *testing.T) {
	if got := Seasons()[0]; got != Spring() {
		t.Fatalf("the year opens on %q", got)
	}
}

func TestEverySeasonIsThreeLines(t *testing.T) {
	for _, poem := range Seasons() {
		if got := len(strings.Split(poem, "\n")); got != 3 {
			t.Errorf("lines = %d, want 3: %q", got, poem)
		}
	}
}
