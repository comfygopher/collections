package coll

import (
	"reflect"
	"testing"
)

type comfyMapIntBuilder[C mapInternal[int, int]] struct {
}

func (lcb *comfyMapIntBuilder[C]) Empty() C {
	return lcb.make([]Pair[int, int]{}).(C)
}

func (lcb *comfyMapIntBuilder[C]) One() C {

	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
	}).(C)
}

func (lcb *comfyMapIntBuilder[C]) Two() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
	}).(C)
}

func (lcb *comfyMapIntBuilder[C]) Three() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
		NewPair(3, 333),
	}).(C)
}

func (lcb *comfyMapIntBuilder[C]) SixWithDuplicates() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
		NewPair(3, 333),
		NewPair(4, 111),
		NewPair(5, 222),
		NewPair(6, 333),
	}).(C)
}

func (lcb *comfyMapIntBuilder[C]) extractUnderlyingSlice(c C) any {
	return (any(c)).(*comfyMap[int, int]).s
}

func (lcb *comfyMapIntBuilder[C]) extractUnderlyingMap(c C) any {
	return (any(c)).(*comfyMap[int, int]).m
}

func (lcb *comfyMapIntBuilder[C]) extractUnderlyingKp(c C) any {
	return (any(c)).(*comfyMap[int, int]).kp
}

func (lcb *comfyMapIntBuilder[C]) extractUnderlyingValsCount(_ C) any {
	return nil
}

func (lcb *comfyMapIntBuilder[C]) make(items []Pair[int, int]) mapInternal[int, int] {
	coll := &comfyMap[int, int]{
		s:  []Pair[int, int](nil),
		m:  make(map[int]Pair[int, int]),
		kp: make(map[int]int),
	}

	for _, pair := range items {
		coll.set(pair)
	}

	return coll
}

func TestNewMap(t *testing.T) {
	t.Run("NewMap[int, int]()", func(t *testing.T) {
		intMap := NewMap[int, int]()
		if intMap == nil {
			t.Error("NewMap[int, int]() returned nil")
		}
		if !reflect.DeepEqual(intMap, &comfyMap[int, int]{
			s:  []Pair[int, int](nil),
			m:  make(map[int]Pair[int, int]),
			kp: make(map[int]int),
		}) {
			t.Error("NewMap[int, int]() did not return a comfyMap[int, int]")
		}
	})
}

func TestNewMapFrom(t *testing.T) {
	t.Run("NewMapFrom[int, int]()", func(t *testing.T) {
		intMap := NewMapFrom[int, int]([]Pair[int, int]{
			NewPair(1, 111),
			NewPair(2, 222),
			NewPair(3, 333),
		})
		if intMap == nil {
			t.Error("NewMapFrom[int, int]() returned nil")
		}
		if !reflect.DeepEqual(intMap, &comfyMap[int, int]{
			s: []Pair[int, int]{
				NewPair(1, 111),
				NewPair(2, 222),
				NewPair(3, 333),
			},
			m: map[int]Pair[int, int]{
				1: NewPair(1, 111),
				2: NewPair(2, 222),
				3: NewPair(3, 333),
			},
			kp: map[int]int{
				1: 0,
				2: 1,
				3: 2,
			},
		}) {
			t.Error("NewMapFrom[int, int]() did not return a comfyMap[int, int]")
		}
	})
}

func Test_comfyMap_Append(t *testing.T) {
	testMapAppend(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
	testMapAppendRef(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_AppendColl(t *testing.T) {
	testMapAppendColl(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Apply(t *testing.T) {
	testMapApply(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_At(t *testing.T) {
	testMapAt(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_AtOrDefault(t *testing.T) {
	testMapAtOrDefault(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Clear(t *testing.T) {
	testMapClear(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Contains(t *testing.T) {
	testMapContains(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Count(t *testing.T) {
	testMapCount(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Each(t *testing.T) {
	testMapEach(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_EachRev(t *testing.T) {
	testMapEachRev(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_EachRevUntil(t *testing.T) {
	testMapEachRevUntil(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_EachUntil(t *testing.T) {
	testMapEachUntil(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Find(t *testing.T) {
	testMapFind(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_FindLast(t *testing.T) {
	testMapFindLast(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Fold(t *testing.T) {
	testMapFold(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Get(t *testing.T) {
	testMapGet(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_GetOrDefault(t *testing.T) {
	testMapGetOrDefault(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Has(t *testing.T) {
	testMapHas(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Head(t *testing.T) {
	testMapHead(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_HeadOrDefault(t *testing.T) {
	testMapHeadOrDefault(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_IsEmpty(t *testing.T) {
	testMapIsEmpty(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Keys(t *testing.T) {
	testMapKeys(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
	testMapKeysBreak(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_KeysToSlice(t *testing.T) {
	testMapKeysToSlice(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_KeyValues(t *testing.T) {
	testMapKeyValuesBreak(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Len(t *testing.T) {
	testMapLen(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Prepend(t *testing.T) {
	testMapPrepend(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Reduce(t *testing.T) {
	testMapReduce(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Remove(t *testing.T) {
	testMapRemove(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_RemoveAt(t *testing.T) {
	testMapRemoveAt(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_RemoveMany(t *testing.T) {
	testMapRemoveMany(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_RemoveMatching(t *testing.T) {
	testMapRemoveMatching(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Reverse(t *testing.T) {
	testMapReverse(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Search(t *testing.T) {
	testMapSearch(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_SearchRev(t *testing.T) {
	testMapSearchRev(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Set(t *testing.T) {
	testMapSet(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_SetMany(t *testing.T) {
	testMapSetMany(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Sort(t *testing.T) {
	testMapSort(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Tail(t *testing.T) {
	testMapTail(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_TailOrDefault(t *testing.T) {
	testMapTailOrDefault(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_Values(t *testing.T) {
	testMapValues(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
	testMapValuesBreak(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
	testMapValuesRef(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyMap_copy(t *testing.T) {
	testMapCopy(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
	testMapCopyDontPreserveRef(t, &comfyMapIntBuilder[mapInternal[int, int]]{})
}
