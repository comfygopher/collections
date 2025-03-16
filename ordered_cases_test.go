package coll

import (
	"reflect"
	"testing"
)

type orderedIntArgs = testArgs[orderedInternal[int], int]
type orderedTestCase = testCase[orderedInternal[int], int]
type orderedCollIntBuilder = testCollectionBuilder[orderedInternal[int]]

type orderedIntPairArgs = testArgs[orderedInternal[Pair[int, int]], Pair[int, int]]
type orderedPairTestCase = testCase[orderedInternal[Pair[int, int]], Pair[int, int]]
type orderedCollIntPairBuilder = testPairCollectionBuilder[orderedInternal[Pair[int, int]]]

func getEachRevCases(t *testing.T, builder orderedCollIntBuilder) []*orderedTestCase {

	// eachRevOnEmptyListCase:

	eachRevOnEmptyListCase := &orderedTestCase{
		name: "EachRev() on empty collection",
		coll: builder.Empty(),
	}
	eachRevOnEmptyListCase.args = orderedIntArgs{
		visit: func(i int, v int) {
			t.Error("EachRev() called on empty collection")
		},
	}

	// eachRevOnOneItemCase:

	eachRevOnOneItemCase := &orderedTestCase{
		name:  "EachRev() on one-item collection",
		coll:  builder.One(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0},
		want2: []int{111},
	}
	eachRevOnOneItemCase.args = orderedIntArgs{
		visit: func(i int, v int) {
			if i != 0 || v != 111 {
				t.Error("EachRev() called with wrong values")
			}
			eachRevOnOneItemCase.got1 = append(eachRevOnOneItemCase.got1.([]int), i)
			eachRevOnOneItemCase.got2 = append(eachRevOnOneItemCase.got2.([]int), v)
		},
	}

	// eachRevOnThreeCase:

	eachRevOnThreeCase := &orderedTestCase{
		name:  "EachRev() on three-item collection",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{2, 1, 0},
		want2: []int{333, 222, 111},
	}

	eachRevOnThreeCase.args = orderedIntArgs{
		visit: func(i int, v int) {
			if i < 0 || i > 2 || v < 111 || v > 333 {
				t.Error("EachRev() called with wrong values")
			}
			eachRevOnThreeCase.got1 = append(eachRevOnThreeCase.got1.([]int), i)
			eachRevOnThreeCase.got2 = append(eachRevOnThreeCase.got2.([]int), v)
		},
	}

	// put the cases together:

	return []*orderedTestCase{
		eachRevOnEmptyListCase,
		eachRevOnOneItemCase,
		eachRevOnThreeCase,
	}
}

func testEachRev(t *testing.T, builder orderedCollIntBuilder) {
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

func getEachRevUntilCases(t *testing.T, builder orderedCollIntBuilder) []*orderedTestCase {

	// eachRevUntilOnEmptyListCase:

	eachRevUntilOnEmptyListCase := &orderedTestCase{
		name: "EachRevUntil() on empty collection",
		coll: builder.Empty(),
	}

	eachRevUntilOnEmptyListCase.args = orderedIntArgs{
		predicate: func(i int, v int) bool {
			t.Error("EachRevUntil() called on empty collection")
			return true
		},
	}

	// eachRevUntilOnOneItemCase:

	eachRevUntilOnOneItemCase := &orderedTestCase{
		name:  "EachRevUntil() on one-item collection",
		coll:  builder.One(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{0},
		want2: []int{111},
	}

	eachRevUntilOnOneItemCase.args = orderedIntArgs{
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

	eachRevUntilFinishMiddleCase := &orderedTestCase{
		name:  "EachRevUntil() finish in middle",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{2, 1},
		want2: []int{333, 222},
	}

	eachRevUntilFinishMiddleCase.args = orderedIntArgs{
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

	eachRevUntilAllThreeCase := &orderedTestCase{
		name:  "EachRevUntil() on three-item collection",
		coll:  builder.Three(),
		got1:  []int{},
		got2:  []int{},
		want1: []int{2, 1, 0},
		want2: []int{333, 222, 111},
	}

	eachRevUntilAllThreeCase.args = orderedIntArgs{
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

	return []*orderedTestCase{
		eachRevUntilOnEmptyListCase,
		eachRevUntilOnOneItemCase,
		eachRevUntilFinishMiddleCase,
		eachRevUntilAllThreeCase,
	}
}

func testEachRevUntil(t *testing.T, builder orderedCollIntBuilder) {
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

func getFindLastCases(builder orderedCollIntBuilder) []*orderedTestCase {
	return []*orderedTestCase{
		{
			name: "FindLast() on empty collection",
			coll: builder.Empty(),
			args: orderedIntArgs{
				predicate:    func(i int, v int) bool { return v == 111 },
				defaultValue: -1,
			},
			want1: -1,
		},
		{
			name: "FindLast() on one-item collection",
			coll: builder.One(),
			args: orderedIntArgs{
				predicate:    func(i int, v int) bool { return v == 111 },
				defaultValue: -1,
			},
			want1: 111,
		},
		{
			name: "FindLast() on three-item collection",
			coll: builder.Three(),
			args: orderedIntArgs{
				predicate:    func(i int, v int) bool { return v == 222 },
				defaultValue: -1,
			},
			want1: 222,
		},
		{
			name: "FindLast() on three-item collection, not found",
			coll: builder.Three(),
			args: orderedIntArgs{
				predicate:    func(i int, v int) bool { return v == 999 },
				defaultValue: -1,
			},
			want1: -1,
		},
		{
			name: "FindLast() last one",
			coll: builder.Three(),
			args: orderedIntArgs{
				predicate:    func(i int, v int) bool { return true },
				defaultValue: -1,
			},
			want1: 333,
		},
		{
			name: "FindLast() on three-item collection, compare indexes and values",
			coll: builder.Three(),
			args: orderedIntArgs{predicate: func(i int, v int) bool {
				if i == 2 && v == 333 {
					return true
				}
				return false
			}},
			want1: 333,
		},
	}
}

func testFindLast(t *testing.T, builder orderedCollIntBuilder) {
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

func getFindLastCasesWithDupes(builder orderedCollIntPairBuilder) []*orderedPairTestCase {
	return []*orderedPairTestCase{
		{
			name: "FindLast() on six-item collection, first one",
			coll: builder.SixWithDuplicates(),
			args: orderedIntPairArgs{
				predicate:    func(i int, p Pair[int, int]) bool { return true },
				defaultValue: nil,
			},
			want1: NewPair(6, 333),
		},
		{
			name: "FindLast() on six-item collection, second one",
			coll: builder.SixWithDuplicates(),
			args: orderedIntPairArgs{
				predicate:    func(i int, p Pair[int, int]) bool { return p.Val() == 222 },
				defaultValue: nil,
			},
			want1: NewPair(5, 222),
		},
		{
			name: "FindLast() on six-item collection, not found",
			coll: builder.SixWithDuplicates(),
			args: orderedIntPairArgs{
				predicate:    func(i int, p Pair[int, int]) bool { return p.Val() == 999 },
				defaultValue: nil,
			},
			want1: nil,
		},
	}
}

func testFindLastWithDupes(t *testing.T, builder orderedCollIntPairBuilder) {
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

func getFoldRevCases(builder orderedCollIntBuilder) []*orderedTestCase {
	return []*orderedTestCase{
		{
			name: "FoldRev() on empty collection",
			coll: builder.Empty(),
			args: orderedIntArgs{
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
			args: orderedIntArgs{
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
			args: orderedIntArgs{
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
			args: orderedIntArgs{
				reducer: func(acc int, i int, current int) int {
					return acc*(i+1) + current
				},
				initial: 100,
			},
			want1: ((100*3+333)*2 + 222) + 111,
		},
	}
}

func testFoldRev(t *testing.T, builder orderedCollIntBuilder) {
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

func getHeadCases(builder orderedCollIntBuilder) []orderedTestCase {
	return []orderedTestCase{
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

func testHead(t *testing.T, builder orderedCollIntBuilder) {
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

func getHeadOrDefaultCases(builder orderedCollIntBuilder) []orderedTestCase {
	return []orderedTestCase{
		{
			name:  "HeadOrDefault() on empty collection",
			coll:  builder.Empty(),
			args:  orderedIntArgs{defaultValue: -1},
			want1: -1,
		},
		{
			name:  "HeadOrDefault() on one-item collection",
			coll:  builder.One(),
			args:  orderedIntArgs{defaultValue: -1},
			want1: 111,
		},
		{
			name:  "HeadOrDefault() on three-item collection",
			coll:  builder.Three(),
			args:  orderedIntArgs{defaultValue: -1},
			want1: 111,
		},
	}
}

func testHeadOrDefault(t *testing.T, builder orderedCollIntBuilder) {
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

func getSearchRevCases(builder orderedCollIntBuilder) []*orderedTestCase {

	searchRevOnEmptyCollectionCase := &orderedTestCase{
		name:  "SearchRev() on empty collection",
		coll:  builder.Empty(),
		want1: 0,
		want2: false,
		want3: nil,
		got3:  nil,
	}

	searchRevOnEmptyCollectionCase.args = orderedIntArgs{predicate: func(i int, v int) bool {
		if v == 111 {
			searchRevOnEmptyCollectionCase.got3 = i
			return true
		}
		return false
	}}

	searchRevOnOneItemCollectionCase := &orderedTestCase{
		name:  "SearchRev() on one-item collection",
		coll:  builder.One(),
		want1: 111,
		want2: true,
		want3: 0,
		got3:  nil,
	}

	searchRevOnOneItemCollectionCase.args = orderedIntArgs{predicate: func(i int, v int) bool {
		if v == 111 {
			searchRevOnOneItemCollectionCase.got3 = i
			return true
		}
		return false
	}}

	searchRevOnThreeItemCollectionCase := &orderedTestCase{
		name:  "SearchRev() on three-item collection",
		coll:  builder.Three(),
		want1: 222,
		want2: true,
		want3: 1,
		got3:  nil,
	}

	searchRevOnThreeItemCollectionCase.args = orderedIntArgs{predicate: func(i int, v int) bool {
		if v == 222 {
			searchRevOnThreeItemCollectionCase.got3 = i
			return true
		}
		return false
	}}

	searchRevOnOneItemCollectionCaseNotFound := &orderedTestCase{
		name:  "SearchRev() on one-item collection, not found",
		coll:  builder.One(),
		want1: 0,
		want2: false,
		want3: nil,
		got3:  nil,
	}

	searchRevOnOneItemCollectionCaseNotFound.args = orderedIntArgs{predicate: func(i int, v int) bool {
		if v == 999 {
			searchRevOnOneItemCollectionCaseNotFound.got3 = i
			return true
		}
		return false
	}}

	searchRevOnThreeItemCollectionCaseAllFound := &orderedTestCase{
		name:  "SearchRev() on three-item collection, all found",
		coll:  builder.Three(),
		want1: 333,
		want2: true,
		want3: 2,
		got3:  nil,
	}

	searchRevOnThreeItemCollectionCaseAllFound.args = orderedIntArgs{predicate: func(i int, v int) bool {
		searchRevOnThreeItemCollectionCaseAllFound.got3 = i
		return true
	}}

	searchRevOnSixItemCollectionCaseFirstFound := &orderedTestCase{
		name:  "SearchRev() on six-item collection, found first occurrence",
		coll:  builder.SixWithDuplicates(),
		want1: 111,
		want2: true,
		want3: 3,
		got3:  nil,
	}

	searchRevOnSixItemCollectionCaseFirstFound.args = orderedIntArgs{predicate: func(i int, v int) bool {
		if v == 111 && i == 3 {
			searchRevOnSixItemCollectionCaseFirstFound.got3 = i
			return true
		}
		return false
	}}

	searchRevOnSixItemCollectionCaseLastFound := &orderedTestCase{
		name:  "SearchRev() on six-item collection, found second occurrence",
		coll:  builder.SixWithDuplicates(),
		want1: 111,
		want2: true,
		want3: 0,
		got3:  nil,
	}

	searchRevOnSixItemCollectionCaseLastFound.args = orderedIntArgs{predicate: func(i int, v int) bool {
		if v == 111 && i == 0 {
			searchRevOnSixItemCollectionCaseLastFound.got3 = i
			return true
		}
		return false
	}}

	return []*orderedTestCase{
		searchRevOnEmptyCollectionCase,
		searchRevOnOneItemCollectionCase,
		searchRevOnThreeItemCollectionCase,
		searchRevOnOneItemCollectionCaseNotFound,
		searchRevOnThreeItemCollectionCaseAllFound,
		searchRevOnSixItemCollectionCaseFirstFound,
		searchRevOnSixItemCollectionCaseLastFound,
	}
}

func testSearchRev(t *testing.T, builder orderedCollIntBuilder) {
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

func getSearchRevPairCases(builder orderedCollIntPairBuilder) []*orderedPairTestCase {
	return []*orderedPairTestCase{
		{
			name: "SearchRev() on six-item collection, found first occurrence",
			coll: builder.SixWithDuplicates(),
			args: orderedIntPairArgs{predicate: func(i int, v Pair[int, int]) bool {
				return v.Val() == 111
			}},
			want1: NewPair(4, 111),
			want2: true,
		},
		{
			name: "SearchRev() on six-item collection, found first occurrence",
			coll: builder.SixWithDuplicates(),
			args: orderedIntPairArgs{predicate: func(i int, v Pair[int, int]) bool {
				return v.Val() == 222
			}},
			want1: NewPair(5, 222),
			want2: true,
		},
		{
			name: "SearchRev() on six-item collection, found first occurrence",
			coll: builder.SixWithDuplicates(),
			args: orderedIntPairArgs{predicate: func(i int, v Pair[int, int]) bool {
				return v.Val() == 333
			}},
			want1: NewPair(6, 333),
			want2: true,
		},
	}
}

func testSearchRevPair(t *testing.T, builder orderedCollIntPairBuilder) {
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

func getReduceRevCases(t *testing.T, builder orderedCollIntBuilder) []*orderedTestCase {
	return []*orderedTestCase{
		{
			name: "Reduce() on empty collection",
			coll: builder.Empty(),
			args: orderedIntArgs{reducer: func(acc int, _ int, current int) int {
				return acc*10 + current
			}},
			want1: 0,
			want2: ErrEmptyCollection,
		},
		{
			name: "Fold() on one-item collection",
			coll: builder.One(),
			args: orderedIntArgs{reducer: func(acc int, _ int, current int) int {
				return acc*10 + current
			}},
			want1: 111,
			want2: nil,
		},
		{
			name: "Fold() on three-item collection",
			coll: builder.Three(),
			args: orderedIntArgs{reducer: func(acc int, _ int, current int) int {
				return acc*10 + current
			}},
			want1: 35631,
			want2: nil,
		},
	}
}

func testReduceRev(t *testing.T, builder orderedCollIntBuilder) {
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

func getTailCases(builder orderedCollIntBuilder) []orderedTestCase {
	return []orderedTestCase{
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

func testTail(t *testing.T, builder orderedCollIntBuilder) {
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

func getTailOrDefaultCases(builder orderedCollIntBuilder) []orderedTestCase {
	return []orderedTestCase{
		{
			name:  "TailOrDefault() on empty collection",
			coll:  builder.Empty(),
			args:  orderedIntArgs{defaultValue: -1},
			want1: -1,
		},
		{
			name:  "TailOrDefault() on one-item collection",
			coll:  builder.One(),
			args:  orderedIntArgs{defaultValue: -1},
			want1: 111,
		},
		{
			name:  "TailOrDefault() on three-item collection",
			coll:  builder.Three(),
			args:  orderedIntArgs{defaultValue: -1},
			want1: 333,
		},
	}
}

func testTailOrDefault(t *testing.T, builder orderedCollIntBuilder) {
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
