package travel

import "testing"

func TestRouteStartsInKyoto(t *testing.T) {
	if got := Route()[0].Name; got != "Kyoto" {
		t.Fatalf("first stop = %q, want Kyoto", got)
	}
}

func TestEveryStopNamesItsPrefecture(t *testing.T) {
	for _, stop := range Route() {
		if stop.Prefecture == "" {
			t.Errorf("%s names no prefecture", stop.Name)
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
	if got := r[len(r)-1].Name; got != "Osaka" {
		t.Fatalf("last stop = %q, want Osaka — the airport bus leaves from Namba", got)
	}
}

func TestRouteHandsBackACopy(t *testing.T) {
	Route()[0].Name = "Himeji"
	if got := Route()[0].Name; got != "Kyoto" {
		t.Fatalf("first stop = %q after a caller wrote to the result, want Kyoto", got)
	}
}

func TestNaraIsADayTripFromKyoto(t *testing.T) {
	for _, stop := range Route() {
		switch stop.Name {
		case "Nara":
			if stop.Nights != 0 {
				t.Errorf("Nara = %d nights, want 0 — the bed stays in Kyoto", stop.Nights)
			}
		case "Kyoto":
			if stop.Nights != 4 {
				t.Errorf("Kyoto = %d nights, want 4 — it holds the Nara day too", stop.Nights)
			}
		}
	}
}
