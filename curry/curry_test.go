package curry

import "testing"

func TestRecipeEndsWithRoux(t *testing.T) {
	r := Recipe()
	if last := r[len(r)-1].Name; last != "roux" {
		t.Fatalf("last ingredient = %q, want roux", last)
	}
}

func TestEveryIngredientIsWeighed(t *testing.T) {
	for _, ing := range Recipe() {
		if ing.Grams <= 0 {
			t.Errorf("%s: grams = %d, want a positive weight", ing.Name, ing.Grams)
		}
	}
}

// Below roughly one part roux to seven of water the sauce will not coat rice.
// The box's own table says 1:6.7; the seed recipe shipped at 1:8.
func TestRouxCarriesTheWater(t *testing.T) {
	var water, roux int
	for _, ing := range Recipe() {
		switch ing.Name {
		case "water":
			water = ing.Grams
		case "roux":
			roux = ing.Grams
		}
	}
	if roux*7 < water {
		t.Fatalf("%dg of roux against %dg of water is thinner than 1:7", roux, water)
	}
}
