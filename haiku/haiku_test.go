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
		t.Fatalf("the year opens on %q", got.Text)
	}
}

func TestEveryPoemIsThreeLines(t *testing.T) {
	for _, p := range Seasons() {
		if got := len(strings.Split(p.Text, "\n")); got != 3 {
			t.Errorf("%s: lines = %d, want 3", p.Author, got)
		}
	}
}

func TestEveryPoemNamesItsPoet(t *testing.T) {
	for _, p := range Seasons() {
		if p.Author == "" {
			t.Errorf("no poet: %q", p.Text)
		}
	}
}
