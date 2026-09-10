// Package travel is the third independently versioned module of the harness.
// Its version line is the tag prefix travel/ — and travel/onsen is a nested
// line of its own, so a file under travel/onsen/ belongs to onsen, not to
// travel (the longest declared prefix wins).
package travel

// Stop is one town on the itinerary. Nights is what you sleep there, so a day
// trip is a stop with zero and the bed that night belongs to the stop before
// it.
type Stop struct {
	Name       string
	Prefecture string
	Nights     int
}

// Itinerary is a trip in visiting order. It is the value the package hands out
// and the receiver every query hangs off, so a second itinerary needs no
// second set of package-level functions.
type Itinerary struct {
	Stops []Stop
}

// Kansai builds its stops fresh on every call: the itinerary it returns,
// slice included, belongs to the caller.
func Kansai() Itinerary {
	return Itinerary{Stops: []Stop{
		{Name: "Kyoto", Prefecture: "Kyoto", Nights: 4},
		{Name: "Nara", Prefecture: "Nara", Nights: 0},
		{Name: "Kobe", Prefecture: "Hyogo", Nights: 1},
		{Name: "Osaka", Prefecture: "Osaka", Nights: 2},
	}}
}

// Duration counts nights, not days: the trip spans one more calendar day than
// this.
func (it Itinerary) Duration() int {
	total := 0
	for _, stop := range it.Stops {
		total += stop.Nights
	}
	return total
}

// SleepTowns is the town of each night in order, one entry per night: a stop
// with four nights repeats four times and a day trip appears not at all.
func (it Itinerary) SleepTowns() []string {
	out := make([]string, 0, it.Duration())
	for _, stop := range it.Stops {
		for range stop.Nights {
			out = append(out, stop.Name)
		}
	}
	return out
}
