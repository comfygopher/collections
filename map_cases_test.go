package coll

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type baseMapIntArgs = testArgs[Map[int, int], Pair[int, int]]
type baseMapTestCase = testCase[Map[int, int], Pair[int, int]]
type baseMapCollIntBuilder = testCollectionBuilder[Map[int, int], Pair[int, int]]

func getMapAtCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "At(0) on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{index: 0},
			want1: nil,
			want2: false,
		},
		{
			name:  "At(0) on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{index: 0},
			want1: NewPair(1, 111),
			want2: true,
		},
		{
			name:  "At(1) on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 1},
			want1: NewPair(2, 222),
			want2: true,
		},
		{
			name:  "At(3) on three-item collection out of bounds",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 3},
			want1: nil,
			want2: false,
		},
		{
			name:  "At(-1) on three-item collection negative index",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: -1},
			want1: nil,
			want2: false,
		},
	}
}

func testMapAt(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapAtCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.At(tt.args.index)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("At() got1 = %v, want1 = %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("At() got2 = %v, want2 = %v", got2, tt.want2)
			}
		})
	}
}

func getMapAtOrDefaultCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "AtOrDefault(0) on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{index: 0, defaultValue: NewPair(999, 999)},
			want1: NewPair(999, 999),
		},
		{
			name:  "AtOrDefault(0) on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{index: 0, defaultValue: nil},
			want1: NewPair(1, 111),
		},
		{
			name:  "AtOrDefault(1) on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 1, defaultValue: nil},
			want1: NewPair(2, 222),
		},
		{
			name:  "AtOrDefault(3) on three-item collection out of bounds",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 3, defaultValue: NewPair(999, 999)},
			want1: NewPair(999, 999),
		},
		{
			name:  "AtOrDefault(-1) on three-item collection negative index",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: -1, defaultValue: NewPair(999, 999)},
			want1: NewPair(999, 999),
		},
		{
			name:  "AtOrDefault(1) on one-item collection out of bounds",
			coll:  builder.One(),
			args:  baseMapIntArgs{index: 1, defaultValue: NewPair(999, 999)},
			want1: NewPair(999, 999),
		},
	}
}

func testMapAtOrDefault(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapAtOrDefaultCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1 := tt.coll.AtOrDefault(tt.args.index, tt.args.defaultValue)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("At() got1 = %v, want1 = %v", got1, tt.want1)
			}
		})
	}
}

func getMapClearCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
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

func testMapClear(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapClearCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Clear()
			actualSlice := []Pair[int, int]{}
			copy(actualSlice, tt.coll.ToSlice())
			actualMap := tt.coll.ToMap()
			if !reflect.DeepEqual(actualSlice, []Pair[int, int]{}) {
				t.Errorf("Clear() did not clear slice correctly")
			}
			if !reflect.DeepEqual(actualMap, map[int]int{}) {
				t.Errorf("Clear() did not clear map correctly")
			}
		})
	}
}

func getMapContainsCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Contains() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return p.Val() == 1 }},
			want1: false,
		},
		{
			name:  "Contains() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return p.Val() == 111 }},
			want1: true,
		},
		{
			name:  "Contains() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return p.Val() == 222 }},
			want1: true,
		},
		{
			name:  "Contains() on three-item collection, all false",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return false }},
			want1: false,
		},
		{
			name:  "Contains() on three-item collection, all true",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return true }},
			want1: true,
		},
	}
}

func testMapContains(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapContainsCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.coll.Contains(tt.args.predicate); got != tt.want1 {
				t.Errorf("Contains() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapCountCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Count() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return p.Val() == 1 }},
			want1: 0,
		},
		{
			name:  "Count() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return p.Val() == 111 }},
			want1: 1,
		},
		{
			name:  "Count() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return p.Val() == 222 }},
			want1: 1,
		},
		{
			name:  "Count() on three-item collection, all false",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return false }},
			want1: 0,
		},
		{
			name:  "Count() on three-item collection, all true",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return true }},
			want1: 3,
		},
	}
}

func testMapCount(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapCountCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.Count(tt.args.predicate)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Count() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapEachCases(t *testing.T, builder baseMapCollIntBuilder) []*baseMapTestCase {
	// eachOnEmptyListCase:

	eachOnEmptyListCase := &baseMapTestCase{
		name: "Each() on empty collection",
		coll: builder.Empty(),
	}
	eachOnEmptyListCase.args = baseMapIntArgs{
		visit: func(i int, p Pair[int, int]) {
			t.Error("Each() called on empty collection")
		},
	}

	// eachOnOneItemCase:

	eachOnOneItemCase := &baseMapTestCase{
		name:  "Each() on one-item collection",
		coll:  builder.One(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0},
		want2: []int{111},
	}

	eachOnOneItemCase.args = baseMapIntArgs{
		visit: func(i int, p Pair[int, int]) {
			if i != 0 || p.Val() != 111 {
				t.Error("Each() called with wrong values")
			}
			eachOnOneItemCase.got1 = append(eachOnOneItemCase.got1.([]int), i)
			eachOnOneItemCase.got2 = append(eachOnOneItemCase.got2.([]int), p.Val())
		},
	}

	// eachOnEmptyListCase:

	eachOnThreeCase := &baseMapTestCase{
		name:  "Each() on three-item collection",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0, 1, 2},
		want2: []int{111, 222, 333},
	}

	eachOnThreeCase.args = baseMapIntArgs{
		visit: func(i int, p Pair[int, int]) {
			if i < 0 || i > 2 || p.Val() < 111 || p.Val() > 333 {
				t.Error("Each() called with wrong values")
			}
			eachOnThreeCase.got1 = append(eachOnThreeCase.got1.([]int), i)
			eachOnThreeCase.got2 = append(eachOnThreeCase.got2.([]int), p.Val())
		},
	}

	// put the cases together:

	return []*baseMapTestCase{
		eachOnEmptyListCase,
		eachOnOneItemCase,
		eachOnThreeCase,
	}
}

func testMapEach(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapEachCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Each(tt.args.visit)
			if tt.got1 != nil && !reflect.DeepEqual(tt.got1, tt.want1) {
				t.Errorf("Each() called with wrong indices: %v, want1 = %v", tt.got1, tt.want1)
			}
			if tt.got2 != nil && !reflect.DeepEqual(tt.got2, tt.want2) {
				t.Errorf("Each() called with wrong values: %v, want1 = %v", tt.got2, tt.want2)
			}
		})
	}
}

func getMapEachRevCases(t *testing.T, builder baseMapCollIntBuilder) []*baseMapTestCase {

	// eachRevOnEmptyListCase:

	eachRevOnEmptyListCase := &baseMapTestCase{
		name: "EachRev() on empty collection",
		coll: builder.Empty(),
	}
	eachRevOnEmptyListCase.args = baseMapIntArgs{
		visit: func(i int, p Pair[int, int]) {
			t.Error("EachRev() called on empty collection")
		},
	}

	// eachRevOnOneItemCase:

	eachRevOnOneItemCase := &baseMapTestCase{
		name:  "EachRev() on one-item collection",
		coll:  builder.One(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0},
		want2: []int{111},
	}

	eachRevOnOneItemCase.args = baseMapIntArgs{
		visit: func(i int, p Pair[int, int]) {
			if i != 0 || p.Val() != 111 {
				t.Error("EachRev() called with wrong values")
			}
			eachRevOnOneItemCase.got1 = append(eachRevOnOneItemCase.got1.([]int), i)
			eachRevOnOneItemCase.got2 = append(eachRevOnOneItemCase.got2.([]int), p.Val())
		},
	}

	// eachRevOnThreeCase:

	eachRevOnThreeCase := &baseMapTestCase{
		name:  "EachRev() on three-item collection",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{2, 1, 0},
		want2: []int{333, 222, 111},
	}

	eachRevOnThreeCase.args = baseMapIntArgs{
		visit: func(i int, p Pair[int, int]) {
			if i < 0 || i > 2 || p.Val() < 111 || p.Val() > 333 {
				t.Error("EachRev() called with wrong values")
			}
			eachRevOnThreeCase.got1 = append(eachRevOnThreeCase.got1.([]int), i)
			eachRevOnThreeCase.got2 = append(eachRevOnThreeCase.got2.([]int), p.Val())
		},
	}

	// put the cases together:

	return []*baseMapTestCase{
		eachRevOnEmptyListCase,
		eachRevOnOneItemCase,
		eachRevOnThreeCase,
	}
}

func testMapEachRev(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapEachRevCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.EachRev(tt.args.visit)
			if tt.got1 != nil && !reflect.DeepEqual(tt.got1, tt.want1) {
				t.Errorf("EachRev() called with wrong indices: %v, want1 = %v", tt.got1, tt.want1)
			}
			if tt.got2 != nil && !reflect.DeepEqual(tt.got2, tt.want2) {
				t.Errorf("EachRev() called with wrong values: %v, want1 = %v", tt.got2, tt.want2)
			}
		})
	}
}

func getMapEachUntilCases(t *testing.T, builder baseMapCollIntBuilder) []*baseMapTestCase {

	// eachUntilOnEmptyListCase:

	eachUntilOnEmptyListCase := &baseMapTestCase{
		name: "EachUntil() on empty collection",
		coll: builder.Empty(),
	}

	eachUntilOnEmptyListCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			t.Error("EachUntil() called on empty collection")
			return true
		},
	}

	// eachUntilOnOneItemCase:

	eachUntilOnOneItemCase := &baseMapTestCase{
		name:  "EachUntil() on one-item collection",
		coll:  builder.One(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0},
		want2: []int{111},
	}

	eachUntilOnOneItemCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if i != 0 || p.Val() != 111 {
				t.Error("EachUntil() called with wrong values")
			}
			eachUntilOnOneItemCase.got1 = append(eachUntilOnOneItemCase.got1.([]int), i)
			eachUntilOnOneItemCase.got2 = append(eachUntilOnOneItemCase.got2.([]int), p.Val())
			return true
		},
	}

	// eachUntilFinishMiddleCase:

	eachUntilFinishMiddleCase := &baseMapTestCase{
		name:  "EachUntil() finish in middle",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0, 1},
		want2: []int{111, 222},
	}

	eachUntilFinishMiddleCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if i < 0 || i > 2 || p.Val() < 111 || p.Val() > 333 {
				t.Error("EachUntil() called with wrong values")
			}
			eachUntilFinishMiddleCase.got1 = append(eachUntilFinishMiddleCase.got1.([]int), i)
			eachUntilFinishMiddleCase.got2 = append(eachUntilFinishMiddleCase.got2.([]int), p.Val())
			stop := i >= 1 && p.Val() >= 222
			cont := !stop
			return cont
		},
	}

	// eachUntilAllThreeCase:

	eachUntilAllThreeCase := &baseMapTestCase{
		name:  "EachUntil() on three-item collection",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0, 1, 2},
		want2: []int{111, 222, 333},
	}

	eachUntilAllThreeCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if i < 0 || i > 2 || p.Val() < 111 || p.Val() > 333 {
				t.Error("EachUntil() called with wrong values")
			}
			eachUntilAllThreeCase.got1 = append(eachUntilAllThreeCase.got1.([]int), i)
			eachUntilAllThreeCase.got2 = append(eachUntilAllThreeCase.got2.([]int), p.Val())
			return true
		},
	}

	// put the cases together:

	return []*baseMapTestCase{
		eachUntilOnEmptyListCase,
		eachUntilOnOneItemCase,
		eachUntilFinishMiddleCase,
		eachUntilAllThreeCase,
	}
}

func testMapEachUntil(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapEachUntilCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.EachUntil(tt.args.predicate)
			if tt.got1 != nil && !reflect.DeepEqual(tt.got1, tt.want1) {
				t.Errorf("EachUntil() called with wrong indices: %v, want1 = %v", tt.got1, tt.want1)
			}
			if tt.got2 != nil && !reflect.DeepEqual(tt.got2, tt.want2) {
				t.Errorf("EachUntil() called with wrong values: %v, want1 = %v", tt.got2, tt.want2)
			}
		})
	}
}

func getMapEachRevUntilCases(t *testing.T, builder baseMapCollIntBuilder) []*baseMapTestCase {

	// eachRevUntilOnEmptyListCase:

	eachRevUntilOnEmptyListCase := &baseMapTestCase{
		name: "EachRevUntil() on empty collection",
		coll: builder.Empty(),
	}

	eachRevUntilOnEmptyListCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			t.Error("EachRevUntil() called on empty collection")
			return true
		},
	}

	// eachRevUntilOnOneItemCase:

	eachRevUntilOnOneItemCase := &baseMapTestCase{
		name:  "EachRevUntil() on one-item collection",
		coll:  builder.One(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0},
		want2: []int{111},
	}

	eachRevUntilOnOneItemCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if i != 0 || p.Val() != 111 {
				t.Error("EachRevUntil() called with wrong values")
			}
			eachRevUntilOnOneItemCase.got1 = append(eachRevUntilOnOneItemCase.got1.([]int), i)
			eachRevUntilOnOneItemCase.got2 = append(eachRevUntilOnOneItemCase.got2.([]int), p.Val())
			return true
		},
	}

	// eachRevUntilFinishMiddleCase:

	eachRevUntilFinishMiddleCase := &baseMapTestCase{
		name:  "EachRevUntil() finish in middle",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{2, 1},
		want2: []int{333, 222},
	}

	eachRevUntilFinishMiddleCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if i < 0 || i > 2 || p.Val() < 111 || p.Val() > 333 {
				t.Error("EachRevUntil() called with wrong values")
			}
			eachRevUntilFinishMiddleCase.got1 = append(eachRevUntilFinishMiddleCase.got1.([]int), i)
			eachRevUntilFinishMiddleCase.got2 = append(eachRevUntilFinishMiddleCase.got2.([]int), p.Val())
			stop := i <= 1 && p.Val() <= 222
			cont := !stop
			return cont
		},
	}

	// eachRevUntilAllThreeCase:

	eachRevUntilAllThreeCase := &baseMapTestCase{
		name:  "EachRevUntil() on three-item collection",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{2, 1, 0},
		want2: []int{333, 222, 111},
	}

	eachRevUntilAllThreeCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if i < 0 || i > 2 || p.Val() < 111 || p.Val() > 333 {
				t.Error("EachRevUntil() called with wrong values")
			}
			eachRevUntilAllThreeCase.got1 = append(eachRevUntilAllThreeCase.got1.([]int), i)
			eachRevUntilAllThreeCase.got2 = append(eachRevUntilAllThreeCase.got2.([]int), p.Val())
			return true
		},
	}

	// put the cases together:

	return []*baseMapTestCase{
		eachRevUntilOnEmptyListCase,
		eachRevUntilOnOneItemCase,
		eachRevUntilFinishMiddleCase,
		eachRevUntilAllThreeCase,
	}
}

func testMapEachRevUntil(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapEachRevUntilCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.EachRevUntil(tt.args.predicate)
			if tt.got1 != nil && !reflect.DeepEqual(tt.got1, tt.want1) {
				t.Errorf("EachRevUntil() called with wrong indices: %v, want1 = %v", tt.got1, tt.want1)
			}
			if tt.got2 != nil && !reflect.DeepEqual(tt.got2, tt.want2) {
				t.Errorf("EachRevUntil() called with wrong values: %v, want1 = %v", tt.got2, tt.want2)
			}
		})
	}
}

func getMapFindCases(t *testing.T, builder baseMapCollIntBuilder) []*baseMapTestCase {

	// findOnEmptyListCase:
	findOnEmptyListCase := &baseMapTestCase{
		name:  "Find() on empty collection",
		coll:  builder.Empty(),
		want1: nil,
	}

	findOnEmptyListCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if p.Val() == 111 {
				findOnEmptyListCase.want1 = p
				findOnEmptyListCase.want2 = true
				return true
			}
			return false
		},
		defaultValue: nil,
	}

	// findOnOneItemCase:
	findOnOneItemCase := &baseMapTestCase{
		name:  "Find() on one-item collection",
		coll:  builder.One(),
		want1: NewPair(1, 111),
	}

	findOnOneItemCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if p.Val() == 111 {
				findOnOneItemCase.want1 = p
				findOnOneItemCase.want2 = true
				return true
			}
			return false
		},
		defaultValue: nil,
	}

	// findOnThreeItemCase:
	findOnThreeItemCase := &baseMapTestCase{
		name:  "Find() on three-item collection",
		coll:  builder.Three(),
		want1: NewPair(2, 333),
	}

	findOnThreeItemCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if p.Val() == 333 {
				findOnThreeItemCase.want1 = p
				findOnThreeItemCase.want2 = true
				return true
			}
			return false
		},
		defaultValue: nil,
	}

	// findFirstOnSixWithDupes:
	findFirstOnSixWithDupes := &baseMapTestCase{
		name:  "Find() first on six-item collection",
		coll:  builder.SixWithDuplicates(),
		want1: NewPair(0, 111),
	}
	findFirstOnSixWithDupes.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if p.Val() == 111 {
				findFirstOnSixWithDupes.want1 = p
				findFirstOnSixWithDupes.want2 = true
				return true
			}
			return false
		},
		defaultValue: nil,
	}

	// findSecondOnSixWithDupes:
	findSecondOnSixWithDupes := &baseMapTestCase{
		name:  "Find() second on six-item collection",
		coll:  builder.SixWithDuplicates(),
		want1: NewPair(1, 222),
	}
	findSecondOnSixWithDupes.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if p.Val() == 222 {
				findSecondOnSixWithDupes.want1 = p
				findSecondOnSixWithDupes.want2 = true
				return true
			}
			return false
		},
		defaultValue: nil,
	}

	// put the cases together:

	return []*baseMapTestCase{
		findOnEmptyListCase,
		findOnOneItemCase,
		findOnThreeItemCase,
		findFirstOnSixWithDupes,
		findSecondOnSixWithDupes,
	}
}

func testMapFind(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapFindCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.Find(tt.args.predicate, tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Find() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapFindLastCases(t *testing.T, builder baseMapCollIntBuilder) []*baseMapTestCase {

	// findLastOnEmptyListCase:
	findLastOnEmptyListCase := &baseMapTestCase{
		name:  "FindLast() on empty collection",
		coll:  builder.Empty(),
		want1: nil,
	}

	findLastOnEmptyListCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if p.Val() == 111 {
				findLastOnEmptyListCase.want1 = p
				findLastOnEmptyListCase.want2 = true
				return true
			}
			return false
		},
		defaultValue: nil,
	}

	// findLastOnOneItemCase:
	findLastOnOneItemCase := &baseMapTestCase{
		name:  "FindLast() on one-item collection",
		coll:  builder.One(),
		want1: NewPair(1, 111),
	}

	findLastOnOneItemCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if p.Val() == 111 {
				findLastOnOneItemCase.want1 = p
				findLastOnOneItemCase.want2 = true
				return true
			}
			return false
		},
		defaultValue: nil,
	}

	// findLastOnThreeItemCase:
	findLastOnThreeItemCase := &baseMapTestCase{
		name:  "FindLast() on three-item collection",
		coll:  builder.Three(),
		want1: NewPair(3, 333),
	}

	findLastOnThreeItemCase.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if p.Val() == 333 {
				findLastOnThreeItemCase.want1 = p
				findLastOnThreeItemCase.want2 = true
				return true
			}
			return false
		},
		defaultValue: nil,
	}

	// findFirstOnSixWithDupes:
	findFirstOnSixWithDupes := &baseMapTestCase{
		name:  "FindLast() first on six-item collection",
		coll:  builder.SixWithDuplicates(),
		want1: NewPair(3, 111),
	}
	findFirstOnSixWithDupes.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if p.Val() == 111 {
				findFirstOnSixWithDupes.want1 = p
				findFirstOnSixWithDupes.want2 = true
				return true
			}
			return false
		},
		defaultValue: nil,
	}

	// findSecondOnSixWithDupes:
	findSecondOnSixWithDupes := &baseMapTestCase{
		name:  "FindLast() second on six-item collection",
		coll:  builder.SixWithDuplicates(),
		want1: NewPair(4, 222),
	}
	findSecondOnSixWithDupes.args = baseMapIntArgs{
		predicate: func(i int, p Pair[int, int]) bool {
			if p.Val() == 222 {
				findSecondOnSixWithDupes.want1 = p
				findSecondOnSixWithDupes.want2 = true
				return true
			}
			return false
		},
		defaultValue: nil,
	}

	// put the cases together:

	return []*baseMapTestCase{
		findLastOnEmptyListCase,
		findLastOnOneItemCase,
		findLastOnThreeItemCase,
		findFirstOnSixWithDupes,
		findSecondOnSixWithDupes,
	}
}

func testMapFindLast(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapFindLastCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.FindLast(tt.args.predicate, tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("FindLast() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapFoldCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name: "Fold() on empty collection",
			coll: builder.Empty(),
			args: baseMapIntArgs{
				reducer: func(acc Pair[int, int], i int, current Pair[int, int]) Pair[int, int] {
					return acc
				},
				initial: nil,
			},
			want1: nil,
		},
		{
			name: "Fold() on one-item collection",
			coll: builder.One(),
			args: baseMapIntArgs{
				reducer: func(acc Pair[int, int], i int, current Pair[int, int]) Pair[int, int] {
					return current
				},
				initial: nil,
			},
			want1: NewPair(1, 111),
		},
		{
			name: "Fold() on three-item collection",
			coll: builder.Three(),
			args: baseMapIntArgs{
				reducer: func(acc Pair[int, int], i int, current Pair[int, int]) Pair[int, int] {
					if acc == nil {
						return current
					}
					return NewPair(acc.Key()+current.Key(), acc.Val()+current.Val())
				},
				initial: nil,
			},
			want1: NewPair(6, 666),
		},
		{
			name: "Fold() on six-item collection",
			coll: builder.SixWithDuplicates(),
			args: baseMapIntArgs{
				reducer: func(acc Pair[int, int], i int, current Pair[int, int]) Pair[int, int] {
					return NewPair(acc.Key()+current.Key(), acc.Val()+current.Val())
				},
				initial: NewPair(0, 0),
			},
			want1: NewPair(21, 1332),
		},
	}
}

func testMapFold(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapFoldCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.Fold(tt.args.reducer, tt.args.initial)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Fold() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapGetCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Get() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{key: 1},
			want1: 0,
			want2: false,
		},
		{
			name:  "Get() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{key: 1},
			want1: 111,
			want2: true,
		},
		{
			name:  "Get() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{key: 2},
			want1: 222,
			want2: true,
		},
		{
			name:  "Get() on three-item collection, not found",
			coll:  builder.Three(),
			args:  baseMapIntArgs{key: 999},
			want1: 0,
			want2: false,
		},
	}
}

func testMapGet(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapGetCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.Get(tt.args.key)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Get() got1 = %v, want1 = %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Get() got2 = %v, want2 = %v", got2, tt.want2)
			}
		})
	}
}

func getMapGetOrDefaultCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "GetOrDefault() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{key: 1, defaultRawValue: 999},
			want1: 999,
		},
		{
			name:  "GetOrDefault() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{key: 1, defaultRawValue: 999},
			want1: 111,
		},
		{
			name:  "GetOrDefault() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{key: 2, defaultRawValue: 999},
			want1: 222,
		},
		{
			name:  "GetOrDefault() on three-item collection, not found",
			coll:  builder.Three(),
			args:  baseMapIntArgs{key: 999, defaultRawValue: 999},
			want1: 999,
		},
	}
}

func testMapGetOrDefault(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapGetOrDefaultCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1 := tt.coll.GetOrDefault(tt.args.key, tt.args.defaultRawValue.(int))
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetOrDefault() got1 = %v, want1 = %v", got1, tt.want1)
			}
		})
	}
}

func getMapHeadCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Head() on empty collection",
			coll:  builder.Empty(),
			want1: nil,
			want2: false,
		},
		{
			name:  "Head() on one-item collection",
			coll:  builder.One(),
			want1: NewPair(1, 111),
			want2: true,
		},
		{
			name:  "Head() on three-item collection",
			coll:  builder.Three(),
			want1: NewPair(1, 111),
			want2: true,
		},
	}
}

func testMapHead(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapHeadCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.Head()
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Head() got1 = %v, want1 = %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Head() got2 = %v, want2 = %v", got2, tt.want2)
			}
		})
	}
}

func getMapHeadOrDefaultCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "HeadOrDefault() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{defaultValue: NewPair(999, 999)},
			want1: NewPair(999, 999),
		},
		{
			name:  "HeadOrDefault() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{defaultValue: NewPair(999, 999)},
			want1: NewPair(1, 111),
		},
		{
			name:  "HeadOrDefault() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{defaultValue: NewPair(999, 999)},
			want1: NewPair(1, 111),
		},
	}
}

func testMapHeadOrDefault(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapHeadOrDefaultCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1 := tt.coll.HeadOrDefault(tt.args.defaultValue)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("HeadOrDefault() got1 = %v, want1 = %v", got1, tt.want1)
			}
		})
	}
}

func getMapIsEmptyCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "IsEmpty() on empty collection",
			coll:  builder.Empty(),
			want1: true,
		},
		{
			name:  "IsEmpty() on one-item collection",
			coll:  builder.One(),
			want1: false,
		},
		{
			name:  "IsEmpty() on three-item collection",
			coll:  builder.Three(),
			want1: false,
		},
	}
}

func testMapIsEmpty(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapIsEmptyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.IsEmpty()
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("IsEmpty() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapKeysCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Keys() on empty collection",
			coll:  builder.Empty(),
			want1: []int{},
		},
		{
			name:  "Keys() on one-item collection",
			coll:  builder.One(),
			want1: []int{1},
		},
		{
			name:  "Keys() on three-item collection",
			coll:  builder.Three(),
			want1: []int{1, 2, 3},
		},
	}
}

func testMapKeys(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapKeysCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			for k := range tt.coll.Keys() {
				got = append(got, k)
			}
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Keys() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapKeysBreakCases(builder baseMapCollIntBuilder) []*baseMapTestCase {
	return []*baseMapTestCase{
		{
			name: "Keys() on three-item collection, break immediately",
			coll: builder.Three(),
			args: baseMapIntArgs{
				intPredicate: func(k int) bool {
					return false
				},
			},
			want1: []int{},
		},
		{
			name: "Keys() on three-item collection, break at middle",
			coll: builder.Three(),
			args: baseMapIntArgs{
				intPredicate: func(k int) bool {
					return k < 2
				},
			},
			want1: []int{1},
		},
		{
			name: "Keys() on three-item collection, break after middle",
			coll: builder.Three(),
			args: baseMapIntArgs{
				intPredicate: func(k int) bool {
					return k <= 2
				},
			},
			want1: []int{1, 2},
		},
	}
}

func testMapKeysBreak(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapKeysBreakCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			for k := range tt.coll.Keys() {
				if !tt.args.intPredicate(k) {
					break
				}
				got = append(got, k)
			}
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Keys() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func testMapKeysToSlice(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapKeysCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			for _, k := range tt.coll.KeysToSlice() {
				got = append(got, k)
			}
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("KeysToSlice() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapKeyValuesBreakCases(builder baseMapCollIntBuilder) []*baseMapTestCase {
	return []*baseMapTestCase{
		{
			name: "KeyValues() on three-item collection, break immediately",
			coll: builder.Three(),
			args: baseMapIntArgs{
				intPredicate: func(_ int) bool {
					return false
				},
			},
			want1: []int{},
			want2: []int{},
		},
		{
			name: "KeyValues() on three-item collection, break at middle",
			coll: builder.Three(),
			args: baseMapIntArgs{
				intPredicate: func(v int) bool {
					return v < 222
				},
			},
			want1: []int{1},
			want2: []int{111},
		},
		{
			name: "KeyValues() on three-item collection, break after middle",
			coll: builder.Three(),
			args: baseMapIntArgs{
				intPredicate: func(v int) bool {
					return v <= 222
				},
			},
			want1: []int{1, 2},
			want2: []int{111, 222},
		},
		{
			name: "KeyValues() on three-item collection,, don't break",
			coll: builder.Three(),
			args: baseMapIntArgs{
				intPredicate: func(_ int) bool {
					return true
				},
			},
			want1: []int{1, 2, 3},
			want2: []int{111, 222, 333},
		},
	}
}

func testMapKeyValuesBreak(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapKeyValuesBreakCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1 := []int{}
			got2 := []int{}
			for k, v := range tt.coll.KeyValues() {
				if !tt.args.intPredicate(v) {
					break
				}
				got1 = append(got1, k)
				got2 = append(got2, v)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("KeyValues() got1 = %v, want1 = %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("KeyValues() got2 = %v, want2 = %v", got2, tt.want2)
			}
		})
	}
}

func getMapLenCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Len() on empty collection",
			coll:  builder.Empty(),
			want1: 0,
		},
		{
			name:  "Len() on one-item collection",
			coll:  builder.One(),
			want1: 1,
		},
		{
			name:  "Len() on three-item collection",
			coll:  builder.Three(),
			want1: 3,
		},
	}
}

func testMapLen(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapLenCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.Len()
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Len() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapReduceCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	sumReducer := func(acc Pair[int, int], i int, current Pair[int, int]) Pair[int, int] {
		return NewPair(acc.Key()+current.Key(), acc.Val()+current.Val())
	}

	return []baseMapTestCase{
		{
			name:  "Reduce() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{reducer: sumReducer},
			want1: nil,
			want2: ErrEmptyCollection,
		},
		{
			name:  "Reduce() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{reducer: sumReducer},
			want1: NewPair(1, 111),
			want2: nil,
		},
		{
			name:  "Reduce() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{reducer: sumReducer},
			want1: NewPair(6, 666),
			want2: nil,
		},
	}
}

func testMapReduce(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapReduceCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.Reduce(tt.args.reducer)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Reduce() got1 = %v, want1 = %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Reduce() got2 = %v, want2 = %v", got2, tt.want2)
			}
		})
	}
}

func getMapRemoveAtCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "RemoveAt() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{index: 0},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			err:   ErrOutOfBounds,
		},
		{
			name:  "RemoveAt() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{index: 0},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: NewPair(1, 111),
		},
		{
			name:  "RemoveAt() on three-item collection at beginning",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 0},
			want1: []Pair[int, int]{NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{2: 222, 3: 333},
			want3: NewPair(1, 111),
		},
		{
			name:  "RemoveAt() on three-item collection in the middle",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 1},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(3, 333)},
			want2: map[int]int{1: 111, 3: 333},
			want3: NewPair(2, 222),
		},
		{
			name:  "RemoveAt() on three-item collection at end",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 2},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222)},
			want2: map[int]int{1: 111, 2: 222},
			want3: NewPair(3, 333),
		},
		{
			name:  "RemoveAt() on three-item collection out of bounds",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 4},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: nil,
			err:   ErrOutOfBounds,
		},
		{
			name:  "RemoveAt() on three-item collection negative index",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: -1},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: nil,
			err:   ErrOutOfBounds,
		},
	}
}

func testMapRemoveAt(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapRemoveAtCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			removed, err := tt.coll.RemoveAt(tt.args.index)
			actualSlice := tt.coll.ToSlice()
			actualMap := tt.coll.ToMap()
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("RemoveAt() did not remove correctly from slice")
			}
			if !reflect.DeepEqual(actualMap, tt.want2) {
				t.Errorf("RemoveAt() did not remove correctly from map")
			}
			if !reflect.DeepEqual(removed, tt.want3) {
				fmt.Println(removed)
				fmt.Println(tt.want3)
				t.Errorf("RemoveAt() did not return removed value correctly")
			}
			if tt.err != nil {
				if !errors.Is(err, tt.err) {
					t.Errorf("RemoveAt() returned wrong error: %v, want error: %v", err, tt.err)
				}
			}
		})
	}
}
