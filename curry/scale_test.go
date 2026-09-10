package curry

import "testing"

func TestForTheBasePotIsTheRecipe(t *testing.T) {
	base := Recipe()
	for i, ing := range For(baseServings) {
		if ing != base[i] {
			t.Errorf("For(%d)[%d] = %+v, want %+v", baseServings, i, ing, base[i])
		}
	}
}

func TestForDoublesEveryWeight(t *testing.T) {
	base := Recipe()
	for i, ing := range For(2 * baseServings) {
		if want := 2 * base[i].Grams; ing.Grams != want {
			t.Errorf("%s: %dg for a double pot, want %dg", ing.Name, ing.Grams, want)
		}
	}
}

// Scaling must not reword the recipe; only the numbers move.
func TestForKeepsThePrep(t *testing.T) {
	base := Recipe()
	for i, ing := range For(9) {
		if ing.Prep != base[i].Prep || ing.Name != base[i].Name {
			t.Errorf("For(9)[%d] = %s/%q, want %s/%q", i, ing.Name, ing.Prep, base[i].Name, base[i].Prep)
		}
	}
}

func TestForRefusesAPotBelowOnePlate(t *testing.T) {
	for _, plates := range []int{0, -1, -2} {
		if got := For(plates); got != nil {
			t.Errorf("For(%d) = %+v, want nil", plates, got)
		}
	}
}

// The small pot is where the blend disappears; cayenne is the first to go.
func TestNoIngredientScalesAwayToNothing(t *testing.T) {
	for plates := 1; plates <= 12; plates++ {
		for _, ing := range For(plates) {
			if ing.Grams < 1 {
				t.Errorf("For(%d): %s came out at %dg", plates, ing.Name, ing.Grams)
			}
		}
	}
}
