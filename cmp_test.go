package coll

import (
	"fmt"
	"testing"
)

func Test_valuesCounter(t *testing.T) {
	t.Run("Increment(), Decrement(), and Count()", func(t *testing.T) {
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
		testCases := []struct {
			value    int
			expected int
		}{
			{1, 2},
			{2, 1},
			{3, 3},
			{4, 4},
			{5, 0}, // Non-existent value
		}

		for _, tc := range testCases {
			t.Run(fmt.Sprintf("Count(%d)", tc.value), func(t *testing.T) {
				if count := vc.Count(tc.value); count != tc.expected {
					t.Errorf("Count(%d) = %v, want %v", tc.value, count, tc.expected)
				}
			})
		}
	})
}
