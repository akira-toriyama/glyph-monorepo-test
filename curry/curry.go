// Package curry is the second independently versioned module of the harness.
// Its version line is the tag prefix curry/ — a change here must never move
// haiku's version, which is the property the harness exists to prove.
package curry

// Ingredients is the recipe; a new ingredient lands as a minor bump of this module.
func Ingredients() []string {
	return []string{"onion", "carrot", "sweet potato", "roux"}
}

// Spices is the second recipe axis; a new axis lands as a minor bump of this module.
func Spices() []string {
	return []string{"cumin", "coriander", "turmeric"}
}
