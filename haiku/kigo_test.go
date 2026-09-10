package haiku

import "testing"

func TestByKigoFindsTheWinterRain(t *testing.T) {
	got, ok := ByKigo("first winter rain")
	if !ok {
		t.Fatal("the first winter rain is filed under nothing")
	}
	if got != Winter() {
		t.Fatalf("found %q", got.Text)
	}
}

func TestByKigoMissesAWordNoPoemCarries(t *testing.T) {
	if got, ok := ByKigo("cherry blossom"); ok {
		t.Fatalf("no poem here carries the cherry blossom, but %q answered", got.Text)
	}
}
