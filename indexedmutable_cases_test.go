package coll

import (
	"errors"
	"reflect"
	"testing"
)

type indexedMutableIntArgs = testArgs[indexedMutableInternal[int], int]
type indexedMutableTestCase = testCase[indexedMutableInternal[int], int]
type indexedMutableCollIntBuilder = testCollectionBuilder[indexedMutableInternal[int]]

func getRemoveAtCases(builder indexedMutableCollIntBuilder) []indexedMutableTestCase {
	return []indexedMutableTestCase{
		{
			name:  "RemoveAt() on empty collection",
			coll:  builder.Empty(),
			args:  indexedMutableIntArgs{index: 0},
			want1: []int(nil),
			want2: 0,
			want3: map[int]int{},
			err:   ErrOutOfBounds,
		},
		{
			name:  "RemoveAt() on one-item collection",
			coll:  builder.One(),
			args:  indexedMutableIntArgs{index: 0},
			want1: []int(nil),
			want2: 111,
			want3: map[int]int{},
		},
		{
			name:  "RemoveAt() on three-item collection at beginning",
			coll:  builder.Three(),
			args:  indexedMutableIntArgs{index: 0},
			want1: []int{222, 333},
			want2: 111,
			want3: map[int]int{222: 1, 333: 1},
		},
		{
			name:  "RemoveAt() on three-item collection at end",
			coll:  builder.Three(),
			args:  indexedMutableIntArgs{index: 2},
			want1: []int{111, 222},
			want3: map[int]int{111: 1, 222: 1},
			want2: 333,
		},
		{
			name:  "RemoveAt() on three-item collection",
			coll:  builder.Three(),
			args:  indexedMutableIntArgs{index: 1},
			want1: []int{111, 333},
			want3: map[int]int{111: 1, 333: 1},
			want2: 222,
		},
		{
			name:  "RemoveAt() on three-item collection out of bounds",
			coll:  builder.Three(),
			args:  indexedMutableIntArgs{index: 4},
			want1: []int{111, 222, 333},
			want2: 0,
			want3: map[int]int{111: 1, 222: 1, 333: 1},
			err:   ErrOutOfBounds,
		},
		{
			name:  "RemoveAt() on three-item collection negative index",
			coll:  builder.Three(),
			args:  indexedMutableIntArgs{index: -1},
			want1: []int{111, 222, 333},
			want2: 0,
			want3: map[int]int{111: 1, 222: 1, 333: 1},
			err:   ErrOutOfBounds,
		},
		{
			name:  "RemoveAt() on six-item with duplicates",
			coll:  builder.SixWithDuplicates(),
			args:  indexedMutableIntArgs{index: 2},
			want1: []int{111, 222, 111, 222, 333},
			want2: 333,
			want3: map[int]int{111: 2, 222: 2, 333: 1},
		},
	}
}

func testRemoveAt(t *testing.T, builder indexedMutableCollIntBuilder) {
	cases := getRemoveAtCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			removed, err := tt.coll.RemoveAt(tt.args.index)
			if tt.err != nil {
				if err == nil {
					t.Errorf("RemoveAt() did not return error")
				}
				if !errors.Is(err, tt.err) {
					t.Errorf("RemoveAt() returned wrong error: %v, want error: %v", err, tt.err)
				}
			}

			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("RemoveAt() resulted in: %v, but wanted = %v", actualSlice, tt.want1)
			}
			if removed != tt.want2 {
				t.Errorf("RemoveAt() removed wrong value: %v, but wanted = %v", removed, tt.want2)
			}
			if actualVC != nil && !reflect.DeepEqual(actualVC, tt.want3) {
				t.Errorf("RemoveAt() did not remove correctly from values counter")
			}
		})
	}
}

func getSortCases(builder indexedMutableCollIntBuilder) []indexedMutableTestCase {
	return []indexedMutableTestCase{
		{
			name:  "Sort() on empty collection",
			coll:  builder.Empty(),
			args:  indexedMutableIntArgs{comparer: func(a, b int) int { return a - b }},
			want1: []int(nil),
		},
		{
			name:  "Sort() on one-item collection",
			coll:  builder.One(),
			args:  indexedMutableIntArgs{comparer: func(a, b int) int { return a - b }},
			want1: []int{111},
		},
		{
			name:  "Sort() on three-item collection",
			coll:  builder.Three(),
			args:  indexedMutableIntArgs{comparer: func(a, b int) int { return a - b }},
			want1: []int{111, 222, 333},
		},
		{
			name:  "Sort() on three-item collection, reverse",
			coll:  builder.Three(),
			args:  indexedMutableIntArgs{comparer: func(a, b int) int { return b - a }},
			want1: []int{333, 222, 111},
		},
	}
}

func testSort(t *testing.T, builder indexedMutableCollIntBuilder) {
	cases := getSortCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Sort(tt.args.comparer)

			actualSlice := builder.extractUnderlyingSlice(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Sort() resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
		})
	}
}
