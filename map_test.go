package coll

import (
	"reflect"
	"testing"
)

type comfyMapIntBuilder[C Base[Pair[int, int]]] struct {
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

func (lcb *comfyMapIntBuilder[C]) make(items []Pair[int, int]) Base[Pair[int, int]] {
	coll := &comfyMap[int, int]{
		s:  items,
		m:  make(map[int]Pair[int, int]),
		kp: make(map[int]int),
	}

	for i, pair := range items {
		coll.m[pair.Key()] = pair
		coll.kp[pair.Key()] = i
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

func Test_comfyMap_Apply(t *testing.T) {
	testMapApply(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_At(t *testing.T) {
	testMapAt(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_AtOrDefault(t *testing.T) {
	testMapAtOrDefault(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Clear(t *testing.T) {
	testMapClear(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Contains(t *testing.T) {
	testMapContains(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Count(t *testing.T) {
	testMapCount(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Each(t *testing.T) {
	testMapEach(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_EachRev(t *testing.T) {
	testMapEachRev(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_EachRevUntil(t *testing.T) {
	testMapEachRevUntil(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_EachUntil(t *testing.T) {
	testMapEachUntil(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Find(t *testing.T) {
	testMapFind(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_FindLast(t *testing.T) {
	testMapFindLast(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Fold(t *testing.T) {
	testMapFold(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Get(t *testing.T) {
	testMapGet(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_GetOrDefault(t *testing.T) {
	testMapGetOrDefault(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Has(t *testing.T) {
	testMapHas(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Head(t *testing.T) {
	testMapHead(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_HeadOrDefault(t *testing.T) {
	testMapHeadOrDefault(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_IsEmpty(t *testing.T) {
	testMapIsEmpty(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Keys(t *testing.T) {
	testMapKeys(t, &comfyMapIntBuilder[Map[int, int]]{})
	testMapKeysBreak(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_KeysToSlice(t *testing.T) {
	testMapKeysToSlice(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_KeyValues(t *testing.T) {
	testMapKeyValuesBreak(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Len(t *testing.T) {
	testMapLen(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Reduce(t *testing.T) {
	testMapReduce(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Remove(t *testing.T) {
	testMapRemove(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_RemoveAt(t *testing.T) {
	testMapRemoveAt(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_RemoveMany(t *testing.T) {
	testMapRemoveMany(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_RemoveMatching(t *testing.T) {
	testMapRemoveMatching(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Reverse(t *testing.T) {
	testMapReverse(t, &comfyMapIntBuilder[Map[int, int]]{})
}

func Test_comfyMap_Set(t *testing.T) {
	testMapSet(t, &comfyMapIntBuilder[Map[int, int]]{})
}
