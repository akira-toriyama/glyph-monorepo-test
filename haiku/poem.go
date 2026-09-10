package haiku

// Poem is one entry of the saijiki. Text is the three lines joined with "\n";
// Author is the poet's name surname first, the order Japanese writes it in and
// the form every entry here uses.
type Poem struct {
	Text   string
	Author string
}
