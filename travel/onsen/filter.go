package onsen

import "strings"

// Hotter returns the baths whose source is above minC, in the guide's order.
// The result is a fresh slice: filtering in place over baths[:0] writes the
// answer over the guide.
func Hotter(minC int) []Bath {
	out := make([]Bath, 0, len(baths))
	for _, b := range baths {
		if b.SourceC > minC {
			out = append(out, b)
		}
	}
	return out
}

// WithWater returns the baths whose analysis board posts kind. A board names
// the water in a phrase ("acidic sulphur"), so kind matches part of one; an
// empty kind is a caller with nothing to ask, not a request for every bath.
func WithWater(kind string) []Bath {
	want := strings.ToLower(strings.TrimSpace(kind))
	if want == "" {
		return nil
	}
	out := make([]Bath, 0, len(baths))
	for _, b := range baths {
		if strings.Contains(strings.ToLower(b.Water), want) {
			out = append(out, b)
		}
	}
	return out
}
