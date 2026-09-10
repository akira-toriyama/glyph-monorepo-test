package onsen

import "strings"

// Find returns the bath posted under name, folding case: a guest types what
// the sign said, and signs do not agree on it.
func Find(name string) (Bath, bool) {
	for _, b := range baths {
		if strings.EqualFold(b.Name, name) {
			return b, true
		}
	}
	return Bath{}, false
}
