package onsen

import "testing"

func TestBathsAreNamed(t *testing.T) {
	if len(Baths()) == 0 {
		t.Fatal("no baths")
	}
}
