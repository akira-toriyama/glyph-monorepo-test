package curry

import "time"

// Step is one instruction. Minutes is elapsed time, not attention: Stir marks
// the stretches a cook cannot walk away from.
type Step struct {
	Text    string
	Minutes int
	Stir    bool
}

// Method is the pot from cold pan to serving. Two orderings are load-bearing:
// the spices meet hot oil before any water, and the roux goes in off the heat —
// stirred into a boiling pot it seizes into lumps that never dissolve.
func Method() []Step {
	return []Step{
		{Text: "sweat the onion in oil until it slumps and turns gold", Minutes: 12, Stir: true},
		{Text: "bloom the ground spices in the hot oil until they smell of more than dust", Minutes: 1, Stir: true},
		{Text: "turn the carrot and sweet potato through the oil", Minutes: 3, Stir: true},
		{Text: "pour in boiling water and bring the pot back up", Minutes: 2},
		{Text: "skim the grey foam off the surface", Minutes: 2, Stir: true},
		{Text: "cover and simmer until a skewer slides through the carrot", Minutes: 20},
		{Text: "off the heat, melt the roux in a ladle at a time", Minutes: 4, Stir: true},
		{Text: "back on the lowest flame, thicken uncovered", Minutes: 8, Stir: true},
		{Text: "rest off the heat before serving", Minutes: 10},
	}
}

// Duration counts the rest as well as the cooking: a cook plans backwards from
// the table, and ten minutes off the heat is still ten minutes.
func Duration() time.Duration {
	var d time.Duration
	for _, s := range Method() {
		d += time.Duration(s.Minutes) * time.Minute
	}
	return d
}

// HandsOn is the share of Duration the kitchen cannot be left for.
func HandsOn() time.Duration {
	var d time.Duration
	for _, s := range Method() {
		if s.Stir {
			d += time.Duration(s.Minutes) * time.Minute
		}
	}
	return d
}

// Reheat is the day-two pot, which is the one worth making a big batch for:
// the starch keeps swelling overnight, so the sauce comes out of the fridge
// thicker than it went in. Never at a boil — roux is an emulsion, and a hard
// boil splits its fat back out onto the surface.
func Reheat() []Step {
	return []Step{
		{Text: "loosen the cold pot with a splash of water", Minutes: 1, Stir: true},
		{Text: "cover and warm it through on the lowest flame", Minutes: 12},
		{Text: "taste for salt — a night in the fridge flattens it", Minutes: 2, Stir: true},
	}
}
