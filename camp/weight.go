package camp

// PackWeight totals the list. Item.Grams is a packed weight, so the sum is what
// the pack reads on the scale before food and water go in.
func PackWeight() int {
	total := 0
	for _, item := range Gear() {
		total += item.Grams
	}
	return total
}
