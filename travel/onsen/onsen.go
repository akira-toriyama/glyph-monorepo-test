// Package onsen is a version line NESTED inside travel/: its tags are
// travel/onsen/vX.Y.Z, and a change here moves onsen alone — travel's line
// does not hear of it. That is the longest-prefix rule the harness proves.
package onsen

// Bath is one hot spring. Water is the 泉質 phrase a bathhouse posts on its
// analysis board, not a chemical formula.
type Bath struct {
	Name       string
	Prefecture string
	Water      string
}

// baths is the guide itself. Adding a spring here is the minor bump this line
// exists to demonstrate.
var baths = []Bath{
	{Name: "Kusatsu", Prefecture: "Gunma", Water: "acidic sulphur"},
	{Name: "Beppu", Prefecture: "Oita", Water: "simple alkaline"},
	{Name: "Noboribetsu", Prefecture: "Aomori", Water: "sulphur"},
}

// All returns the guide in printing order.
func All() []Bath {
	return baths
}

// Baths lists the names alone, the shape the guide had before a spring was
// more than a word.
func Baths() []string {
	names := make([]string, 0, len(baths))
	for _, b := range baths {
		names = append(names, b.Name)
	}
	return names
}
