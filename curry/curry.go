// Package curry is the second independently versioned module of the harness.
// Its version line is the tag prefix curry/ — a change here must never move
// haiku's version, which is the property the harness exists to prove.
package curry

// Ingredient is one line of the pot. Grams is the only unit here: cups measure
// nothing reproducible, and roux boxes have changed block size twice.
type Ingredient struct {
	Name  string
	Grams int
	Prep  string
}

// Recipe is four plates, listed in the order the ingredients go into the pot.
// Water is an ingredient because the roux is weighed against it, not guessed.
func Recipe() []Ingredient {
	return []Ingredient{
		{Name: "onion", Grams: 400, Prep: "sliced along the grain"},
		{Name: "carrot", Grams: 200, Prep: "rolling-cut"},
		{Name: "sweet potato", Grams: 300, Prep: "3cm cubes"},
		{Name: "water", Grams: 800},
		{Name: "roux", Grams: 120, Prep: "broken into squares"},
	}
}

// Ingredients is the shopping view, for callers with no scale in the kitchen.
func Ingredients() []string {
	r := Recipe()
	names := make([]string, len(r))
	for i, ing := range r {
		names[i] = ing.Name
	}
	return names
}
