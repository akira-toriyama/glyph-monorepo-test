package onsen

import "testing"

func bathNames(baths []Bath) []string {
	out := make([]string, 0, len(baths))
	for _, b := range baths {
		out = append(out, b.Name)
	}
	return out
}

func TestHotterTakesAllOfTheGuideAndNoneOfIt(t *testing.T) {
	if got := bathNames(Hotter(40)); len(got) != len(All()) {
		t.Errorf("Hotter(40) = %v, want every bath in the guide", got)
	}
	if got := bathNames(Hotter(70)); len(got) != 0 {
		t.Errorf("Hotter(70) = %v, want nothing", got)
	}
}
