package coll

import (
	"reflect"
	"testing"
)

type orderedMutableIntArgs = testArgs[CmpMutable[int], int]
type orderedMutableTestCase = testCase[CmpMutable[int], int]
type orderedMutableCollIntBuilder = testCollectionBuilder[CmpMutable[int]]

func getRemoveValuesCases(builder orderedMutableCollIntBuilder) []orderedMutableTestCase {
	return []orderedMutableTestCase{
		{
			name:  "RemoveValues() on empty collection",
			coll:  builder.Empty(),
			args:  orderedMutableIntArgs{value: 1},
			want1: []int(nil),
			want2: map[int]int{},
		},
		{
			name:  "RemoveValues() on one-item collection",
			coll:  builder.One(),
			args:  orderedMutableIntArgs{value: 111},
			want1: []int(nil),
			want2: map[int]int{},
		},
		{
			name:  "RemoveValues() on three-item collection - first item",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{value: 111},
			want1: []int{222, 333},
			want2: map[int]int{222: 1, 333: 1},
		},
		{
			name:  "RemoveValues() on three-item collection - second item",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{value: 222},
			want1: []int{111, 333},
			want2: map[int]int{111: 1, 333: 1},
		},
		{
			name:  "RemoveValues() on three-item collection - third item",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{value: 333},
			want1: []int{111, 222},
			want2: map[int]int{111: 1, 222: 1},
		},
		{
			name:  "RemoveValues() on three-item collection, not found",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{value: 999},
			want1: []int{111, 222, 333},
			want2: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "RemoveValues() on six-item collection, 2 `111` found ",
			coll:  builder.SixWithDuplicates(),
			args:  orderedMutableIntArgs{value: 111},
			want1: []int{222, 333, 222, 333},
			want2: map[int]int{222: 2, 333: 2},
		},
		{
			name:  "RemoveValues() on six-item collection, 2 `222` found ",
			coll:  builder.SixWithDuplicates(),
			args:  orderedMutableIntArgs{value: 222},
			want1: []int{111, 333, 111, 333},
			want2: map[int]int{111: 2, 333: 2},
		},
		{
			name:  "RemoveValues() on six-item collection, 2 `333` found ",
			coll:  builder.SixWithDuplicates(),
			args:  orderedMutableIntArgs{value: 333},
			want1: []int{111, 222, 111, 222},
			want2: map[int]int{111: 2, 222: 2},
		},
	}
}

func testRemoveValues(t *testing.T, builder orderedMutableCollIntBuilder) {
	cases := getRemoveValuesCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.RemoveValues(tt.args.value)
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("RemoveValues() resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
			if !reflect.DeepEqual(actualVC, tt.want2) {
				t.Errorf("RemoveValues() did not remove correctly from values counter")
			}
		})
	}
}

func getSortAscCases[C any](builder testCollectionBuilder[C]) []testCase[C, int] {

	sortOnEmptyCollectionCase := testCase[C, int]{
		name:  "SortAsc() on empty collection",
		coll:  builder.Empty(),
		want1: []int(nil),
		want2: map[int]int{},
	}

	sortOnOneItemCollectionCase := testCase[C, int]{
		name:  "SortAsc() on one-item collection",
		coll:  builder.One(),
		want1: []int{111},
		want2: map[int]int{111: 1},
	}

	sortOnThreeItemCollectionCase := testCase[C, int]{
		name:  "SortAsc() on three-item collection",
		coll:  builder.Three(),
		want1: []int{111, 222, 333},
		want2: map[int]int{111: 1, 222: 1, 333: 1},
	}

	sortOnSixItemCollectionCase := testCase[C, int]{
		name:  "SortAsc() on six-item collection",
		coll:  builder.SixWithDuplicates(),
		want1: []int{111, 111, 222, 222, 333, 333},
		want2: map[int]int{111: 2, 222: 2, 333: 2},
	}

	sortOnThreeItemCollectionReversedCase := testCase[C, int]{
		name:  "SortAsc() on three-item collection reversed",
		coll:  builder.ThreeRev(),
		want1: []int{111, 222, 333},
		want2: map[int]int{111: 1, 222: 1, 333: 1},
	}

	sortOnSixItemCollectionReversedCase := testCase[C, int]{
		name:  "SortAsc() on six-item collection reversed",
		coll:  builder.SixWithDuplicates(),
		want1: []int{111, 111, 222, 222, 333, 333},
		want2: map[int]int{111: 2, 222: 2, 333: 2},
	}

	return []testCase[C, int]{
		sortOnEmptyCollectionCase,
		sortOnOneItemCollectionCase,
		sortOnThreeItemCollectionCase,
		sortOnSixItemCollectionCase,
		sortOnThreeItemCollectionReversedCase,
		sortOnSixItemCollectionReversedCase,
	}
}

func testSortAsc[C cmpMutableInternal[int]](t *testing.T, builder testCollectionBuilder[C]) {
	cases := getSortAscCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.SortAsc()
			actualSlice := builder.extractRawValues(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("SortAsc() resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
			if !reflect.DeepEqual(actualVC, tt.want2) {
				t.Errorf("SortAsc() did not sort correctly from values counter")
			}
		})
	}
}

func getSortDescCases[C any](builder testCollectionBuilder[C]) []testCase[C, int] {
	sortOnEmptyCollectionCase := testCase[C, int]{
		name:  "SortDesc() on empty collection",
		coll:  builder.Empty(),
		want1: []int(nil),
		want2: map[int]int{},
	}

	sortOnOneItemCollectionCase := testCase[C, int]{
		name:  "SortDesc() on one-item collection",
		coll:  builder.One(),
		want1: []int{111},
		want2: map[int]int{111: 1},
	}

	sortOnThreeItemCollectionCase := testCase[C, int]{
		name:  "SortDesc() on three-item collection",
		coll:  builder.Three(),
		want1: []int{333, 222, 111},
		want2: map[int]int{111: 1, 222: 1, 333: 1},
	}

	sortOnSixItemCollectionCase := testCase[C, int]{
		name:  "SortDesc() on six-item collection",
		coll:  builder.SixWithDuplicates(),
		want1: []int{333, 333, 222, 222, 111, 111},
		want2: map[int]int{111: 2, 222: 2, 333: 2},
	}

	sortOnThreeItemCollectionReversedCase := testCase[C, int]{
		name:  "SortDesc() on three-item collection reversed",
		coll:  builder.ThreeRev(),
		want1: []int{333, 222, 111},
		want2: map[int]int{111: 1, 222: 1, 333: 1},
	}

	sortOnSixItemCollectionReversedCase := testCase[C, int]{
		name:  "SortDesc() on six-item collection reversed",
		coll:  builder.SixWithDuplicates(),
		want1: []int{333, 333, 222, 222, 111, 111},
		want2: map[int]int{111: 2, 222: 2, 333: 2},
	}

	return []testCase[C, int]{
		sortOnEmptyCollectionCase,
		sortOnOneItemCollectionCase,
		sortOnThreeItemCollectionCase,
		sortOnSixItemCollectionCase,
		sortOnThreeItemCollectionReversedCase,
		sortOnSixItemCollectionReversedCase,
	}
}

func testSortDesc[C cmpMutableInternal[int]](t *testing.T, builder testCollectionBuilder[C]) {
	cases := getSortDescCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.SortDesc()
			actualSlice := builder.extractRawValues(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("SortAsc() resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
			if !reflect.DeepEqual(actualVC, tt.want2) {
				t.Errorf("SortAsc() did not sort correctly from values counter")
			}
		})
	}
}
