// Package camp is the fourth independently versioned module of the harness.
// Its version line is the tag prefix camp/.
package camp

import "slices"

// Item is one line of the packing list. Grams is what actually leaves the
// house: poles, pegs and the stuff sack are already inside it.
type Item struct {
	Name     string
	Grams    int
	Category Category
}

// list is packing order — what goes into the pack first is listed first. A new
// item lands as a minor bump of this module, a swapped one as a patch.
var list = []Item{
	{Name: "tent", Grams: 2400, Category: Shelter},
	{Name: "sleeping mat", Grams: 480, Category: Sleep},
	{Name: "sleeping bag", Grams: 900, Category: Sleep},
	{Name: "gas stove", Grams: 210, Category: Kitchen},
	{Name: "pot", Grams: 340, Category: Kitchen},
	{Name: "mug", Grams: 120, Category: Kitchen},
	{Name: "water filter", Grams: 350, Category: Kitchen},
}

// Gear is the whole list, shelter through kitchen. The result is a copy: list
// is package state, and a caller that sorts or truncates what it gets back
// would otherwise repack the list for everyone else.
func Gear() []Item {
	return slices.Clone(list)
}
