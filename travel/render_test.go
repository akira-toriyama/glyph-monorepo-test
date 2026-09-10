package travel

import (
	"strings"
	"testing"
)

func rows(t *testing.T) []string {
	t.Helper()
	out := strings.Split(strings.TrimSuffix(Kansai().String(), "\n"), "\n")
	if len(out) != len(Kansai().Stops) {
		t.Fatalf("%d rows for %d stops", len(out), len(Kansai().Stops))
	}
	return out
}

func TestSheetSaysNaraIsADayTrip(t *testing.T) {
	if got := rows(t)[1]; !strings.Contains(got, "day trip") {
		t.Fatalf("Nara's row = %q, want it to read day trip", got)
	}
}

func TestSheetNamesEveryTownAndItsPrefecture(t *testing.T) {
	sheet := rows(t)
	for i, stop := range Kansai().Stops {
		if !strings.Contains(sheet[i], stop.Name) || !strings.Contains(sheet[i], "("+stop.Prefecture+")") {
			t.Errorf("row %d = %q, want %s (%s)", i, sheet[i], stop.Name, stop.Prefecture)
		}
	}
}

func TestTheRideOnARowIsTheOneYouArriveOn(t *testing.T) {
	sheet := rows(t)
	if strings.Contains(sheet[0], "in on") {
		t.Errorf("Kyoto's row = %q, want no ride in — the trip starts there", sheet[0])
	}
	if last := sheet[len(sheet)-1]; !strings.Contains(last, "JR special rapid") {
		t.Fatalf("Osaka's row = %q, want the JR special rapid it arrives on", last)
	}
}
