package onsen

import "testing"

func TestFindFoldsCase(t *testing.T) {
	b, ok := Find("kusatsu")
	if !ok {
		t.Fatal(`Find("kusatsu") found nothing`)
	}
	if b.Name != "Kusatsu" {
		t.Fatalf("Find returned %q", b.Name)
	}
}

func TestFindMissesQuietly(t *testing.T) {
	if b, ok := Find("Atami"); ok {
		t.Fatalf("Find returned %+v; Atami is not in the guide", b)
	}
}

func TestFindTakesTheSignboardSuffix(t *testing.T) {
	for _, q := range []string{"Kusatsu Onsen", "  kusatsu-onsen ", "KUSATSU"} {
		b, ok := Find(q)
		if !ok || b.Name != "Kusatsu" {
			t.Errorf("Find(%q) = %+v, %v; want Kusatsu", q, b, ok)
		}
	}
}
