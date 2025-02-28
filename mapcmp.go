package coll

//lint:file-ignore U1000

import (
	"cmp"
	"iter"
	"slices"
)

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

type comfyCmpMap[K comparable, V cmp.Ordered] struct {
	m         map[K]Pair[K, V]
	s         []Pair[K, V]
	kp        map[K]int
	valsCount map[V]int
}

func (c *comfyCmpMap[K, V]) Append(p ...Pair[K, V]) {
	keys := []K(nil)
	for _, pair := range p {
		keys = append(keys, pair.Key())
	}
	c.RemoveMany(keys)
	for _, pair := range p {
		c.set(pair)
	}
}

func (c *comfyCmpMap[K, V]) AppendColl(coll Linear[Pair[K, V]]) {
	c.Append(coll.ToSlice()...)
}

func (c *comfyCmpMap[K, V]) At(i int) (V, bool) {
	if i < 0 || i >= len(c.s) {
		var v V
		return v, true
	}

	return c.s[i].Val(), false
}

func (c *comfyCmpMap[K, V]) AtOrDefault(i int, defaultValue V) V {
	if i < 0 || i >= len(c.s) {
		return defaultValue
	}

	return c.s[i].Val()
}

func (c *comfyCmpMap[K, V]) Clear() {
	c.m = make(map[K]Pair[K, V])
	c.s = make([]Pair[K, V], 0)
}

func (c *comfyCmpMap[K, V]) Contains(predicate Predicate[V]) bool {
	return comfyContains[Indexed[V], V](c, predicate)
}

func (c *comfyCmpMap[K, V]) Count(predicate Predicate[V]) int {
	return comfyCount[Indexed[V], V](c, predicate)
}

func (c *comfyCmpMap[K, V]) Each(f Visitor[V]) {
	for i, pair := range c.s {
		f(i, pair.Val())
	}
}

func (c *comfyCmpMap[K, V]) EachRev(f Visitor[V]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		f(i, c.s[i].Val())
	}
}

func (c *comfyCmpMap[K, V]) EachRevUntil(f Predicate[V]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if !f(i, c.s[i].Val()) {
			return
		}
	}
}

func (c *comfyCmpMap[K, V]) EachUntil(f Predicate[V]) {
	for i, pair := range c.s {
		if !f(i, pair.Val()) {
			return
		}
	}
}

func (c *comfyCmpMap[K, V]) Find(predicate Predicate[V], defaultValue V) V {
	return comfyFind[Indexed[V], V](c, predicate, defaultValue)
}

func (c *comfyCmpMap[K, V]) FindLast(predicate Predicate[V], defaultValue V) V {
	return comfyFindLast[Indexed[V], V](c, predicate, defaultValue)
}

func (c *comfyCmpMap[K, V]) Fold(reducer Reducer[V], initial V) V {
	return comfyFold(c, reducer, initial)
}

func (c *comfyCmpMap[K, V]) Head() (V, bool) {
	if len(c.s) == 0 {
		var v V
		return v, false
	}

	return c.s[0].Val(), true
}

func (c *comfyCmpMap[K, V]) HeadOrDefault(defaultValue V) V {
	if len(c.s) == 0 {
		return defaultValue
	}

	return c.s[0].Val()
}

func (c *comfyCmpMap[K, V]) IsEmpty() bool {
	return len(c.s) == 0
}

func (c *comfyCmpMap[K, V]) Len() int {
	return len(c.s)
}

func (c *comfyCmpMap[K, V]) Reduce(reducer Reducer[V]) (V, error) {
	return comfyReduce(c, reducer)
}

func (c *comfyCmpMap[K, V]) RemoveAt(idx int) error {
	if idx < 0 || idx >= len(c.s) {
		return ErrOutOfBounds
	}

	c.s = append(c.s[:idx], c.s[idx+1:]...)
	delete(c.m, c.s[idx].Key())

	return nil
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

func (c *comfyCmpMap[K, V]) Search(predicate Predicate[V]) (V, bool) {
	for i, pair := range c.s {
		if predicate(i, pair.Val()) {
			return pair.Val(), true
		}
	}

	var v V
	return v, false
}

func (c *comfyCmpMap[K, V]) SearchRev(predicate Predicate[V]) (V, bool) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if predicate(i, c.s[i].Val()) {
			return c.s[i].Val(), true
		}
	}

	var v V
	return v, false
}

func (c *comfyCmpMap[K, V]) Tail() (V, bool) {
	if len(c.s) == 0 {
		var v V
		return v, false
	}

	return c.s[len(c.s)-1].Val(), true
}

func (c *comfyCmpMap[K, V]) TailOrDefault(defaultValue V) V {
	if len(c.s) == 0 {
		return defaultValue
	}

	return c.s[len(c.s)-1].Val()
}

func (c *comfyCmpMap[K, V]) ToSlice() []V {
	return slices.Collect(c.Values())
}

func (c *comfyCmpMap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, pair := range c.s {
			if !yield(pair.Val()) {
				break
			}
		}
	}
}

// Mutable[V any] interface implementation:

func (c *comfyCmpMap[K, V]) Apply(f Mapper[Pair[K, V]]) {
	for i, pair := range c.s {
		mapped := f(i, pair)
		c.s[i] = mapped
		c.m[pair.Key()] = mapped
	}
}

func (c *comfyCmpMap[K, V]) RemoveMatching(predicate Predicate[V]) {
	c.RemoveMatchingKV(func(idx int, _ K, value V) bool {
		return predicate(idx, value)
	})
}

// Map[K comparable, V any] interface implementation:

func (c *comfyCmpMap[K, V]) Has(k K) bool {
	_, ok := c.m[k]
	return ok
}

//func (c *comfyCmpMap[K, V]) ContainsKV(predicate KVPredicate[K, V]) bool {
//	return comfyContainsKV[Map[K, V], K, V](c, predicate)
//}

func (c *comfyCmpMap[K, V]) EachKV(f KVVistor[K, V]) {
	for i, pair := range c.s {
		f(i, pair.Key(), pair.Val())
	}
}

func (c *comfyCmpMap[K, V]) EachKVUntil(f KVPredicate[K, V]) {
	for i, pair := range c.s {
		if !f(i, pair.Key(), pair.Val()) {
			return
		}
	}
}

func (c *comfyCmpMap[K, V]) FindKV(predicate KVPredicate[K, V], defaultValue V) V {
	for i, pair := range c.s {
		if predicate(i, pair.Key(), pair.Val()) {
			return pair.Val()
		}
	}

	return defaultValue
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

func (c *comfyCmpMap[K, V]) Sort(cmp func(a, b V) int) {
	slices.SortFunc(c.s, func(a, b Pair[K, V]) int {
		return cmp(a.Val(), b.Val())
	})
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

// Cmp[V] interface implementation:

func (c *comfyCmpMap[K, V]) ContainsValue(v V) bool {
	_, ok := c.valsCount[v]
	return ok
}

func (c *comfyCmpMap[K, V]) CountValues(v V) int {
	return c.valsCount[v]
}

func (c *comfyCmpMap[K, V]) HasValue(v V) bool {
	return c.ContainsValue(v)
}

func (c *comfyCmpMap[K, V]) IndexOf(v V) (int, error) {
	// TODO: remove iteration when structure contains value => position map
	for i, current := range c.s {
		if current.Val() == v {
			return i, nil
		}
	}

	return -1, ErrValueNotFound
}

func (c *comfyCmpMap[K, V]) LastIndexOf(v V) (int, error) {
	// TODO: remove iteration when structure contains value => position map
	for i := len(c.s) - 1; i >= 0; i-- {
		if c.s[i].Val() == v {
			return i, nil
		}
	}

	return -1, ErrValueNotFound
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

//nolint:unused
func (c *comfyCmpMap[K, V]) copy() Base[V] {
	newCm := &comfyCmpMap[K, V]{
		s: make([]Pair[K, V], 0),
		m: make(map[K]Pair[K, V]),
	}
	for _, pair := range c.s {
		newCm.set(pair)
	}

	return newCm
}
