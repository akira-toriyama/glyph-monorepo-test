package curry

// Recipe is written for this many plates; every other pot is a ratio of it.
const baseServings = 4

// For weighs the recipe out for plates, and returns nil below one: integer
// scaling is happy to hand back negative weights, and there is no pot for zero.
// Prep travels unchanged — a 3cm cube is a 3cm cube whether the pot feeds two
// or twelve.
func For(plates int) []Ingredient {
	if plates < 1 {
		return nil
	}
	base := Recipe()
	scaled := make([]Ingredient, len(base))
	for i, ing := range base {
		scaled[i] = ing
		// Round up. Truncating divided 2g of cayenne to nothing at one plate,
		// and a spice that rounds away is a different curry.
		scaled[i].Grams = (ing.Grams*plates + baseServings - 1) / baseServings
	}
	return scaled
}
