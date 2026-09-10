package curry

// Step is one instruction. Minutes is elapsed time, not attention: Stir marks
// the stretches a cook cannot walk away from.
type Step struct {
	Text    string
	Minutes int
	Stir    bool
}

// Method is the pot from cold pan to serving. The roux goes in off the heat —
// stirred into a boiling pot it seizes into lumps that never dissolve.
func Method() []Step {
	return []Step{
		{Text: "sweat the onion in oil until it slumps and turns gold", Minutes: 12, Stir: true},
		{Text: "turn the carrot and sweet potato through the oil", Minutes: 3, Stir: true},
		{Text: "pour in boiling water and bring the pot back up", Minutes: 2},
		{Text: "skim the grey foam off the surface", Minutes: 2, Stir: true},
		{Text: "cover and simmer until a skewer slides through the carrot", Minutes: 20},
		{Text: "off the heat, melt the roux in a ladle at a time", Minutes: 4, Stir: true},
		{Text: "back on the lowest flame, thicken uncovered", Minutes: 8, Stir: true},
		{Text: "rest off the heat before serving", Minutes: 10},
	}
}
