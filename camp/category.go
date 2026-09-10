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
	"water filter": Kitchen,
}

// CategoryOf reports the bag an item rides in. ok is false for a name that is
// not on the list.
func CategoryOf(item string) (c Category, ok bool) {
	c, ok = category[item]
	return c, ok
}

// In returns the items filed under c, in list order. It walks Gear rather than
// the category map because ranging a map is unordered, and the result is read
// as a packing sequence.
func In(c Category) []string {
	var items []string
	for _, item := range Gear() {
		if category[item] == c {
			items = append(items, item)
		}
	}
	return items
}
