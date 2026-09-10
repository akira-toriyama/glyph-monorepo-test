package curry

// Rice is deliberately outside Recipe: it cooks in its own pot, on its own
// clock, and a cook with a rice cooker already has this handled.
type Rice struct {
	DryGrams    int
	WaterGrams  int
	SoakMinutes int
}

// RiceFor weighs the rice for plates, and returns the zero Rice below one —
// the same refusal For makes. 90g of dry short-grain a plate is a scale
// reading, not a plastic measure: a rice-cooker cup is 150g and feeds closer
// to two. The water is the SOAKED figure: grain that has drunk for half an
// hour needs 1.1 times its dry weight, and the 1.2 on the bag assumes it goes
// in dry.
func RiceFor(plates int) Rice {
	if plates < 1 {
		return Rice{}
	}
	dry := 90 * plates
	return Rice{
		DryGrams:    dry,
		WaterGrams:  dry * 11 / 10,
		SoakMinutes: 30,
	}
}
