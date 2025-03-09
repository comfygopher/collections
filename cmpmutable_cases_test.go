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
			want1: []int{},
		},
		{
			name:  "RemoveValues() on one-item collection",
			coll:  builder.One(),
			args:  orderedMutableIntArgs{value: 111},
			want1: []int{},
		},
		{
			name:  "RemoveValues() on three-item collection - first item",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{value: 111},
			want1: []int{222, 333},
		},
		{
			name:  "RemoveValues() on three-item collection - second item",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{value: 222},
			want1: []int{111, 333},
		},
		{
			name:  "RemoveValues() on three-item collection - third item",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{value: 333},
			want1: []int{111, 222},
		},
		{
			name:  "RemoveValues() on three-item collection, not found",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{value: 999},
			want1: []int{111, 222, 333},
		},
		{
			name:  "RemoveValues() on six-item collection, 2 `111` found ",
			coll:  builder.SixWithDuplicates(),
			args:  orderedMutableIntArgs{value: 111},
			want1: []int{222, 333, 222, 333},
		},
		{
			name:  "RemoveValues() on six-item collection, 2 `222` found ",
			coll:  builder.SixWithDuplicates(),
			args:  orderedMutableIntArgs{value: 222},
			want1: []int{111, 333, 111, 333},
		},
		{
			name:  "RemoveValues() on six-item collection, 2 `333` found ",
			coll:  builder.SixWithDuplicates(),
			args:  orderedMutableIntArgs{value: 333},
			want1: []int{111, 222, 111, 222},
		},
	}
}

func testRemoveValues(t *testing.T, builder orderedMutableCollIntBuilder) {
	cases := getRemoveValuesCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.RemoveValues(tt.args.value)
			slice := tt.coll.(Base[int]).ToSlice()
			if !reflect.DeepEqual(slice, tt.want1) {
				t.Errorf("RemoveValues() resulted in: %v, but wanted %v", slice, tt.want1)
			}
		})
	}
}

func getSortAscCases[C any](builder testCollectionBuilder[C]) []testCase[C, int] {

	sortOnEmptyCollectionCase := testCase[C, int]{
		name:  "SortAsc() on empty collection",
		coll:  builder.Empty(),
		want1: []int{},
	}

	sortOnOneItemCollectionCase := testCase[C, int]{
		name:  "SortAsc() on one-item collection",
		coll:  builder.One(),
		want1: []int{111},
	}

	sortOnThreeItemCollectionCase := testCase[C, int]{
		name:  "SortAsc() on three-item collection",
		coll:  builder.Three(),
		want1: []int{111, 222, 333},
	}

	sortOnSixItemCollectionCase := testCase[C, int]{
		name:  "SortAsc() on six-item collection",
		coll:  builder.SixWithDuplicates(),
		want1: []int{111, 111, 222, 222, 333, 333},
	}

	sortOnThreeItemCollectionReversedCase := testCase[C, int]{
		name:  "SortAsc() on three-item collection reversed",
		coll:  builder.ThreeRev(),
		want1: []int{111, 222, 333},
	}

	sortOnSixItemCollectionReversedCase := testCase[C, int]{
		name:  "SortAsc() on six-item collection reversed",
		coll:  builder.SixWithDuplicates(),
		want1: []int{111, 111, 222, 222, 333, 333},
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
			slice := builder.extractRawValues(tt.coll)
			if !reflect.DeepEqual(slice, tt.want1) {
				t.Errorf("SortAsc() resulted in: %v, but wanted %v", slice, tt.want1)
			}
		})
	}
}

func getSortDescCases[C any](builder testCollectionBuilder[C]) []testCase[C, int] {
	sortOnEmptyCollectionCase := testCase[C, int]{
		name:  "SortDesc() on empty collection",
		coll:  builder.Empty(),
		want1: []int{},
	}

	sortOnOneItemCollectionCase := testCase[C, int]{
		name:  "SortDesc() on one-item collection",
		coll:  builder.One(),
		want1: []int{111},
	}

	sortOnThreeItemCollectionCase := testCase[C, int]{
		name:  "SortDesc() on three-item collection",
		coll:  builder.Three(),
		want1: []int{333, 222, 111},
	}

	sortOnSixItemCollectionCase := testCase[C, int]{
		name:  "SortDesc() on six-item collection",
		coll:  builder.SixWithDuplicates(),
		want1: []int{333, 333, 222, 222, 111, 111},
	}

	sortOnThreeItemCollectionReversedCase := testCase[C, int]{
		name:  "SortDesc() on three-item collection reversed",
		coll:  builder.ThreeRev(),
		want1: []int{333, 222, 111},
	}

	sortOnSixItemCollectionReversedCase := testCase[C, int]{
		name:  "SortDesc() on six-item collection reversed",
		coll:  builder.SixWithDuplicates(),
		want1: []int{333, 333, 222, 222, 111, 111},
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
			slice := builder.extractRawValues(tt.coll)
			if !reflect.DeepEqual(slice, tt.want1) {
				t.Errorf("SortDesc() resulted in: %v, but wanted %v", slice, tt.want1)
			}
		})
	}
}
