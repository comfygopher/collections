package coll

import (
	"reflect"
	"testing"
)

func getAppendRefCases(builder baseMapCollIntBuilder) []*baseMapTestCase {

	appendOneOnEmptyCaseValue := NewPair(1, 111)
	appendOneOnEmptyCase := &baseMapTestCase{
		name:  "Append() on empty collection",
		coll:  builder.Empty(),
		args:  baseMapIntArgs{value: appendOneOnEmptyCaseValue},
		want1: []Pair[int, int]{NewPair(1, 999)},
		want2: map[int]int{1: 999},
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
		want2: map[int]int{1: 111, 10: 999},
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
	cases := getAppendRefCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.values != nil {
				tt.coll.Append(tt.args.values...)
			} else {
				tt.coll.Append(tt.args.value)
			}
			tt.modify()
			actualSlice := tt.coll.ToSlice()
			actualMap := tt.coll.ToMap()
			if !reflect.DeepEqual(actualSlice, tt.want1) {
				t.Errorf("Append() did not append correctly to slice")
			}
			if !reflect.DeepEqual(actualMap, tt.want2) {
				t.Errorf("Append() did not append correctly to map")
			}
		})
	}
}
