package travel

import (
	"slices"
	"testing"
)

func TestKansaiStartsInKyoto(t *testing.T) {
	if got := Kansai().Stops[0].Name; got != "Kyoto" {
		t.Fatalf("first stop = %q, want Kyoto", got)
	}
}

func TestEveryStopNamesItsPrefecture(t *testing.T) {
	for _, stop := range Kansai().Stops {
		if stop.Prefecture == "" {
			t.Errorf("%s names no prefecture", stop.Name)
		}
	}
}

func TestDurationSumsTheNights(t *testing.T) {
	if got := Kansai().Duration(); got != 7 {
		t.Fatalf("Duration() = %d nights, want 7", got)
	}
}

func TestOsakaIsLast(t *testing.T) {
	stops := Kansai().Stops
	if got := stops[len(stops)-1].Name; got != "Osaka" {
		t.Fatalf("last stop = %q, want Osaka — the airport bus leaves from Namba", got)
	}
}

func TestNaraIsADayTripFromKyoto(t *testing.T) {
	for _, stop := range Kansai().Stops {
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

func TestKansaiHandsBackAFreshItinerary(t *testing.T) {
	Kansai().Stops[0].Name = "Himeji"
	if got := Kansai().Stops[0].Name; got != "Kyoto" {
		t.Fatalf("first stop = %q after a caller wrote to the result, want Kyoto", got)
	}
}

func TestSleepTownsHasOneEntryPerNight(t *testing.T) {
	it := Kansai()
	towns := it.SleepTowns()
	if len(towns) != it.Duration() {
		t.Fatalf("%d sleep towns for %d nights", len(towns), it.Duration())
	}
	if slices.Contains(towns, "Nara") {
		t.Error("Nara is a day trip; the bed that night is in Kyoto")
	}
	if towns[0] != "Kyoto" || towns[len(towns)-1] != "Osaka" {
		t.Fatalf("nights run %q to %q, want Kyoto to Osaka", towns[0], towns[len(towns)-1])
	}
}
