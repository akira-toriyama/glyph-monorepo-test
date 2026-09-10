package haiku

// Spring is Buson's spring sea.
func Spring() string {
	return "the spring sea\nrising and falling, rising\nand falling all day"
}

// Summer is Bashō's summer grasses.
func Summer() string {
	return "the summer grasses —\nall that remains\nof the warriors' dreams"
}

// Autumn is Bashō's old pond, carried from the bootstrap commit unchanged.
func Autumn() string {
	return "an old silent pond\na frog jumps into the pond —\nsplash! silence again"
}

// Winter is Bashō's first winter rain.
func Winter() string {
	return "first winter rain—\neven the monkey seems to want\na little straw coat"
}

// Seasons hands back the year in calendar order, which is not the order the
// seasons were added in.
func Seasons() []string {
	return []string{Spring(), Summer(), Autumn(), Winter()}
}
