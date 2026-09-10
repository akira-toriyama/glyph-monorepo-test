package onsen

import "strings"

// Find returns the bath posted under name, in whatever form the guest read it
// off a sign or a timetable.
func Find(name string) (Bath, bool) {
	q := normalize(name)
	for _, b := range baths {
		if normalize(b.Name) == q {
			return b, true
		}
	}
	return Bath{}, false
}

// normalize drops what a signboard adds and the guide does not carry: the case,
// the surrounding space, and the word Onsen itself.
func normalize(name string) string {
	n := strings.ToLower(strings.TrimSpace(name))
	n = strings.TrimSuffix(n, "onsen")
	return strings.Trim(n, " -")
}
