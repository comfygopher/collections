package coll

import (
	"iter"
	"maps"
	"slices"
)

type comfyMap[K comparable, V any] struct {
	s []Pair[K, V]
	m map[K]Pair[K, V]
}

// NewMap creates a new Map instance.
func NewMap[K comparable, V any]() Map[K, V] {
	return &comfyMap[K, V]{
		s: make([]Pair[K, V], 0),
		m: make(map[K]Pair[K, V]),
	}
}

// NewMapFrom creates a new Map instance and copies elemnts from given map.
func NewMapFrom[K comparable, V any](m map[K]V) Map[K, V] {
	cm := &comfyMap[K, V]{
		m: make(map[K]Pair[K, V]),
	}
	for k, v := range m {
		pair := NewPair(k, v)
		cm.m[k] = pair
		cm.s = append(cm.s, pair)
	}

	return cm
}

func (c *comfyMap[K, V]) At(i int) (Pair[K, V], bool) {
	if i < 0 || i >= len(c.s) {
		return NilPair[K, V](), true
	}

	return c.s[i], false
}

func (c *comfyMap[K, V]) AtOrDefault(i int, defaultValue Pair[K, V]) Pair[K, V] {
	if i < 0 || i >= len(c.s) {
		return defaultValue
	}

	return c.s[i]
}

func (c *comfyMap[K, V]) Clear() {
	c.m = make(map[K]Pair[K, V])
	c.s = make([]Pair[K, V], 0)
}

func (c *comfyMap[K, V]) Contains(predicate Predicate[Pair[K, V]]) bool {
	return comfyContains[Indexed[Pair[K, V]], Pair[K, V]](c, predicate)
}

func (c *comfyMap[K, V]) Count(predicate Predicate[Pair[K, V]]) int {
	panic("not implemented")
	//return comfyCount[Indexed[Pair[K, V]], V](c, predicate)
}

func (c *comfyMap[K, V]) Each(f Visitor[Pair[K, V]]) {
	for i, pair := range c.s {
		f(i, pair)
	}
}

func (c *comfyMap[K, V]) EachRev(f Visitor[Pair[K, V]]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		f(i, c.s[i])
	}
}

func (c *comfyMap[K, V]) EachRevUntil(f Predicate[Pair[K, V]]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if !f(i, c.s[i]) {
			return
		}
	}
}

func (c *comfyMap[K, V]) EachUntil(f Predicate[Pair[K, V]]) {
	for i, pair := range c.s {
		if !f(i, pair) {
			return
		}
	}
}

func (c *comfyMap[K, V]) Find(predicate Predicate[Pair[K, V]], defaultValue Pair[K, V]) Pair[K, V] {
	panic("not implemented")
	//return comfyFind[Indexed[V], V](c, predicate, defaultValue)
}

func (c *comfyMap[K, V]) FindLast(predicate Predicate[Pair[K, V]], defaultValue Pair[K, V]) Pair[K, V] {
	panic("not implemented")
	//return comfyFindLast[Indexed[V], V](c, predicate, defaultValue)
}

func (c *comfyMap[K, V]) Fold(reducer Reducer[Pair[K, V]], initial Pair[K, V]) Pair[K, V] {
	panic("not implemented")
	//return comfyFold(c, reducer, initial)
}

func (c *comfyMap[K, V]) Get(k K) (V, bool) {
	pair, ok := c.m[k]
	if !ok {
		var v V
		return v, false
	}

	return pair.Val(), true
}

func (c *comfyMap[K, V]) GetOrDefault(k K, defaultValue V) (V, bool) {
	pair, ok := c.m[k]
	if !ok {
		return defaultValue, false
	}

	return pair.Val(), true
}

func (c *comfyMap[K, V]) Head() (Pair[K, V], bool) {
	if len(c.s) == 0 {
		return NilPair[K, V](), false
	}

	return c.s[0], true
}

func (c *comfyMap[K, V]) HeadOrDefault(defaultValue Pair[K, V]) Pair[K, V] {
	if len(c.s) == 0 {
		return defaultValue
	}

	return c.s[0]
}

func (c *comfyMap[K, V]) IsEmpty() bool {
	return len(c.s) == 0
}

func (c *comfyMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for _, pair := range c.s {
			if !yield(pair.Key()) {
				break
			}
		}
	}
}

func (c *comfyMap[K, V]) KeysToSlice() []K {
	return slices.Collect(c.Keys())
}

func (c *comfyMap[K, V]) Len() int {
	return len(c.s)
}

func (c *comfyMap[K, V]) KeyValues() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, pair := range c.s {
			if !yield(pair.Key(), pair.Val()) {
				break
			}
		}
	}
}

func (c *comfyMap[K, V]) Reduce(reducer Reducer[Pair[K, V]]) (Pair[K, V], error) {
	panic("not implemented")
	//return comfyReduce(c, reducer)
}

func (c *comfyMap[K, V]) RemoveAt(idx int) error {
	if idx < 0 || idx >= len(c.s) {
		return ErrOutOfBounds
	}

	c.s = append(c.s[:idx], c.s[idx+1:]...)
	delete(c.m, c.s[idx].Key())

	return nil
}

func (c *comfyMap[K, V]) Reverse() {
	newS := make([]Pair[K, V], 0)
	newM := make(map[K]Pair[K, V])
	for i := len(c.s) - 1; i >= 0; i-- {
		newS = append(newS, c.s[i])
		newM[c.s[i].Key()] = c.s[i]
	}
	c.s = newS
	c.m = newM
}

func (c *comfyMap[K, V]) Search(predicate Predicate[Pair[K, V]]) (Pair[K, V], bool) {
	for i, pair := range c.s {
		if predicate(i, pair) {
			return pair, true
		}
	}

	return NilPair[K, V](), false
}

func (c *comfyMap[K, V]) SearchRev(predicate Predicate[Pair[K, V]]) (Pair[K, V], bool) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if predicate(i, c.s[i]) {
			return c.s[i], true
		}
	}

	return NilPair[K, V](), false
}

func (c *comfyMap[K, V]) Tail() (Pair[K, V], bool) {
	if len(c.s) == 0 {
		return NilPair[K, V](), false
	}

	return c.s[len(c.s)-1], true
}

func (c *comfyMap[K, V]) TailOrDefault(defaultValue Pair[K, V]) Pair[K, V] {
	if len(c.s) == 0 {
		return defaultValue
	}

	return c.s[len(c.s)-1]
}

func (c *comfyMap[K, V]) ToMap() map[K]V {
	return maps.Collect(c.KeyValues())
}

func (c *comfyMap[K, V]) ToSlice() []Pair[K, V] {
	return slices.Collect(c.Values())
}

func (c *comfyMap[K, V]) Values() iter.Seq[Pair[K, V]] {
	return func(yield func(Pair[K, V]) bool) {
		for _, pair := range c.s {
			if !yield(pair) {
				break
			}
		}
	}
}

// Mutable[V any] interface implementation:

func (c *comfyMap[K, V]) Apply(f Mapper[Pair[K, V]]) {
	for i, pair := range c.s {
		mapped := f(i, pair)
		c.s[i] = mapped
		c.m[pair.Key()] = mapped
	}
}

func (c *comfyMap[K, V]) RemoveMatching(predicate Predicate[Pair[K, V]]) {
	newS := make([]Pair[K, V], 0)
	newM := make(map[K]Pair[K, V])
	for i, pair := range c.s {
		if !predicate(i, pair) {
			newS = append(newS, pair)
			newM[pair.Key()] = pair
		}
	}

	c.s = newS
	c.m = newM
}

// Map[K comparable, V any] interface implementation:

func (c *comfyMap[K, V]) Has(k K) bool {
	_, ok := c.m[k]
	return ok
}

func (c *comfyMap[K, V]) SetAll(im map[K]V) {
	for k, v := range im {
		c.set(NewPair(k, v))
	}
}

func (c *comfyMap[K, V]) Sort(cmp func(a, b V) int) {
	slices.SortFunc(c.s, func(a, b Pair[K, V]) int {
		return cmp(a.Val(), b.Val())
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

// Private:

func (c *comfyMap[K, V]) set(pair Pair[K, V]) {
	if _, ok := c.m[pair.Key()]; !ok {
		c.s = append(c.s, pair)
		c.m[pair.Key()] = pair
		return
	}

	// TODO: remove iteration when structure contains key => position map
	for i, current := range c.s {
		if current.Key() == pair.Key() {
			c.m[pair.Key()] = pair
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
		if current.Key() == k {
			c.s = append(c.s[:i], c.s[i+1:]...)
			delete(c.m, k)
			return
		}
	}
}

//nolint:unused
func (c *comfyMap[K, V]) copy() Base[Pair[K, V]] {
	newCm := &comfyMap[K, V]{
		s: make([]Pair[K, V], 0),
		m: make(map[K]Pair[K, V]),
	}
	for _, pair := range c.s {
		newCm.set(pair)
	}

	return newCm
}
