package travel

// Leg is the rail hop from one stop to the next. Minutes is scheduled time on
// the train alone — no transfer, no wait for the next departure.
type Leg struct {
	From    string
	To      string
	Line    string
	Minutes int
}

// hops is the rail table, keyed by "<from>>><to>" so any itinerary over the
// same towns reads the same services. A pair missing here is not an error: the
// leg is still reported, with no Line and no Minutes.
var hops = map[string]Leg{
	"Kyoto>>>Nara": {Line: "Kintetsu limited express", Minutes: 35},
	"Nara>>>Kobe":  {Line: "Hanshin Namba Line through service", Minutes: 80},
	"Kobe>>>Osaka": {Line: "JR special rapid", Minutes: 21},
}

// Legs walks the itinerary and looks up the service between each pair of stops.
func (it Itinerary) Legs() []Leg {
	out := make([]Leg, 0, max(len(it.Stops)-1, 0))
	for i := 0; i+1 < len(it.Stops); i++ {
		from, to := it.Stops[i].Name, it.Stops[i+1].Name
		leg := hops[from+">>>"+to]
		leg.From, leg.To = from, to
		out = append(out, leg)
	}
	return out
}

// TravelTime is time on trains only: platform transfers and the wait for the
// next departure are nobody's schedule but the traveller's.
func (it Itinerary) TravelTime() int {
	total := 0
	for _, leg := range it.Legs() {
		total += leg.Minutes
	}
	return total
}
