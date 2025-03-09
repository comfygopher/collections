package coll

import (
	"reflect"
	"testing"
)

type comfyCmpSeqIntBuilder[C any] struct {
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

func (lcb *comfyCmpSeqIntBuilder[C]) ThreeRev() C {
	return lcb.make([]int{333, 222, 111}).(C)
}

func (lcb *comfyCmpSeqIntBuilder[C]) SixWithDuplicates() C {
	return lcb.make([]int{111, 222, 333, 111, 222, 333}).(C)
}

func (lcb *comfyCmpSeqIntBuilder[C]) extractRawValues(coll C) any {
	s := lcb.extractUnderlyingSlice(coll).([]int)
	if s == nil {
		return nil
	}
	return s
}

func (lcb *comfyCmpSeqIntBuilder[C]) extractUnderlyingSlice(coll C) any {
	return (any(coll)).(*comfyCmpSeq[int]).s
}

func (lcb *comfyCmpSeqIntBuilder[C]) extractUnderlyingMap(_ C) any {
	return nil
}

func (lcb *comfyCmpSeqIntBuilder[C]) extractUnderlyingKp(_ C) any {
	return nil
}

func (lcb *comfyCmpSeqIntBuilder[C]) extractUnderlyingValsCount(_ C) any {
	return nil
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
	testAppendOne(t, &comfyCmpSeqIntBuilder[linearMutableInternal[int]]{})
	testAppendMany(t, &comfyCmpSeqIntBuilder[linearMutableInternal[int]]{})
}

func Test_comfyCmpSeq_AppendColl(t *testing.T) {
	testAppendColl(t, &comfyCmpSeqIntBuilder[linearMutableInternal[int]]{})
}

func Test_comfyCmpSeq_Apply(t *testing.T) {
	testApply(t, &comfyCmpSeqIntBuilder[mutableInternal[int]]{})
}

func Test_comfyCmpSeq_At(t *testing.T) {
	testAt(t, &comfyCmpSeqIntBuilder[indexedInternal[int]]{})
}

func Test_comfyCmpSeq_AtOrDefault(t *testing.T) {
	testAtOrDefault(t, &comfyCmpSeqIntBuilder[indexedInternal[int]]{})
}

func Test_comfyCmpSeq_Clear(t *testing.T) {
	testClear(t, &comfyCmpSeqIntBuilder[mutableInternal[int]]{})
}

func Test_comfyCmpSeq_Contains(t *testing.T) {
	testContains(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_ContainsValue(t *testing.T) {
	testContainsValue(t, &comfyCmpSeqIntBuilder[cmpBaseInternal[int, int]]{})
}

func Test_comfyCmpSeq_Count(t *testing.T) {
	testCount(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_CountValues(t *testing.T) {
	testCountValues(t, &comfyCmpSeqIntBuilder[cmpBaseInternal[int, int]]{})
}

func Test_comfyCmpSeq_Each(t *testing.T) {
	testEach(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_EachRev(t *testing.T) {
	testEachRev(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_EachRevUntil(t *testing.T) {
	testEachRevUntil(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_EachUntil(t *testing.T) {
	testEachUntil(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_Find(t *testing.T) {
	testFind(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_FindLast(t *testing.T) {
	testFindLast(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_Head(t *testing.T) {
	testHead(t, &comfyCmpSeqIntBuilder[linearInternal[int]]{})
}

func Test_comfyCmpSeq_HeadOrDefault(t *testing.T) {
	testHeadOrDefault(t, &comfyCmpSeqIntBuilder[linearInternal[int]]{})
}

func Test_comfyCmpSeq_IndexOf(t *testing.T) {
	testIndexOf(t, &comfyCmpSeqIntBuilder[cmpBaseInternal[int, int]]{})
}

func Test_comfyCmpSeq_InsertAt(t *testing.T) {
	testInsertAt(t, &comfyCmpSeqIntBuilder[listInternal[int]]{})
}

func Test_comfyCmpSeq_IsEmpty(t *testing.T) {
	testIsEmpty(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_LastIndexOf(t *testing.T) {
	testLastIndexOf(t, &comfyCmpSeqIntBuilder[cmpBaseInternal[int, int]]{})
}

func Test_comfyCmpSeq_Len(t *testing.T) {
	testLen(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_Max(t *testing.T) {
	testMax(t, &comfyCmpSeqIntBuilder[cmpBaseInternal[int, int]]{})
}

func Test_comfyCmpSeq_Min(t *testing.T) {
	testMin(t, &comfyCmpSeqIntBuilder[cmpBaseInternal[int, int]]{})
}

func Test_comfyCmpSeq_Prepend(t *testing.T) {
	testPrependOne(t, &comfyCmpSeqIntBuilder[linearMutableInternal[int]]{})
	testPrependMany(t, &comfyCmpSeqIntBuilder[linearMutableInternal[int]]{})
}

func Test_comfyCmpSeq_Reduce(t *testing.T) {
	testReduce(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_RemoveAt(t *testing.T) {
	testRemoveAt(t, &comfyCmpSeqIntBuilder[indexedMutableInternal[int]]{})
}

func Test_comfyCmpSeq_RemoveMatching(t *testing.T) {
	testRemoveMatching(t, &comfyCmpSeqIntBuilder[mutableInternal[int]]{})
}

func Test_comfyCmpSeq_RemoveValues(t *testing.T) {
	testRemoveValues(t, &comfyCmpSeqIntBuilder[CmpMutable[int]]{})
}

func Test_comfyCmpSeq_Reverse(t *testing.T) {
	testReverse(t, &comfyCmpSeqIntBuilder[linearMutableInternal[int]]{})
	testReverseTwice(t, &comfyCmpSeqIntBuilder[linearMutableInternal[int]]{})
}

func Test_comfyCmpSeq_Search(t *testing.T) {
	testSearch(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_SearchRev(t *testing.T) {
	testSearchRev(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_Sort(t *testing.T) {
	testSort(t, &comfyCmpSeqIntBuilder[indexedMutableInternal[int]]{})
}

func Test_comfyCmpSeq_SortAsc(t *testing.T) {
	testSortAsc(t, &comfyCmpSeqIntBuilder[cmpBaseMutableInternal[int, int]]{})
}

func Test_comfyCmpSeq_SortDesc(t *testing.T) {
	testSortDesc(t, &comfyCmpSeqIntBuilder[CmpMutable[int]]{})
}

func Test_comfyCmpSeq_Sum(t *testing.T) {
	testSum(t, &comfyCmpSeqIntBuilder[cmpBaseInternal[int, int]]{})
}

func Test_comfyCmpSeq_Tail(t *testing.T) {
	testTail(t, &comfyCmpSeqIntBuilder[linearInternal[int]]{})
}

func Test_comfyCmpSeq_TailOrDefault(t *testing.T) {
	testTailOrDefault(t, &comfyCmpSeqIntBuilder[linearInternal[int]]{})
}

func Test_comfyCmpSeq_ToSlice(t *testing.T) {
	testToSlice(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_Values(t *testing.T) {
	testValues(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
	testValuesBreak(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_copy(t *testing.T) {
	testCopy(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
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
