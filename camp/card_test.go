package camp

import (
	"strings"
	"testing"
)

func TestCardOpensWithTheTent(t *testing.T) {
	if first := strings.SplitN(Card(), "\n", 2)[0]; !strings.HasPrefix(first, "tent ") {
		t.Fatalf("card opens %q, want the tent", first)
	}
}

func TestCardClosesWithTheTotal(t *testing.T) {
	lines := strings.Split(strings.TrimRight(Card(), "\n"), "\n")
	if got, want := len(lines), len(Gear())+1; got != want {
		t.Fatalf("card has %d lines, want one per item plus the total (%d)", got, want)
	}
	last := lines[len(lines)-1]
	if !strings.HasPrefix(last, "total") || !strings.Contains(last, "4.3 kg") {
		t.Fatalf("last line = %q, want the total at 4.3 kg", last)
	}
}
