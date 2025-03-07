package coll

import (
	"errors"
	"testing"
)

func getContainsValueCases[C any](builder testCollectionBuilder[C, int]) []testCase[C, int] {
	return []testCase[C, int]{
		{
			name:  "ContainsValue() on empty collection",
			coll:  builder.Empty(),
			args:  testArgs[C, int]{value: 1},
			want1: false,
		},
		{
			name:  "ContainsValue() on one-item collection",
			coll:  builder.One(),
			args:  testArgs[C, int]{value: 111},
			want1: true,
		},
		{
			name:  "ContainsValue() on three-item collection - first val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 111},
			want1: true,
		},
		{
			name:  "ContainsValue() on three-item collection - second val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 222},
			want1: true,
		},
		{
			name:  "ContainsValue() on three-item collection - third val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 333},
			want1: true,
		},
		{
			name:  "ContainsValue() on three-item collection, not found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 999},
			want1: false,
		},
	}
}

func testContainsValue[C cmpInternal[int]](t *testing.T, builder testCollectionBuilder[C, int]) {
	cases := getContainsValueCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.coll.ContainsValue(tt.args.value); got != tt.want1 {
				t.Errorf("ContainsValue() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getCountValuesCases[C any](builder testCollectionBuilder[C, int]) []testCase[C, int] {
	return []testCase[C, int]{
		{
			name:  "CountValues() on empty collection",
			coll:  builder.Empty(),
			args:  testArgs[C, int]{value: 1},
			want1: 0,
		},
		{
			name:  "CountValues() on one-item collection",
			coll:  builder.One(),
			args:  testArgs[C, int]{value: 111},
			want1: 1,
		},
		{
			name:  "CountValues() on three-item collection - first val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 111},
			want1: 1,
		},
		{
			name:  "CountValues() on three-item collection - second val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 222},
			want1: 1,
		},
		{
			name:  "CountValues() on three-item collection - third val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 333},
			want1: 1,
		},
		{
			name:  "CountValues() on three-item collection, not found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 999},
			want1: 0,
		},
		{
			name:  "CountValues() on six-item collection, 2 `111` found ",
			coll:  builder.SixWithDuplicates(),
			args:  testArgs[C, int]{value: 111},
			want1: 2,
		},
		{
			name:  "CountValues() on six-item collection, 2 `222` found ",
			coll:  builder.SixWithDuplicates(),
			args:  testArgs[C, int]{value: 222},
			want1: 2,
		},
		{
			name:  "CountValues() on six-item collection, 2 `333` found ",
			coll:  builder.SixWithDuplicates(),
			args:  testArgs[C, int]{value: 333},
			want1: 2,
		},
	}
}

func testCountValues[C cmpInternal[int]](t *testing.T, builder testCollectionBuilder[C, int]) {
	cases := getCountValuesCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.CountValues(tt.args.value)
			if got != tt.want1 {
				t.Errorf("CountValues() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getIndexOfCases[C any](builder testCollectionBuilder[C, int]) []testCase[C, int] {
	return []testCase[C, int]{
		{
			name:  "IndexOf() on empty collection",
			coll:  builder.Empty(),
			args:  testArgs[C, int]{value: 1},
			want1: -1,
			want2: false,
			err:   ErrValueNotFound,
		},
		{
			name:  "IndexOf() on one-item collection",
			coll:  builder.One(),
			args:  testArgs[C, int]{value: 111},
			want1: 0,
			want2: true,
		},
		{
			name:  "IndexOf() on three-item collection - first val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 111},
			want1: 0,
			want2: true,
		},
		{
			name:  "IndexOf() on three-item collection - second val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 222},
			want1: 1,
			want2: true,
		},
		{
			name:  "IndexOf() on three-item collection - third val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 333},
			want1: 2,
			want2: true,
		},
		{
			name:  "IndexOf() on three-item collection, not found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 999},
			want1: -1,
			want2: false,
			err:   ErrValueNotFound,
		},
		{
			name:  "IndexOf() on six-item collection, first `111` found ",
			coll:  builder.SixWithDuplicates(),
			args:  testArgs[C, int]{value: 111},
			want1: 0,
			want2: true,
		},
		{
			name:  "IndexOf() on six-item collection, first `222` found ",
			coll:  builder.SixWithDuplicates(),
			args:  testArgs[C, int]{value: 222},
			want1: 1,
			want2: true,
		},
		{
			name:  "IndexOf() on six-item collection, first `333` found ",
			coll:  builder.SixWithDuplicates(),
			args:  testArgs[C, int]{value: 333},
			want1: 2,
			want2: true,
		},
	}
}

func testIndexOf[C cmpInternal[int]](t *testing.T, builder testCollectionBuilder[C, int]) {
	cases := getIndexOfCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.IndexOf(tt.args.value)
			if got1 != tt.want1 {
				t.Errorf("IndexOf() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("IndexOf() got2 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func getLastIndexOfCases[C any](builder testCollectionBuilder[C, int]) []testCase[C, int] {
	return []testCase[C, int]{
		{
			name:  "LastIndexOf() on empty collection",
			coll:  builder.Empty(),
			args:  testArgs[C, int]{value: 1},
			want1: -1,
			want2: false,
		},
		{
			name:  "LastIndexOf() on one-item collection",
			coll:  builder.One(),
			args:  testArgs[C, int]{value: 111},
			want1: 0,
			want2: true,
		},
		{
			name:  "LastIndexOf() on three-item collection - first val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 111},
			want1: 0,
			want2: true,
		},
		{
			name:  "LastIndexOf() on three-item collection - second val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 222},
			want1: 1,
			want2: true,
		},
		{
			name:  "LastIndexOf() on three-item collection - third val found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 333},
			want1: 2,
			want2: true,
		},
		{
			name:  "LastIndexOf() on three-item collection, not found",
			coll:  builder.Three(),
			args:  testArgs[C, int]{value: 999},
			want1: -1,
			want2: false,
		},
		{
			name:  "LastIndexOf() on six-item collection, last `111` found ",
			coll:  builder.SixWithDuplicates(),
			args:  testArgs[C, int]{value: 111},
			want1: 3,
			want2: true,
		},
		{
			name:  "LastIndexOf() on six-item collection, last `222` found ",
			coll:  builder.SixWithDuplicates(),
			args:  testArgs[C, int]{value: 222},
			want1: 4,
			want2: true,
		},
		{
			name:  "LastIndexOf() on six-item collection, last `333` found ",
			coll:  builder.SixWithDuplicates(),
			args:  testArgs[C, int]{value: 333},
			want1: 5,
			want2: true,
		},
	}
}

func testLastIndexOf[C cmpInternal[int]](t *testing.T, builder testCollectionBuilder[C, int]) {
	cases := getLastIndexOfCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.LastIndexOf(tt.args.value)
			if got1 != tt.want1 {
				t.Errorf("LastIndexOf() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("LastIndexOf() got2 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func getMaxCases[C any](builder testCollectionBuilder[C, int]) []testCase[C, int] {
	return []testCase[C, int]{
		{
			name:  "Max() on empty collection",
			coll:  builder.Empty(),
			want1: 0,
			err:   ErrEmptyCollection,
		},
		{
			name:  "Max() on one-item collection",
			coll:  builder.One(),
			want1: 111,
		},
		{
			name:  "Max() on three-item collection",
			coll:  builder.Three(),
			want1: 333,
		},
		{
			name:  "Max() on six-item collection",
			coll:  builder.SixWithDuplicates(),
			want1: 333,
		},
	}
}

func testMax[C cmpInternal[int]](t *testing.T, builder testCollectionBuilder[C, int]) {
	cases := getMaxCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.coll.Max()
			if err != nil {
				if !errors.Is(err, tt.err) {
					t.Errorf("Max() error = %v, wantErr %v", err, tt.err)
				}
			} else if got != tt.want1 {
				t.Errorf("Max() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getMinCases[C any](builder testCollectionBuilder[C, int]) []testCase[C, int] {
	return []testCase[C, int]{
		{
			name:  "Min() on empty collection",
			coll:  builder.Empty(),
			want1: 0,
			err:   ErrEmptyCollection,
		},
		{
			name:  "Min() on one-item collection",
			coll:  builder.One(),
			want1: 111,
		},
		{
			name:  "Min() on three-item collection",
			coll:  builder.Three(),
			want1: 111,
		},
		{
			name:  "Min() on six-item collection",
			coll:  builder.SixWithDuplicates(),
			want1: 111,
		},
	}
}

func testMin[C cmpInternal[int]](t *testing.T, builder testCollectionBuilder[C, int]) {
	cases := getMinCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.coll.Min()
			if err != nil {
				if !errors.Is(err, tt.err) {
					t.Errorf("Min() error = %v, wantErr %v", err, tt.err)
				}
			} else if got != tt.want1 {
				t.Errorf("Min() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getSumCases[C any](builder testCollectionBuilder[C, int]) []testCase[C, int] {
	return []testCase[C, int]{
		{
			name:  "Sum() on empty collection",
			coll:  builder.Empty(),
			want1: 0,
		},
		{
			name:  "Sum() on one-item collection",
			coll:  builder.One(),
			want1: 111,
		},
		{
			name:  "Sum() on three-item collection",
			coll:  builder.Three(),
			want1: 666,
		},
		{
			name:  "Sum() on six-item collection",
			coll:  builder.SixWithDuplicates(),
			want1: 1332,
		},
	}
}

func testSum[C cmpInternal[int]](t *testing.T, builder testCollectionBuilder[C, int]) {
	cases := getSumCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.Sum()
			if got != tt.want1 {
				t.Errorf("Sum() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}
