// Package travel is the third independently versioned module of the harness.
// Its version line is the tag prefix travel/ — and travel/onsen is a nested
// line of its own, so a file under travel/onsen/ belongs to onsen, not to
// travel (the longest declared prefix wins).
package travel

// Route is the itinerary; a new stop lands as a minor bump of this module.
func Route() []string {
	return []string{"Kyoto", "Nara", "Osaka"}
}
