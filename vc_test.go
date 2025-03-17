package coll

import "testing"

func Test_valuesCounter_Set_WithZeroValue(t *testing.T) {
	vc := newValuesCounter[int]()
	vc.Set(1, 0)
	if len(vc.counter) != 0 {
		t.Error("Set() did not remove the value when count is 0")
	}
}
