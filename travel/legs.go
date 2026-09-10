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
	"Kyoto>>>Nara":  {Line: "Kintetsu limited express", Minutes: 35},
	"Nara>>>Kobe":   {Line: "Kintetsu Nara Line", Minutes: 80},
	"Kobe>>>Osaka":  {Line: "JR special rapid", Minutes: 21},
	"Osaka>>>Kyoto": {Line: "JR special rapid", Minutes: 28},
}

// Legs walks the itinerary and looks up the service between each pair of stops.
func (it Itinerary) Legs() []Leg {
	out := make([]Leg, 0, len(it.Stops))
	for i := range it.Stops {
		from, to := it.Stops[i].Name, it.Stops[(i+1)%len(it.Stops)].Name
		leg := hops[from+">>>"+to]
		leg.From, leg.To = from, to
		out = append(out, leg)
	}
	return out
}
