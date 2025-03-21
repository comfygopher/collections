// nolint:,unused
package coll

import (
	"reflect"
	"testing"
)

type comfyCmpSeqIntBuilder[C any] struct {
}

func (lcb *comfyCmpSeqIntBuilder[C]) Empty() C {
	return lcb.make([]int(nil)).(C)
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

func (lcb *comfyCmpSeqIntBuilder[C]) FromValues(values []any) C {
	c := lcb.make([]int{})
	for _, v := range values {
		c.Append(v.(int))
	}
	return c.(C)
}

func (lcb *comfyCmpSeqIntBuilder[C]) extractRawValues(coll C) any {
	return lcb.extractUnderlyingSlice(coll)
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

func (lcb *comfyCmpSeqIntBuilder[C]) extractUnderlyingValsCount(coll C) any {
	vc := (any(coll)).(*comfyCmpSeq[int]).vc.counter
	if vc == nil {
		panic("Could not extract Values Counter from comfyCmpSeq")
	}
	return vc
}

func (lcb *comfyCmpSeqIntBuilder[C]) make(items []int) orderedMutableInternal[int] {
	coll := &comfyCmpSeq[int]{
		s:  items,
		vc: newValuesCounter[int](),
	}

	for _, v := range items {
		coll.vc.Increment(v)
	}

	return coll
}

func TestNewCmpSequence(t *testing.T) {
	intSeq := NewCmpSequence[int]()
	if intSeq == nil {
		t.Error("NewCmpSequence[int]() returned nil")
	}
	if !reflect.DeepEqual(intSeq, &comfyCmpSeq[int]{s: []int(nil), vc: newValuesCounter[int]()}) {
		t.Error("NewCmpSequence[int]() did not return a comfyCmpSeq[int]")
	}

	stringSeq := NewCmpSequence[string]()
	if stringSeq == nil {
		t.Error("NewCmpSequence[string]() returned nil")
	}
	if !reflect.DeepEqual(stringSeq, &comfyCmpSeq[string]{s: []string(nil), vc: newValuesCounter[string]()}) {
		t.Error("NewCmpSequence[int]() did not return a comfyCmpSeq[int]")
	}
}

func TestNewCmpSequenceFrom(t *testing.T) {
	intSlice := []int{1, 2, 3}
	intSeq := NewCmpSequenceFrom[int](intSlice)
	if intSeq == nil {
		t.Error("NewCmpSequence[int]() returned nil")
	}
	wantIntVC := &valuesCounter[int]{
		counter: map[int]int{
			1: 1,
			2: 1,
			3: 1,
		},
	}
	if !reflect.DeepEqual(intSeq, &comfyCmpSeq[int]{s: intSlice, vc: wantIntVC}) {
		t.Error("NewCmpSequence[int]() did not return a comfyCmpSeq[int]")
	}

	stringSlice := []string{"a", "b", "c"}
	stringSeq := NewCmpSequenceFrom[string](stringSlice)
	if stringSeq == nil {
		t.Error("NewCmpSequence[string]() returned nil")
	}
	wantStringVC := &valuesCounter[string]{
		counter: map[string]int{
			"a": 1,
			"b": 1,
			"c": 1,
		},
	}
	if !reflect.DeepEqual(stringSeq, &comfyCmpSeq[string]{s: stringSlice, vc: wantStringVC}) {
		t.Error("NewCmpSequence[int]() did not return a comfyCmpSeq[int]")
	}
}

func Test_comfyCmpSeq_Append_one(t *testing.T) {
	testAppendOne(t, &comfyCmpSeqIntBuilder[orderedMutableInternal[int]]{})
	testAppendMany(t, &comfyCmpSeqIntBuilder[orderedMutableInternal[int]]{})
}

func Test_comfyCmpSeq_AppendColl(t *testing.T) {
	testAppendColl(t, &comfyCmpSeqIntBuilder[orderedMutableInternal[int]]{})
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

func Test_comfyCmpSeq_ContainsValue(t *testing.T) {
	testContainsValue(t, &comfyCmpSeqIntBuilder[cmpBaseInternal[int, int]]{})
}

func Test_comfyCmpSeq_CountValues(t *testing.T) {
	testCountValues(t, &comfyCmpSeqIntBuilder[cmpBaseInternal[int, int]]{})
}

func Test_comfyCmpSeq_HasValue(t *testing.T) {
	testHasValue(t, &comfyCmpSeqIntBuilder[cmpBaseInternal[int, int]]{})
}

func Test_comfyCmpSeq_IndexOf(t *testing.T) {
	testIndexOf(t, &comfyCmpSeqIntBuilder[cmpBaseInternal[int, int]]{})
}

func Test_comfyCmpSeq_IndexOfInvalidState(t *testing.T) {
	t.Run("invalid internal state", func(t *testing.T) {
		coll := NewCmpSequenceFrom[int]([]int{1, 2, 3})
		coll.(*comfyCmpSeq[int]).vc.Increment(4)
		defer func() {
			r := recover()
			if r == nil {
				t.Error("IndexOf() did not panic")
			}
			if r != "invalid internal state of comfyCmpSeq" {
				t.Errorf("IndexOf() panicked with wrong error: %v", r)
			}
		}()
		coll.IndexOf(4)
	})
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

func Test_comfyCmpSeq_LastIndexOfInvalidState(t *testing.T) {
	t.Run("invalid internal state", func(t *testing.T) {
		coll := NewCmpSequenceFrom[int]([]int{1, 2, 3})
		coll.(*comfyCmpSeq[int]).vc.Increment(4)
		defer func() {
			r := recover()
			if r == nil {
				t.Error("LastIndexOf() did not panic")
			}
			if r != "invalid internal state of comfyCmpSeq" {
				t.Errorf("LastIndexOf() panicked with wrong error: %v", r)
			}
		}()
		coll.LastIndexOf(4)
	})
}

func Test_comfyCmpSeq_Len(t *testing.T) {
	testLen(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_Prepend(t *testing.T) {
	testPrependOne(t, &comfyCmpSeqIntBuilder[orderedMutableInternal[int]]{})
	testPrependMany(t, &comfyCmpSeqIntBuilder[orderedMutableInternal[int]]{})
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
	testReverse(t, &comfyCmpSeqIntBuilder[orderedMutableInternal[int]]{})
	testReverseTwice(t, &comfyCmpSeqIntBuilder[orderedMutableInternal[int]]{})
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

func Test_comfyCmpSeq_Values(t *testing.T) {
	testValues(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
	testValuesBreak(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_ValuesRev(t *testing.T) {
	testValuesRev(t, &comfyCmpSeqIntBuilder[orderedInternal[int]]{})
	testValuesRevBreak(t, &comfyCmpSeqIntBuilder[orderedInternal[int]]{})
}

func Test_comfyCmpSeq_copy(t *testing.T) {
	testCopy(t, &comfyCmpSeqIntBuilder[baseInternal[int]]{})
}

func Test_comfyCmpSeq_copy_pointer(t *testing.T) {
	c1 := &comfyCmpSeq[int]{
		s: []int{123, 234, 345},
		vc: &valuesCounter[int]{
			counter: map[int]int{
				123: 1,
				234: 1,
				345: 1,
			},
		},
	}

	c2 := c1.copy()

	t.Run("copy() creates a new instance", func(t *testing.T) {
		if c1 == c2 {
			t.Error("copy() did not create a new instance")
		}
	})

	t.Run("copy() creates a deep copy", func(t *testing.T) {
		c1.s[0] = 999
		c2s := c2.(*comfyCmpSeq[int]).s
		for _, v := range c2s {
			if v == 999 {
				t.Error("copy() did not create a deep copy")
			}
		}
	})
}
