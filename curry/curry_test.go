package curry

import "testing"

func TestIngredientsEndWithRoux(t *testing.T) {
	in := Ingredients()
	if in[len(in)-1] != "roux" {
		t.Fatalf("last ingredient = %q, want roux", in[len(in)-1])
	}
}
