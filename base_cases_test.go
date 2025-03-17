package coll

import (
	"reflect"
	"slices"
	"testing"
)

type baseIntArgs = testArgs[baseInternal[int], int]
type baseTestCase = testCase[baseInternal[int], int]
type baseCollIntBuilder = testCollectionBuilder[baseInternal[int]]

func getIsEmptyCases(builder baseCollIntBuilder) []baseTestCase {
	return []baseTestCase{
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

func testIsEmpty(t *testing.T, builder baseCollIntBuilder) {
	cases := getIsEmptyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.IsEmpty()
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("IsEmpty() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getLenCases(builder baseCollIntBuilder) []baseTestCase {
	return []baseTestCase{
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

func testLen(t *testing.T, builder baseCollIntBuilder) {
	cases := getLenCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.Len()
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Len() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getValuesCases(builder baseCollIntBuilder) []*baseTestCase {
	return []*baseTestCase{
		{
			name:  "Values() on empty collection",
			coll:  builder.Empty(),
			want1: []int(nil),
		},
		{
			name:  "Values() on one-item collection",
			coll:  builder.One(),
			want1: []int{111},
		},
		{
			name:  "Values() on three-item collection",
			coll:  builder.Three(),
			want1: []int{111, 222, 333},
		},
	}
}

func testValues(t *testing.T, builder baseCollIntBuilder) {
	cases := getValuesCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Collect(tt.coll.Values())
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Values() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getValuesBreakCases(builder baseCollIntBuilder) []*baseTestCase {
	return []*baseTestCase{
		{
			name: "Values() on three-item collection, break",
			coll: builder.Three(),
			args: baseIntArgs{
				predicate: func(v int) bool {
					return v < 222
				},
			},
			want1: []int{111},
		},
		{
			name: "Values() on three-item collection, break",
			coll: builder.Three(),
			args: baseIntArgs{
				predicate: func(v int) bool {
					return v <= 222
				},
			},
			want1: []int{111, 222},
		},
	}
}

func testValuesBreak(t *testing.T, builder baseCollIntBuilder) {
	cases := getValuesBreakCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			for v := range tt.coll.Values() {
				if !tt.args.predicate(v) {
					break
				}
				got = append(got, v)
			}
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Values() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getCopyCases(builder baseCollIntBuilder) []*baseTestCase {
	return []*baseTestCase{
		{
			name: "Copy() on empty collection",
			coll: builder.Empty(),
		},
		{
			name: "Copy() on one-item collection",
			coll: builder.One(),
		},
		{
			name: "Copy() on three-item collection",
			coll: builder.Three(),
		},
	}
}

func testCopy(t *testing.T, builder baseCollIntBuilder) {
	cases := getCopyCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.copy()
			if !reflect.DeepEqual(got, tt.coll) {
				t.Errorf("Copy() = %v, want1 %v", got, tt.coll)
			}
		})
	}
}
