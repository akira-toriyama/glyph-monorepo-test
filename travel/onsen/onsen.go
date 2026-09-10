// Package onsen is a version line NESTED inside travel/: its tags are
// travel/onsen/vX.Y.Z, and a change here moves onsen alone — travel's line
// does not hear of it. That is the longest-prefix rule the harness proves.
//
// From v1 on the surface is stable: Bath's fields keep their meaning, a new
// field is a minor, and dropping a reader is a major. The guide's contents are
// not part of that promise — a spring is corrected in a patch.
package onsen

import "slices"

// Bath is one hot spring. Water is the 泉質 phrase a bathhouse posts on its
// analysis board, not a chemical formula. SourceC is the temperature at the
// spring head; the tub is cooler, sometimes by twenty degrees.
type Bath struct {
	Name       string
	Prefecture string
	Water      string
	SourceC    int
}

// baths is the guide itself, kept hottest first — nothing sorts it, so a row
// whose temperature is corrected has to move. Adding a spring here is the
// minor bump this line exists to demonstrate.
var baths = []Bath{
	{Name: "Beppu", Prefecture: "Oita", Water: "sodium chloride", SourceC: 60},
	{Name: "Noboribetsu", Prefecture: "Hokkaido", Water: "sulphur", SourceC: 45},
	{Name: "Kusatsu", Prefecture: "Gunma", Water: "acidic sulphur", SourceC: 51},
}

// All returns a copy: the guide is package state, and a caller that sorts what
// it gets must not reorder it for the next caller.
func All() []Bath {
	return slices.Clone(baths)
}
