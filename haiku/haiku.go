package haiku

var (
	spring = Poem{
		Text:   "the spring sea\nrising and falling, rising\nand falling all day",
		Author: "Yosa Buson",
	}

	summer = Poem{
		Text:   "the summer grasses —\nall that remains\nof the warriors' dreams",
		Author: "Matsuo Bashō",
	}

	autumn = Poem{
		Text:   "an old silent pond\na frog jumps into the pond —\nsplash! silence again",
		Author: "Matsuo Bashō",
	}

	winter = Poem{
		Text:   "first winter rain—\neven the monkey seems to want\na little straw coat",
		Author: "Matsuo Bashō",
	}

	newYear = Poem{
		Text:   "in my first dream\nI saw my home village —\nand I wept",
		Author: "Kobayashi Issa",
	}
)

// year is calendar order, which is not the order the poems were added in.
var year = []Poem{spring, summer, autumn, winter, newYear}

func Spring() Poem { return spring }

func Summer() Poem { return summer }

func Autumn() Poem { return autumn }

func Winter() Poem { return winter }

// NewYear is the fifth volume of a saijiki: it stands beside the four seasons
// rather than inside winter.
func NewYear() Poem { return newYear }

// Seasons hands back the whole year in one call.
func Seasons() []Poem { return year }
