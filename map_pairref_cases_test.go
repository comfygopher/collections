package coll

import (
	"reflect"
	"testing"
)

func getMapAppendRefCases(builder baseMapCollIntBuilder) []*baseMapTestCase {

	appendOneOnEmptyCaseValue := NewPair(1, 111)
	appendOneOnEmptyCase := &baseMapTestCase{
		name:  "Append() on empty collection",
		coll:  builder.Empty(),
		args:  baseMapIntArgs{value: appendOneOnEmptyCaseValue},
		want1: []Pair[int, int]{NewPair(1, 999)},
		want2: map[int]Pair[int, int]{1: NewPair(1, 999)},
		modify: func() {
			appendOneOnEmptyCaseValue.SetVal(999)
		},
	}

	appendOneOnOneItemCaseValue := NewPair(10, 111)
	appendOneOnOneItemCase := &baseMapTestCase{
		name:  "Append() on one-item collection",
		coll:  builder.One(),
		args:  baseMapIntArgs{value: appendOneOnOneItemCaseValue},
		want1: []Pair[int, int]{NewPair(1, 111), NewPair(10, 999)},
		want2: map[int]Pair[int, int]{1: NewPair(1, 111), 10: NewPair(10, 999)},
		modify: func() {
			appendOneOnOneItemCaseValue.SetVal(999)
		},
	}

	return []*baseMapTestCase{
		appendOneOnEmptyCase,
		appendOneOnOneItemCase,
	}
}

func testMapAppendRef(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapAppendRefCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.values != nil {
				tt.coll.Append(tt.args.values...)
			} else {
				tt.coll.Append(tt.args.value)
			}
			tt.modify()

			actualSlice := builder.extractUnderlyingSlice(tt.coll)
			actualMap := builder.extractUnderlyingMap(tt.coll)

			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Append() did not append correctly to slice")
			}
			if !reflect.DeepEqual(actualMap, tt.want2) {
				t.Errorf("Append() did not append correctly to map")
			}
		})
	}
}

func getMapValuesRefCases(builder baseMapCollIntBuilder) []*baseMapTestCase {
	threeItemModifyValuesCase := &baseMapTestCase{
		name: "Values() on three-item collection, modify values",
		args: baseMapIntArgs{
			values: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
		},
		want1: []Pair[int, int]{NewPair(1, 1111), NewPair(2, 1222), NewPair(3, 1333)},
	}
	threeItemModifyValuesCase.collBuilder = func() mapInternal[int, int] {
		return builder.FromValues(threeItemModifyValuesCase.args.valuesAsAnySlice())
	}
	threeItemModifyValuesCase.args.visit = func(_ int, p Pair[int, int]) {
		p.SetVal(p.Val() + 1000)
	}

	return []*baseMapTestCase{
		threeItemModifyValuesCase,
	}
}

func testMapValuesRef(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapValuesRefCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			tt.coll = tt.collBuilder()

			actualSliceBeforeModification := builder.extractUnderlyingSlice(tt.coll)
			if !reflect.DeepEqual(actualSliceBeforeModification, tt.args.values) {
				t.Errorf("Values() did not modify values correctly in slice")
			}

			for v := range tt.coll.Values() {
				tt.args.visit(-1, v)
			}

			actualSliceAfterModification := builder.extractUnderlyingSlice(tt.coll)
			if !reflect.DeepEqual(actualSliceAfterModification, tt.want1) {
				t.Errorf("Values() did not modify values correctly in slice")
			}
		})
	}
}

func getMapCopyRefCases(builder baseMapCollIntBuilder) []*baseMapTestCase {
	threeItemModifyValuesCase := &baseMapTestCase{
		name: "Values() on three-item collection, modify values",
		args: baseMapIntArgs{
			values: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
		},
		want1: []Pair[int, int]{NewPair(1, 111), NewPair(2, 222), NewPair(3, 333)},
	}
	threeItemModifyValuesCase.collBuilder = func() mapInternal[int, int] {
		return builder.FromValues(threeItemModifyValuesCase.args.valuesAsAnySlice())
	}
	threeItemModifyValuesCase.args.visit = func(_ int, p Pair[int, int]) {
		p.SetVal(p.Val() + 1000)
	}

	return []*baseMapTestCase{
		threeItemModifyValuesCase,
	}
}

func testMapCopyDontPreserveRef(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapCopyRefCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			tt.coll = tt.collBuilder()

			actualSliceBeforeModification := builder.extractUnderlyingSlice(tt.coll)

			if !reflect.DeepEqual(actualSliceBeforeModification, tt.args.values) {
				t.Errorf("Values() did not modify values correctly in slice")
			}

			copiedMap := tt.coll.copy().(mapInternal[int, int])
			for v := range copiedMap.Values() {
				tt.args.visit(-1, v)
			}

			actualSliceAfterModification := builder.extractUnderlyingSlice(tt.coll)
			if !reflect.DeepEqual(actualSliceAfterModification, tt.want1) {
				t.Errorf("Values() wanted not modified values = %v, but got = %v", tt.want1, actualSliceAfterModification)
			}
		})
	}
}
