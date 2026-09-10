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

func TestOsakaIsLast(t *testing.T) {
	r := Route()
	if got := r[len(r)-1]; got != "Osaka" {
		t.Fatalf("last stop = %q, want Osaka — the airport bus leaves from Namba", got)
	}
}

func TestRouteHandsBackACopy(t *testing.T) {
	Route()[0] = "Himeji"
	if got := Route()[0]; got != "Kyoto" {
		t.Fatalf("first stop = %q after a caller wrote to the result, want Kyoto", got)
	}
}
