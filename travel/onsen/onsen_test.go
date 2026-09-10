package onsen

import "testing"

func TestEveryBathIsPlacedAndTyped(t *testing.T) {
	want := map[string]Bath{
		"Kusatsu":     {Name: "Kusatsu", Prefecture: "Gunma", Water: "acidic sulphur", SourceC: 51},
		"Beppu":       {Name: "Beppu", Prefecture: "Oita", Water: "sodium chloride", SourceC: 60},
		"Noboribetsu": {Name: "Noboribetsu", Prefecture: "Hokkaido", Water: "sulphur", SourceC: 45},
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

// TestV1SurfaceHolds is the 1.0 declaration in machine-readable form. Baths
// left with v2; the declarations below stop compiling if one of the readers
// still under the promise moves, and the body checks they answer over one guide.
func TestV1SurfaceHolds(t *testing.T) {
	var (
		all    func() []Bath             = All
		find   func(string) (Bath, bool) = Find
		hotter func(int) []Bath          = Hotter
	)
	for _, b := range all() {
		if _, ok := find(b.Name); !ok {
			t.Errorf("Find misses %s, which All returns", b.Name)
		}
	}
	if got := hotter(0); len(got) != len(all()) {
		t.Errorf("Hotter(0) returns %d baths, want the whole guide", len(got))
	}
}

func TestTheGuideReadsHottestFirst(t *testing.T) {
	got := All()
	for i := 1; i < len(got); i++ {
		if got[i-1].SourceC < got[i].SourceC {
			t.Fatalf("%s (%d°C) prints before %s (%d°C)",
				got[i-1].Name, got[i-1].SourceC, got[i].Name, got[i].SourceC)
		}
	}
}
