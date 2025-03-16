package coll

import (
	"reflect"
	"slices"
	"testing"
)

type baseIntArgs = testArgs[baseInternal[int], int]
type baseTestCase = testCase[baseInternal[int], int]
type baseCollIntBuilder = testCollectionBuilder[baseInternal[int]]

type baseIntPairArgs = testArgs[baseInternal[Pair[int, int]], Pair[int, int]]
type baseIntPairTestCase = testCase[baseInternal[Pair[int, int]], Pair[int, int]]
type baseCollIntPairBuilder = testPairCollectionBuilder[baseInternal[Pair[int, int]]]

func getContainsCases(builder baseCollIntBuilder) []baseTestCase {
	return []baseTestCase{
		{
			name:  "Contains() on empty collection",
			coll:  builder.Empty(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return v == 1 }},
			want1: false,
		},
		{
			name:  "Contains() on one-item collection",
			coll:  builder.One(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return v == 111 }},
			want1: true,
		},
		{
			name:  "Contains() on three-item collection",
			coll:  builder.Three(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return v == 222 }},
			want1: true,
		},
		{
			name:  "Contains() on three-item collection, all false",
			coll:  builder.Three(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return false }},
			want1: false,
		},
		{
			name:  "Contains() on three-item collection, not found",
			coll:  builder.Three(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return v == 999 }},
			want1: false,
		},
	}
}

func testContains(t *testing.T, builder baseCollIntBuilder) {
	cases := getContainsCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.coll.Contains(tt.args.predicate); got != tt.want1 {
				t.Errorf("Contains() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getCountCases(builder baseCollIntBuilder) []baseTestCase {
	return []baseTestCase{
		{
			name:  "Count() on empty collection",
			coll:  builder.Empty(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return v == 1 }},
			want1: 0,
		},
		{
			name:  "Count() on one-item collection",
			coll:  builder.One(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return v == 111 }},
			want1: 1,
		},
		{
			name:  "Count() on three-item collection",
			coll:  builder.Three(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return v == 222 }},
			want1: 1,
		},
		{
			name:  "Count() on three-item collection, all false",
			coll:  builder.Three(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return false }},
			want1: 0,
		},
		{
			name:  "Count() on three-item collection, all true",
			coll:  builder.Three(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return true }},
			want1: 3,
		},
		{
			name:  "Count() on three-item collection, some mod 2",
			coll:  builder.Three(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return v%2 == 0 }},
			want1: 1,
		},
		{
			name:  "Count() on three-item collection, some not mod 2",
			coll:  builder.Three(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return v%2 != 0 }},
			want1: 2,
		},
		{
			name:  "Count() on three-item collection, 2 found",
			coll:  builder.Three(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return v == 111 }},
			want1: 1,
		},
		{
			name:  "Count() on three-item collection, not found",
			coll:  builder.Three(),
			args:  baseIntArgs{predicate: func(i int, v int) bool { return v == 999 }},
			want1: 0,
		},
		{
			name: "Count() on three-item collection, compare indexes and values",
			coll: builder.Three(),
			args: baseIntArgs{predicate: func(i int, v int) bool {
				if i == 0 && v == 111 {
					return true
				}
				if i == 1 && v == 222 {
					return true
				}
				if i == 2 && v == 333 {
					return true
				}
				return false
			}},
			want1: 3,
		},
	}
}

func testCount(t *testing.T, builder baseCollIntBuilder) {
	cases := getCountCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.Count(tt.args.predicate)
			if got != tt.want1 {
				t.Errorf("Count() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getEachCases(t *testing.T, builder baseCollIntBuilder) []*baseTestCase {

	// eachOnEmptyListCase:

	eachOnEmptyListCase := &baseTestCase{
		name: "Each() on empty collection",
		coll: builder.Empty(),
	}
	eachOnEmptyListCase.args = baseIntArgs{
		visit: func(i int, v int) {
			t.Error("Each() called on empty collection")
		},
	}

	// eachOnOneItemCase:

	eachOnOneItemCase := &baseTestCase{
		name:  "Each() on one-item collection",
		coll:  builder.One(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0},
		want2: []int{111},
	}

	eachOnOneItemCase.args = baseIntArgs{
		visit: func(i int, v int) {
			if i != 0 || v != 111 {
				t.Error("Each() called with wrong values")
			}
			eachOnOneItemCase.got1 = append(eachOnOneItemCase.got1.([]int), i)
			eachOnOneItemCase.got2 = append(eachOnOneItemCase.got2.([]int), v)
		},
	}

	// eachOnEmptyListCase:

	eachOnThreeCase := &baseTestCase{
		name:  "Each() on three-item collection",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0, 1, 2},
		want2: []int{111, 222, 333},
	}

	eachOnThreeCase.args = baseIntArgs{
		visit: func(i int, v int) {
			if i < 0 || i > 2 || v < 111 || v > 333 {
				t.Error("Each() called with wrong values")
			}
			eachOnThreeCase.got1 = append(eachOnThreeCase.got1.([]int), i)
			eachOnThreeCase.got2 = append(eachOnThreeCase.got2.([]int), v)
		},
	}

	// put the cases together:

	return []*baseTestCase{
		eachOnEmptyListCase,
		eachOnOneItemCase,
		eachOnThreeCase,
	}
}

func testEach(t *testing.T, builder baseCollIntBuilder) {
	cases := getEachCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.Each(tt.args.visit)
			if tt.got1 != nil && !reflect.DeepEqual(tt.got1, tt.want1) {
				t.Errorf("Each() called with wrong indices: %v, want1 %v", tt.got1, tt.want1)
			}
			if tt.got2 != nil && !reflect.DeepEqual(tt.got2, tt.want2) {
				t.Errorf("Each() called with wrong values: %v, want1 %v", tt.got2, tt.want2)
			}
		})
	}
}

func getEachUntilCases(t *testing.T, builder baseCollIntBuilder) []*baseTestCase {

	// eachUntilOnEmptyListCase:

	eachUntilOnEmptyListCase := &baseTestCase{
		name: "EachUntil() on empty collection",
		coll: builder.Empty(),
	}

	eachUntilOnEmptyListCase.args = baseIntArgs{
		predicate: func(i int, v int) bool {
			t.Error("EachUntil() called on empty collection")
			return true
		},
	}

	// eachUntilOnOneItemCase:

	eachUntilOnOneItemCase := &baseTestCase{
		name:  "EachUntil() on one-item collection",
		coll:  builder.One(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0},
		want2: []int{111},
	}

	eachUntilOnOneItemCase.args = baseIntArgs{
		predicate: func(i int, v int) bool {
			if i != 0 || v != 111 {
				t.Error("EachUntil() called with wrong values")
			}
			eachUntilOnOneItemCase.got1 = append(eachUntilOnOneItemCase.got1.([]int), i)
			eachUntilOnOneItemCase.got2 = append(eachUntilOnOneItemCase.got2.([]int), v)
			return true
		},
	}

	// eachUntilFinishMiddleCase:

	eachUntilFinishMiddleCase := &baseTestCase{
		name:  "EachUntil() finish in middle",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0, 1},
		want2: []int{111, 222},
	}

	eachUntilFinishMiddleCase.args = baseIntArgs{
		predicate: func(i int, v int) bool {
			if i < 0 || i > 2 || v < 111 || v > 333 {
				t.Error("EachUntil() called with wrong values")
			}
			eachUntilFinishMiddleCase.got1 = append(eachUntilFinishMiddleCase.got1.([]int), i)
			eachUntilFinishMiddleCase.got2 = append(eachUntilFinishMiddleCase.got2.([]int), v)
			stop := i >= 1 && v >= 222
			cont := !stop
			return cont
		},
	}

	// eachUntilAllThreeCase:

	eachUntilAllThreeCase := &baseTestCase{
		name:  "EachUntil() on three-item collection",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0, 1, 2},
		want2: []int{111, 222, 333},
	}

	eachUntilAllThreeCase.args = baseIntArgs{
		predicate: func(i int, v int) bool {
			if i < 0 || i > 2 || v < 111 || v > 333 {
				t.Error("EachUntil() called with wrong values")
			}
			eachUntilAllThreeCase.got1 = append(eachUntilAllThreeCase.got1.([]int), i)
			eachUntilAllThreeCase.got2 = append(eachUntilAllThreeCase.got2.([]int), v)
			return true
		},
	}

	// put the cases together:

	return []*baseTestCase{
		eachUntilOnEmptyListCase,
		eachUntilOnOneItemCase,
		eachUntilFinishMiddleCase,
		eachUntilAllThreeCase,
	}
}

func testEachUntil(t *testing.T, builder baseCollIntBuilder) {
	cases := getEachUntilCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.EachUntil(tt.args.predicate)
			if tt.got1 != nil && !reflect.DeepEqual(tt.got1, tt.want1) {
				t.Errorf("EachUntil() called with wrong indices: %v, want1 %v", tt.got1, tt.want1)
			}
			if tt.got2 != nil && !reflect.DeepEqual(tt.got2, tt.want2) {
				t.Errorf("EachUntil() called with wrong values: %v, want1 %v", tt.got2, tt.want2)
			}
		})
	}
}

func getFindCases(builder baseCollIntBuilder) []*baseTestCase {
	return []*baseTestCase{
		{
			name: "Find() on empty collection",
			coll: builder.Empty(),
			args: baseIntArgs{
				predicate:    func(i int, v int) bool { return v == 1 },
				defaultValue: -1,
			},
			want1: -1,
		},
		{
			name: "Find() on one-item collection",
			coll: builder.One(),
			args: baseIntArgs{
				predicate:    func(i int, v int) bool { return v == 111 },
				defaultValue: -1,
			},
			want1: 111,
		},
		{
			name: "Find() on three-item collection",
			coll: builder.Three(),
			args: baseIntArgs{
				predicate:    func(i int, v int) bool { return v == 222 },
				defaultValue: -1,
			},
			want1: 222,
		},
		{
			name: "Find() on three-item collection, not found",
			coll: builder.Three(),
			args: baseIntArgs{
				predicate:    func(i int, v int) bool { return v == 999 },
				defaultValue: -1,
			},
			want1: -1,
		},
		{
			name: "Find() first one",
			coll: builder.Three(),
			args: baseIntArgs{
				predicate:    func(i int, v int) bool { return true },
				defaultValue: -1,
			},
			want1: 111,
		},
		{
			name: "Find() on three-item collection, compare indexes and values",
			coll: builder.Three(),
			args: baseIntArgs{predicate: func(i int, v int) bool {
				if i == 0 && v == 111 {
					return true
				}
				return false
			}},
			want1: 111,
		},
	}
}

func testFind(t *testing.T, builder baseCollIntBuilder) {
	cases := getFindCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.Find(tt.args.predicate, tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Find() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getFindCasesWithDupes(builder baseCollIntPairBuilder) []*baseIntPairTestCase {
	return []*baseIntPairTestCase{
		{
			name: "Find() on six-item collection, first one",
			coll: builder.SixWithDuplicates(),
			args: baseIntPairArgs{
				predicate:    func(i int, p Pair[int, int]) bool { return true },
				defaultValue: nil,
			},
			want1: NewPair(1, 111),
		},
		{
			name: "Find() on six-item collection, second one",
			coll: builder.SixWithDuplicates(),
			args: baseIntPairArgs{
				predicate:    func(i int, p Pair[int, int]) bool { return p.Val() == 222 },
				defaultValue: nil,
			},
			want1: NewPair(2, 222),
		},
		{
			name: "Find() on six-item collection, not found",
			coll: builder.SixWithDuplicates(),
			args: baseIntPairArgs{
				predicate:    func(i int, p Pair[int, int]) bool { return p.Val() == 999 },
				defaultValue: nil,
			},
			want1: nil,
		},
	}
}

func testFindWithDupes(t *testing.T, builder baseCollIntPairBuilder) {
	cases := getFindCasesWithDupes(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.Find(tt.args.predicate, tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Find() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getFoldCases(builder baseCollIntBuilder) []*baseTestCase {
	return []*baseTestCase{
		{
			name: "Fold() on empty collection",
			coll: builder.Empty(),
			args: baseIntArgs{
				reducer: func(acc int, i int, current int) int {
					return acc*10 + current
				},
				initial: 100,
			},
			want1: 100,
		},
		{
			name: "Fold() on one-item collection",
			coll: builder.One(),
			args: baseIntArgs{
				reducer: func(acc int, i int, current int) int {
					return acc*10 + current
				},
				initial: 0,
			},
			want1: 111,
		},
		{
			name: "Fold() on three-item collection",
			coll: builder.Three(),
			args: baseIntArgs{
				reducer: func(acc int, i int, current int) int {
					return acc*10 + current
				},
				initial: 100,
			},
			want1: 113653,
		},
		{
			name: "Fold() on three-item collection, include index",
			coll: builder.Three(),
			args: baseIntArgs{
				reducer: func(acc int, i int, current int) int {
					return acc*(i+1) + current
				},
				initial: 100,
			},
			want1: ((100+111)*2+222)*3 + 333,
		},
	}
}

func testFold(t *testing.T, builder baseCollIntBuilder) {
	cases := getFoldCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.Fold(tt.args.reducer, tt.args.initial)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("Reduce() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

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

func getReduceCases(t *testing.T, builder baseCollIntBuilder) []*baseTestCase {
	return []*baseTestCase{
		{
			name: "Reduce() on empty collection",
			coll: builder.Empty(),
			args: baseIntArgs{reducer: func(acc int, _ int, current int) int {
				return acc*10 + current
			}},
			want1: 0,
			want2: ErrEmptyCollection,
		},
		{
			name: "Fold() on one-item collection",
			coll: builder.One(),
			args: baseIntArgs{reducer: func(acc int, _ int, current int) int {
				return acc*10 + current
			}},
			want1: 111,
			want2: nil,
		},
		{
			name: "Fold() on three-item collection",
			coll: builder.Three(),
			args: baseIntArgs{reducer: func(acc int, _ int, current int) int {
				return acc*10 + current
			}},
			want1: 13653,
			want2: nil,
		},
	}
}

func testReduce(t *testing.T, builder baseCollIntBuilder) {
	cases := getReduceCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.Reduce(tt.args.reducer)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Reduce() = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Reduce() = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func getSearchCases(builder baseCollIntBuilder) []*baseTestCase {

	searchOnEmptyCollectionCase := &baseTestCase{
		name:  "Search() on empty collection",
		coll:  builder.Empty(),
		want1: 0,
		want2: false,
		want3: nil,
		got3:  nil,
	}

	searchOnEmptyCollectionCase.args = baseIntArgs{predicate: func(i int, v int) bool {
		if v == 1 {
			searchOnEmptyCollectionCase.got3 = i
			return true
		}
		return false
	}}

	searchOnOneItemCollectionCase := &baseTestCase{
		name:  "Search() on one-item collection",
		coll:  builder.One(),
		want1: 111,
		want2: true,
		want3: 0,
		got3:  nil,
	}

	searchOnOneItemCollectionCase.args = baseIntArgs{predicate: func(i int, v int) bool {
		if v == 111 {
			searchOnOneItemCollectionCase.got3 = i
			return true
		}
		return false
	}}

	searchOnThreeItemCollectionCase := &baseTestCase{
		name:  "Search() on three-item collection",
		coll:  builder.Three(),
		want1: 222,
		want2: true,
		want3: 1,
		got3:  nil,
	}

	searchOnThreeItemCollectionCase.args = baseIntArgs{predicate: func(i int, v int) bool {
		if v == 222 {
			searchOnThreeItemCollectionCase.got3 = i
			return true
		}
		return false
	}}

	searchOnOneItemCollectionCaseNotFound := &baseTestCase{
		name:  "Search() on one-item collection, not found",
		coll:  builder.One(),
		want1: 0,
		want2: false,
		want3: nil,
		got3:  nil,
	}

	searchOnOneItemCollectionCaseNotFound.args = baseIntArgs{predicate: func(i int, v int) bool {
		if v == 999 {
			searchOnOneItemCollectionCaseNotFound.got3 = i
			return true
		}
		return false
	}}

	searchOnThreeItemCollectionCaseAllFound := &baseTestCase{
		name:  "Search() on three-item collection, all found",
		coll:  builder.Three(),
		want1: 111,
		want2: true,
		want3: 0,
		got3:  nil,
	}

	searchOnThreeItemCollectionCaseAllFound.args = baseIntArgs{predicate: func(i int, v int) bool {
		searchOnThreeItemCollectionCaseAllFound.got3 = i
		return true
	}}

	searchOnSixItemCollectionCaseFirstFound := &baseTestCase{
		name:  "Search() on six-item collection, found first occurrence",
		coll:  builder.SixWithDuplicates(),
		want1: 111,
		want2: true,
		want3: 0,
		got3:  nil,
	}

	searchOnSixItemCollectionCaseFirstFound.args = baseIntArgs{predicate: func(i int, v int) bool {
		if v == 111 && i == 0 {
			searchOnSixItemCollectionCaseFirstFound.got3 = i
			return true
		}
		return false
	}}

	searchOnSixItemCollectionCaseLastFound := &baseTestCase{
		name:  "Search() on six-item collection, found second occurrence",
		coll:  builder.SixWithDuplicates(),
		want1: 111,
		want2: true,
		want3: 3,
		got3:  nil,
	}

	searchOnSixItemCollectionCaseLastFound.args = baseIntArgs{predicate: func(i int, v int) bool {
		if v == 111 && i == 3 {
			searchOnSixItemCollectionCaseLastFound.got3 = i
			return true
		}
		return false
	}}

	return []*baseTestCase{
		searchOnEmptyCollectionCase,
		searchOnOneItemCollectionCase,
		searchOnThreeItemCollectionCase,
		searchOnOneItemCollectionCaseNotFound,
		searchOnThreeItemCollectionCaseAllFound,
		searchOnSixItemCollectionCaseFirstFound,
		searchOnSixItemCollectionCaseLastFound,
	}
}

func testSearch(t *testing.T, builder baseCollIntBuilder) {
	cases := getSearchCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.Search(tt.args.predicate)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Search() got1 = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Search() got2 = %v, want2 %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(tt.got3, tt.want3) {
				t.Errorf("Search() got3 = %v, want3 %v", tt.got3, tt.want3)
			}
		})
	}
}

func getSearchPairCases(builder baseCollIntPairBuilder) []*baseIntPairTestCase {
	return []*baseIntPairTestCase{
		{
			name: "Search() pair on six-item collection, found first occurrence",
			coll: builder.SixWithDuplicates(),
			args: baseIntPairArgs{predicate: func(i int, v Pair[int, int]) bool {
				return v.Val() == 111
			}},
			want1: NewPair(1, 111),
			want2: true,
		},
		{
			name: "Search() pair on six-item collection, found first occurrence",
			coll: builder.SixWithDuplicates(),
			args: baseIntPairArgs{predicate: func(i int, v Pair[int, int]) bool {
				return v.Val() == 222
			}},
			want1: NewPair(2, 222),
			want2: true,
		},
		{
			name: "Search() pair on six-item collection, found first occurrence",
			coll: builder.SixWithDuplicates(),
			args: baseIntPairArgs{predicate: func(i int, v Pair[int, int]) bool {
				return v.Val() == 333
			}},
			want1: NewPair(3, 333),
			want2: true,
		},
	}
}

func testSearchPair(t *testing.T, builder baseCollIntPairBuilder) {
	cases := getSearchPairCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.Search(tt.args.predicate)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Search() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Search() got2 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func getToSliceCases(builder baseCollIntBuilder) []*baseTestCase {
	return []*baseTestCase{
		{
			name:  "ToSlice() on empty collection",
			coll:  builder.Empty(),
			want1: []int{},
		},
		{
			name:  "ToSlice() on one-item collection",
			coll:  builder.One(),
			want1: []int{111},
		},
		{
			name:  "ToSlice() on three-item collection",
			coll:  builder.Three(),
			want1: []int{111, 222, 333},
		},
	}
}

func testToSlice(t *testing.T, builder baseCollIntBuilder) {
	cases := getToSliceCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.ToSlice()
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("ToSlice() = %v, want1 %v", got, tt.want1)
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
				predicate: func(i int, v int) bool {
					return v < 222
				},
			},
			want1: []int{111},
		},
		{
			name: "Values() on three-item collection, break",
			coll: builder.Three(),
			args: baseIntArgs{
				predicate: func(i int, v int) bool {
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
				if !tt.args.predicate(-1, v) {
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
