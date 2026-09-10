// Package onsen is a version line NESTED inside travel/: its tags are
// travel/onsen/vX.Y.Z, and a change here moves onsen alone — travel's line
// does not hear of it. That is the longest-prefix rule the harness proves.
package onsen

// Baths is the list of hot springs; a new bath lands as a minor bump of this module.
func Baths() []string {
	return []string{"Kusatsu", "Beppu"}
}
