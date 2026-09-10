package curry

import (
	"strings"
	"testing"
	"time"
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

// The point of the split: a pot with no long unattended stretch cannot be
// started while the guests are already in the house.
func TestThePotCooksAloneForALongStretch(t *testing.T) {
	if alone := Duration() - HandsOn(); alone < 20*time.Minute {
		t.Fatalf("unattended %s of %s", alone, Duration())
	}
}

func TestDurationCountsTheRest(t *testing.T) {
	var upToTheLastStir time.Duration
	for _, s := range Method() {
		if !strings.HasPrefix(s.Text, "rest") {
			upToTheLastStir += time.Duration(s.Minutes) * time.Minute
		}
	}
	if Duration() <= upToTheLastStir {
		t.Fatalf("Duration %s stops at the last stir", Duration())
	}
}

// Ground spice is fat-soluble and tastes of dust until it hits hot oil; once
// the water is in, blooming it is no longer possible.
func TestTheSpicesBloomBeforeTheWater(t *testing.T) {
	bloom, water := -1, -1
	for i, s := range Method() {
		switch {
		case strings.HasPrefix(s.Text, "bloom"):
			bloom = i
		case strings.HasPrefix(s.Text, "pour"):
			water = i
		}
	}
	if bloom < 0 || water < 0 {
		t.Fatalf("the method lost the bloom (%d) or the water (%d)", bloom, water)
	}
	if bloom > water {
		t.Errorf("bloom at step %d, water at step %d", bloom, water)
	}
}

// Day two has to be quick or the big batch has no point.
func TestReheatIsQuickerThanTheCook(t *testing.T) {
	var d time.Duration
	for _, s := range Reheat() {
		if s.Minutes <= 0 {
			t.Errorf("reheat step %q takes %d minutes", s.Text, s.Minutes)
		}
		d += time.Duration(s.Minutes) * time.Minute
	}
	if d >= Duration() {
		t.Fatalf("reheating costs %s against a cook of %s", d, Duration())
	}
}
