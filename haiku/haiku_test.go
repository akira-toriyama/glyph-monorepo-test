package haiku

import (
	"slices"
	"strings"
	"testing"
)

func TestSeasonsHoldsEverySeason(t *testing.T) {
	if got := len(Seasons()); got != 5 {
		t.Fatalf("poems = %d, want 5", got)
	}
	if !slices.Contains(Seasons(), NewYear()) {
		t.Error("the new year is filed nowhere")
	}
}

func TestSeasonsOpensOnTheNewYear(t *testing.T) {
	if got := Seasons()[0]; got != NewYear() {
		t.Fatalf("the book opens on %q", got.Text)
	}
}

func TestSeasonsCannotBeReordered(t *testing.T) {
	first := Seasons()[0]
	got := Seasons()
	got[0] = Poem{Text: "vandalism"}
	if Seasons()[0] != first {
		t.Fatalf("the year now opens on %q", Seasons()[0].Text)
	}
}

func TestEveryPoemNamesItsKigo(t *testing.T) {
	filed := map[string]string{}
	for _, p := range Seasons() {
		if p.Kigo == "" {
			t.Errorf("no season word: %q", p.Text)
		}
		if other, dup := filed[p.Kigo]; dup {
			t.Errorf("%q files both %q and %q", p.Kigo, other, p.Text)
		}
		filed[p.Kigo] = p.Text
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
