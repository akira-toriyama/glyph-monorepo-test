package curry

import "testing"

// Rice scales straight off the plate count, not off baseServings: nobody eats
// three quarters of a bowl.
func TestRiceScalesPerPlate(t *testing.T) {
	one, four := RiceFor(1), RiceFor(4)
	if four.DryGrams != 4*one.DryGrams {
		t.Errorf("four plates = %dg dry, one plate = %dg", four.DryGrams, one.DryGrams)
	}
	if four.SoakMinutes != one.SoakMinutes {
		t.Errorf("the soak moved with the pot size: %d against %d", four.SoakMinutes, one.SoakMinutes)
	}
}

func TestRiceRefusesAPotBelowOnePlate(t *testing.T) {
	for _, plates := range []int{0, -1} {
		if got := RiceFor(plates); got != (Rice{}) {
			t.Errorf("RiceFor(%d) = %+v, want the zero pot", plates, got)
		}
	}
}

func TestRiceDrinksMoreThanItWeighs(t *testing.T) {
	r := RiceFor(2)
	if r.WaterGrams <= r.DryGrams {
		t.Errorf("%dg of water on %dg of rice", r.WaterGrams, r.DryGrams)
	}
}

// The soak and the water ratio are one decision, not two: soaked grain at 1.2
// steams to porridge, dry grain at 1.1 comes out chalky.
func TestSoakedRiceTakesElevenPartsWater(t *testing.T) {
	for _, plates := range []int{1, 2, 4, 7} {
		r := RiceFor(plates)
		if r.SoakMinutes == 0 {
			t.Fatalf("RiceFor(%d) dropped the soak", plates)
		}
		if r.WaterGrams*10 != r.DryGrams*11 {
			t.Errorf("RiceFor(%d): %dg of water on %dg of rice", plates, r.WaterGrams, r.DryGrams)
		}
	}
}
