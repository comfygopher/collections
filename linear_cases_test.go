package coll

import (
	"reflect"
	"testing"
)

type linearIntArgs = testArgs[linearInternal[int], int]
type linearTestCase = testCase[linearInternal[int], int]
type linearCollIntBuilder = testCollectionBuilder[linearInternal[int]]

type linearIntPairArgs = testArgs[linearInternal[Pair[int, int]], Pair[int, int]]
type linearPairTestCase = testCase[linearInternal[Pair[int, int]], Pair[int, int]]
type linearCollIntPairBuilder = testPairCollectionBuilder[linearInternal[Pair[int, int]]]

func getEachRevCases(t *testing.T, builder linearCollIntBuilder) []*linearTestCase {

	// eachRevOnEmptyListCase:

	eachRevOnEmptyListCase := &linearTestCase{
		name: "EachRev() on empty collection",
		coll: builder.Empty(),
	}
	eachRevOnEmptyListCase.args = linearIntArgs{
		visit: func(i int, v int) {
			t.Error("EachRev() called on empty collection")
		},
	}

	// eachRevOnOneItemCase:

	eachRevOnOneItemCase := &linearTestCase{
		name:  "EachRev() on one-item collection",
		coll:  builder.One(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0},
		want2: []int{111},
	}
	eachRevOnOneItemCase.args = linearIntArgs{
		visit: func(i int, v int) {
			if i != 0 || v != 111 {
				t.Error("EachRev() called with wrong values")
			}
			eachRevOnOneItemCase.got1 = append(eachRevOnOneItemCase.got1.([]int), i)
			eachRevOnOneItemCase.got2 = append(eachRevOnOneItemCase.got2.([]int), v)
		},
	}

	// eachRevOnThreeCase:

	eachRevOnThreeCase := &linearTestCase{
		name:  "EachRev() on three-item collection",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{2, 1, 0},
		want2: []int{333, 222, 111},
	}

	eachRevOnThreeCase.args = linearIntArgs{
		visit: func(i int, v int) {
			if i < 0 || i > 2 || v < 111 || v > 333 {
				t.Error("EachRev() called with wrong values")
			}
			eachRevOnThreeCase.got1 = append(eachRevOnThreeCase.got1.([]int), i)
			eachRevOnThreeCase.got2 = append(eachRevOnThreeCase.got2.([]int), v)
		},
	}

	// put the cases together:

	return []*linearTestCase{
		eachRevOnEmptyListCase,
		eachRevOnOneItemCase,
		eachRevOnThreeCase,
	}
}

func testEachRev(t *testing.T, builder linearCollIntBuilder) {
	cases := getEachRevCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.EachRev(tt.args.visit)
			if tt.got1 != nil && !reflect.DeepEqual(tt.got1, tt.want1) {
				t.Errorf("EachRev() called with wrong indices: %v, want1 %v", tt.got1, tt.want1)
			}
			if tt.got2 != nil && !reflect.DeepEqual(tt.got2, tt.want2) {
				t.Errorf("EachRev() called with wrong values: %v, want1 %v", tt.got2, tt.want2)
			}
		})
	}
}

func getEachRevUntilCases(t *testing.T, builder linearCollIntBuilder) []*linearTestCase {

	// eachRevUntilOnEmptyListCase:

	eachRevUntilOnEmptyListCase := &linearTestCase{
		name: "EachRevUntil() on empty collection",
		coll: builder.Empty(),
	}

	eachRevUntilOnEmptyListCase.args = linearIntArgs{
		predicate: func(i int, v int) bool {
			t.Error("EachRevUntil() called on empty collection")
			return true
		},
	}

	// eachRevUntilOnOneItemCase:

	eachRevUntilOnOneItemCase := &linearTestCase{
		name:  "EachRevUntil() on one-item collection",
		coll:  builder.One(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0},
		want2: []int{111},
	}

	eachRevUntilOnOneItemCase.args = linearIntArgs{
		predicate: func(i int, v int) bool {
			if i != 0 || v != 111 {
				t.Error("EachRevUntil() called with wrong values")
			}
			eachRevUntilOnOneItemCase.got1 = append(eachRevUntilOnOneItemCase.got1.([]int), i)
			eachRevUntilOnOneItemCase.got2 = append(eachRevUntilOnOneItemCase.got2.([]int), v)
			return true
		},
	}

	// eachRevUntilFinishMiddleCase:

	eachRevUntilFinishMiddleCase := &linearTestCase{
		name:  "EachRevUntil() finish in middle",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{2, 1},
		want2: []int{333, 222},
	}

	eachRevUntilFinishMiddleCase.args = linearIntArgs{
		predicate: func(i int, v int) bool {
			if i < 0 || i > 2 || v < 111 || v > 333 {
				t.Error("EachRevUntil() called with wrong values")
			}
			eachRevUntilFinishMiddleCase.got1 = append(eachRevUntilFinishMiddleCase.got1.([]int), i)
			eachRevUntilFinishMiddleCase.got2 = append(eachRevUntilFinishMiddleCase.got2.([]int), v)
			stop := i <= 1 && v <= 222
			cont := !stop
			return cont
		},
	}

	// eachRevUntilAllThreeCase:

	eachRevUntilAllThreeCase := &linearTestCase{
		name:  "EachRevUntil() on three-item collection",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{2, 1, 0},
		want2: []int{333, 222, 111},
	}

	eachRevUntilAllThreeCase.args = linearIntArgs{
		predicate: func(i int, v int) bool {
			if i < 0 || i > 2 || v < 111 || v > 333 {
				t.Error("EachRevUntil() called with wrong values")
			}
			eachRevUntilAllThreeCase.got1 = append(eachRevUntilAllThreeCase.got1.([]int), i)
			eachRevUntilAllThreeCase.got2 = append(eachRevUntilAllThreeCase.got2.([]int), v)
			return true
		},
	}

	// put the cases together:

	return []*linearTestCase{
		eachRevUntilOnEmptyListCase,
		eachRevUntilOnOneItemCase,
		eachRevUntilFinishMiddleCase,
		eachRevUntilAllThreeCase,
	}
}

func testEachRevUntil(t *testing.T, builder linearCollIntBuilder) {
	cases := getEachRevUntilCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.coll.EachRevUntil(tt.args.predicate)
			if tt.got1 != nil && !reflect.DeepEqual(tt.got1, tt.want1) {
				t.Errorf("EachRevUntil() called with wrong indices: %v, want1 %v", tt.got1, tt.want1)
			}
			if tt.got2 != nil && !reflect.DeepEqual(tt.got2, tt.want2) {
				t.Errorf("EachRevUntil() called with wrong values: %v, want1 %v", tt.got2, tt.want2)
			}
		})
	}
}

func getFindLastCases(builder linearCollIntBuilder) []*linearTestCase {
	return []*linearTestCase{
		{
			name: "FindLast() on empty collection",
			coll: builder.Empty(),
			args: linearIntArgs{
				predicate:    func(i int, v int) bool { return v == 111 },
				defaultValue: -1,
			},
			want1: -1,
		},
		{
			name: "FindLast() on one-item collection",
			coll: builder.One(),
			args: linearIntArgs{
				predicate:    func(i int, v int) bool { return v == 111 },
				defaultValue: -1,
			},
			want1: 111,
		},
		{
			name: "FindLast() on three-item collection",
			coll: builder.Three(),
			args: linearIntArgs{
				predicate:    func(i int, v int) bool { return v == 222 },
				defaultValue: -1,
			},
			want1: 222,
		},
		{
			name: "FindLast() on three-item collection, not found",
			coll: builder.Three(),
			args: linearIntArgs{
				predicate:    func(i int, v int) bool { return v == 999 },
				defaultValue: -1,
			},
			want1: -1,
		},
		{
			name: "FindLast() last one",
			coll: builder.Three(),
			args: linearIntArgs{
				predicate:    func(i int, v int) bool { return true },
				defaultValue: -1,
			},
			want1: 333,
		},
		{
			name: "FindLast() on three-item collection, compare indexes and values",
			coll: builder.Three(),
			args: linearIntArgs{predicate: func(i int, v int) bool {
				if i == 2 && v == 333 {
					return true
				}
				return false
			}},
			want1: 333,
		},
	}
}

func testFindLast(t *testing.T, builder linearCollIntBuilder) {
	cases := getFindLastCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.FindLast(tt.args.predicate, tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("FindLast() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getFindLastCasesWithDupes(builder linearCollIntPairBuilder) []*linearPairTestCase {
	return []*linearPairTestCase{
		{
			name: "FindLast() on six-item collection, first one",
			coll: builder.SixWithDuplicates(),
			args: linearIntPairArgs{
				predicate:    func(i int, p Pair[int, int]) bool { return true },
				defaultValue: nil,
			},
			want1: NewPair(6, 333),
		},
		{
			name: "FindLast() on six-item collection, second one",
			coll: builder.SixWithDuplicates(),
			args: linearIntPairArgs{
				predicate:    func(i int, p Pair[int, int]) bool { return p.Val() == 222 },
				defaultValue: nil,
			},
			want1: NewPair(5, 222),
		},
		{
			name: "FindLast() on six-item collection, not found",
			coll: builder.SixWithDuplicates(),
			args: linearIntPairArgs{
				predicate:    func(i int, p Pair[int, int]) bool { return p.Val() == 999 },
				defaultValue: nil,
			},
			want1: nil,
		},
	}
}

func testFindLastWithDupes(t *testing.T, builder linearCollIntPairBuilder) {
	cases := getFindLastCasesWithDupes(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.FindLast(tt.args.predicate, tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("FindLast() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getFoldRevCases(builder linearCollIntBuilder) []*linearTestCase {
	return []*linearTestCase{
		{
			name: "FoldRev() on empty collection",
			coll: builder.Empty(),
			args: linearIntArgs{
				reducer: func(acc int, _ int, current int) int {
					return acc*10 + current
				},
				initial: 100,
			},
			want1: 100,
		},
		{
			name: "FoldRev() on one-item collection",
			coll: builder.One(),
			args: linearIntArgs{
				reducer: func(acc int, _ int, current int) int {
					return acc*10 + current
				},
				initial: 100,
			},
			want1: 1111,
		},
		{
			name: "FoldRev() on three-item collection",
			coll: builder.Three(),
			args: linearIntArgs{
				reducer: func(acc int, _ int, current int) int {
					return acc*10 + current
				},
				initial: 100,
			},
			want1: 135631,
		},
		{
			name: "FoldRev() on three-item collection, include index",
			coll: builder.Three(),
			args: linearIntArgs{
				reducer: func(acc int, i int, current int) int {
					return acc*(i+1) + current
				},
				initial: 100,
			},
			want1: ((100*3+333)*2 + 222) + 111,
		},
	}
}

func testFoldRev(t *testing.T, builder linearCollIntBuilder) {
	cases := getFoldRevCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.FoldRev(tt.args.reducer, tt.args.initial)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("FoldRev() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getHeadCases(builder linearCollIntBuilder) []linearTestCase {
	return []linearTestCase{
		{
			name:  "Head() on empty collection",
			coll:  builder.Empty(),
			want1: 0,
			want2: false,
		},
		{
			name:  "Head() on one-item collection",
			coll:  builder.One(),
			want1: 111,
			want2: true,
		},
		{
			name:  "Head() on three-item collection",
			coll:  builder.Three(),
			want1: 111,
			want2: true,
		},
	}
}

func testHead(t *testing.T, builder linearCollIntBuilder) {
	cases := getHeadCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.Head()
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Head() got1 = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Head() got2 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func getHeadOrDefaultCases(builder linearCollIntBuilder) []linearTestCase {
	return []linearTestCase{
		{
			name:  "HeadOrDefault() on empty collection",
			coll:  builder.Empty(),
			args:  linearIntArgs{defaultValue: -1},
			want1: -1,
		},
		{
			name:  "HeadOrDefault() on one-item collection",
			coll:  builder.One(),
			args:  linearIntArgs{defaultValue: -1},
			want1: 111,
		},
		{
			name:  "HeadOrDefault() on three-item collection",
			coll:  builder.Three(),
			args:  linearIntArgs{defaultValue: -1},
			want1: 111,
		},
	}
}

func testHeadOrDefault(t *testing.T, builder linearCollIntBuilder) {
	cases := getHeadOrDefaultCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.HeadOrDefault(tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("HeadOrDefault() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getSearchRevCases(builder linearCollIntBuilder) []*linearTestCase {

	searchRevOnEmptyCollectionCase := &linearTestCase{
		name:  "SearchRev() on empty collection",
		coll:  builder.Empty(),
		want1: 0,
		want2: false,
		want3: nil,
		got3:  nil,
	}

	searchRevOnEmptyCollectionCase.args = linearIntArgs{predicate: func(i int, v int) bool {
		if v == 111 {
			searchRevOnEmptyCollectionCase.got3 = i
			return true
		}
		return false
	}}

	searchRevOnOneItemCollectionCase := &linearTestCase{
		name:  "SearchRev() on one-item collection",
		coll:  builder.One(),
		want1: 111,
		want2: true,
		want3: 0,
		got3:  nil,
	}

	searchRevOnOneItemCollectionCase.args = linearIntArgs{predicate: func(i int, v int) bool {
		if v == 111 {
			searchRevOnOneItemCollectionCase.got3 = i
			return true
		}
		return false
	}}

	searchRevOnThreeItemCollectionCase := &linearTestCase{
		name:  "SearchRev() on three-item collection",
		coll:  builder.Three(),
		want1: 222,
		want2: true,
		want3: 1,
		got3:  nil,
	}

	searchRevOnThreeItemCollectionCase.args = linearIntArgs{predicate: func(i int, v int) bool {
		if v == 222 {
			searchRevOnThreeItemCollectionCase.got3 = i
			return true
		}
		return false
	}}

	searchRevOnOneItemCollectionCaseNotFound := &linearTestCase{
		name:  "SearchRev() on one-item collection, not found",
		coll:  builder.One(),
		want1: 0,
		want2: false,
		want3: nil,
		got3:  nil,
	}

	searchRevOnOneItemCollectionCaseNotFound.args = linearIntArgs{predicate: func(i int, v int) bool {
		if v == 999 {
			searchRevOnOneItemCollectionCaseNotFound.got3 = i
			return true
		}
		return false
	}}

	searchRevOnThreeItemCollectionCaseAllFound := &linearTestCase{
		name:  "SearchRev() on three-item collection, all found",
		coll:  builder.Three(),
		want1: 333,
		want2: true,
		want3: 2,
		got3:  nil,
	}

	searchRevOnThreeItemCollectionCaseAllFound.args = linearIntArgs{predicate: func(i int, v int) bool {
		searchRevOnThreeItemCollectionCaseAllFound.got3 = i
		return true
	}}

	searchRevOnSixItemCollectionCaseFirstFound := &linearTestCase{
		name:  "SearchRev() on six-item collection, found first occurrence",
		coll:  builder.SixWithDuplicates(),
		want1: 111,
		want2: true,
		want3: 3,
		got3:  nil,
	}

	searchRevOnSixItemCollectionCaseFirstFound.args = linearIntArgs{predicate: func(i int, v int) bool {
		if v == 111 && i == 3 {
			searchRevOnSixItemCollectionCaseFirstFound.got3 = i
			return true
		}
		return false
	}}

	searchRevOnSixItemCollectionCaseLastFound := &linearTestCase{
		name:  "SearchRev() on six-item collection, found second occurrence",
		coll:  builder.SixWithDuplicates(),
		want1: 111,
		want2: true,
		want3: 0,
		got3:  nil,
	}

	searchRevOnSixItemCollectionCaseLastFound.args = linearIntArgs{predicate: func(i int, v int) bool {
		if v == 111 && i == 0 {
			searchRevOnSixItemCollectionCaseLastFound.got3 = i
			return true
		}
		return false
	}}

	return []*linearTestCase{
		searchRevOnEmptyCollectionCase,
		searchRevOnOneItemCollectionCase,
		searchRevOnThreeItemCollectionCase,
		searchRevOnOneItemCollectionCaseNotFound,
		searchRevOnThreeItemCollectionCaseAllFound,
		searchRevOnSixItemCollectionCaseFirstFound,
		searchRevOnSixItemCollectionCaseLastFound,
	}
}

func testSearchRev(t *testing.T, builder linearCollIntBuilder) {
	cases := getSearchRevCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.SearchRev(tt.args.predicate)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SearchRev() got1 = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("SearchRev() got2 = %v, want2 %v", got2, tt.want2)
			}
			if !reflect.DeepEqual(tt.got3, tt.want3) {
				t.Errorf("Search() got3 = %v, want3 %v", tt.got3, tt.want3)
			}
		})
	}
}

func getSearchRevPairCases(builder linearCollIntPairBuilder) []*linearPairTestCase {
	return []*linearPairTestCase{
		{
			name: "SearchRev() on six-item collection, found first occurrence",
			coll: builder.SixWithDuplicates(),
			args: linearIntPairArgs{predicate: func(i int, v Pair[int, int]) bool {
				return v.Val() == 111
			}},
			want1: NewPair(4, 111),
			want2: true,
		},
		{
			name: "SearchRev() on six-item collection, found first occurrence",
			coll: builder.SixWithDuplicates(),
			args: linearIntPairArgs{predicate: func(i int, v Pair[int, int]) bool {
				return v.Val() == 222
			}},
			want1: NewPair(5, 222),
			want2: true,
		},
		{
			name: "SearchRev() on six-item collection, found first occurrence",
			coll: builder.SixWithDuplicates(),
			args: linearIntPairArgs{predicate: func(i int, v Pair[int, int]) bool {
				return v.Val() == 333
			}},
			want1: NewPair(6, 333),
			want2: true,
		},
	}
}

func testSearchRevPair(t *testing.T, builder linearCollIntPairBuilder) {
	cases := getSearchRevPairCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.SearchRev(tt.args.predicate)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SearchRev() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("SearchRev() got2 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func getReduceRevCases(t *testing.T, builder linearCollIntBuilder) []*linearTestCase {
	return []*linearTestCase{
		{
			name: "Reduce() on empty collection",
			coll: builder.Empty(),
			args: linearIntArgs{reducer: func(acc int, _ int, current int) int {
				return acc*10 + current
			}},
			want1: 0,
			want2: ErrEmptyCollection,
		},
		{
			name: "Fold() on one-item collection",
			coll: builder.One(),
			args: linearIntArgs{reducer: func(acc int, _ int, current int) int {
				return acc*10 + current
			}},
			want1: 111,
			want2: nil,
		},
		{
			name: "Fold() on three-item collection",
			coll: builder.Three(),
			args: linearIntArgs{reducer: func(acc int, _ int, current int) int {
				return acc*10 + current
			}},
			want1: 35631,
			want2: nil,
		},
	}
}

func testReduceRev(t *testing.T, builder linearCollIntBuilder) {
	cases := getReduceRevCases(t, builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.ReduceRev(tt.args.reducer)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Reduce() = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Reduce() = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func getTailCases(builder linearCollIntBuilder) []linearTestCase {
	return []linearTestCase{
		{
			name:  "Tail() on empty collection",
			coll:  builder.Empty(),
			want1: 0,
			want2: false,
		},
		{
			name:  "Tail() on one-item collection",
			coll:  builder.One(),
			want1: 111,
			want2: true,
		},
		{
			name:  "Tail() on three-item collection",
			coll:  builder.Three(),
			want1: 333,
			want2: true,
		},
	}
}

func testTail(t *testing.T, builder linearCollIntBuilder) {
	cases := getTailCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.Tail()
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Tail() got1 = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Tail() got2 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func getTailOrDefaultCases(builder linearCollIntBuilder) []linearTestCase {
	return []linearTestCase{
		{
			name:  "TailOrDefault() on empty collection",
			coll:  builder.Empty(),
			args:  linearIntArgs{defaultValue: -1},
			want1: -1,
		},
		{
			name:  "TailOrDefault() on one-item collection",
			coll:  builder.One(),
			args:  linearIntArgs{defaultValue: -1},
			want1: 111,
		},
		{
			name:  "TailOrDefault() on three-item collection",
			coll:  builder.Three(),
			args:  linearIntArgs{defaultValue: -1},
			want1: 333,
		},
	}
}

func testTailOrDefault(t *testing.T, builder linearCollIntBuilder) {
	cases := getTailOrDefaultCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.TailOrDefault(tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("TailOrDefault() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}
