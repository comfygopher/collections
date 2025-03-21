package coll

import (
	"reflect"
	"testing"
)

type mutableIntTestArgs = testArgs[mutableInternal[int], int]
type mutableIntTestCase = testCase[mutableInternal[int], int]
type mutableIntTestBuilder = testCollectionBuilder[mutableInternal[int]]

func getApplyCases(builder mutableIntTestBuilder) []mutableIntTestCase {
	return []mutableIntTestCase{
		{
			name:  "Apply() on empty collection",
			coll:  builder.Empty(),
			args:  mutableIntTestArgs{mapper: func(v int) int { return 10 + v }},
			want1: []int(nil),
			want2: map[int]int{},
		},
		{
			name:  "Apply() on one-item collection",
			coll:  builder.One(),
			args:  mutableIntTestArgs{mapper: func(v int) int { return v * 10 }},
			want1: []int{1110},
			want2: map[int]int{1110: 1},
		},
		{
			name:  "Apply() on three-item collection",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{mapper: func(v int) int { return v * 10 }},
			want1: []int{1110, 2220, 3330},
			want2: map[int]int{1110: 1, 2220: 1, 3330: 1},
		},
		{
			name:  "Apply() on six-item collection",
			coll:  builder.SixWithDuplicates(),
			args:  mutableIntTestArgs{mapper: func(v int) int { return v * 2 }},
			want1: []int{222, 444, 666, 222, 444, 666},
			want2: map[int]int{222: 2, 444: 2, 666: 2},
		},
	}
}

func testApply(t *testing.T, builder mutableIntTestBuilder) {
	cases := getApplyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Apply(tt.args.mapper)

			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Apply() resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
			if actualVC != nil && !reflect.DeepEqual(actualVC, tt.want2) {
				t.Errorf("Apply() did not append correctly from values counter")
			}
		})
	}
}

func getClearCases(builder mutableIntTestBuilder) []mutableIntTestCase {
	return []mutableIntTestCase{
		{
			name:  "Clear() on empty collection",
			coll:  builder.Empty(),
			want1: []int(nil),
			want2: map[int]int{},
		},
		{
			name:  "Clear() on one-item collection",
			coll:  builder.One(),
			want1: []int(nil),
			want2: map[int]int{},
		},
		{
			name:  "Clear() on three-item collection",
			coll:  builder.Three(),
			want1: []int(nil),
			want2: map[int]int{},
		},
	}
}

func testClear(t *testing.T, builder mutableIntTestBuilder) {
	cases := getClearCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Clear()

			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Clear() resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
			if actualVC != nil && !reflect.DeepEqual(actualVC, tt.want2) {
				t.Errorf("Clear() did not append correctly from values counter")
			}
		})
	}
}

func getRemoveMatchingCases(builder mutableIntTestBuilder) []mutableIntTestCase {
	return []mutableIntTestCase{
		{
			name:  "RemoveMatching() on empty collection",
			coll:  builder.Empty(),
			args:  mutableIntTestArgs{predicate: func(v int) bool { return true }},
			want1: []int(nil),
			want2: map[int]int{},
			want3: 0,
		},
		{
			name:  "RemoveMatching() on one-item collection",
			coll:  builder.One(),
			args:  mutableIntTestArgs{predicate: func(v int) bool { return v == 111 }},
			want1: []int(nil),
			want2: map[int]int{},
			want3: 1,
		},
		{
			name:  "RemoveMatching() on three-item collection",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{predicate: func(v int) bool { return v == 222 }},
			want1: []int{111, 333},
			want2: map[int]int{111: 1, 333: 1},
			want3: 1,
		},
		{
			name:  "RemoveMatching() on three-item collection, all false",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{predicate: func(v int) bool { return false }},
			want1: []int{111, 222, 333},
			want2: map[int]int{111: 1, 222: 1, 333: 1},
			want3: 0,
		},
		{
			name:  "RemoveMatching() on three-item collection, all true",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{predicate: func(v int) bool { return true }},
			want1: []int(nil),
			want2: map[int]int{},
			want3: 3,
		},
		{
			name:  "RemoveMatching() on three-item collection, some mod 2",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{predicate: func(v int) bool { return v%2 == 0 }},
			want1: []int{111, 333},
			want2: map[int]int{111: 1, 333: 1},
			want3: 1,
		},
		{
			name:  "RemoveMatching() on three-item collection, some not mod 2",
			coll:  builder.Three(),
			args:  mutableIntTestArgs{predicate: func(v int) bool { return v%2 != 0 }},
			want1: []int{222},
			want2: map[int]int{222: 1},
			want3: 2,
		},
	}
}

func testRemoveMatching(t *testing.T, builder mutableIntTestBuilder) {
	cases := getRemoveMatchingCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			count := tt.coll.RemoveMatching(tt.args.predicate)

			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("RemoveMatching() did not remove correctly")
			}
			if actualVC != nil && !reflect.DeepEqual(actualVC, tt.want2) {
				t.Errorf("RemoveMatching() did not remove correctly from values counter")
			}
			if count != tt.want3 {
				t.Errorf("RemoveMatching() returned wrong count: %v, but wanted = %v", count, tt.want3)
			}
		})
	}
}
