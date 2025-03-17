package coll

import (
	"reflect"
	"slices"
	"testing"
)

type orderedIntArgs = testArgs[orderedInternal[int], int]
type orderedTestCase = testCase[orderedInternal[int], int]
type orderedCollIntBuilder = testCollectionBuilder[orderedInternal[int]]

func getValuesRevCases(builder orderedCollIntBuilder) []orderedTestCase {
	return []orderedTestCase{
		{
			name:  "ValuesRev() on empty collection",
			coll:  builder.Empty(),
			want1: []int(nil),
		},
		{
			name:  "ValuesRev() on one-item collection",
			coll:  builder.One(),
			want1: []int{111},
		},
		{
			name:  "ValuesRev() on three-item collection",
			coll:  builder.Three(),
			want1: []int{333, 222, 111},
		},
	}
}

func testValuesRev(t *testing.T, builder orderedCollIntBuilder) {
	cases := getValuesRevCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Collect(tt.coll.ValuesRev())
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("ValuesRev() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getValuesRevBreakCases(builder orderedCollIntBuilder) []*orderedTestCase {
	return []*orderedTestCase{
		{
			name: "ValuesRev() on three-item collection, break immediately",
			coll: builder.Three(),
			args: orderedIntArgs{
				predicate: func(v int) bool {
					return false
				},
			},
			want1: []int(nil),
		},
		{
			name: "ValuesRev() on three-item collection, break at middle",
			coll: builder.Three(),
			args: orderedIntArgs{
				predicate: func(v int) bool {
					return v > 222
				},
			},
			want1: []int{333},
		},
		{
			name: "ValuesRev() on three-item collection, break after middle",
			coll: builder.Three(),
			args: orderedIntArgs{
				predicate: func(v int) bool {
					return v >= 222
				},
			},
			want1: []int{333, 222},
		},
	}
}

func testValuesRevBreak(t *testing.T, builder orderedCollIntBuilder) {
	cases := getValuesRevBreakCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := []int(nil)
			for v := range tt.coll.ValuesRev() {
				if !tt.args.predicate(v) {
					break
				}
				got = append(got, v)
			}
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("ValuesRev() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}
