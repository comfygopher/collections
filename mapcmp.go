package coll

//lint:file-ignore U1000

import (
	"cmp"
	"iter"
	"slices"
)

type comfyCmpMap[K comparable, V cmp.Ordered] struct {
	s         []Pair[K, V]
	m         map[K]Pair[K, V]
	kp        map[K]int
	valsCount map[V]int
}

// NewCmpMap creates a new CmpMap instance.
//func NewCmpMap[K comparable, V cmp.Cmp]() CmpMap[Pair[K, V], K, V] {
//	return &comfyCmpMap[K, V]{
//		m:         make(map[K]*kvPair[K, V]),
//		valsCount: make(map[V]int),
//	}
//}

// NewCmpMapFrom creates a new CmpMap instance from a map.
//func NewCmpMapFrom[K comparable, V cmp.Cmp](m map[K]V) CmpMap[K, V] {
//	cm := NewCmpMap[K, V]()
//	for k, v := range m {
//		cm.Set(k, v)
//	}
//
//	return cm
//}

func (c *comfyCmpMap[K, V]) Append(p ...Pair[K, V]) {
	comfyAppendMap(c, p...)
}

func (c *comfyCmpMap[K, V]) AppendColl(coll Linear[Pair[K, V]]) {
	c.Append(coll.ToSlice()...)
}

func (c *comfyCmpMap[K, V]) Apply(f Mapper[Pair[K, V]]) {
	newS := []Pair[K, V](nil)
	newM := make(map[K]Pair[K, V])
	newKP := make(map[K]int)
	newValsCount := make(map[V]int)

	for i, pair := range c.s {
		mapped := f(i, pair)
		c.s[i] = mapped
		c.m[pair.Key()] = mapped

		if pair.Val() == mapped.Val() {
			continue
		}
		c.valsCount[pair.Val()]--
		if _, ok := c.valsCount[mapped.Val()]; !ok {
			c.valsCount[pair.Val()] = 1
		} else {
			c.valsCount[pair.Val()]++
		}
	}

	c.s = newS
	c.m = newM
	c.kp = newKP
	c.valsCount = newValsCount
}

func (c *comfyCmpMap[K, V]) At(i int) (p Pair[K, V], found bool) {
	if i < 0 || i >= len(c.s) {
		return nil, false
	}
	return c.s[i], true
}

func (c *comfyCmpMap[K, V]) AtOrDefault(i int, defaultValue Pair[K, V]) Pair[K, V] {
	if i < 0 || i >= len(c.s) {
		return defaultValue
	}
	return c.s[i]
}

func (c *comfyCmpMap[K, V]) Clear() {
	c.s = []Pair[K, V](nil)
	c.m = make(map[K]Pair[K, V])
	c.kp = make(map[K]int)
	c.valsCount = make(map[V]int)
}

func (c *comfyCmpMap[K, V]) Contains(predicate Predicate[Pair[K, V]]) bool {
	return comfyContains[Indexed[Pair[K, V]], Pair[K, V]](c, predicate)
}

func (c *comfyCmpMap[K, V]) ContainsValue(v V) bool {
	panic("not implemented")
	//_, ok := c.valsCount[v]
	//return ok
}

func (c *comfyCmpMap[K, V]) Count(predicate Predicate[Pair[K, V]]) int {
	return comfyCount[Indexed[Pair[K, V]], Pair[K, V]](c, predicate)
}

func (c *comfyCmpMap[K, V]) CountValues(v V) int {
	panic("not implemented")
	//return c.valsCount[v]
}

func (c *comfyCmpMap[K, V]) Each(f Visitor[Pair[K, V]]) {
	for i, pair := range c.s {
		f(i, pair)
	}
}

func (c *comfyCmpMap[K, V]) EachRev(f Visitor[Pair[K, V]]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		f(i, c.s[i])
	}
}

func (c *comfyCmpMap[K, V]) EachRevUntil(f Predicate[Pair[K, V]]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if !f(i, c.s[i]) {
			return
		}
	}
}

func (c *comfyCmpMap[K, V]) EachUntil(f Predicate[Pair[K, V]]) {
	for i, pair := range c.s {
		if !f(i, pair) {
			return
		}
	}
}

func (c *comfyCmpMap[K, V]) Find(predicate Predicate[Pair[K, V]], defaultValue Pair[K, V]) Pair[K, V] {
	panic("not implemented")
	//return comfyFind[Indexed[Pair[K, V]]](c, predicate, defaultValue)
}

func (c *comfyCmpMap[K, V]) FindLast(predicate Predicate[Pair[K, V]], defaultValue Pair[K, V]) Pair[K, V] {
	panic("not implemented")
	//return comfyFindLast[Indexed[Pair[K, V]]](c, predicate, defaultValue)
}

func (c *comfyCmpMap[K, V]) Fold(reducer Reducer[Pair[K, V]], initial Pair[K, V]) Pair[K, V] {
	return comfyFold(c, reducer, initial)
}

func (c *comfyCmpMap[K, V]) Get(k K) (V, bool) {
	panic("not implemented")
}

func (c *comfyCmpMap[K, V]) GetOrDefault(k K, defaultValue V) V {
	panic("not implemented")
}

func (c *comfyCmpMap[K, V]) HasValue(v V) bool {
	return c.ContainsValue(v)
}

func (c *comfyCmpMap[K, V]) Head() (Pair[K, V], bool) {
	if len(c.s) == 0 {
		return nil, false
	}
	return c.s[0], true
}

func (c *comfyCmpMap[K, V]) HeadOrDefault(defaultValue Pair[K, V]) Pair[K, V] {
	if len(c.s) == 0 {
		return defaultValue
	}
	return c.s[0]
}

func (c *comfyCmpMap[K, V]) IndexOf(v V) (int, error) {
	panic("not implemented")
	// TODO: remove iteration when structure contains value => position map
	//for i, current := range c.s {
	//	if current.Val() == v {
	//		return i, nil
	//	}
	//}
	//
	//return -1, ErrValueNotFound
}

func (c *comfyCmpMap[K, V]) IsEmpty() bool {
	return len(c.s) == 0
}

func (c *comfyCmpMap[K, V]) Keys() iter.Seq[K] {
	panic("not implemented")
}

func (c *comfyCmpMap[K, V]) KeysToSlice() []K {
	panic("not implemented")
}

func (c *comfyCmpMap[K, V]) KeyValues() iter.Seq2[K, V] {
	panic("not implemented")
}

func (c *comfyCmpMap[K, V]) LastIndexOf(v V) (int, error) {
	panic("not implemented")
	// TODO: remove iteration when structure contains value => position map
	//for i := len(c.s) - 1; i >= 0; i-- {
	//	if c.s[i].Val() == v {
	//		return i, nil
	//	}
	//}
	//
	//return -1, ErrValueNotFound
}

func (c *comfyCmpMap[K, V]) Len() int {
	return len(c.s)
}

func (c *comfyCmpMap[K, V]) Prepend(p ...Pair[K, V]) {
	panic("not implemented")
}

func (c *comfyCmpMap[K, V]) Reduce(reducer Reducer[Pair[K, V]]) (Pair[K, V], error) {
	return comfyReduce(c, reducer)
}

func (c *comfyCmpMap[K, V]) RemoveAt(idx int) (removed Pair[K, V], err error) {
	panic("not implemented")
	//if idx < 0 || idx >= len(c.s) {
	//	return ErrOutOfBounds
	//}
	//
	//c.s = append(c.s[:idx], c.s[idx+1:]...)
	//delete(c.m, c.s[idx].Key())
	//
	//return nil
}

func (c *comfyCmpMap[K, V]) RemoveMatching(predicate Predicate[Pair[K, V]]) {
	panic("not implemented")
	//c.RemoveMatchingKV(func(idx int, _ K, value V) bool {
	//	return predicate(idx, value)
	//})
}

func (c *comfyCmpMap[K, V]) Reverse() {
	newS := make([]Pair[K, V], 0)
	newM := make(map[K]Pair[K, V])
	for i := len(c.s) - 1; i >= 0; i-- {
		newS = append(newS, c.s[i])
		newM[c.s[i].Key()] = c.s[i]
	}
	c.s = newS
	c.m = newM
}

func (c *comfyCmpMap[K, V]) Search(predicate Predicate[Pair[K, V]]) (Pair[K, V], bool) {
	for i, pair := range c.s {
		if predicate(i, pair) {
			return pair, true
		}
	}
	return nil, false
}

func (c *comfyCmpMap[K, V]) SearchRev(predicate Predicate[Pair[K, V]]) (Pair[K, V], bool) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if predicate(i, c.s[i]) {
			return c.s[i], true
		}
	}
	return nil, false
}

func (c *comfyCmpMap[K, V]) SetMany(s []Pair[K, V]) {
	panic("not implemented")
}

func (c *comfyCmpMap[K, V]) Tail() (Pair[K, V], bool) {
	if len(c.s) == 0 {
		return nil, false
	}
	return c.s[len(c.s)-1], true
}

func (c *comfyCmpMap[K, V]) TailOrDefault(defaultValue Pair[K, V]) Pair[K, V] {
	if len(c.s) == 0 {
		return defaultValue
	}
	return c.s[len(c.s)-1]
}

func (c *comfyCmpMap[K, V]) ToMap() map[K]V {
	panic("not implemented")
}

func (c *comfyCmpMap[K, V]) ToSlice() []Pair[K, V] {
	return slices.Collect(c.Values())
}

func (c *comfyCmpMap[K, V]) Values() iter.Seq[Pair[K, V]] {
	panic("not implemented")
	//return func(yield func(V) bool) {
	//	for _, pair := range c.s {
	//		if !yield(pair.Val()) {
	//			break
	//		}
	//	}
	//}
}

// Map[K comparable, V any] interface implementation:

func (c *comfyCmpMap[K, V]) Has(k K) bool {
	_, ok := c.m[k]
	return ok
}

// TODO
//func (c *comfyCmpMap[K, V]) ReduceKV(reducer KVReducer[K, V], initialKey K, initialValue V) (K, V) {
//	return comfyReduceKV(c, reducer, initialKey, initialValue)
//}

func (c *comfyCmpMap[K, V]) RemoveMatchingKV(predicate KVPredicate[K, V]) {
	newS := make([]Pair[K, V], 0)
	newM := make(map[K]Pair[K, V])
	for i, pair := range c.s {
		if !predicate(i, pair.Key(), pair.Val()) {
			newS = append(newS, pair)
			newM[pair.Key()] = pair
		}
	}

	c.s = newS
	c.m = newM
}

func (c *comfyCmpMap[K, V]) Set(k K, v V) {
	c.set(NewPair(k, v))
}

func (c *comfyCmpMap[K, V]) SetAll(im map[K]V) {
	for k, v := range im {
		c.set(NewPair(k, v))
	}
}

func (c *comfyCmpMap[K, V]) Sort(cmp PairComparator[K, V]) {
	panic("not implemented")
}

func (c *comfyCmpMap[K, V]) Remove(k K) {
	c.remove(k)
}

func (c *comfyCmpMap[K, V]) RemoveMany(keys []K) {
	if len(keys) == 0 {
		return
	}
	if len(keys) == 1 {
		c.remove(keys[0])
		return
	}

	newS := make([]Pair[K, V], 0)
	newM := make(map[K]Pair[K, V])

	for _, pair := range c.s {
		for _, k := range keys {
			if pair.Key() != k {
				newS = append(newS, pair)
				newM[pair.Key()] = pair
			}
		}
	}

	c.s = newS
	c.m = newM
}

// TODO
//func (c *comfyCmpMap[K, V]) Max() (V, error) {
//	return comfyMax[Cmp[V], V](c)
//}
//
//func (c *comfyCmpMap[K, V]) Min() (V, error) {
//	return comfyMin[Cmp[V], V](c)
//}
//
//func (c *comfyCmpMap[K, V]) Sum() V {
//	return comfySum[Cmp[V], V](c)
//}

// Private:

func (c *comfyCmpMap[K, V]) set(pair Pair[K, V]) {
	if _, ok := c.m[pair.Key()]; !ok {
		c.s = append(c.s, pair)
		c.m[pair.Key()] = pair
		c.valsCount[pair.Val()] = 1
		return
	}

	// TODO: remove iteration when structure contains key => position map
	for i, current := range c.s {
		if current.Key() == pair.Key() {
			c.m[pair.Key()] = pair
			c.s[i] = pair
			c.valsCount[pair.Val()]++
			return
		}
	}
}

func (c *comfyCmpMap[K, V]) prependAll(pairs []Pair[K, V]) {
	panic("not implemented")
}

func (c *comfyCmpMap[K, V]) remove(k K) {
	if _, ok := c.m[k]; !ok {
		return
	}

	// TODO: remove iteration when structure contains key => position map
	for i, current := range c.s {
		if current.Key() == k {
			c.s = append(c.s[:i], c.s[i+1:]...)
			delete(c.m, k)
			c.valsCount[current.Val()]--
			return
		}
	}
}

func (c *comfyCmpMap[K, V]) removeMany(keys []K) {
	panic("not implemented")
}

func (c *comfyCmpMap[K, V]) setMany(pairs []Pair[K, V]) {
	for _, pair := range pairs {
		c.set(pair)
	}
}

//nolint:unused
func (c *comfyCmpMap[K, V]) copy() mapInternal[K, V] {
	newCm := &comfyCmpMap[K, V]{
		s: make([]Pair[K, V], 0),
		m: make(map[K]Pair[K, V]),
	}
	for _, pair := range c.s {
		newCm.set(pair)
	}

	return newCm
}
