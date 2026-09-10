package camp

// grams carries one entry per name Gear returns; TestEveryItemIsWeighed holds
// the two in step, which is why PackWeight may read a miss as zero.
var grams = map[string]int{
	"tent":         1900,
	"sleeping bag": 900,
	"gas stove":    210,
}

// WeightGrams reports what an item adds to the pack. ok is false for a name
// that is not on the list.
func WeightGrams(item string) (g int, ok bool) {
	g, ok = grams[item]
	return g, ok
}

// PackWeight totals the list. The weights are packed weights, so the sum is
// what the pack reads on the scale before food and water go in.
func PackWeight() int {
	total := 0
	for _, item := range Gear() {
		total += grams[item]
	}
	return total
}
