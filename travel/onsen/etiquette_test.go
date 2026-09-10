package onsen

import (
	"strings"
	"testing"
)

func TestEtiquetteWashesBeforeItRinses(t *testing.T) {
	rules := Etiquette()
	if len(rules) != 4 {
		t.Fatalf("Etiquette posts %d rules, want 4", len(rules))
	}
	wash, rinse := -1, -1
	for i, r := range rules {
		switch {
		case strings.Contains(r, "wash"):
			wash = i
		case strings.Contains(r, "rinse"):
			rinse = i
		}
	}
	if wash < 0 || rinse < 0 || wash > rinse {
		t.Fatalf("wash is rule %d and rinse is rule %d, in %v", wash+1, rinse+1, rules)
	}
}
