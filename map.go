package coll

import (
	"iter"
	"slices"
)

type comfyMap[K comparable, V any] struct {
	s []*kvPair[K, V]
	m map[K]*kvPair[K, V]
}

type kvPair[K comparable, V any] struct {
	k K
	v V
}

// NewMap creates a new Map instance.
//func NewMap[K comparable, V any]() Map[Pair[K, V], K, V] {
//	return &comfyMap[K, V]{
//		s: make([]*kvPair[K, V], 0),
//		m: make(map[K]*kvPair[K, V]),
//	}
//}

// NewMapFrom creates a new Map instance and copies elemnts from given map.
//func NewMapFrom[K comparable, V any](m map[K]V) Map[K, V] {
//	cm := &comfyMap[K, V]{
//		m: make(map[K]*kvPair[K, V]),
//	}
//	for k, v := range m {
//		pair := &kvPair[K, V]{k, v}
//		cm.m[k] = pair
//		cm.s = append(cm.s, pair)
//	}
//
//	return cm
//}

// Indexed[V any] interface implementation:

func (c *comfyMap[K, V]) At(i int) (V, bool) {
	if i < 0 || i >= len(c.s) {
		var v V
		return v, true
	}

	return c.s[i].v, false
}

func (c *comfyMap[K, V]) AtOrDefault(i int, defaultValue V) V {
	if i < 0 || i >= len(c.s) {
		return defaultValue
	}

	return c.s[i].v
}

func (c *comfyMap[K, V]) Clear() {
	c.m = make(map[K]*kvPair[K, V])
	c.s = make([]*kvPair[K, V], 0)
}

func (c *comfyMap[K, V]) Contains(predicate Predicate[V]) bool {
	return comfyContains[Indexed[V], V](c, predicate)
}

func (c *comfyMap[K, V]) Count(predicate Predicate[V]) int {
	return comfyCount[Indexed[V], V](c, predicate)
}

func (c *comfyMap[K, V]) Each(f Visitor[V]) {
	for i, pair := range c.s {
		f(i, pair.v)
	}
}

func (c *comfyMap[K, V]) EachRev(f Visitor[V]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		f(i, c.s[i].v)
	}
}

func (c *comfyMap[K, V]) EachRevUntil(f Predicate[V]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if !f(i, c.s[i].v) {
			return
		}
	}
}

func (c *comfyMap[K, V]) EachUntil(f Predicate[V]) {
	for i, pair := range c.s {
		if !f(i, pair.v) {
			return
		}
	}
}

func (c *comfyMap[K, V]) Find(predicate Predicate[V], defaultValue V) V {
	return comfyFind[Indexed[V], V](c, predicate, defaultValue)
}

func (c *comfyMap[K, V]) FindLast(predicate Predicate[V], defaultValue V) V {
	return comfyFindLast[Indexed[V], V](c, predicate, defaultValue)
}

func (c *comfyMap[K, V]) Fold(reducer Reducer[V], initial V) V {
	return comfyFold(c, reducer, initial)
}

func (c *comfyMap[K, V]) Head() (V, bool) {
	if len(c.s) == 0 {
		var v V
		return v, false
	}

	return c.s[0].v, true
}

func (c *comfyMap[K, V]) HeadOrDefault(defaultValue V) V {
	if len(c.s) == 0 {
		return defaultValue
	}

	return c.s[0].v
}

func (c *comfyMap[K, V]) IsEmpty() bool {
	return len(c.s) == 0
}

func (c *comfyMap[K, V]) Len() int {
	return len(c.s)
}

func (c *comfyMap[K, V]) Reduce(reducer Reducer[V]) (V, error) {
	return comfyReduce(c, reducer)
}

func (c *comfyMap[K, V]) RemoveAt(idx int) error {
	if idx < 0 || idx >= len(c.s) {
		return ErrOutOfBounds
	}

	c.s = append(c.s[:idx], c.s[idx+1:]...)
	delete(c.m, c.s[idx].k)

	return nil
}

func (c *comfyMap[K, V]) Reverse() {
	newS := make([]*kvPair[K, V], 0)
	newM := make(map[K]*kvPair[K, V])
	for i := len(c.s) - 1; i >= 0; i-- {
		newS = append(newS, c.s[i])
		newM[c.s[i].k] = c.s[i]
	}
	c.s = newS
	c.m = newM
}

func (c *comfyMap[K, V]) Search(predicate Predicate[V]) (V, bool) {
	for i, pair := range c.s {
		if predicate(i, pair.v) {
			return pair.v, true
		}
	}

	var v V
	return v, false
}

func (c *comfyMap[K, V]) SearchRev(predicate Predicate[V]) (V, bool) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if predicate(i, c.s[i].v) {
			return c.s[i].v, true
		}
	}

	var v V
	return v, false
}

func (c *comfyMap[K, V]) Tail() (V, bool) {
	if len(c.s) == 0 {
		var v V
		return v, false
	}

	return c.s[len(c.s)-1].v, true
}

func (c *comfyMap[K, V]) TailOrDefault(defaultValue V) V {
	if len(c.s) == 0 {
		return defaultValue
	}

	return c.s[len(c.s)-1].v
}

func (c *comfyMap[K, V]) ToSlice() []V {
	return slices.Collect(c.Values())
}

func (c *comfyMap[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, pair := range c.s {
			if !yield(pair.v) {
				break
			}
		}
	}
}

// Mutable[V any] interface implementation:

func (c *comfyMap[K, V]) Apply(f Mapper[V]) {
	for i, pair := range c.s {
		mapped := f(i, pair.v)
		c.s[i].v = mapped
		c.m[pair.k].v = mapped
	}
}

func (c *comfyMap[K, V]) RemoveMatching(predicate Predicate[V]) {
	c.RemoveMatchingKV(func(idx int, _ K, value V) bool {
		return predicate(idx, value)
	})
}

// Map[K comparable, V any] interface implementation:

func (c *comfyMap[K, V]) Has(k K) bool {
	_, ok := c.m[k]
	return ok
}

// TODO
//func (c *comfyMap[K, V]) ContainsKV(predicate KVPredicate[K, V]) bool {
//	return comfyContainsKV[Map[K, V], K, V](c, predicate)
//}

func (c *comfyMap[K, V]) EachKV(f KVVistor[K, V]) {
	for i, pair := range c.s {
		f(i, pair.k, pair.v)
	}
}

func (c *comfyMap[K, V]) EachKVUntil(f KVPredicate[K, V]) {
	for i, pair := range c.s {
		if !f(i, pair.k, pair.v) {
			return
		}
	}
}

func (c *comfyMap[K, V]) FindKV(predicate KVPredicate[K, V], defaultValue V) V {
	for i, pair := range c.s {
		if predicate(i, pair.k, pair.v) {
			return pair.v
		}
	}

	return defaultValue
}

// TODO
//func (c *comfyMap[K, V]) ReduceKV(reducer KVReducer[K, V], initialKey K, initialValue V) (K, V) {
//	return comfyReduceKV(c, reducer, initialKey, initialValue)
//}

func (c *comfyMap[K, V]) RemoveMatchingKV(predicate KVPredicate[K, V]) {
	newS := make([]*kvPair[K, V], 0)
	newM := make(map[K]*kvPair[K, V])
	for i, pair := range c.s {
		if !predicate(i, pair.k, pair.v) {
			newS = append(newS, pair)
			newM[pair.k] = pair
		}
	}

	c.s = newS
	c.m = newM
}

func (c *comfyMap[K, V]) AppendKV(k K, v V) {
	c.set(&kvPair[K, V]{k, v})
}

func (c *comfyMap[K, V]) SetAll(im map[K]V) {
	for k, v := range im {
		c.set(&kvPair[K, V]{k, v})
	}
}

func (c *comfyMap[K, V]) Sort(cmp func(a, b V) int) {
	slices.SortFunc(c.s, func(a, b *kvPair[K, V]) int {
		return cmp(a.v, b.v)
	})
}

func (c *comfyMap[K, V]) Remove(k K) {
	c.remove(k)
}

func (c *comfyMap[K, V]) RemoveMany(keys []K) {
	if len(keys) == 0 {
		return
	}
	if len(keys) == 1 {
		c.remove(keys[0])
		return
	}

	newS := make([]*kvPair[K, V], 0)
	newM := make(map[K]*kvPair[K, V])

	for _, pair := range c.s {
		for _, k := range keys {
			if pair.k != k {
				newS = append(newS, pair)
				newM[pair.k] = pair
			}
		}
	}

	c.s = newS
	c.m = newM
}

// Private:

func (c *comfyMap[K, V]) set(pair *kvPair[K, V]) {
	if _, ok := c.m[pair.k]; !ok {
		c.s = append(c.s, pair)
		c.m[pair.k] = pair
		return
	}

	// TODO: remove iteration when structure contains key => position map
	for i, current := range c.s {
		if current.k == pair.k {
			c.m[pair.k] = pair
			c.s[i] = pair
			return
		}
	}
}

func (c *comfyMap[K, V]) remove(k K) {
	if _, ok := c.m[k]; !ok {
		return
	}

	// TODO: remove iteration when structure contains key => position map
	for i, current := range c.s {
		if current.k == k {
			c.s = append(c.s[:i], c.s[i+1:]...)
			delete(c.m, k)
			return
		}
	}
}

//nolint:unused
func (c *comfyMap[K, V]) copy() Base[V] {
	newCm := &comfyMap[K, V]{
		s: make([]*kvPair[K, V], 0),
		m: make(map[K]*kvPair[K, V]),
	}
	for _, pair := range c.s {
		newCm.set(pair)
	}

	return newCm
}
