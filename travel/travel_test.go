package travel

import "testing"

func TestRouteStartsInKyoto(t *testing.T) {
	if got := Route()[0]; got != "Kyoto" {
		t.Fatalf("first stop = %q, want Kyoto", got)
	}
}

func TestEveryStopIsInTheNightsTable(t *testing.T) {
	for _, stop := range Route() {
		if _, ok := nights[stop]; !ok {
			t.Errorf("%s has no entry in the nights table", stop)
		}
	}
}

func TestDurationSumsTheNights(t *testing.T) {
	if got := Duration(); got != 7 {
		t.Fatalf("Duration() = %d nights, want 7", got)
	}
}
