package coll

import (
	"reflect"
	"testing"
)

type comfySeqIntBuilder[C baseInternal[int]] struct {
}

func (lcb *comfySeqIntBuilder[C]) Empty() C {
	return lcb.make([]int(nil)).(C)
}

func (lcb *comfySeqIntBuilder[C]) One() C {
	return lcb.make([]int{111}).(C)
}

func (lcb *comfySeqIntBuilder[C]) Two() C {
	return lcb.make([]int{123, 234}).(C)
}

func (lcb *comfySeqIntBuilder[C]) Three() C {
	return lcb.make([]int{111, 222, 333}).(C)
}

func (lcb *comfySeqIntBuilder[C]) ThreeRev() C {
	return lcb.make([]int{333, 222, 111}).(C)
}

func (lcb *comfySeqIntBuilder[C]) SixWithDuplicates() C {
	return lcb.make([]int{111, 222, 333, 111, 222, 333}).(C)
}

func (lcb *comfySeqIntBuilder[C]) FromValues(values []any) C {
	c := lcb.make([]int{})
	for _, v := range values {
		c.Append(v.(int))
	}
	return c.(C)
}

func (lcb *comfySeqIntBuilder[C]) make(items []int) linearMutableInternal[int] {
	coll := &comfySeq[int]{
		s: items,
	}

	return coll
}

func (lcb *comfySeqIntBuilder[C]) extractRawValues(c C) any {
	return lcb.extractUnderlyingSlice(c)
}

func (lcb *comfySeqIntBuilder[C]) extractUnderlyingSlice(c C) any {
	return (any(c)).(*comfySeq[int]).s
}

func (lcb *comfySeqIntBuilder[C]) extractUnderlyingMap(_ C) any {
	return nil
}

func (lcb *comfySeqIntBuilder[C]) extractUnderlyingKp(_ C) any {
	return nil
}

func (lcb *comfySeqIntBuilder[C]) extractUnderlyingValsCount(_ C) any {
	return nil
}

type comfySeqPairBuilder[C Base[Pair[int, int]]] struct {
}

func (lcb *comfySeqPairBuilder[C]) SixWithDuplicates() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
		NewPair(3, 333),
		NewPair(4, 111),
		NewPair(5, 222),
		NewPair(6, 333),
	}).(C)
}

func (lcb *comfySeqPairBuilder[C]) make(items []Pair[int, int]) Base[Pair[int, int]] {
	coll := &comfySeq[Pair[int, int]]{
		s: items,
	}

	return coll
}

func TestNewSequence(t *testing.T) {
	t.Run("NewSequence[int]()", func(t *testing.T) {
		intSeq := NewSequence[int]()
		if intSeq == nil {
			t.Error("NewSequence[int]() returned nil")
		}
		if !reflect.DeepEqual(intSeq, &comfySeq[int]{s: make([]int, 0)}) {
			t.Error("NewSequence[int]() did not return a comfySeq[int]")
		}

		stringSeq := NewSequence[string]()
		if stringSeq == nil {
			t.Error("NewSequence[string]() returned nil")
		}
		if !reflect.DeepEqual(stringSeq, &comfySeq[string]{s: make([]string, 0)}) {
			t.Error("NewSequence[int]() did not return a comfySeq[int]")
		}
	})
}

func TestNewSequenceFrom(t *testing.T) {
	t.Run("NewSequenceFrom[int]()", func(t *testing.T) {
		intSlice := []int{1, 2, 3}
		intSeq := NewSequenceFrom(intSlice)
		if intSeq == nil {
			t.Error("NewSequence[int]() returned nil")
		}
		if !reflect.DeepEqual(intSeq, &comfySeq[int]{s: intSlice}) {
			t.Error("NewSequence[int]() did not return a comfySeq[int]")
		}

		stringSlice := []string{"a", "b", "c"}
		stringSeq := NewSequenceFrom[string](stringSlice)
		if stringSeq == nil {
			t.Error("NewSequence[string]() returned nil")
		}
		if !reflect.DeepEqual(stringSeq, &comfySeq[string]{s: stringSlice}) {
			t.Error("NewSequence[int]() did not return a comfySeq[int]")
		}
	})
}

func Test_comfySeq_Append(t *testing.T) {
	testAppendOne(t, &comfySeqIntBuilder[linearMutableInternal[int]]{})
	testAppendMany(t, &comfySeqIntBuilder[linearMutableInternal[int]]{})
}

func Test_comfySeq_AppendColl(t *testing.T) {
	testAppendColl(t, &comfySeqIntBuilder[linearMutableInternal[int]]{})
}

func Test_comfySeq_Apply(t *testing.T) {
	testApply(t, &comfySeqIntBuilder[mutableInternal[int]]{})
}

func Test_comfySeq_At(t *testing.T) {
	testAt(t, &comfySeqIntBuilder[indexedInternal[int]]{})
}

func Test_comfySeq_AtOrDefault(t *testing.T) {
	testAtOrDefault(t, &comfySeqIntBuilder[indexedInternal[int]]{})
}

func Test_comfySeq_Clear(t *testing.T) {
	testClear(t, &comfySeqIntBuilder[mutableInternal[int]]{})
}

func Test_comfySeq_Contains(t *testing.T) {
	testContains(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_Count(t *testing.T) {
	testCount(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_Each(t *testing.T) {
	testEach(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_EachRev(t *testing.T) {
	testEachRev(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_EachRevUntil(t *testing.T) {
	testEachRevUntil(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_EachUntil(t *testing.T) {
	testEachUntil(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_Find(t *testing.T) {
	testFind(t, &comfySeqIntBuilder[baseInternal[int]]{})
	testFindWithDupes(t, &comfySeqPairBuilder[baseInternal[Pair[int, int]]]{})
}

func Test_comfySeq_FindLast(t *testing.T) {
	testFindLast(t, &comfySeqIntBuilder[baseInternal[int]]{})
	testFindLastWithDupes(t, &comfySeqPairBuilder[baseInternal[Pair[int, int]]]{})
}

func Test_comfySeq_Fold(t *testing.T) {
	testFold(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_Head(t *testing.T) {
	testHead(t, &comfySeqIntBuilder[linearInternal[int]]{})
}

func Test_comfySeq_HeadOrDefault(t *testing.T) {
	testHeadOrDefault(t, &comfySeqIntBuilder[linearInternal[int]]{})
}

func Test_comfySeq_InsertAt(t *testing.T) {
	testInsertAt(t, &comfySeqIntBuilder[listInternal[int]]{})
}

func Test_comfySeq_IsEmpty(t *testing.T) {
	testIsEmpty(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_Len(t *testing.T) {
	testLen(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_Prepend(t *testing.T) {
	testPrependOne(t, &comfySeqIntBuilder[linearMutableInternal[int]]{})
	testPrependMany(t, &comfySeqIntBuilder[linearMutableInternal[int]]{})
}

func Test_comfySeq_Reduce(t *testing.T) {
	testReduce(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_RemoveAt(t *testing.T) {
	testRemoveAt(t, &comfySeqIntBuilder[indexedMutableInternal[int]]{})
}

func Test_comfySeq_RemoveMatching(t *testing.T) {
	testRemoveMatching(t, &comfySeqIntBuilder[mutableInternal[int]]{})
}

func Test_comfySeq_Reverse(t *testing.T) {
	testReverse(t, &comfySeqIntBuilder[linearMutableInternal[int]]{})
	testReverseTwice(t, &comfySeqIntBuilder[linearMutableInternal[int]]{})
}

func Test_comfySeq_Search(t *testing.T) {
	testSearch(t, &comfySeqIntBuilder[baseInternal[int]]{})
	testSearchPair(t, &comfySeqPairBuilder[baseInternal[Pair[int, int]]]{})
}

func Test_comfySeq_SearchRev(t *testing.T) {
	testSearchRev(t, &comfySeqIntBuilder[baseInternal[int]]{})
	testSearchRevPair(t, &comfySeqPairBuilder[baseInternal[Pair[int, int]]]{})
}

func Test_comfySeq_Sort(t *testing.T) {
	testSort(t, &comfySeqIntBuilder[indexedMutableInternal[int]]{})
}

func Test_comfySeq_Tail(t *testing.T) {
	testTail(t, &comfySeqIntBuilder[linearInternal[int]]{})
}

func Test_comfySeq_TailOrDefault(t *testing.T) {
	testTailOrDefault(t, &comfySeqIntBuilder[linearInternal[int]]{})
}

func Test_comfySeq_ToSlice(t *testing.T) {
	testToSlice(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_Values(t *testing.T) {
	testValues(t, &comfySeqIntBuilder[baseInternal[int]]{})
	testValuesBreak(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_copy(t *testing.T) {
	testCopy(t, &comfySeqIntBuilder[baseInternal[int]]{})
}

func Test_comfySeq_copy_pointer(t *testing.T) {
	c1 := &comfySeq[int]{s: []int{123, 234, 345}}
	c2 := c1.copy()

	t.Run("copy() creates a new instance", func(t *testing.T) {
		if c1 == c2 {
			t.Error("copy() did not create a new instance")
		}
	})

	t.Run("copy() creates a deep copy", func(t *testing.T) {
		c1.s[0] = 999
		for v := range c2.(*comfySeq[int]).s {
			if v == 999 {
				t.Error("copy() did not create a deep copy")
			}
		}
	})
}
