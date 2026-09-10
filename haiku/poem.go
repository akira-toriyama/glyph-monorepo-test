package haiku

// Poem is one entry of the saijiki. Text is the three lines joined with "\n";
// Author is the poet's name surname first, the order Japanese writes it in and
// the form every entry here uses.
type Poem struct {
	Text   string
	Author string

	// Kigo is the season word a saijiki files the poem under, in English and
	// as one phrase: it is the index key, not a description of the poem.
	Kigo string
}
