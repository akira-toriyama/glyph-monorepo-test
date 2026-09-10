// Package travel is the third independently versioned module of the harness.
// Its version line is the tag prefix travel/ — and travel/onsen is a nested
// line of its own, so a file under travel/onsen/ belongs to onsen, not to
// travel (the longest declared prefix wins).
package travel

import "slices"

// route is the itinerary in visiting order; nights is keyed by the same names,
// so a stop added to one has to land in the other.
var route = []string{"Kyoto", "Nara", "Kobe", "Osaka"}

var nights = map[string]int{
	"Kyoto": 3,
	"Nara":  1,
	"Kobe":  1,
	"Osaka": 2,
}

// Route is the itinerary; a new stop lands as a minor bump of this module.
// The result is the caller's to sort, reverse or truncate.
func Route() []string {
	return slices.Clone(route)
}

// Nights answers 0 both for a stop nobody sleeps at and for a name that was
// never on the route.
func Nights(stop string) int {
	return nights[stop]
}

// Duration counts nights, not days: the trip spans one more calendar day than
// this.
func Duration() int {
	total := 0
	for _, stop := range route {
		total += nights[stop]
	}
	return total
}
