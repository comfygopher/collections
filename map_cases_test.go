package coll

import (
	"errors"
	"reflect"
	"slices"
	"testing"
)

func i2iToPairs(m any) map[int]Pair[int, int] {
	mp := m.(map[int]int)
	pairs := make(map[int]Pair[int, int])
	for k, v := range mp {
		pairs[k] = NewPair(k, v)
	}
	return pairs
}

type baseMapIntArgs = testArgs[mapInternal[int, int], Pair[int, int]]
type baseMapTestCase = testCase[mapInternal[int, int], Pair[int, int]]
type baseMapCollIntBuilder = testCollectionBuilder[mapInternal[int, int]]

func getAppendCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Append() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{value: NewPair(1, 111)},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "Append() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{value: NewPair(2, 222)},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222)},
			want2: map[int]int{1: 111, 2: 222},
			want3: map[int]int{1: 0, 2: 1},
			want4: map[int]int{111: 1, 222: 1},
		},
		{
			name:  "Append() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{value: NewPair(4, 444)},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333), NewPair(4, 444)},
			want2: map[int]int{1: 111, 2: 222, 3: 333, 4: 444},
			want3: map[int]int{1: 0, 2: 1, 3: 2, 4: 3},
			want4: map[int]int{111: 1, 222: 1, 333: 1, 444: 1},
		},
		{
			name:  "Append() on three-item collection - duplicate key",
			coll:  builder.Three(),
			args:  baseMapIntArgs{value: NewPair(2, 999)},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(3, 333), NewPair(2, 999)},
			want2: map[int]int{1: 111, 3: 333, 2: 999},
			want3: map[int]int{1: 0, 3: 1, 2: 2},
			want4: map[int]int{111: 1, 333: 1, 999: 1},
		},
		{
			name:  "Append() many on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{values: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)}},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "Append() many on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{values: []Pair[int, int]{NewPair(2, 222), NewPair(3, 333)}},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "Append() many on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{values: []Pair[int, int]{NewPair(4, 444), NewPair(5, 555)}},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333), NewPair(4, 444), NewPair(5, 555)},
			want2: map[int]int{1: 111, 2: 222, 3: 333, 4: 444, 5: 555},
			want3: map[int]int{1: 0, 2: 1, 3: 2, 4: 3, 5: 4},
			want4: map[int]int{111: 1, 222: 1, 333: 1, 444: 1, 555: 1},
		},
		{
			name: "Append() many on three-item collection with duplicates",
			coll: builder.Three(),
			args: baseMapIntArgs{values: []Pair[int, int]{
				NewPair(4, 444),
				NewPair(5, 555),
				NewPair(6, 111),
				NewPair(7, 444),
			}},
			want1: []Pair[int, int]{
				NewPair(1, 111),
				NewPair(2, 222),
				NewPair(3, 333),
				NewPair(4, 444),
				NewPair(5, 555),
				NewPair(6, 111),
				NewPair(7, 444),
			},
			want3: map[int]int{1: 0, 2: 1, 3: 2, 4: 3, 5: 4, 6: 5, 7: 6},
			want2: map[int]int{1: 111, 2: 222, 3: 333, 4: 444, 5: 555, 6: 111, 7: 444},
			want4: map[int]int{111: 2, 222: 1, 333: 1, 444: 2, 555: 1},
		},
	}
}

func testMapAppend(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getAppendCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.values != nil {
				tt.coll.Append(tt.args.values...)
			} else {
				tt.coll.Append(tt.args.value)
			}
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKp := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Append() did not append correctly to slice")
			}
			if !reflect.DeepEqual(actualMap, i2iToPairs(tt.want2)) {
				t.Errorf("Append() did not append correctly to map")
			}
			if !reflect.DeepEqual(actualKp, tt.want3) {
				t.Errorf("Append() did not append correctly to kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("Append() did not append correctly to values counter")
				}
			}
		})
	}
}

func getMapAppendCollCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Append() empty collection on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{coll: builder.Empty()},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "Append() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{coll: builder.One()},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "Append() empty collection on one item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{coll: builder.Empty()},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "Append() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{coll: builder.One()},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "Append() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{coll: builder.Two()},
			want1: []Pair[int, int]{NewPair(3, 333), NewPair(1, 111), NewPair(2, 222)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: map[int]int{1: 1, 2: 2, 3: 0},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "Append() empty collection on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{coll: builder.Empty()},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name: "Append() empty collection on three-item collection with duplicates",
			coll: builder.Three(),
			args: baseMapIntArgs{coll: builder.SixWithDuplicates()},
			want1: []Pair[int, int]{
				NewPair(1, 111),
				NewPair(2, 222),
				NewPair(3, 333),
				NewPair(4, 111),
				NewPair(5, 222),
				NewPair(6, 333),
			},
			want2: map[int]int{1: 111, 2: 222, 3: 333, 4: 111, 5: 222, 6: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2, 4: 3, 5: 4, 6: 5},
			want4: map[int]int{111: 2, 222: 2, 333: 2},
		},
	}
}

func testMapAppendColl(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapAppendCollCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.AppendColl(tt.args.coll)
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKP := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Append() did not append correctly to slice")
			}
			if !reflect.DeepEqual(actualMap, i2iToPairs(tt.want2)) {
				t.Errorf("Append() did not append correctly to map")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("Append() did not append correctly to kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("Append() did not append correctly to values counter")
				}
			}
		})
	}
}

func getMapApplyCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Apply() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{mapper: func(p Pair[int, int]) Pair[int, int] { return NewPair(p.Key()+10, p.Val()+1) }},
			want1: []Pair[int, int](nil),
			want2: map[int]Pair[int, int]{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "Apply() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{mapper: func(p Pair[int, int]) Pair[int, int] { return NewPair(p.Key()+10, p.Val()+1) }},
			want1: []Pair[int, int]{NewPair(11, 112)},
			want2: map[int]Pair[int, int]{11: NewPair(11, 112)},
			want3: map[int]int{11: 0},
			want4: map[int]int{112: 1},
		},
		{
			name:  "Apply() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{mapper: func(p Pair[int, int]) Pair[int, int] { return NewPair(p.Key()+10, p.Val()+1) }},
			want1: []Pair[int, int]{NewPair(11, 112), NewPair(12, 223), NewPair(13, 334)},
			want2: map[int]Pair[int, int]{11: NewPair(11, 112), 12: NewPair(12, 223), 13: NewPair(13, 334)},
			want3: map[int]int{11: 0, 12: 1, 13: 2},
			want4: map[int]int{112: 1, 223: 1, 334: 1},
		},
		{
			name: "Apply() on six-item collection with duplicates",
			coll: builder.SixWithDuplicates(),
			args: baseMapIntArgs{mapper: func(p Pair[int, int]) Pair[int, int] { return NewPair(p.Key()+10, p.Val()+1) }},
			want1: []Pair[int, int]{
				NewPair(11, 112),
				NewPair(12, 223),
				NewPair(13, 334),
				NewPair(14, 112),
				NewPair(15, 223),
				NewPair(16, 334),
			},
			want2: map[int]Pair[int, int]{
				11: NewPair(11, 112),
				12: NewPair(12, 223),
				13: NewPair(13, 334),
				14: NewPair(14, 112),
				15: NewPair(15, 223),
				16: NewPair(16, 334),
			},
			want3: map[int]int{11: 0, 12: 1, 13: 2, 14: 3, 15: 4, 16: 5},
			want4: map[int]int{112: 2, 223: 2, 334: 2},
		},
	}
}

func testMapApply(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapApplyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Apply(tt.args.mapper)
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKP := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Apply() did not apply correctly to slice")
			}
			if !reflect.DeepEqual(actualMap, tt.want2) {
				t.Errorf("Apply() did not apply correctly to map")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("Apply() did not apply correctly to kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("Apply() did not apply correctly to values counter")
				}
			}
		})
	}
}

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
			name:  "Clear() on empty collection",
			coll:  builder.Empty(),
			want1: []Pair[int, int](nil),
			want2: make(map[int]Pair[int, int]),
			want3: make(map[int]int),
			want4: make(map[int]int),
		},
		{
			name:  "Clear() on one-item collection",
			coll:  builder.One(),
			want1: []Pair[int, int](nil),
			want2: make(map[int]Pair[int, int]),
			want3: make(map[int]int),
			want4: make(map[int]int),
		},
		{
			name:  "Clear() on three-item collection",
			coll:  builder.Three(),
			want1: []Pair[int, int](nil),
			want2: make(map[int]Pair[int, int]),
			want3: make(map[int]int),
			want4: make(map[int]int),
		},
	}
}

func testMapClear(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapClearCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Clear()
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKP := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Clear() did not clear slice correctly")
			}
			if !reflect.DeepEqual(actualMap, tt.want2) {
				t.Errorf("Clear() did not clear map correctly")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("Clear() did not clear kp correctly")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("Clear() did not clear values counter correctly")
				}
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

func getMapHasCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Has() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{key: 1},
			want1: false,
		},
		{
			name:  "Has() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{key: 1},
			want1: true,
		},
		{
			name:  "Has() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{key: 2},
			want1: true,
		},
		{
			name:  "Has() on three-item collection, not found",
			coll:  builder.Three(),
			args:  baseMapIntArgs{key: 999},
			want1: false,
		},
	}
}

func testMapHas(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapHasCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.Has(tt.args.key)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Has() = %v, want1 = %v", got, tt.want1)
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

func getMapPrependCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Prepend() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{value: NewPair(1, 111)},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]Pair[int, int]{1: NewPair(1, 111)},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "Prepend() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{value: NewPair(2, 222)},
			want1: []Pair[int, int]{NewPair(2, 222), NewPair(1, 111)},
			want2: map[int]Pair[int, int]{2: NewPair(2, 222), 1: NewPair(1, 111)},
			want3: map[int]int{2: 0, 1: 1},
			want4: map[int]int{222: 1, 111: 1},
		},
		{
			name: "Prepend() on three-item collection",
			coll: builder.Three(),
			args: baseMapIntArgs{value: NewPair(4, 444)},
			want1: []Pair[int, int]{
				NewPair(4, 444),
				NewPair(1, 111),
				NewPair(2, 222),
				NewPair(3, 333),
			},
			want2: map[int]Pair[int, int]{
				4: NewPair(4, 444),
				1: NewPair(1, 111),
				2: NewPair(2, 222),
				3: NewPair(3, 333),
			},
			want3: map[int]int{4: 0, 1: 1, 2: 2, 3: 3},
			want4: map[int]int{444: 1, 111: 1, 222: 1, 333: 1},
		},
		{
			name: "Prepend() on three-item collection - duplicate key",
			coll: builder.Three(),
			args: baseMapIntArgs{value: NewPair(2, 999)},
			want1: []Pair[int, int]{
				NewPair(2, 999),
				NewPair(1, 111),
				NewPair(3, 333),
			},
			want2: map[int]Pair[int, int]{
				2: NewPair(2, 999),
				1: NewPair(1, 111),
				3: NewPair(3, 333),
			},
			want3: map[int]int{2: 0, 1: 1, 3: 2},
			want4: map[int]int{999: 1, 111: 1, 333: 1},
		},
		{
			name: "Prepend() many on empty collection",
			coll: builder.Empty(),
			args: baseMapIntArgs{
				values: []Pair[int, int]{
					NewPair(1, 111),
					NewPair(2, 222),
					NewPair(3, 333),
				},
			},
			want1: []Pair[int, int]{
				NewPair(1, 111),
				NewPair(2, 222),
				NewPair(3, 333),
			},
			want2: map[int]Pair[int, int]{
				1: NewPair(1, 111),
				2: NewPair(2, 222),
				3: NewPair(3, 333),
			},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name: "Prepend() many on one-item collection",
			coll: builder.One(),
			args: baseMapIntArgs{
				values: []Pair[int, int]{
					NewPair(2, 222),
					NewPair(3, 333),
				},
			},
			want1: []Pair[int, int]{
				NewPair(2, 222),
				NewPair(3, 333),
				NewPair(1, 111),
			},
			want2: map[int]Pair[int, int]{
				2: NewPair(2, 222),
				3: NewPair(3, 333),
				1: NewPair(1, 111),
			},
			want3: map[int]int{2: 0, 3: 1, 1: 2},
			want4: map[int]int{222: 1, 333: 1, 111: 1},
		},
		{
			name: "Prepend() many on three-item collection",
			coll: builder.Three(),
			args: baseMapIntArgs{
				values: []Pair[int, int]{
					NewPair(4, 444),
					NewPair(5, 555),
				},
			},
			want1: []Pair[int, int]{
				NewPair(4, 444),
				NewPair(5, 555),
				NewPair(1, 111),
				NewPair(2, 222),
				NewPair(3, 333),
			},
			want2: map[int]Pair[int, int]{
				4: NewPair(4, 444),
				5: NewPair(5, 555),
				1: NewPair(1, 111),
				2: NewPair(2, 222),
				3: NewPair(3, 333),
			},
			want3: map[int]int{4: 0, 5: 1, 1: 2, 2: 3, 3: 4},
			want4: map[int]int{444: 1, 555: 1, 111: 1, 222: 1, 333: 1},
		},
		{
			name: "Prepend() many on three-item collection - duplicate key",
			coll: builder.Three(),
			args: baseMapIntArgs{
				values: []Pair[int, int]{
					NewPair(2, 999),
					NewPair(3, 999),
				},
			},
			want1: []Pair[int, int]{
				NewPair(2, 999),
				NewPair(3, 999),
				NewPair(1, 111),
			},
			want2: map[int]Pair[int, int]{
				2: NewPair(2, 999),
				3: NewPair(3, 999),
				1: NewPair(1, 111),
			},
			want3: map[int]int{2: 0, 3: 1, 1: 2},
			want4: map[int]int{999: 2, 111: 1},
		},
		{
			name: "Prepend() many on three-item collection - duplicates",
			coll: builder.Three(),
			args: baseMapIntArgs{
				values: []Pair[int, int]{
					NewPair(1, 111),
					NewPair(2, 222),
					NewPair(3, 333),
					NewPair(4, 111),
					NewPair(5, 222),
					NewPair(6, 333),
				},
			},
			want1: []Pair[int, int]{
				NewPair(1, 111),
				NewPair(2, 222),
				NewPair(3, 333),
				NewPair(4, 111),
				NewPair(5, 222),
				NewPair(6, 333),
			},
			want2: map[int]Pair[int, int]{
				1: NewPair(1, 111),
				2: NewPair(2, 222),
				3: NewPair(3, 333),
				4: NewPair(4, 111),
				5: NewPair(5, 222),
				6: NewPair(6, 333),
			},
			want3: map[int]int{1: 0, 2: 1, 3: 2, 4: 3, 5: 4, 6: 5},
			want4: map[int]int{111: 2, 222: 2, 333: 2},
		},
	}
}

func testMapPrepend(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapPrependCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.values != nil {
				tt.coll.Prepend(tt.args.values...)
			} else {
				tt.coll.Prepend(tt.args.value)
			}
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKP := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Prepend() did not prepend correctly to slice")
			}
			if !reflect.DeepEqual(actualMap, tt.want2) {
				t.Errorf("Prepend() did not prepend correctly to map")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("Prepend() did not prepend correctly to kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("Prepend() did not prepend correctly to values counter")
				}
			}
		})
	}
}

func getMapRemoveCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Remove() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{key: 1},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "Remove() on one-item collection - found",
			coll:  builder.One(),
			args:  baseMapIntArgs{key: 1},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "Remove() on one-item collection - not found",
			coll:  builder.One(),
			args:  baseMapIntArgs{key: 2},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "Remove() on three-item collection - first item",
			coll:  builder.Three(),
			args:  baseMapIntArgs{key: 1},
			want1: []Pair[int, int]{NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{2: 222, 3: 333},
			want3: map[int]int{2: 0, 3: 1},
			want4: map[int]int{222: 1, 333: 1},
		},
		{
			name:  "Remove() on three-item collection - middle item",
			coll:  builder.Three(),
			args:  baseMapIntArgs{key: 2},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(3, 333)},
			want2: map[int]int{1: 111, 3: 333},
			want3: map[int]int{1: 0, 3: 1},
			want4: map[int]int{111: 1, 333: 1},
		},
		{
			name:  "Remove() on three-item collection - last item",
			coll:  builder.Three(),
			args:  baseMapIntArgs{key: 3},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222)},
			want2: map[int]int{1: 111, 2: 222},
			want3: map[int]int{1: 0, 2: 1},
			want4: map[int]int{111: 1, 222: 1},
		},
		{
			name:  "Remove() on three-item collection - not found",
			coll:  builder.Three(),
			args:  baseMapIntArgs{key: 999},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "Remove() on six-item collection",
			coll:  builder.SixWithDuplicates(),
			args:  baseMapIntArgs{key: 1},
			want1: []Pair[int, int]{NewPair(2, 222), NewPair(3, 333), NewPair(4, 111), NewPair(5, 222), NewPair(6, 333)},
			want2: map[int]int{2: 222, 3: 333, 4: 111, 5: 222, 6: 333},
			want3: map[int]int{2: 0, 3: 1, 4: 2, 5: 3, 6: 4},
			want4: map[int]int{222: 2, 333: 2, 111: 1},
		},
	}
}

func testMapRemove(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapRemoveCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Remove(tt.args.key)
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKP := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Remove() did not remove correctly from slice")
			}
			if !reflect.DeepEqual(actualMap, i2iToPairs(tt.want2)) {
				t.Errorf("Remove() did not remove correctly from map")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("Remove() did not remove correctly from kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("Remove() did not remove correctly from values counter")
				}
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
			want4: map[int]int{},
			err:   ErrOutOfBounds,
		},
		{
			name:  "RemoveAt() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{index: 0},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: NewPair(1, 111),
			want4: map[int]int{},
		},
		{
			name:  "RemoveAt() on three-item collection at beginning",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 0},
			want1: []Pair[int, int]{NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{2: 222, 3: 333},
			want3: NewPair(1, 111),
			want4: map[int]int{222: 1, 333: 1},
		},
		{
			name:  "RemoveAt() on three-item collection in the middle",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 1},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(3, 333)},
			want2: map[int]int{1: 111, 3: 333},
			want3: NewPair(2, 222),
			want4: map[int]int{111: 1, 333: 1},
		},
		{
			name:  "RemoveAt() on three-item collection at end",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 2},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222)},
			want2: map[int]int{1: 111, 2: 222},
			want3: NewPair(3, 333),
			want4: map[int]int{111: 1, 222: 1},
		},
		{
			name:  "RemoveAt() on three-item collection out of bounds",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: 4},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: nil,
			want4: map[int]int{111: 1, 222: 1, 333: 1},
			err:   ErrOutOfBounds,
		},
		{
			name:  "RemoveAt() on three-item collection negative index",
			coll:  builder.Three(),
			args:  baseMapIntArgs{index: -1},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: nil,
			want4: map[int]int{111: 1, 222: 1, 333: 1},
			err:   ErrOutOfBounds,
		},
		{
			name:  "RemoveAt() on six-item with duplicates",
			coll:  builder.SixWithDuplicates(),
			args:  baseMapIntArgs{index: 2},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(4, 111), NewPair(5, 222), NewPair(6, 333)},
			want2: map[int]int{1: 111, 2: 222, 4: 111, 5: 222, 6: 333},
			want3: NewPair(3, 333),
			want4: map[int]int{111: 2, 222: 2, 333: 1},
		},
	}
}

func testMapRemoveAt(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapRemoveAtCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			removed, err := tt.coll.RemoveAt(tt.args.index)
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("RemoveAt() did not remove correctly from slice")
			}
			if !reflect.DeepEqual(actualMap, i2iToPairs(tt.want2)) {
				t.Errorf("RemoveAt() did not remove correctly from map")
			}
			if !reflect.DeepEqual(removed, tt.want3) {
				t.Errorf("RemoveAt() did not return removed value correctly")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("RemoveAt() did not remove correctly from values counter")
				}
			}
			if tt.err != nil {
				if !errors.Is(err, tt.err) {
					t.Errorf("RemoveAt() returned wrong error: %v, want error: %v", err, tt.err)
				}
			}
		})
	}
}

func getMapRemoveManyCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "RemoveMany() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{keys: []int{1, 2, 3}},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "RemoveMany() on empty collection - empty keys",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{keys: []int{}},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "RemoveMany() on one-item collection - empty keys",
			coll:  builder.One(),
			args:  baseMapIntArgs{keys: []int{}},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "RemoveMany() on one-item collection - all found",
			coll:  builder.One(),
			args:  baseMapIntArgs{keys: []int{1}},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "RemoveMany() on one-item collection - not found",
			coll:  builder.One(),
			args:  baseMapIntArgs{keys: []int{2}},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "RemoveMany() on one-item collection - some found",
			coll:  builder.One(),
			args:  baseMapIntArgs{keys: []int{1, 2}},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "RemoveMany() on three-item collection - all found",
			coll:  builder.Three(),
			args:  baseMapIntArgs{keys: []int{1, 2, 3}},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "RemoveMany() on three-item collection - not found",
			coll:  builder.Three(),
			args:  baseMapIntArgs{keys: []int{4, 5, 6}},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "RemoveMany() on three-item collection - some found",
			coll:  builder.Three(),
			args:  baseMapIntArgs{keys: []int{1, 2, 4}},
			want1: []Pair[int, int]{NewPair(3, 333)},
			want2: map[int]int{3: 333},
			want3: map[int]int{3: 0},
			want4: map[int]int{333: 1},
		},
		{
			name:  "RemoveMany() on three-item collection - some found, some not",
			coll:  builder.Three(),
			args:  baseMapIntArgs{keys: []int{1, 2, 4, 5}},
			want1: []Pair[int, int]{NewPair(3, 333)},
			want2: map[int]int{3: 333},
			want3: map[int]int{3: 0},
			want4: map[int]int{333: 1},
		},
		{
			name:  "RemoveMany() on three-item collection - some found, some not, some duplicate",
			coll:  builder.Three(),
			args:  baseMapIntArgs{keys: []int{1, 2, 4, 5, 1}},
			want1: []Pair[int, int]{NewPair(3, 333)},
			want2: map[int]int{3: 333},
			want3: map[int]int{3: 0},
			want4: map[int]int{333: 1},
		},
		{
			name:  "RemoveMany() on three-item collection - all found, some duplicate",
			coll:  builder.Three(),
			args:  baseMapIntArgs{keys: []int{1, 2, 3, 1}},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "RemoveMany() on three-item collection - all found, some duplicate, some not",
			coll:  builder.Three(),
			args:  baseMapIntArgs{keys: []int{1, 2, 3, 1, 4}},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "RemoveMany() on six-item collection with duplicates - some found at the beginning",
			coll:  builder.SixWithDuplicates(),
			args:  baseMapIntArgs{keys: []int{1, 2, 3}},
			want1: []Pair[int, int]{NewPair(4, 111), NewPair(5, 222), NewPair(6, 333)},
			want2: map[int]int{4: 111, 5: 222, 6: 333},
			want3: map[int]int{4: 0, 5: 1, 6: 2},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "RemoveMany() on six-item collection with duplicates - some found at the end",
			coll:  builder.SixWithDuplicates(),
			args:  baseMapIntArgs{keys: []int{4, 5, 6}},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
	}
}

func testMapRemoveMany(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapRemoveManyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.RemoveMany(tt.args.keys)
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKP := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("RemoveMany() did not remove correctly from slice")
			}
			if !reflect.DeepEqual(actualMap, i2iToPairs(tt.want2)) {
				t.Errorf("RemoveMany() did not remove correctly from map")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("RemoveMany() did not remove correctly from kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("RemoveMany() did not remove correctly from values counter")
				}
			}
		})
	}
}

func getMapRemoveMatchingCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "RemoveMatching() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{predicate: func(p Pair[int, int]) bool { return true }},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
			want5: 0,
		},
		{
			name:  "RemoveMatching() on one-item collection, found",
			coll:  builder.One(),
			args:  baseMapIntArgs{predicate: func(p Pair[int, int]) bool { return p.Val() == 111 }},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
			want5: 1,
		},
		{
			name:  "RemoveMatching() on one-item collection, not found",
			coll:  builder.One(),
			args:  baseMapIntArgs{predicate: func(p Pair[int, int]) bool { return p.Val() == 222 }},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
			want5: 0,
		},
		{
			name:  "RemoveMatching() on three-item collection, found all",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(p Pair[int, int]) bool { return true }},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
			want5: 3,
		},
		{
			name:  "RemoveMatching() on three-item collection, found none",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(p Pair[int, int]) bool { return false }},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
			want5: 0,
		},
		{
			name:  "RemoveMatching() on three-item collection, found first",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(p Pair[int, int]) bool { return p.Val() == 111 }},
			want1: []Pair[int, int]{NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{2: 222, 3: 333},
			want3: map[int]int{2: 0, 3: 1},
			want4: map[int]int{222: 1, 333: 1},
			want5: 1,
		},
		{
			name:  "RemoveMatching() on three-item collection, found middle",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(p Pair[int, int]) bool { return p.Val() == 222 }},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(3, 333)},
			want2: map[int]int{1: 111, 3: 333},
			want3: map[int]int{1: 0, 3: 1},
			want4: map[int]int{111: 1, 333: 1},
			want5: 1,
		},
		{
			name:  "RemoveMatching() on three-item collection, found last",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(p Pair[int, int]) bool { return p.Val() == 333 }},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222)},
			want2: map[int]int{1: 111, 2: 222},
			want3: map[int]int{1: 0, 2: 1},
			want4: map[int]int{111: 1, 222: 1},
			want5: 1,
		},
		{
			name:  "RemoveMatching() on three-item collection, found even",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(p Pair[int, int]) bool { return p.Val()%2 != 0 }},
			want1: []Pair[int, int]{NewPair(2, 222)},
			want2: map[int]int{2: 222},
			want3: map[int]int{2: 0},
			want4: map[int]int{222: 1},
			want5: 2,
		},
	}
}

func testMapRemoveMatching(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapRemoveMatchingCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			count := tt.coll.RemoveMatching(tt.args.predicate)
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKP := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("RemoveMatching() did not remove correctly from slice")
			}
			if !reflect.DeepEqual(actualMap, i2iToPairs(tt.want2)) {
				t.Errorf("RemoveMatching() did not remove correctly from map")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("RemoveMatching() did not remove correctly from kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("RemoveMatching() did not remove correctly from values counter")
				}
			}
			if count != tt.want5 {
				t.Errorf("RemoveMatching() returned wrong count: %v, but wanted = %v", count, tt.want5)
			}
		})
	}
}

func getMapReverseCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:     "Reverse() on empty collection",
			coll:     builder.Empty(),
			want1:    []Pair[int, int](nil),
			want2:    map[int]int{},
			want3:    map[int]int{},
			want4:    map[int]int{},
			metaInt1: 1,
		},
		{
			name:     "Reverse() on one-item collection",
			coll:     builder.One(),
			want1:    []Pair[int, int]{NewPair(1, 111)},
			want2:    map[int]int{1: 111},
			want3:    map[int]int{1: 0},
			want4:    map[int]int{111: 1},
			metaInt1: 1,
		},
		{
			name:     "Reverse() on three-item collection",
			coll:     builder.Three(),
			want1:    []Pair[int, int]{NewPair(3, 333), NewPair(2, 222), NewPair(1, 111)},
			want2:    map[int]int{3: 333, 2: 222, 1: 111},
			want3:    map[int]int{1: 2, 2: 1, 3: 0},
			want4:    map[int]int{111: 1, 222: 1, 333: 1},
			metaInt1: 1,
		},
		{
			name:     "Reverse() twice on three-item collection",
			coll:     builder.Three(),
			want1:    []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2:    map[int]int{1: 111, 2: 222, 3: 333},
			want3:    map[int]int{1: 0, 2: 1, 3: 2},
			want4:    map[int]int{111: 1, 222: 1, 333: 1},
			metaInt1: 2,
		},
		{
			name:     "Reverse() on six-item collection",
			coll:     builder.SixWithDuplicates(),
			want1:    []Pair[int, int]{NewPair(6, 333), NewPair(5, 222), NewPair(4, 111), NewPair(3, 333), NewPair(2, 222), NewPair(1, 111)},
			want2:    map[int]int{6: 333, 5: 222, 4: 111, 3: 333, 2: 222, 1: 111},
			want3:    map[int]int{1: 5, 2: 4, 3: 3, 4: 2, 5: 1, 6: 0},
			want4:    map[int]int{111: 2, 222: 2, 333: 2},
			metaInt1: 1,
		},
	}
}

func testMapReverse(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapReverseCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.metaInt1; i++ {
				tt.coll.Reverse()
			}
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKP := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Reverse() did not reverse correctly from slice")
			}
			if !reflect.DeepEqual(actualMap, i2iToPairs(tt.want2)) {
				t.Errorf("Reverse() did not reverse correctly from map")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("Reverse() did not reverse correctly from kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("Reverse() did not reverse correctly from values counter")
				}
			}
		})
	}
}

func getMapSetCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Set() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{value: NewPair(1, 111)},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "Set() on one-item collection - replace",
			coll:  builder.One(),
			args:  baseMapIntArgs{value: NewPair(1, 999)},
			want1: []Pair[int, int]{NewPair(1, 999)},
			want2: map[int]int{1: 999},
			want3: map[int]int{1: 0},
			want4: map[int]int{999: 1},
		},
		{
			name:  "Set() on one item collection - add",
			coll:  builder.One(),
			args:  baseMapIntArgs{value: NewPair(2, 222)},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222)},
			want2: map[int]int{1: 111, 2: 222},
			want3: map[int]int{1: 0, 2: 1},
			want4: map[int]int{111: 1, 222: 1},
		},
		{
			name:  "Set() on three-item collection - replace",
			coll:  builder.Three(),
			args:  baseMapIntArgs{value: NewPair(2, 999)},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 999), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 999, 3: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 999: 1, 333: 1},
		},
		{
			name:  "Set() on three-item collection - add",
			coll:  builder.Three(),
			args:  baseMapIntArgs{value: NewPair(4, 444)},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333), NewPair(4, 444)},
			want2: map[int]int{1: 111, 2: 222, 3: 333, 4: 444},
			want3: map[int]int{1: 0, 2: 1, 3: 2, 4: 3},
			want4: map[int]int{111: 1, 222: 1, 333: 1, 444: 1},
		},
	}
}

func testMapSet(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapSetCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Set(tt.args.value.Key(), tt.args.value.Val())
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKP := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Set() did not set correctly to slice")
			}
			if !reflect.DeepEqual(actualMap, i2iToPairs(tt.want2)) {
				t.Errorf("Set() did not set correctly to map")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("Set() did not set correctly to kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("Set() did not set correctly to values counter")
				}
			}
		})
	}
}

func getMapSetManyCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "SetAll() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{values: []Pair[int, int]{NewPair(1, 111)}},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "SetAll() on one-item collection - replace",
			coll:  builder.One(),
			args:  baseMapIntArgs{values: []Pair[int, int]{NewPair(1, 999)}},
			want1: []Pair[int, int]{NewPair(1, 999)},
			want2: map[int]int{1: 999},
			want3: map[int]int{1: 0},
			want4: map[int]int{999: 1},
		},
		{
			name:  "SetAll() on one item collection - add",
			coll:  builder.One(),
			args:  baseMapIntArgs{values: []Pair[int, int]{NewPair(2, 222)}},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222)},
			want2: map[int]int{1: 111, 2: 222},
			want3: map[int]int{1: 0, 2: 1},
			want4: map[int]int{111: 1, 222: 1},
		},
		{
			name:  "SetAll() on three-item collection - replace",
			coll:  builder.Three(),
			args:  baseMapIntArgs{values: []Pair[int, int]{NewPair(2, 999)}},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 999), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 999, 3: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 999: 1, 333: 1},
		},
		{
			name:  "SetAll() on three-item collection - add",
			coll:  builder.Three(),
			args:  baseMapIntArgs{values: []Pair[int, int]{NewPair(4, 444)}},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333), NewPair(4, 444)},
			want2: map[int]int{1: 111, 2: 222, 3: 333, 4: 444},
			want3: map[int]int{1: 0, 2: 1, 3: 2, 4: 3},
			want4: map[int]int{111: 1, 222: 1, 333: 1, 444: 1},
		},
		{
			name:  "SetAll() on three-item collection - replace and add",
			coll:  builder.Three(),
			args:  baseMapIntArgs{values: []Pair[int, int]{NewPair(2, 999), NewPair(4, 444)}},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 999), NewPair(3, 333), NewPair(4, 444)},
			want2: map[int]int{1: 111, 2: 999, 3: 333, 4: 444},
			want3: map[int]int{1: 0, 2: 1, 3: 2, 4: 3},
			want4: map[int]int{111: 1, 999: 1, 333: 1, 444: 1},
		},
		{
			name: "SetAll() on six-item collection - replace and add",
			coll: builder.SixWithDuplicates(),
			args: baseMapIntArgs{values: []Pair[int, int]{NewPair(2, 999), NewPair(7, 444), NewPair(8, 333)}},
			want1: []Pair[int, int]{
				NewPair(1, 111),
				NewPair(2, 999),
				NewPair(3, 333),
				NewPair(4, 111),
				NewPair(5, 222),
				NewPair(6, 333),
				NewPair(7, 444),
				NewPair(8, 333),
			},
			want2: map[int]int{1: 111, 2: 999, 3: 333, 4: 111, 5: 222, 6: 333, 7: 444, 8: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2, 4: 3, 5: 4, 6: 5, 7: 6, 8: 7},
			want4: map[int]int{111: 2, 999: 1, 333: 3, 444: 1, 222: 1},
		},
	}
}

func testMapSetMany(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapSetManyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.SetMany(tt.args.values)
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKP := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("SetAll() did not set correctly to slice")
			}
			if !reflect.DeepEqual(actualMap, i2iToPairs(tt.want2)) {
				t.Errorf("SetAll() did not set correctly to map")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("SetAll() did not set correctly to kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("SetAll() did not set correctly to values counter")
				}
			}
		})
	}
}

func getMapSortCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Sort() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{comparer: func(a, b Pair[int, int]) int { return a.Key() - b.Key() }},
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "Sort() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{comparer: func(a, b Pair[int, int]) int { return a.Key() - b.Key() }},
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "Sort() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{comparer: func(a, b Pair[int, int]) int { return a.Key() - b.Key() }},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "Sort() on three-item collection - reverse",
			coll:  builder.Three(),
			args:  baseMapIntArgs{comparer: func(a, b Pair[int, int]) int { return b.Key() - a.Key() }},
			want1: []Pair[int, int]{NewPair(3, 333), NewPair(2, 222), NewPair(1, 111)},
			want2: map[int]int{3: 333, 2: 222, 1: 111},
			want3: map[int]int{1: 2, 2: 1, 3: 0},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:     "Sort() on three-item collection - reverse twice",
			coll:     builder.Three(),
			args:     baseMapIntArgs{comparer: func(a, b Pair[int, int]) int { return b.Key() - a.Key() }},
			metaInt1: 2,
			want1:    []Pair[int, int]{NewPair(3, 333), NewPair(2, 222), NewPair(1, 111)},
			want2:    map[int]int{3: 333, 2: 222, 1: 111},
			want3:    map[int]int{1: 2, 2: 1, 3: 0},
			want4:    map[int]int{111: 1, 222: 1, 333: 1},
		},
		{
			name:  "Sort() on six-item collection",
			coll:  builder.SixWithDuplicates(),
			args:  baseMapIntArgs{comparer: func(a, b Pair[int, int]) int { return a.Key() - b.Key() }},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333), NewPair(4, 111), NewPair(5, 222), NewPair(6, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333, 4: 111, 5: 222, 6: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2, 4: 3, 5: 4, 6: 5},
			want4: map[int]int{111: 2, 222: 2, 333: 2},
		},
		{
			name:  "Sort() on six-item collection - reverse",
			coll:  builder.SixWithDuplicates(),
			args:  baseMapIntArgs{comparer: func(a, b Pair[int, int]) int { return b.Key() - a.Key() }},
			want1: []Pair[int, int]{NewPair(6, 333), NewPair(5, 222), NewPair(4, 111), NewPair(3, 333), NewPair(2, 222), NewPair(1, 111)},
			want2: map[int]int{6: 333, 5: 222, 4: 111, 3: 333, 2: 222, 1: 111},
			want3: map[int]int{1: 5, 2: 4, 3: 3, 4: 2, 5: 1, 6: 0},
			want4: map[int]int{111: 2, 222: 2, 333: 2},
		},
		{
			name:     "Sort() on six-item collection - reverse twice",
			coll:     builder.SixWithDuplicates(),
			args:     baseMapIntArgs{comparer: func(a, b Pair[int, int]) int { return b.Key() - a.Key() }},
			metaInt1: 2,
			want1:    []Pair[int, int]{NewPair(6, 333), NewPair(5, 222), NewPair(4, 111), NewPair(3, 333), NewPair(2, 222), NewPair(1, 111)},
			want2:    map[int]int{6: 333, 5: 222, 4: 111, 3: 333, 2: 222, 1: 111},
			want3:    map[int]int{1: 5, 2: 4, 3: 3, 4: 2, 5: 1, 6: 0},
			want4:    map[int]int{111: 2, 222: 2, 333: 2},
		},
	}
}

func testMapSort(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapSortCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			times := 1
			if tt.metaInt1 > 0 {
				times = tt.metaInt1
			}
			for i := 0; i < times; i++ {
				tt.coll.Sort(tt.args.comparer)
			}
			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)
			actualKP := builder.extractUnderlyingKp(tt.coll)
			actualVC := builder.extractUnderlyingValsCount(tt.coll)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Sort() did not sort correctly to slice")
			}
			if !reflect.DeepEqual(actualMap, i2iToPairs(tt.want2)) {
				t.Errorf("Sort() did not sort correctly to map")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("Sort() did not sort correctly to kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("Sort() did not sort correctly to values counter")
				}
			}
		})
	}
}

func getMapValuesCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Values() on empty collection",
			coll:  builder.Empty(),
			want1: []Pair[int, int](nil),
		},
		{
			name:  "Values() on one-item collection",
			coll:  builder.One(),
			want1: []Pair[int, int]{NewPair(1, 111)},
		},
		{
			name:  "Values() on three-item collection",
			coll:  builder.Three(),
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
		},
	}
}

func testMapValues(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapValuesCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Collect(tt.coll.Values())
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Values() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapValuesWithBreakCases(builder baseMapCollIntBuilder) []*baseMapTestCase {
	return []*baseMapTestCase{
		{
			name: "Values() on three-item collection, break immediately",
			coll: builder.Three(),
			args: baseMapIntArgs{
				predicate: func(p Pair[int, int]) bool {
					return false
				},
			},
			want1: []Pair[int, int](nil),
		},
		{
			name: "Values() on three-item collection, break at middle",
			coll: builder.Three(),
			args: baseMapIntArgs{
				predicate: func(p Pair[int, int]) bool {
					return p.Key() < 2
				},
			},
			want1: []Pair[int, int]{NewPair(1, 111)},
		},
		{
			name: "Values() on three-item collection, break after middle",
			coll: builder.Three(),
			args: baseMapIntArgs{
				predicate: func(p Pair[int, int]) bool {
					return p.Key() <= 2
				},
			},
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222)},
		},
	}
}

func testMapValuesBreak(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapValuesWithBreakCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := []Pair[int, int](nil)
			for v := range tt.coll.Values() {
				if !tt.args.predicate(v) {
					break
				}
				got = append(got, v)
			}
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Values() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapValuesRevCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "ValuesRev() on empty collection",
			coll:  builder.Empty(),
			want1: []Pair[int, int](nil),
		},
		{
			name:  "ValuesRev() on one-item collection",
			coll:  builder.One(),
			want1: []Pair[int, int]{NewPair(1, 111)},
		},
		{
			name:  "ValuesRev() on three-item collection",
			coll:  builder.Three(),
			want1: []Pair[int, int]{NewPair(3, 333), NewPair(2, 222), NewPair(1, 111)},
		},
	}
}

func testMapValuesRev(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapValuesRevCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Collect(tt.coll.ValuesRev())
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("ValuesRev() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapValuesRevBreakCases(builder baseMapCollIntBuilder) []*baseMapTestCase {
	return []*baseMapTestCase{
		{
			name: "ValuesRev() on three-item collection, break immediately",
			coll: builder.Three(),
			args: baseMapIntArgs{
				predicate: func(p Pair[int, int]) bool {
					return false
				},
			},
			want1: []Pair[int, int](nil),
		},
		{
			name: "ValuesRev() on three-item collection, break at middle",
			coll: builder.Three(),
			args: baseMapIntArgs{
				predicate: func(p Pair[int, int]) bool {
					return p.Key() > 2
				},
			},
			want1: []Pair[int, int]{NewPair(3, 333)},
		},
		{
			name: "ValuesRev() on three-item collection, break after middle",
			coll: builder.Three(),
			args: baseMapIntArgs{
				predicate: func(p Pair[int, int]) bool {
					return p.Key() >= 2
				},
			},
			want1: []Pair[int, int]{NewPair(3, 333), NewPair(2, 222)},
		},
	}
}

func testMapValuesRevBreak(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapValuesRevBreakCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := []Pair[int, int](nil)
			for v := range tt.coll.ValuesRev() {
				if !tt.args.predicate(v) {
					break
				}
				got = append(got, v)
			}
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("ValuesRev() = %v, want1 = %v", got, tt.want1)
			}
		})
	}
}

func getMapCopyCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Copy() on empty collection",
			coll:  builder.Empty(),
			want1: []Pair[int, int](nil),
			want2: map[int]int{},
			want3: map[int]int{},
			want4: map[int]int{},
		},
		{
			name:  "Copy() on one-item collection",
			coll:  builder.One(),
			want1: []Pair[int, int]{NewPair(1, 111)},
			want2: map[int]int{1: 111},
			want3: map[int]int{1: 0},
			want4: map[int]int{111: 1},
		},
		{
			name:  "Copy() on three-item collection",
			coll:  builder.Three(),
			want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
			want2: map[int]int{1: 111, 2: 222, 3: 333},
			want3: map[int]int{1: 0, 2: 1, 3: 2},
			want4: map[int]int{111: 1, 222: 1, 333: 1},
		},
	}
}

func testMapCopy(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapCopyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.copy().(mapInternal[int, int])
			actualSlice := builder.extractUnderlyingSlice(got)
			actualMap := builder.extractUnderlyingMap(got)
			actualKP := builder.extractUnderlyingKp(got)
			actualVC := builder.extractUnderlyingValsCount(got)
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Copy() did not copy correctly to slice")
			}
			if !reflect.DeepEqual(actualMap, i2iToPairs(tt.want2)) {
				t.Errorf("copy() did not copy correctly to map")
			}
			if !reflect.DeepEqual(actualKP, tt.want3) {
				t.Errorf("copy() did not copy correctly to kp")
			}
			if actualVC != nil {
				if !reflect.DeepEqual(actualVC, tt.want4) {
					t.Errorf("copy() did not copy correctly to values counter")
				}
			}
		})
	}
}
