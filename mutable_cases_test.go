package coll

import (
	"reflect"
	"testing"
)

type mutableIntTestArgs = testArgs[mutableInternal[int], int]
type mutableIntTestCase = testCase[mutableInternal[int], int]
type mutableIntTestBuilder = testCollectionBuilder[mutableInternal[int], int]

func getApplyCases(builder mutableIntTestBuilder) []mutableIntTestCase {
	return []mutableIntTestCase{
		{
			name:  "Apply() on empty collection",
			coll:  builder.Empty(),
			args:  mutableIntTestArgs{mapper: func(i int, v int) int { return i + v }},
			want1: []int{},
		},
		{
			name:  "Apply() on one-item collection",
			coll:  builder.One(),
			args:  mutableIntTestArgs{mapper: func(i int, v int) int { return i * v }},
			want1: []int{0},
		},
		{
			name:  "Apply() on three-item collection",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{mapper: func(i int, v int) int { return i + v }},
			want1: []int{111, 223, 335},
		},
	}
}

func testApply(t *testing.T, builder mutableIntTestBuilder) {
	cases := getApplyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Apply(tt.args.mapper)
			if !reflect.DeepEqual(tt.coll.ToSlice(), tt.want1) {
				t.Errorf("Apply() resulted in: %v, but wanted %v", tt.coll.ToSlice(), tt.want1)
			}
		})
	}
}

func getClearCases(builder mutableIntTestBuilder) []mutableIntTestCase {
	return []mutableIntTestCase{
		{
			name: "Clear() on empty collection",
			coll: builder.Empty(),
		},
		{
			name: "Clear() on one-item collection",
			coll: builder.One(),
		},
		{
			name: "Clear() on three-item collection",
			coll: builder.Three(),
		},
	}
}

func testClear(t *testing.T, builder mutableIntTestBuilder) {
	cases := getClearCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Clear()
			if !reflect.DeepEqual(tt.coll.ToSlice(), []int{}) {
				t.Errorf("Clear() did not clear correctly")
			}
		})
	}
}

func getRemoveMatchingCases(builder mutableIntTestBuilder) []mutableIntTestCase {
	return []mutableIntTestCase{
		{
			name:  "RemoveMatching() on empty collection",
			coll:  builder.Empty(),
			args:  mutableIntTestArgs{predicate: func(i int, v int) bool { return true }},
			want1: []int{},
		},
		{
			name:  "RemoveMatching() on one-item collection",
			coll:  builder.One(),
			args:  mutableIntTestArgs{predicate: func(i int, v int) bool { return v == 111 }},
			want1: []int{},
		},
		{
			name:  "RemoveMatching() on three-item collection",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{predicate: func(i int, v int) bool { return v == 222 }},
			want1: []int{111, 333},
		},
		{
			name:  "RemoveMatching() on three-item collection, all false",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{predicate: func(i int, v int) bool { return false }},
			want1: []int{111, 222, 333},
		},
		{
			name:  "RemoveMatching() on three-item collection, all true",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{predicate: func(i int, v int) bool { return true }},
			want1: []int{},
		},
		{
			name:  "RemoveMatching() on three-item collection, some mod 2",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{predicate: func(i int, v int) bool { return v%2 == 0 }},
			want1: []int{111, 333},
		},
		{
			name:  "RemoveMatching() on three-item collection, some not mod 2",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{predicate: func(i int, v int) bool { return v%2 != 0 }},
			want1: []int{222},
		},
		{
			name:  "RemoveMatching() on three-item collection, using index",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{predicate: func(i int, v int) bool { return i == 0 || i == 1 || i == 2 }},
			want1: []int{},
		},
	}
}

func testRemoveMatching(t *testing.T, builder mutableIntTestBuilder) {
	cases := getRemoveMatchingCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.RemoveMatching(tt.args.predicate)
			if !reflect.DeepEqual(tt.coll.ToSlice(), tt.want1) {
				t.Errorf("RemoveMatching() did not remove correctly")
			}
		})
	}
}
