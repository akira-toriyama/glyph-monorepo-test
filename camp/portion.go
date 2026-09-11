package camp

// Portion is a serving size in grams.
type Portion int

// Halve returns half the portion, rounded down.
func (p Portion) Halve() Portion { return p / 2 }
