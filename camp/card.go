package camp

import (
	"fmt"
	"strings"
)

// Card renders the list the way it is read at the door: packing order, one
// item per line, the total last. The name column is padded to the widest name
// so the weights line up under each other.
func Card() string {
	items := Gear()
	width := len("total")
	for _, item := range items {
		if len(item.Name) > width {
			width = len(item.Name)
		}
	}
	var b strings.Builder
	for _, item := range items {
		fmt.Fprintf(&b, "%-*s  %s\n", width, item.Name, weight(item.Grams))
	}
	fmt.Fprintf(&b, "%-*s  %s\n", width, "total", weight(PackWeight()))
	return b.String()
}

// weight is the card's only formatting rule; the card is the whole of this
// module's user-facing text. Kilograms once an item reaches one, whole grams
// below it — a stove at "0.2 kg" tells the reader nothing it can pack around.
func weight(g int) string {
	if g >= 1000 {
		return fmt.Sprintf("%.1f kg", float64(g)/1000)
	}
	return fmt.Sprintf("%d g", g)
}
