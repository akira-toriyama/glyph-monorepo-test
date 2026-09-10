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
// The four ground spices are single-figure weights on purpose: they are what
// separates this from the sauce that comes out of the box alone.
func Recipe() []Ingredient {
	return []Ingredient{
		{Name: "onion", Grams: 400, Prep: "sliced along the grain"},
		{Name: "carrot", Grams: 200, Prep: "rolling-cut"},
		{Name: "sweet potato", Grams: 300, Prep: "3cm cubes"},
		{Name: "water", Grams: 800},
		{Name: "cumin", Grams: 6, Prep: "ground"},
		{Name: "coriander", Grams: 8, Prep: "ground"},
		{Name: "turmeric", Grams: 4, Prep: "ground"},
		{Name: "cayenne", Grams: 2, Prep: "ground"},
		{Name: "roux", Grams: 120, Prep: "broken into squares"},
	}
}
