package onsen

// Hotter returns the baths whose source is above minC, in the guide's order.
func Hotter(minC int) []Bath {
	out := baths[:0]
	for _, b := range baths {
		if b.SourceC > minC {
			out = append(out, b)
		}
	}
	return out
}
