package onsen

import "testing"

func TestBathsAreNamed(t *testing.T) {
	if len(Baths()) == 0 {
		t.Fatal("no baths")
	}
}

func TestEveryBathIsPlacedAndTyped(t *testing.T) {
	want := map[string]Bath{
		"Kusatsu":     {Name: "Kusatsu", Prefecture: "Gunma", Water: "acidic sulphur", SourceC: 41},
		"Beppu":       {Name: "Beppu", Prefecture: "Oita", Water: "simple alkaline", SourceC: 60},
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

// TestV1SurfaceHolds is the 1.0 declaration in machine-readable form: the
// declarations below stop compiling if a frozen signature moves, and the body
// checks the readers still answer over one guide.
func TestV1SurfaceHolds(t *testing.T) {
	var (
		all    func() []Bath             = All
		names  func() []string           = Baths
		find   func(string) (Bath, bool) = Find
		hotter func(int) []Bath          = Hotter
	)
	if len(all()) != len(names()) {
		t.Fatalf("All returns %d baths and Baths %d names", len(all()), len(names()))
	}
	for _, b := range all() {
		if _, ok := find(b.Name); !ok {
			t.Errorf("Find misses %s, which All returns", b.Name)
		}
	}
	if got := hotter(0); len(got) != len(all()) {
		t.Errorf("Hotter(0) returns %d baths, want the whole guide", len(got))
	}
}
