package haiku

// saijiki files every poem under the volume its season word belongs to. Read
// it through Of or Poems, both of which copy: a caller handed these slices
// could reorder the book for everyone.
var saijiki = map[Season][]Poem{
	NewYear: {
		{
			Text:   "in my first dream\nI saw my home village —\nand I wept",
			Author: "Kobayashi Issa",
			Kigo:   "first dream",
		},
	},
	Spring: {
		{
			Text:   "the spring sea\nrising and falling, rising\nand falling all day",
			Author: "Yosa Buson",
			Kigo:   "the spring sea",
		},
	},
	Summer: {
		{
			Text:   "the summer grasses —\nall that remains\nof the warriors' dreams",
			Author: "Matsuo Bashō",
			Kigo:   "summer grasses",
		},
	},
	Autumn: {
		{
			Text:   "an old silent pond\na frog jumps into the pond —\nsplash! silence again",
			Author: "Matsuo Bashō",
			Kigo:   "frog",
		},
		{
			Text:   "on a bare branch\na crow has alighted —\nautumn nightfall",
			Author: "Matsuo Bashō",
			Kigo:   "autumn nightfall",
		},
	},
	Winter: {
		{
			Text:   "first winter rain—\neven the monkey seems to want\na little straw coat",
			Author: "Matsuo Bashō",
			Kigo:   "first winter rain",
		},
	},
}
