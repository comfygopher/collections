package coll

import (
	"reflect"
	"testing"
)

type comfyCmpMapIntBuilder[C any] struct {
}

func (lcb *comfyCmpMapIntBuilder[C]) Empty() C {
	return lcb.make([]Pair[int, int]{}).(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) One() C {

	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
	}).(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) Two() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
	}).(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) Three() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
		NewPair(3, 333),
	}).(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) ThreeRev() C {
	return lcb.make([]Pair[int, int]{
		NewPair(10, 333),
		NewPair(20, 222),
		NewPair(30, 111),
	}).(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) SixWithDuplicates() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
		NewPair(3, 333),
		NewPair(4, 111),
		NewPair(5, 222),
		NewPair(6, 333),
	}).(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) FromValues(values []any) C {
	c := lcb.make([]Pair[int, int]{})
	for _, v := range values {
		c.set(v.(Pair[int, int]))
	}
	return c.(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) extractRawValues(c C) any {
	s := lcb.extractUnderlyingSlice(c).([]Pair[int, int])
	if s == nil {
		return []int(nil)
	}
	flat := make([]int, 0, len(s))
	for _, pair := range s {
		flat = append(flat, pair.Val())
	}
	return flat
}

func (lcb *comfyCmpMapIntBuilder[C]) extractUnderlyingSlice(c C) any {
	return (any(c)).(*comfyCmpMap[int, int]).s
}

func (lcb *comfyCmpMapIntBuilder[C]) extractUnderlyingMap(c C) any {
	return (any(c)).(*comfyCmpMap[int, int]).m
}

func (lcb *comfyCmpMapIntBuilder[C]) extractUnderlyingKp(c C) any {
	return (any(c)).(*comfyCmpMap[int, int]).kp
}

func (lcb *comfyCmpMapIntBuilder[C]) extractUnderlyingValsCount(c C) any {
	vc := (any(c)).(*comfyCmpMap[int, int]).vc.counter
	if vc == nil {
		panic("Could not extract Values Counter from comfyCmpMap")
	}
	return vc
}

func (lcb *comfyCmpMapIntBuilder[C]) make(items []Pair[int, int]) mapInternal[int, int] {
	coll := &comfyCmpMap[int, int]{
		s:  []Pair[int, int](nil),
		m:  make(map[int]Pair[int, int]),
		kp: make(map[int]int),
		vc: newValuesCounter[int](),
	}

	for _, pair := range items {
		coll.set(pair)
	}

	return coll
}

func TestNewMapCmp(t *testing.T) {
	t.Run("NewCmpMap[int, int]()", func(t *testing.T) {
		intMap := NewCmpMap[int, int]()
		if intMap == nil {
			t.Error("NewCmpMap[int, int]() returned nil")
		}
		if !reflect.DeepEqual(intMap, &comfyCmpMap[int, int]{
			s:  []Pair[int, int](nil),
			m:  make(map[int]Pair[int, int]),
			kp: make(map[int]int),
			vc: newValuesCounter[int](),
		}) {
			t.Error("NewCmpMap[int, int]() did not return a comfyCmpMap[int, int]")
		}
	})
}

func TestNewMapCmpFrom(t *testing.T) {
	t.Run("NewCmpMapFrom[int, int]()", func(t *testing.T) {
		intMap := NewCmpMapFrom[int, int]([]Pair[int, int]{
			NewPair(1, 111),
			NewPair(2, 222),
			NewPair(3, 333),
			NewPair(4, 333),
		})
		if intMap == nil {
			t.Error("NewCmpMapFrom[int, int]() returned nil")
		}
		if !reflect.DeepEqual(intMap, &comfyCmpMap[int, int]{
			s: []Pair[int, int]{
				NewPair(1, 111),
				NewPair(2, 222),
				NewPair(3, 333),
				NewPair(4, 333),
			},
			m: map[int]Pair[int, int]{
				1: NewPair(1, 111),
				2: NewPair(2, 222),
				3: NewPair(3, 333),
				4: NewPair(4, 333),
			},
			kp: map[int]int{
				1: 0,
				2: 1,
				3: 2,
				4: 3,
			},
			vc: &valuesCounter[int]{
				counter: map[int]int{
					111: 1,
					222: 1,
					333: 2,
				},
			},
		}) {
			t.Error("NewCmpMapFrom[int, int]() did not return a comfyCmpMap[int, int]")
		}
	})
}

func Test_comfyCmpMap_Append(t *testing.T) {
	testMapAppend(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
	testMapAppendRef(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_AppendColl(t *testing.T) {
	testMapAppendColl(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_Apply(t *testing.T) {
	testMapApply(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_At(t *testing.T) {
	testMapAt(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_AtOrDefault(t *testing.T) {
	testMapAtOrDefault(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_Clear(t *testing.T) {
	testMapClear(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_ContainsValue(t *testing.T) {
	testContainsValue(t, &comfyCmpMapIntBuilder[cmpMapBaseInternal[int, int]]{})
}

func Test_comfyCmpMap_CountValues(t *testing.T) {
	testCountValues(t, &comfyCmpMapIntBuilder[cmpMapBaseInternal[int, int]]{})
}

func Test_comfyCmpMap_Get(t *testing.T) {
	testMapGet(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_GetOrDefault(t *testing.T) {
	testMapGetOrDefault(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_Has(t *testing.T) {
	testMapHas(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_HasValue(t *testing.T) {
	testHasValue(t, &comfyCmpMapIntBuilder[cmpMapBaseInternal[int, int]]{})
}

func Test_comfyCmpMap_IndexOf(t *testing.T) {
	testIndexOf(t, &comfyCmpMapIntBuilder[cmpMapBaseInternal[int, int]]{})
}

func Test_comfyCmpMap_IsEmpty(t *testing.T) {
	testMapIsEmpty(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_Keys(t *testing.T) {
	testMapKeys(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
	testMapKeysBreak(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_KeyValues(t *testing.T) {
	testMapKeyValuesBreak(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_LastIndexOf(t *testing.T) {
	testLastIndexOf(t, &comfyCmpMapIntBuilder[cmpMapBaseInternal[int, int]]{})
}

func Test_comfyCmpMap_Len(t *testing.T) {
	testMapLen(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_Prepend(t *testing.T) {
	testMapPrepend(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_Remove(t *testing.T) {
	testMapRemove(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_RemoveAt(t *testing.T) {
	testMapRemoveAt(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_RemoveMany(t *testing.T) {
	testMapRemoveMany(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_RemoveMatching(t *testing.T) {
	testMapRemoveMatching(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_RemoveValues(t *testing.T) {
	testRemoveValues(t, &comfyCmpMapIntBuilder[CmpMutable[int]]{})
}

func Test_comfyCmpMap_Reverse(t *testing.T) {
	testMapReverse(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_Set(t *testing.T) {
	testMapSet(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_SetMany(t *testing.T) {
	testMapSetMany(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_Sort(t *testing.T) {
	testMapSort(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_SortAsc(t *testing.T) {
	testSortAsc(t, &comfyCmpMapIntBuilder[cmpMutableInternal[int]]{})
}

func Test_comfyCmpMap_SortDesc(t *testing.T) {
	testSortDesc(t, &comfyCmpMapIntBuilder[cmpMutableInternal[int]]{})
}

func Test_comfyCmpMap_Values(t *testing.T) {
	testMapValues(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
	testMapValuesBreak(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
	testMapValuesRef(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_ValuesRev(t *testing.T) {
	testMapValuesRev(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
	testMapValuesRevBreak(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_copy(t *testing.T) {
	testMapCopy(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
	testMapCopyDontPreserveRef(t, &comfyCmpMapIntBuilder[mapInternal[int, int]]{})
}

func Test_comfyCmpMap_copy_pointer(t *testing.T) {
	c1 := &comfyCmpMap[int, int]{
		m: map[int]Pair[int, int]{
			1: NewPair(1, 111),
			2: NewPair(2, 222),
			3: NewPair(3, 333),
		},
		vc: &valuesCounter[int]{
			counter: map[int]int{
				111: 1,
				222: 1,
				333: 1,
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
		c1.m[1] = NewPair(1, 999)
		c2m := c2.(*comfyCmpMap[int, int]).m
		for _, pair := range c2m {
			if pair.Val() == 999 {
				t.Error("copy() did not create a deep copy")
			}
		}
	})
}
