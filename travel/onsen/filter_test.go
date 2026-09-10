package onsen

import (
	"slices"
	"testing"
)

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

func TestHotterLeavesTheGuideAlone(t *testing.T) {
	before := All()
	if got := bathNames(Hotter(55)); len(got) != 1 {
		t.Fatalf("Hotter(55) = %v, want one bath", got)
	}
	if after := All(); !slices.Equal(before, after) {
		t.Fatalf("the guide reads %v after one filter, want %v",
			bathNames(after), bathNames(before))
	}
}

func TestWithWaterGathersOneKind(t *testing.T) {
	if got := bathNames(WithWater("acidic sulphur")); len(got) != 1 || got[0] != "Kusatsu" {
		t.Errorf(`WithWater("acidic sulphur") = %v, want [Kusatsu]`, got)
	}
	if got := bathNames(WithWater("brine")); len(got) != 0 {
		t.Errorf(`WithWater("brine") = %v, want nothing`, got)
	}
}

func TestWithWaterReadsPartOfThePhrase(t *testing.T) {
	got := bathNames(WithWater("sulphur"))
	slices.Sort(got)
	if want := []string{"Kusatsu", "Noboribetsu"}; !slices.Equal(got, want) {
		t.Errorf(`WithWater("sulphur") = %v, want %v`, got, want)
	}
	if got := bathNames(WithWater("")); len(got) != 0 {
		t.Errorf(`WithWater("") = %v, want nothing`, got)
	}
}
