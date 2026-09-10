package camp

// Category is the bag an item rides in. The constants are declared in packing
// order: shelter goes to the bottom of the pack, kitchen rides on top.
type Category string

const (
	Shelter Category = "shelter"
	Sleep   Category = "sleep"
	Kitchen Category = "kitchen"
)

// category carries one entry per name Gear returns; TestEveryItemIsFiled holds
// the two in step.
var category = map[string]Category{
	"tent":         Shelter,
	"sleeping mat": Shelter,
	"sleeping bag": Sleep,
	"gas stove":    Kitchen,
}

// CategoryOf reports the bag an item rides in. ok is false for a name that is
// not on the list.
func CategoryOf(item string) (c Category, ok bool) {
	c, ok = category[item]
	return c, ok
}

// In returns the items filed under c.
func In(c Category) []string {
	var items []string
	for item, filed := range category {
		if filed == c {
			items = append(items, item)
		}
	}
	return items
}
