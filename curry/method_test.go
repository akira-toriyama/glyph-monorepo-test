package curry

import (
	"strings"
	"testing"
)

// Roux stirred into a boiling pot seizes and never dissolves, so the step that
// names it has to name the heat too.
func TestRouxMeltsOffTheHeat(t *testing.T) {
	for _, s := range Method() {
		if strings.Contains(s.Text, "roux") && !strings.Contains(s.Text, "off the heat") {
			t.Fatalf("roux step %q leaves the pot on the heat", s.Text)
		}
	}
}

func TestEveryStepTakesTime(t *testing.T) {
	for i, s := range Method() {
		if s.Minutes <= 0 {
			t.Errorf("step %d %q takes %d minutes", i, s.Text, s.Minutes)
		}
	}
}

// Foam is protein that has nowhere to go once the lid is on; it settles back
// into the sauce and turns it grey.
func TestSkimComesBeforeTheLid(t *testing.T) {
	skim, lid := -1, -1
	for i, s := range Method() {
		switch {
		case strings.HasPrefix(s.Text, "skim"):
			skim = i
		case strings.HasPrefix(s.Text, "cover"):
			lid = i
		}
	}
	if skim < 0 || lid < 0 {
		t.Fatalf("the method lost the skim (%d) or the lid (%d)", skim, lid)
	}
	if skim > lid {
		t.Errorf("skim at step %d, lid at step %d", skim, lid)
	}
}

// Cold water off a finished sweat drops the pot under a simmer and the sweat
// has to be paid for twice; the step is a top-up, not a restart.
func TestTheWaterGoesInBoiling(t *testing.T) {
	for _, s := range Method() {
		if strings.Contains(s.Text, "water") && !strings.Contains(s.Text, "boiling") {
			t.Fatalf("water step %q pours it in cold", s.Text)
		}
	}
}
