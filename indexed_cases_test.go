package coll

import (
	"reflect"
	"testing"
)

type indexedIntArgs = testArgs[Indexed[int], int]
type indexedTestCase = testCase[Indexed[int], int]
type indexedCollIntBuilder = testCollectionBuilder[Indexed[int], int]

func getAtCases(builder indexedCollIntBuilder) []indexedTestCase {
	return []indexedTestCase{
		{
			name:  "At(0) on empty collection",
			coll:  builder.Empty(),
			args:  indexedIntArgs{index: 0},
			want1: 0,
			want2: false,
		},
		{
			name:  "At(0) on one-item collection",
			coll:  builder.One(),
			args:  indexedIntArgs{index: 0},
			want1: 111,
			want2: true,
		},
		{
			name:  "At(1) on three-item collection",
			coll:  builder.Three(),
			args:  indexedIntArgs{index: 1},
			want1: 222,
			want2: true,
		},
		{
			name:  "At(3) on three-item collection out of bounds",
			coll:  builder.Three(),
			args:  indexedIntArgs{index: 3},
			want1: 0,
			want2: false,
		},
		{
			name:  "At(-1) on three-item collection negative index",
			coll:  builder.Three(),
			args:  indexedIntArgs{index: -1},
			want1: 0,
			want2: false,
		},
		{
			name:  "At(1) on one-item collection out of bounds",
			coll:  builder.One(),
			args:  indexedIntArgs{index: 1},
			want1: 0,
			want2: false,
		},
	}
}

func testAt(t *testing.T, builder indexedCollIntBuilder) {
	cases := getAtCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.At(tt.args.index)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("At() got1 = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("At() got2 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func getAtOrDefaultCases(builder indexedCollIntBuilder) []indexedTestCase {
	return []indexedTestCase{
		{
			name:  "AtOrDefault(0) on empty collection",
			coll:  builder.Empty(),
			args:  indexedIntArgs{index: 0, defaultValue: -1},
			want1: -1,
		},
		{
			name:  "AtOrDefault(0) on one-item collection",
			coll:  builder.One(),
			args:  indexedIntArgs{index: 0, defaultValue: -1},
			want1: 111,
		},
		{
			name:  "AtOrDefault(1) on three-item collection",
			coll:  builder.Three(),
			args:  indexedIntArgs{index: 1, defaultValue: -1},
			want1: 222,
		},
		{
			name:  "AtOrDefault(2) on three-item collection",
			coll:  builder.Three(),
			args:  indexedIntArgs{index: 2, defaultValue: -1},
			want1: 333,
		},
		{
			name:  "AtOrDefault(3) on three-item collection out of bounds",
			coll:  builder.Three(),
			args:  indexedIntArgs{index: 3, defaultValue: -1},
			want1: -1,
		},
		{
			name:  "AtOrDefault(-1) on three-item collection negative index",
			coll:  builder.Three(),
			args:  indexedIntArgs{index: -1, defaultValue: -1},
			want1: -1,
		},
		{
			name:  "AtOrDefault(1) on one-item collection out of bounds",
			coll:  builder.One(),
			args:  indexedIntArgs{index: 1, defaultValue: -1},
			want1: -1,
		},
	}
}

func testAtOrDefault(t *testing.T, builder indexedCollIntBuilder) {
	cases := getAtOrDefaultCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.AtOrDefault(tt.args.index, tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("AtOrDefault() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}
