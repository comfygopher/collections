package coll

import (
	"reflect"
	"testing"
)

type orderedMutableIntArgs = testArgs[orderedMutableInternal[int], int]
type orderedMutableTestCase = testCase[orderedMutableInternal[int], int]
type orderedMutableCollIntBuilder = testCollectionBuilder[orderedMutableInternal[int]]

func getAppendOneCases(builder orderedMutableCollIntBuilder) []orderedMutableTestCase {
	return []orderedMutableTestCase{
		{
			name:  "Append() on empty collection",
			coll:  builder.Empty(),
			args:  orderedMutableIntArgs{value: 111},
			want1: []int{111},
			want2: map[int]int{111: 1},
		},
		{
			name:  "Append() on one-item collection",
			coll:  builder.One(),
			args:  orderedMutableIntArgs{value: 1},
			want1: []int{111, 1},
			want2: map[int]int{111: 1, 1: 1},
		},
		{
			name:  "Append() on three-item collection",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{value: 4},
			want1: []int{111, 222, 333, 4},
			want2: map[int]int{111: 1, 222: 1, 333: 1, 4: 1},
		},
	}
}

func testAppendOne(t *testing.T, builder orderedMutableCollIntBuilder) {
	cases := getAppendOneCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Append(tt.args.value)

			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Append() one resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
			if actualVC != nil && !reflect.DeepEqual(actualVC, tt.want2) {
				t.Errorf("Append() did not append correctly from values counter")
			}
		})
	}
}

func getAppendManyCases(builder orderedMutableCollIntBuilder) []orderedMutableTestCase {
	return []orderedMutableTestCase{
		{
			name:  "Append() on empty collection",
			coll:  builder.Empty(),
			args:  orderedMutableIntArgs{values: []int{111}},
			want1: []int{111},
			want2: map[int]int{111: 1},
		},
		{
			name:  "Append() on one-item collection",
			coll:  builder.One(),
			args:  orderedMutableIntArgs{values: []int{1}},
			want1: []int{111, 1},
			want2: map[int]int{111: 1, 1: 1},
		},
		{
			name:  "Append() on three-item collection",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{values: []int{4, 5}},
			want1: []int{111, 222, 333, 4, 5},
			want2: map[int]int{111: 1, 222: 1, 333: 1, 4: 1, 5: 1},
		},
		{
			name:  "Append() on none",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{values: []int{}},
			want1: []int{111, 222, 333},
			want2: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "Append() none to empty collection",
			coll:  builder.Empty(),
			args:  orderedMutableIntArgs{values: []int{}},
			want1: []int(nil),
			want2: map[int]int{},
		},
	}
}

func testAppendMany(t *testing.T, builder orderedMutableCollIntBuilder) {
	cases := getAppendManyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Append(tt.args.values...)

			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Append() many resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
			if actualVC != nil && !reflect.DeepEqual(actualVC, tt.want2) {
				t.Errorf("Append() did not append correctly from values counter")
			}
		})
	}
}

func getAppendCollCases(builder orderedMutableCollIntBuilder) []orderedMutableTestCase {
	return []orderedMutableTestCase{
		{
			name:  "Append() on empty collection",
			coll:  builder.Empty(),
			args:  orderedMutableIntArgs{coll: builder.One()},
			want1: []int{111},
			want2: map[int]int{111: 1},
		},
		{
			name:  "Append() on one-item collection",
			coll:  builder.One(),
			args:  orderedMutableIntArgs{coll: builder.One()},
			want1: []int{111, 111},
			want2: map[int]int{111: 2},
		},
		{
			name:  "Append() on three-item collection",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{coll: builder.Two()},
			want1: []int{111, 222, 333, 123, 234},
			want2: map[int]int{111: 1, 222: 1, 333: 1, 123: 1, 234: 1},
		},
		{
			name:  "Append() empty collection on three-item collection",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{coll: builder.Empty()},
			want1: []int{111, 222, 333},
			want2: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "Append() empty collection on empty collection",
			coll:  builder.Empty(),
			args:  orderedMutableIntArgs{coll: builder.Empty()},
			want1: []int(nil),
			want2: map[int]int{},
		},
	}
}

func testAppendColl(t *testing.T, builder orderedMutableCollIntBuilder) {
	cases := getAppendCollCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.AppendColl(tt.args.coll)

			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("AppendColl() resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
			if actualVC != nil && !reflect.DeepEqual(actualVC, tt.want2) {
				t.Errorf("AppendColl() did not append correctly from values counter")
			}
		})
	}
}

func getPrependOneCases(builder orderedMutableCollIntBuilder) []orderedMutableTestCase {
	return []orderedMutableTestCase{
		{
			name:  "Prepend() on empty collection",
			coll:  builder.Empty(),
			args:  orderedMutableIntArgs{value: 1},
			want1: []int{1},
			want2: map[int]int{1: 1},
		},
		{
			name:  "Prepend() on one-item collection",
			coll:  builder.One(),
			args:  orderedMutableIntArgs{value: 1},
			want1: []int{1, 111},
			want2: map[int]int{1: 1, 111: 1},
		},
		{
			name:  "Prepend() on three-item collection",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{value: 4},
			want1: []int{4, 111, 222, 333},
			want2: map[int]int{4: 1, 111: 1, 222: 1, 333: 1},
		},
	}
}

func testPrependOne(t *testing.T, builder orderedMutableCollIntBuilder) {
	cases := getPrependOneCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Prepend(tt.args.value)

			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Prepend() resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
			if actualVC != nil && !reflect.DeepEqual(actualVC, tt.want2) {
				t.Errorf("Prepend() did not append correctly from values counter")
			}
		})
	}
}

func getPrependManyCases(builder orderedMutableCollIntBuilder) []orderedMutableTestCase {
	return []orderedMutableTestCase{
		{
			name:  "Prepend() on empty collection",
			coll:  builder.Empty(),
			args:  orderedMutableIntArgs{values: []int{1}},
			want1: []int{1},
			want2: map[int]int{1: 1},
		},
		{
			name:  "Prepend() on one-item collection",
			coll:  builder.One(),
			args:  orderedMutableIntArgs{values: []int{1}},
			want1: []int{1, 111},
			want2: map[int]int{1: 1, 111: 1},
		},
		{
			name:  "Prepend() on three-item collection",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{values: []int{4, 5}},
			want1: []int{4, 5, 111, 222, 333},
			want2: map[int]int{4: 1, 5: 1, 111: 1, 222: 1, 333: 1},
		},
		{
			name:  "Prepend() on none",
			coll:  builder.Three(),
			args:  orderedMutableIntArgs{values: []int{}},
			want1: []int{111, 222, 333},
			want2: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "Prepend() none to empty collection",
			coll:  builder.Empty(),
			args:  orderedMutableIntArgs{values: []int{}},
			want1: []int(nil),
			want2: map[int]int{},
		},
	}
}

func testPrependMany(t *testing.T, builder orderedMutableCollIntBuilder) {
	cases := getPrependManyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Prepend(tt.args.values...)

			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Prepend() resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
			if actualVC != nil && !reflect.DeepEqual(actualVC, tt.want2) {
				t.Errorf("Prepend() did not append correctly from values counter")
			}
		})
	}
}

func getReverseCases(builder orderedMutableCollIntBuilder) []orderedMutableTestCase {
	return []orderedMutableTestCase{
		{
			name:  "Reverse() on empty collection",
			coll:  builder.Empty(),
			want1: []int(nil),
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

func testReverse(t *testing.T, builder orderedMutableCollIntBuilder) {
	cases := getReverseCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Reverse()

			actualSlice := builder.extractUnderlyingSlice(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Reverse() resulted in: %v, but wanted %v", actualSlice, tt.want1)
			}
		})
	}
}

func testReverseTwice(t *testing.T, builder orderedMutableCollIntBuilder) {
	t.Run("Reverse() twice", func(t *testing.T) {
		coll := builder.Three()
		coll.Reverse()

		actualSlice1 := builder.extractUnderlyingSlice(coll)

		if !reflect.DeepEqual(actualSlice1, []int{333, 222, 111}) {
			t.Errorf("Reverse() twice is not identity")
		}

		coll.Reverse()

		actualSlice2 := builder.extractUnderlyingSlice(coll)

		if !reflect.DeepEqual(actualSlice2, []int{111, 222, 333}) {
			t.Errorf("Reverse() twice is not identity")
		}
	})
}
