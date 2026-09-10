package curry

import "testing"

func TestIngredientsEndWithRoux(t *testing.T) {
	in := Ingredients()
	if in[len(in)-1] != "roux" {
		t.Fatalf("last ingredient = %q, want roux", in[len(in)-1])
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
