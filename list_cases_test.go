package coll

import (
	"errors"
	"reflect"
	"testing"
)

type listIntArgs = testArgs[listInternal[int], int]
type listTestCase = testCase[listInternal[int], int]
type listCollIntBuilder = testCollectionBuilder[listInternal[int]]

func getInsertAtCases(builder listCollIntBuilder) []listTestCase {
	return []listTestCase{
		{
			name:  "InsertAt() on empty collection",
			coll:  builder.Empty(),
			args:  listIntArgs{index: 0, value: 999},
			want1: []int{999},
			want2: map[int]int{999: 1},
		},
		{
			name:  "InsertAt() on one-item collection",
			coll:  builder.One(),
			args:  listIntArgs{index: 0, value: 1},
			want1: []int{1, 111},
			want2: map[int]int{1: 1, 111: 1},
		},
		{
			name:  "InsertAt() on one-item collection at end",
			coll:  builder.One(),
			args:  listIntArgs{index: 1, value: 1},
			want1: []int{111, 1},
			want2: map[int]int{111: 1, 1: 1},
		},
		{
			name:  "InsertAt() on one-item collection at beginning",
			coll:  builder.One(),
			args:  listIntArgs{index: 0, value: 1},
			want1: []int{1, 111},
			want2: map[int]int{1: 1, 111: 1},
		},
		{
			name:  "InsertAt() on three-item collection at beginning",
			coll:  builder.Three(),
			args:  listIntArgs{index: 0, value: 4},
			want1: []int{4, 111, 222, 333},
			want2: map[int]int{4: 1, 111: 1, 222: 1, 333: 1},
		},
		{
			name:  "InsertAt() on three-item collection at end",
			coll:  builder.Three(),
			args:  listIntArgs{index: 3, value: 4},
			want1: []int{111, 222, 333, 4},
			want2: map[int]int{111: 1, 222: 1, 333: 1, 4: 1},
		},
		{
			name:  "InsertAt() on three-item collection",
			coll:  builder.Three(),
			args:  listIntArgs{index: 1, value: 4},
			want1: []int{111, 4, 222, 333},
			want2: map[int]int{111: 1, 4: 1, 222: 1, 333: 1},
		},
		{
			name:  "InsertAt() on six-item collection",
			coll:  builder.SixWithDuplicates(),
			args:  listIntArgs{index: 6, value: 111},
			want1: []int{111, 222, 333, 111, 222, 333, 111},
			want2: map[int]int{111: 3, 222: 2, 333: 2},
		},
		{
			name:  "InsertAt() on six-item collection - duplicate",
			coll:  builder.SixWithDuplicates(),
			args:  listIntArgs{index: 1, value: 111},
			want1: []int{111, 111, 222, 333, 111, 222, 333},
			want2: map[int]int{111: 3, 222: 2, 333: 2},
		},
		{
			name:    "InsertAt() on three-item collection out of bounds",
			coll:    builder.Three(),
			args:    listIntArgs{index: 5, value: 4},
			wantErr: true,
			err:     ErrOutOfBounds,
		},
		{
			name:    "InsertAt() on three-item collection negative index",
			coll:    builder.Three(),
			args:    listIntArgs{index: -1, value: 4},
			wantErr: true,
			err:     ErrOutOfBounds,
		},
		{
			name:    "InsertAt() on empty collection out of bounds",
			coll:    builder.Empty(),
			args:    listIntArgs{index: 1, value: 4},
			wantErr: true,
			err:     ErrOutOfBounds,
		},
		{
			name:    "InsertAt() on empty collection negative index",
			coll:    builder.Empty(),
			args:    listIntArgs{index: -1, value: 4},
			wantErr: true,
			err:     ErrOutOfBounds,
		},
	}
}

func testInsertAt(t *testing.T, builder listCollIntBuilder) {
	cases := getInsertAtCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.coll.InsertAt(tt.args.index, tt.args.value)
			if tt.wantErr {
				if err == nil {
					t.Errorf("InsertAt() did not return error")
				}
				if !errors.Is(err, tt.err) {
					t.Errorf("InsertAt() returned wrong error: %v, want error: %v", err, tt.err)
				}
			} else {
				if err != nil {
					t.Errorf("InsertAt() returned error: %v", err)
				}

				actualSlice := builder.extractUnderlyingSlice(tt.coll)
				actualVC := builder.extractUnderlyingValsCount(tt.coll)

				if !reflect.DeepEqual(actualSlice, tt.want1) {
					t.Errorf("InsertAt() resulted in: %v, but wanted %v", actualSlice, tt.want1)
				}
				if actualVC != nil && !reflect.DeepEqual(actualVC, tt.want2) {
					t.Errorf("InsertAt() did not append correctly from values counter")
				}
			}
		})
	}
}
