// Package travel is the third independently versioned module of the harness.
// Its version line is the tag prefix travel/ — and travel/onsen is a nested
// line of its own, so a file under travel/onsen/ belongs to onsen, not to
// travel (the longest declared prefix wins).
package travel

import "slices"

// Stop is one town on the itinerary. Nights is what you sleep there, so a day
// trip is a stop with zero and the bed that night belongs to the stop before
// it.
type Stop struct {
	Name       string
	Prefecture string
	Nights     int
}

// route is the itinerary in visiting order.
var route = []Stop{
	{Name: "Kyoto", Prefecture: "Kyoto", Nights: 3},
	{Name: "Nara", Prefecture: "Nara", Nights: 1},
	{Name: "Kobe", Prefecture: "Osaka", Nights: 1},
	{Name: "Osaka", Prefecture: "Osaka", Nights: 2},
}

// Route is the itinerary; a new stop lands as a minor bump of this module.
// The result is the caller's to sort, reverse or truncate.
func Route() []Stop {
	return slices.Clone(route)
}

// Duration counts nights, not days: the trip spans one more calendar day than
// this.
func Duration() int {
	total := 0
	for _, stop := range route {
		total += stop.Nights
	}
	return total
}
