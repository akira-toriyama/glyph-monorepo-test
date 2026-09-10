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
