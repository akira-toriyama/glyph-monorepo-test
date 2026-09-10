package haiku

import "slices"

// Season is one volume of the saijiki. The five constants are the whole set; a
// Season a caller invented has nothing filed under it, and Of answers nil.
type Season string

const (
	NewYear Season = "new year"
	Spring  Season = "spring"
	Summer  Season = "summer"
	Autumn  Season = "autumn"
	Winter  Season = "winter"
)

// Year is saijiki order: the new year opens the book, then the four seasons in
// calendar order.
func Year() []Season {
	return []Season{NewYear, Spring, Summer, Autumn, Winter}
}

// Of hands back a fresh slice; the map behind it is package state.
func Of(s Season) []Poem {
	return slices.Clone(saijiki[s])
}

// Poems is every poem, volume by volume, each volume in its own filing order.
func Poems() []Poem {
	var out []Poem
	for _, s := range Year() {
		out = append(out, saijiki[s]...)
	}
	return out
}
