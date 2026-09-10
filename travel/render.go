package travel

import (
	"fmt"
	"strings"
)

// String is the sheet that goes in the bag: one row per stop, the ride printed
// beside the stop it belongs to. Widths are padded for the four Kansai towns,
// so a longer name pushes its own row and nothing else.
func (it Itinerary) String() string {
	legs := it.Legs()
	var b strings.Builder
	for i, stop := range it.Stops {
		fmt.Fprintf(&b, "%-6s %-8s %-9s", stop.Name, "("+stop.Prefecture+")", stay(stop.Nights))
		if i < len(legs) {
			fmt.Fprintf(&b, " in on the %s, %d min", legs[i].Line, legs[i].Minutes)
		}
		b.WriteString("\n")
	}
	return b.String()
}

func stay(nights int) string {
	switch nights {
	case 0:
		return "day trip"
	case 1:
		return "1 night"
	default:
		return fmt.Sprintf("%d nights", nights)
	}
}
