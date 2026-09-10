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
