package haiku

// ByKigo answers with the poem filed under a season word. The word is compared
// whole: matching a substring meant "" answered with whatever stood first, and
// half a word ("rain") found the winter poem by accident.
func ByKigo(word string) (Poem, bool) {
	for _, p := range year {
		if p.Kigo == word {
			return p, true
		}
	}
	return Poem{}, false
}
