package onsen

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
