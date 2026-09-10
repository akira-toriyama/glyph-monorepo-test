package camp

import "slices"

// Missing returns the list's items that packed does not name, in list order. A
// name in packed that the list never asked for is not an error: the pack is
// allowed to hold more than the list.
func Missing(packed []string) []Item {
	var left []Item
	for _, item := range Gear() {
		if !slices.Contains(packed, item.Name) {
			left = append(left, item)
		}
	}
	return left
}
