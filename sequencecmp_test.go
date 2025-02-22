package coll

import (
	"cmp"
	"reflect"
	"testing"
)

type comfyCmpSeqIntBuilder[C Base[int]] struct {
}

func (lcb *comfyCmpSeqIntBuilder[C]) Empty() C {
	return lcb.make([]int{}).(C)
}

func (lcb *comfyCmpSeqIntBuilder[C]) One() C {
	return lcb.make([]int{111}).(C)
}

func (lcb *comfyCmpSeqIntBuilder[C]) Two() C {
	return lcb.make([]int{123, 234}).(C)
}

func (lcb *comfyCmpSeqIntBuilder[C]) Three() C {
	return lcb.make([]int{111, 222, 333}).(C)
}

func (lcb *comfyCmpSeqIntBuilder[C]) SixWithDuplicates() C {
	return lcb.make([]int{111, 222, 333, 111, 222, 333}).(C)
}

func (lcb *comfyCmpSeqIntBuilder[C]) make(items []int) Base[int] {
	coll := &comfyCmpSeq[int]{
		s: items,
	}

	return coll
}

func TestNewCmpSequence(t *testing.T) {
	intSeq := NewCmpSequence[int]()
	if intSeq == nil {
		t.Error("NewCmpSequence[int]() returned nil")
	}
	if !reflect.DeepEqual(intSeq, &comfyCmpSeq[int]{s: make([]int, 0)}) {
		t.Error("NewCmpSequence[int]() did not return a comfyCmpSeq[int]")
	}

	stringSeq := NewCmpSequence[string]()
	if stringSeq == nil {
		t.Error("NewCmpSequence[string]() returned nil")
	}
	if !reflect.DeepEqual(stringSeq, &comfyCmpSeq[string]{s: make([]string, 0)}) {
		t.Error("NewCmpSequence[int]() did not return a comfyCmpSeq[int]")
	}
}

func TestNewCmpSequenceFrom(t *testing.T) {
	intSlice := []int{1, 2, 3}
	intSeq := NewCmpSequenceFrom[int](intSlice)
	if intSeq == nil {
		t.Error("NewSequence[int]() returned nil")
	}
	if !reflect.DeepEqual(intSeq, &comfyCmpSeq[int]{s: intSlice}) {
		t.Error("NewSequence[int]() did not return a comfyCmpSeq[int]")
	}

	stringSlice := []string{"a", "b", "c"}
	stringSeq := NewCmpSequenceFrom[string](stringSlice)
	if stringSeq == nil {
		t.Error("NewSequence[string]() returned nil")
	}
	if !reflect.DeepEqual(stringSeq, &comfyCmpSeq[string]{s: stringSlice}) {
		t.Error("NewSequence[int]() did not return a comfyCmpSeq[int]")
	}
}

func Test_comfyCmpSeq_Append_one(t *testing.T) {
	testAppendOne(t, &comfyCmpSeqIntBuilder[LinearMutable[int]]{})
	testAppendMany(t, &comfyCmpSeqIntBuilder[LinearMutable[int]]{})
}

func Test_comfyCmpSeq_AppendColl(t *testing.T) {
	testAppendColl(t, &comfyCmpSeqIntBuilder[LinearMutable[int]]{})
}

func Test_comfyCmpSeq_Apply(t *testing.T) {
	testApply(t, &comfyCmpSeqIntBuilder[Mutable[int]]{})
}

func Test_comfyCmpSeq_At(t *testing.T) {
	testAt(t, &comfyCmpSeqIntBuilder[Indexed[int]]{})
}

func Test_comfyCmpSeq_AtOrDefault(t *testing.T) {
	testAtOrDefault(t, &comfyCmpSeqIntBuilder[Indexed[int]]{})
}

func Test_comfyCmpSeq_Clear(t *testing.T) {
	testClear(t, &comfyCmpSeqIntBuilder[Mutable[int]]{})
}

func Test_comfyCmpSeq_Contains(t *testing.T) {
	testContains(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_ContainsValue(t *testing.T) {
	type args[V cmp.Ordered] struct {
		val V
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Contains() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{val: 1},
			want: false,
		},
		{
			name: "Contains() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{val: 123},
			want: true,
		},
		{
			name: "Contains() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{val: 234},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ContainsValue(tt.args.val); got != tt.want {
				t.Errorf("ContainsValue() = %v, want1 %v", got, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_Count(t *testing.T) {
	testCount(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_CountValues(t *testing.T) {
	type args[V cmp.Ordered] struct {
		val V
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want int
	}
	tests := []testCase[int]{
		{
			name: "Count() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{val: 1},
			want: 0,
		},
		{
			name: "Count() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{val: 123},
			want: 1,
		},
		{
			name: "Count() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{val: 234},
			want: 1,
		},
		{
			name: "Count() on three item, not found",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{val: 1},
			want: 0,
		},
		{
			name: "Count() on three item, all found",
			c:    comfyCmpSeq[int]{s: []int{123, 123, 123}},
			args: args[int]{val: 123},
			want: 3,
		},
		{
			name: "Count() on three item, none found",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{val: -1},
			want: 0,
		},
		{
			name: "Count() on three item, 2 found",
			c:    comfyCmpSeq[int]{s: []int{123, 123, 345}},
			args: args[int]{val: 123},
			want: 2,
		},
		{
			name: "Count() on three item, some not mod 2 found",
			c:    comfyCmpSeq[int]{s: []int{1, 123, 123}},
			args: args[int]{val: 123},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.CountValues(tt.args.val); got != tt.want {
				t.Errorf("CountValues() = %v, want1 %v", got, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_Each(t *testing.T) {
	testEach(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_EachRev(t *testing.T) {
	testEachRev(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_EachRevUntil(t *testing.T) {
	testEachRevUntil(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_EachUntil(t *testing.T) {
	testEachUntil(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_Find(t *testing.T) {
	testFind(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_FindLast(t *testing.T) {
	testFindLast(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_Head(t *testing.T) {
	testHead(t, &comfyCmpSeqIntBuilder[Linear[int]]{})
}

func Test_comfyCmpSeq_HeadOrDefault(t *testing.T) {
	testHeadOrDefault(t, &comfyCmpSeqIntBuilder[Linear[int]]{})
}

func Test_comfyCmpSeq_InsertAt(t *testing.T) {
	testInsertAt(t, &comfyCmpSeqIntBuilder[List[int]]{})
}

func Test_comfyCmpSeq_IsEmpty(t *testing.T) {
	testIsEmpty(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_Len(t *testing.T) {
	testLen(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_Prepend(t *testing.T) {
	testPrependOne(t, &comfyCmpSeqIntBuilder[LinearMutable[int]]{})
	testPrependMany(t, &comfyCmpSeqIntBuilder[LinearMutable[int]]{})
}

func Test_comfyCmpSeq_Reduce(t *testing.T) {
	testReduce(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_RemoveAt(t *testing.T) {
	testRemoveAt(t, &comfyCmpSeqIntBuilder[IndexedMutable[int]]{})
}

func Test_comfyCmpSeq_RemoveMatching(t *testing.T) {
	testRemoveMatching(t, &comfyCmpSeqIntBuilder[Mutable[int]]{})
}

func Test_comfyCmpSeq_Reverse(t *testing.T) {
	testReverse(t, &comfyCmpSeqIntBuilder[LinearMutable[int]]{})
	testReverseTwice(t, &comfyCmpSeqIntBuilder[LinearMutable[int]]{})
}

func Test_comfyCmpSeq_Search(t *testing.T) {
	testSearch(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_SearchRev(t *testing.T) {
	testSearchRev(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_Sort(t *testing.T) {
	testSort(t, &comfyCmpSeqIntBuilder[IndexedMutable[int]]{})
}

func Test_comfyCmpSeq_Tail(t *testing.T) {
	testTail(t, &comfyCmpSeqIntBuilder[Linear[int]]{})
}

func Test_comfyCmpSeq_TailOrDefault(t *testing.T) {
	testTailOrDefault(t, &comfyCmpSeqIntBuilder[Linear[int]]{})
}

func Test_comfyCmpSeq_ToSlice(t *testing.T) {
	testToSlice(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_Values(t *testing.T) {
	testValues(t, &comfyCmpSeqIntBuilder[Base[int]]{})
	testValuesBreak(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_copy(t *testing.T) {
	testCopy(t, &comfyCmpSeqIntBuilder[Base[int]]{})
}

func Test_comfyCmpSeq_copy_pointer(t *testing.T) {
	c1 := &comfyCmpSeq[int]{s: []int{123, 234, 345}}
	c2 := c1.copy()

	t.Run("copy() creates a new instance", func(t *testing.T) {
		if c1 == c2 {
			t.Error("copy() did not create a new instance")
		}
	})

	t.Run("copy() creates a deep copy", func(t *testing.T) {
		c1.s[0] = 999
		c2s := c2.ToSlice()
		for v := range c2s {
			if v == 999 {
				t.Error("copy() did not create a deep copy")
			}
		}
	})
}

//func Test_comfyCmpSeq_IndexOf(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		v V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		testArgs    testArgs[V]
//		want    int
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.c.IndexOf(tt.testArgs.v)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("IndexOf() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("IndexOf() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_InsertAt(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		coll int
//		v V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		testArgs    testArgs[V]
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := tt.c.InsertAt(tt.testArgs.coll, tt.testArgs.v); (err != nil) != tt.wantErr {
//				t.Errorf("InsertAt() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_IsEmpty(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.IsEmpty(); got != tt.want {
//				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_LastIndexOf(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		v V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		testArgs    testArgs[V]
//		want    int
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.c.LastIndexOf(tt.testArgs.v)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("LastIndexOf() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("LastIndexOf() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Len(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want int
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.Len(); got != tt.want {
//				t.Errorf("Len() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Max(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		want    V
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.c.Max()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Max() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Max() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Min(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		want    V
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.c.Min()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Min() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Min() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Prepend(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		v []V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.Prepend(tt.testArgs.v...)
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Reduce(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		reducer Reducer
//		initial V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//		want V
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.Reduce(tt.testArgs.reducer, tt.testArgs.initial); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Reduce() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_RemoveAt(t *testing.T) {
//	type testArgs struct {
//		coll int
//	}
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		testArgs    testArgs
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := tt.c.RemoveAt(tt.testArgs.coll); (err != nil) != tt.wantErr {
//				t.Errorf("RemoveAt() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_RemoveMatching(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		predicate Predicate
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.RemoveMatching(tt.testArgs.predicate)
//		})
//	}
//}
//
//func Test_comfyCmpSeq_RemoveValues(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		v V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.RemoveValues(tt.testArgs.v)
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Reverse(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.Reverse()
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Search(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		predicate Predicate
//	}
//	type testCase[V cmp.Ordered] struct {
//		name  string
//		c     comfyCmpSeq[V]
//		testArgs  testArgs[V]
//		want  V
//		want1 bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := tt.c.Search(tt.testArgs.predicate)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Search() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("Search() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_SearchRev(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		predicate Predicate
//	}
//	type testCase[V cmp.Ordered] struct {
//		name  string
//		c     comfyCmpSeq[V]
//		testArgs  testArgs[V]
//		want  V
//		want1 bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := tt.c.SearchRev(tt.testArgs.predicate)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("SearchRev() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("SearchRev() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Sort(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		cmp func(a, b V) int
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.Sort(tt.testArgs.cmp)
//		})
//	}
//}
//
//func Test_comfyCmpSeq_SortAsc(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.SortAsc()
//		})
//	}
//}
//
//func Test_comfyCmpSeq_SortDesc(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.SortDesc()
//		})
//	}
//}
//
//func Test_comfyCmpSeq_SortMut(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		cmp func(a, b V) int
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.SortMut(tt.testArgs.cmp)
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Sum(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want V
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.Sum(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Sum() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Tail(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name  string
//		c     comfyCmpSeq[V]
//		want  V
//		want1 bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := tt.c.Tail()
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Tail() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("Tail() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_TailOrDefault(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		defaultValue V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//		want V
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.TailOrDefault(tt.testArgs.defaultValue); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("TailOrDefault() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_ToSlice(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want []V
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.ToSlice(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Values(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want iter.Seq
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.Values(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Values() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_copy(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want Indexed
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.copy(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("copy() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
