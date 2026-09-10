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
