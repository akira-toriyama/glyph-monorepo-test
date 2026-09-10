package haiku

import (
	"strings"
	"testing"
)

func TestEverySeasonIsThreeLines(t *testing.T) {
	for _, poem := range []string{Spring(), Summer(), Autumn(), Winter()} {
		if got := len(strings.Split(poem, "\n")); got != 3 {
			t.Errorf("lines = %d, want 3: %q", got, poem)
		}
	}
}
