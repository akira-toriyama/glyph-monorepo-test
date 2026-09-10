package haiku

import (
	"strings"
	"testing"
)

func TestYearOpensOnTheNewYear(t *testing.T) {
	if got := Year()[0]; got != NewYear {
		t.Fatalf("the book opens on %q", got)
	}
}

func TestEveryVolumeHoldsAPoem(t *testing.T) {
	for _, s := range Year() {
		if len(Of(s)) == 0 {
			t.Errorf("%s: nothing filed", s)
		}
	}
}

func TestOfMissesASeasonNobodyNamed(t *testing.T) {
	if got := Of("monsoon"); got != nil {
		t.Fatalf("the monsoon is no season here, but it holds %d poems", len(got))
	}
}

func TestAVolumeCannotBeReordered(t *testing.T) {
	first := Of(Spring)[0]
	got := Of(Spring)
	got[0] = Poem{Text: "vandalism"}
	if Of(Spring)[0] != first {
		t.Fatalf("spring now opens on %q", Of(Spring)[0].Text)
	}
}

func TestSomeVolumeHoldsMoreThanOnePoem(t *testing.T) {
	for _, s := range Year() {
		if len(Of(s)) > 1 {
			return
		}
	}
	t.Fatal("every volume still holds exactly one poem")
}

func TestAutumnKeepsTheCrow(t *testing.T) {
	if !holdsKigo(Of(Autumn), "autumn nightfall") {
		t.Fatal("the crow is filed under no season")
	}
}

func TestTheFrogIsFiledUnderSpring(t *testing.T) {
	if !holdsKigo(Of(Spring), "frog") {
		t.Error("spring does not hold the old pond")
	}
	if holdsKigo(Of(Autumn), "frog") {
		t.Error("autumn still holds the old pond")
	}
}

func holdsKigo(volume []Poem, kigo string) bool {
	for _, p := range volume {
		if p.Kigo == kigo {
			return true
		}
	}
	return false
}

func TestEveryPoemIsThreeLines(t *testing.T) {
	for _, p := range Poems() {
		if got := len(strings.Split(p.Text, "\n")); got != 3 {
			t.Errorf("%s: lines = %d, want 3", p.Author, got)
		}
	}
}

func TestEveryPoemNamesItsPoetAndKigo(t *testing.T) {
	filed := map[string]string{}
	for _, p := range Poems() {
		if p.Author == "" || p.Kigo == "" {
			t.Errorf("incomplete entry: %q", p.Text)
		}
		if other, dup := filed[p.Kigo]; dup {
			t.Errorf("%q files both %q and %q", p.Kigo, other, p.Text)
		}
		filed[p.Kigo] = p.Text
	}
}
