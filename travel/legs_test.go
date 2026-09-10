package travel

import "testing"

func TestLegsChainFromStopToStop(t *testing.T) {
	it := Kansai()
	legs := it.Legs()
	if got := legs[0].From; got != it.Stops[0].Name {
		t.Fatalf("first leg leaves %q, want %q", got, it.Stops[0].Name)
	}
	for i := 1; i < len(legs); i++ {
		if legs[i].From != legs[i-1].To {
			t.Errorf("leg %d leaves %q but the one before arrives at %q", i, legs[i].From, legs[i-1].To)
		}
	}
}

func TestEveryLegNamesItsService(t *testing.T) {
	for _, leg := range Kansai().Legs() {
		if leg.Line == "" || leg.Minutes == 0 {
			t.Errorf("%s to %s has no service in the table", leg.From, leg.To)
		}
	}
}

func TestKobeToOsakaRidesTheSpecialRapid(t *testing.T) {
	for _, leg := range Kansai().Legs() {
		if leg.From != "Kobe" || leg.To != "Osaka" {
			continue
		}
		if leg.Line != "JR special rapid" || leg.Minutes != 21 {
			t.Fatalf("Kobe to Osaka = %s in %d min, want the JR special rapid in 21", leg.Line, leg.Minutes)
		}
		return
	}
	t.Fatal("no leg from Kobe to Osaka")
}

func TestTravelTimeSumsTheLegs(t *testing.T) {
	it := Kansai()
	want := 0
	for _, leg := range it.Legs() {
		want += leg.Minutes
	}
	if got := it.TravelTime(); got != want {
		t.Fatalf("TravelTime() = %d min, want %d", got, want)
	}
}

func TestLegsStopAtTheLastStop(t *testing.T) {
	it := Kansai()
	legs := it.Legs()
	if len(legs) != len(it.Stops)-1 {
		t.Fatalf("%d legs for %d stops, want one fewer leg than stops", len(legs), len(it.Stops))
	}
	last := legs[len(legs)-1]
	if want := it.Stops[len(it.Stops)-1].Name; last.To != want {
		t.Fatalf("last leg arrives at %q, want %q", last.To, want)
	}
}
