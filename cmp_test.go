package coll

import "testing"

func Test_valuesCounter(t *testing.T) {
	t.Run("Increment() and Count()", func(t *testing.T) {
		vc := newValuesCounter[int]()
		vc.Decrement(1)
		vc.Increment(1)
		vc.Increment(1)
		vc.Increment(2)
		vc.Increment(3)
		vc.Increment(3)
		vc.Increment(3)
		vc.Increment(4)
		vc.Increment(4)
		vc.Increment(4)
		vc.Increment(4)
		if vc.Count(1) != 2 {
			t.Errorf("Count(1) = %v, want %v", vc.Count(1), 2)
		}
		if vc.Count(2) != 1 {
			t.Errorf("Count(2) = %v, want %v", vc.Count(2), 1)
		}
		if vc.Count(3) != 3 {
			t.Errorf("Count(3) = %v, want %v", vc.Count(3), 3)
		}
		if vc.Count(4) != 4 {
			t.Errorf("Count(4) = %v, want %v", vc.Count(4), 4)
		}
		if vc.Count(5) != 0 {
			t.Errorf("Count(5) = %v, want %v", vc.Count(5), 0)
		}
	})
}
