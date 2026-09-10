package curry

// Recipe is written for this many plates; every other pot is a ratio of it.
const baseServings = 4

// For weighs the recipe out for plates. Prep travels unchanged: a 3cm cube is
// a 3cm cube whether the pot feeds two or twelve.
func For(plates int) []Ingredient {
	base := Recipe()
	scaled := make([]Ingredient, len(base))
	for i, ing := range base {
		scaled[i] = ing
		scaled[i].Grams = ing.Grams * plates / baseServings
	}
	return scaled
}
