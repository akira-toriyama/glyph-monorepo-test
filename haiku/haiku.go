package haiku

import "slices"

var (
	newYear = Poem{
		Text:   "in my first dream\nI saw my home village —\nand I wept",
		Author: "Kobayashi Issa",
		Kigo:   "first dream",
	}

	spring = Poem{
		Text:   "the spring sea\nrising and falling, rising\nand falling all day",
		Author: "Yosa Buson",
		Kigo:   "the spring sea",
	}

	summer = Poem{
		Text:   "the summer grasses —\nall that remains\nof the warriors' dreams",
		Author: "Matsuo Bashō",
		Kigo:   "grass",
	}

	autumn = Poem{
		Text:   "an old silent pond\na frog jumps into the pond —\nsplash! silence again",
		Author: "Matsuo Bashō",
		Kigo:   "frog",
	}

	winter = Poem{
		Text:   "first winter rain—\neven the monkey seems to want\na little straw coat",
		Author: "Matsuo Bashō",
		Kigo:   "first winter rain",
	}
)

// year is saijiki order: the new year opens the book, then the four seasons in
// calendar order.
var year = []Poem{newYear, spring, summer, autumn, winter}

// NewYear is the fifth volume of a saijiki: it stands beside the four seasons
// rather than inside winter.
func NewYear() Poem { return newYear }

func Spring() Poem { return spring }

func Summer() Poem { return summer }

func Autumn() Poem { return autumn }

func Winter() Poem { return winter }

// Seasons hands back a fresh slice every call: year is package state, and a
// caller that sorted or appended to the result would otherwise reorder the
// year for every caller after it.
func Seasons() []Poem { return slices.Clone(year) }
