package travel

import "testing"

func TestRouteStartsInKyoto(t *testing.T) {
	if got := Route()[0]; got != "Kyoto" {
		t.Fatalf("first stop = %q, want Kyoto", got)
	}
}
