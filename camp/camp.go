// Package camp is the fourth independently versioned module of the harness.
// Its version line is the tag prefix camp/.
package camp

// Gear is the packing list in packing order — what goes into the pack first is
// listed first. A new item lands as a minor bump of this module, a swapped one
// as a patch.
func Gear() []string {
	return []string{"tent", "sleeping bag", "sleeping mat", "gas stove"}
}
