package camp

import "testing"

func TestGearHasATent(t *testing.T) {
	if got := Gear()[0]; got != "tent" {
		t.Fatalf("first item = %q, want tent", got)
	}
}
