package haiku

import "strings"

// ByKigo answers with the poem filed under a season word.
func ByKigo(word string) (Poem, bool) {
	for _, p := range year {
		if strings.Contains(p.Kigo, word) {
			return p, true
		}
	}
	return Poem{}, false
}
