package camp

import "strings"

// Missing returns the list's items that packed does not name, in list order. A
// name in packed that the list never asked for is not an error: the pack is
// allowed to hold more than the list.
//
// Names match with case folded and surrounding space trimmed — packed is typed
// by hand, not read off Gear.
func Missing(packed []string) []Item {
	have := make(map[string]bool, len(packed))
	for _, name := range packed {
		have[fold(name)] = true
	}
	var left []Item
	for _, item := range Gear() {
		if !have[fold(item.Name)] {
			left = append(left, item)
		}
	}
	return left
}

func fold(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}
