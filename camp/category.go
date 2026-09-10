package camp

// Category is the bag an item rides in. The constants are declared in packing
// order: shelter goes to the bottom of the pack, kitchen rides on top.
type Category string

const (
	Shelter Category = "shelter"
	Sleep   Category = "sleep"
	Kitchen Category = "kitchen"
)

// In returns the items filed under c, in list order. It walks the list rather
// than an index keyed by name because the result is read as a packing
// sequence.
func In(c Category) []Item {
	var items []Item
	for _, item := range Gear() {
		if item.Category == c {
			items = append(items, item)
		}
	}
	return items
}
