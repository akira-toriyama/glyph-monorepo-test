package onsen

import "testing"

func TestBathsAreNamed(t *testing.T) {
	if len(Baths()) == 0 {
		t.Fatal("no baths")
	}
}

func TestEveryBathIsPlacedAndTyped(t *testing.T) {
	want := map[string]Bath{
		"Kusatsu":     {Name: "Kusatsu", Prefecture: "Gunma", Water: "acidic sulphur"},
		"Beppu":       {Name: "Beppu", Prefecture: "Oita", Water: "simple alkaline"},
		"Noboribetsu": {Name: "Noboribetsu", Prefecture: "Aomori", Water: "sulphur"},
	}
	got := All()
	if len(got) != len(want) {
		t.Fatalf("the guide holds %d baths, want %d", len(got), len(want))
	}
	for _, b := range got {
		w, ok := want[b.Name]
		if !ok {
			t.Errorf("%s is not a bath the guide indexes", b.Name)
			continue
		}
		if b != w {
			t.Errorf("%s = %+v, want %+v", b.Name, b, w)
		}
	}
}

func TestAllHandsBackACopy(t *testing.T) {
	got := All()
	if len(got) == 0 {
		t.Fatal("the guide is empty")
	}
	first := got[0]
	got[0] = Bath{Name: "Atami"}
	if again := All()[0]; again != first {
		t.Fatalf("the guide opens on %+v after a caller overwrote its copy, want %+v", again, first)
	}
}
