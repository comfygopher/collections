package coll

import (
	"reflect"
	"testing"
)

type linearMutableIntArgs = testArgs[linearMutableInternal[int], int]
type linearMutableTestCase = testCase[linearMutableInternal[int], int]
type linearMutableCollIntBuilder = testCollectionBuilder[linearMutableInternal[int], int]

func getAppendOneCases(builder linearMutableCollIntBuilder) []linearMutableTestCase {
	return []linearMutableTestCase{
		{
			name:  "Append() on empty collection",
			coll:  builder.Empty(),
			args:  linearMutableIntArgs{value: 111},
			want1: []int{111},
		},
		{
			name:  "Append() on one-item collection",
			coll:  builder.One(),
			args:  linearMutableIntArgs{value: 1},
			want1: []int{111, 1},
		},
		{
			name:  "Append() on three-item collection",
			coll:  builder.Three(),
			args:  linearMutableIntArgs{value: 4},
			want1: []int{111, 222, 333, 4},
		},
	}
}

func testAppendOne(t *testing.T, builder linearMutableCollIntBuilder) {
	cases := getAppendOneCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Append(tt.args.value)
			if !reflect.DeepEqual(tt.coll.ToSlice(), tt.want1) {
				t.Errorf("Append() resulted in: %v, but wanted %v", tt.coll.ToSlice(), tt.want1)
			}
		})
	}
}

func getAppendManyCases(builder linearMutableCollIntBuilder) []linearMutableTestCase {
	return []linearMutableTestCase{
		{
			name:  "Append() on empty collection",
			coll:  builder.Empty(),
			args:  linearMutableIntArgs{values: []int{111}},
			want1: []int{111},
		},
		{
			name:  "Append() on one-item collection",
			coll:  builder.One(),
			args:  linearMutableIntArgs{values: []int{1}},
			want1: []int{111, 1},
		},
		{
			name:  "Append() on three-item collection",
			coll:  builder.Three(),
			args:  linearMutableIntArgs{values: []int{4, 5}},
			want1: []int{111, 222, 333, 4, 5},
		},
		{
			name:  "Append() on none",
			coll:  builder.Three(),
			args:  linearMutableIntArgs{values: []int{}},
			want1: []int{111, 222, 333},
		},
		{
			name:  "Append() none to empty collection",
			coll:  builder.Empty(),
			args:  linearMutableIntArgs{values: []int{}},
			want1: []int{},
		},
	}
}

func testAppendMany(t *testing.T, builder linearMutableCollIntBuilder) {
	cases := getAppendManyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Append(tt.args.values...)
			if !reflect.DeepEqual(tt.coll.ToSlice(), tt.want1) {
				t.Errorf("Append() resulted in: %v, but wanted %v", tt.coll.ToSlice(), tt.want1)
			}
		})
	}
}

func getAppendCollCases(builder linearMutableCollIntBuilder) []linearMutableTestCase {
	return []linearMutableTestCase{
		{
			name:  "Append() on empty collection",
			coll:  builder.Empty(),
			args:  linearMutableIntArgs{coll: builder.One()},
			want1: []int{111},
		},
		{
			name:  "Append() on one-item collection",
			coll:  builder.One(),
			args:  linearMutableIntArgs{coll: builder.One()},
			want1: []int{111, 111},
		},
		{
			name:  "Append() on three-item collection",
			coll:  builder.Three(),
			args:  linearMutableIntArgs{coll: builder.Two()},
			want1: []int{111, 222, 333, 123, 234},
		},
		{
			name:  "Append() empty collection on three-item collection",
			coll:  builder.Three(),
			args:  linearMutableIntArgs{coll: builder.Empty()},
			want1: []int{111, 222, 333},
		},
		{
			name:  "Append() empty collection on empty collection",
			coll:  builder.Empty(),
			args:  linearMutableIntArgs{coll: builder.Empty()},
			want1: []int{},
		},
	}
}

func testAppendColl(t *testing.T, builder linearMutableCollIntBuilder) {
	cases := getAppendCollCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.AppendColl(tt.args.coll)
			if !reflect.DeepEqual(tt.coll.ToSlice(), tt.want1) {
				t.Errorf("Append() resulted in: %v, but wanted %v", tt.coll.ToSlice(), tt.want1)
			}
		})
	}
}

func getPrependOneCases(builder linearMutableCollIntBuilder) []linearMutableTestCase {
	return []linearMutableTestCase{
		{
			name:  "Prepend() on empty collection",
			coll:  builder.Empty(),
			args:  linearMutableIntArgs{value: 1},
			want1: []int{1},
		},
		{
			name:  "Prepend() on one-item collection",
			coll:  builder.One(),
			args:  linearMutableIntArgs{value: 1},
			want1: []int{1, 111},
		},
		{
			name:  "Prepend() on three-item collection",
			coll:  builder.Three(),
			args:  linearMutableIntArgs{value: 4},
			want1: []int{4, 111, 222, 333},
		},
	}
}

func testPrependOne(t *testing.T, builder linearMutableCollIntBuilder) {
	cases := getPrependOneCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Prepend(tt.args.value)
			if !reflect.DeepEqual(tt.coll.ToSlice(), tt.want1) {
				t.Errorf("Prepend() resulted in: %v, but wanted %v", tt.coll.ToSlice(), tt.want1)
			}
		})
	}
}

func getPrependManyCases(builder linearMutableCollIntBuilder) []linearMutableTestCase {
	return []linearMutableTestCase{
		{
			name:  "Prepend() on empty collection",
			coll:  builder.Empty(),
			args:  linearMutableIntArgs{values: []int{1}},
			want1: []int{1},
		},
		{
			name:  "Prepend() on one-item collection",
			coll:  builder.One(),
			args:  linearMutableIntArgs{values: []int{1}},
			want1: []int{1, 111},
		},
		{
			name:  "Prepend() on three-item collection",
			coll:  builder.Three(),
			args:  linearMutableIntArgs{values: []int{4, 5}},
			want1: []int{4, 5, 111, 222, 333},
		},
		{
			name:  "Prepend() on none",
			coll:  builder.Three(),
			args:  linearMutableIntArgs{values: []int{}},
			want1: []int{111, 222, 333},
		},
		{
			name:  "Prepend() none to empty collection",
			coll:  builder.Empty(),
			args:  linearMutableIntArgs{values: []int{}},
			want1: []int{},
		},
	}
}

func testPrependMany(t *testing.T, builder linearMutableCollIntBuilder) {
	cases := getPrependManyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Prepend(tt.args.values...)
			if !reflect.DeepEqual(tt.coll.ToSlice(), tt.want1) {
				t.Errorf("Prepend() resulted in: %v, but wanted %v", tt.coll.ToSlice(), tt.want1)
			}
		})
	}
}

func getReverseCases(builder linearMutableCollIntBuilder) []linearMutableTestCase {
	return []linearMutableTestCase{
		{
			name:  "Reverse() on empty collection",
			coll:  builder.Empty(),
			want1: []int{},
		},
		{
			name:  "Reverse() on one-item collection",
			coll:  builder.One(),
			want1: []int{111},
		},
		{
			name:  "Reverse() on three-item collection",
			coll:  builder.Three(),
			want1: []int{333, 222, 111},
		},
	}
}

func testReverse(t *testing.T, builder linearMutableCollIntBuilder) {
	cases := getReverseCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Reverse()
			if !reflect.DeepEqual(tt.coll.ToSlice(), tt.want1) {
				t.Errorf("Reverse() resulted in: %v, but wanted %v", tt.coll.ToSlice(), tt.want1)
			}
		})
	}
}

func testReverseTwice(t *testing.T, builder linearMutableCollIntBuilder) {
	t.Run("Reverse() twice", func(t *testing.T) {
		coll := builder.Three()
		coll.Reverse()
		if !reflect.DeepEqual(coll.ToSlice(), []int{333, 222, 111}) {
			t.Errorf("Reverse() twice is not identity")
		}
		coll.Reverse()
		if !reflect.DeepEqual(coll.ToSlice(), []int{111, 222, 333}) {
			t.Errorf("Reverse() twice is not identity")
		}
	})
}
