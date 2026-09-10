package haiku

import (
	"strings"
	"testing"
)

func TestAutumnHasThreeLines(t *testing.T) {
	if got := len(strings.Split(Autumn(), "\n")); got != 3 {
		t.Fatalf("lines = %d, want 3", got)
	}
}
