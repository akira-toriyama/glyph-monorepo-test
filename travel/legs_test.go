package travel

import "testing"

func TestLegsChainFromStopToStop(t *testing.T) {
	legs := Legs()
	if got := legs[0].From; got != Route()[0].Name {
		t.Fatalf("first leg leaves %q, want %q", got, Route()[0].Name)
	}
	for i := 1; i < len(legs); i++ {
		if legs[i].From != legs[i-1].To {
			t.Errorf("leg %d leaves %q but the one before arrives at %q", i, legs[i].From, legs[i-1].To)
		}
	}
}

func TestEveryLegNamesItsService(t *testing.T) {
	for _, leg := range Legs() {
		if leg.Line == "" || leg.Minutes == 0 {
			t.Errorf("%s to %s has no service in the table", leg.From, leg.To)
		}
	}
}

func TestKobeToOsakaRidesTheSpecialRapid(t *testing.T) {
	for _, leg := range Legs() {
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
