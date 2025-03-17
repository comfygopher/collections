package coll

import (
	"iter"
	"slices"
)

type comfyMap[K comparable, V any] struct {
	s  []Pair[K, V]
	m  map[K]Pair[K, V]
	kp map[K]int
}

// NewMap creates a new Map instance.
// Note that there is no NewMapFrom constructor, because it would create collection in random order.
func NewMap[K comparable, V any]() Map[K, V] {
	return &comfyMap[K, V]{
		s:  []Pair[K, V](nil),
		m:  make(map[K]Pair[K, V]),
		kp: make(map[K]int),
	}
}

// NewMapFrom creates a new Map instance from a slice of pairs.
func NewMapFrom[K comparable, V any](s []Pair[K, V]) Map[K, V] {
	cm := NewMap[K, V]()
	cm.SetMany(s)
	return cm
}

// Public functions:

func (c *comfyMap[K, V]) Append(p ...Pair[K, V]) {
	comfyAppendMap(c, p...)
}

func (c *comfyMap[K, V]) AppendColl(coll Ordered[Pair[K, V]]) {
	c.Append(slices.Collect(coll.Values())...)
}

func (c *comfyMap[K, V]) Apply(f Mapper[Pair[K, V]]) {
	newS := []Pair[K, V](nil)
	newM := make(map[K]Pair[K, V])
	newKP := make(map[K]int)

	idx := 0
	for _, pair := range c.s {
		mapped := f(pair)
		newS = append(newS, mapped)
		newM[mapped.Key()] = mapped
		newKP[mapped.Key()] = idx
		idx++
	}

	c.s = newS
	c.m = newM
	c.kp = newKP
}

func (c *comfyMap[K, V]) At(i int) (p Pair[K, V], found bool) {
	if i < 0 || i >= len(c.s) {
		return nil, false
	}
	return c.s[i], true
}

func (c *comfyMap[K, V]) AtOrDefault(i int, defaultValue Pair[K, V]) Pair[K, V] {
	if i < 0 || i >= len(c.s) {
		return defaultValue
	}
	return c.s[i]
}

func (c *comfyMap[K, V]) Clear() {
	c.s = []Pair[K, V](nil)
	c.m = make(map[K]Pair[K, V])
	c.kp = make(map[K]int)
}

func (c *comfyMap[K, V]) Get(k K) (V, bool) {
	pair, ok := c.m[k]
	if !ok {
		var v V
		return v, false
	}
	return pair.Val(), true
}

func (c *comfyMap[K, V]) GetOrDefault(k K, defaultValue V) V {
	pair, ok := c.m[k]
	if !ok {
		return defaultValue
	}
	return pair.Val()
}

func (c *comfyMap[K, V]) Has(k K) bool {
	_, ok := c.m[k]
	return ok
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

func (c *comfyMap[K, V]) KeyValues() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, pair := range c.s {
			if !yield(pair.Key(), pair.Val()) {
				break
			}
		}
	}
}

func (c *comfyMap[K, V]) Len() int {
	return len(c.s)
}

func (c *comfyMap[K, V]) Prepend(p ...Pair[K, V]) {
	c.prependAll(p)
}

func (c *comfyMap[K, V]) Remove(k K) {
	c.remove(k)
}

func (c *comfyMap[K, V]) RemoveAt(idx int) (removed Pair[K, V], err error) {
	if removed, c.s, err = sliceRemoveAt(c.s, idx); err != nil {
		return removed, err
	}
	delete(c.m, removed.Key())
	delete(c.kp, removed.Key())
	return removed, nil
}

func (c *comfyMap[K, V]) RemoveMany(keys []K) {
	c.removeMany(keys)
}

func (c *comfyMap[K, V]) RemoveMatching(predicate Predicate[Pair[K, V]]) {
	newS := []Pair[K, V](nil)
	newM := make(map[K]Pair[K, V])
	newKP := make(map[K]int)

	idx := 0
	for _, pair := range c.s {
		if !predicate(pair) {
			newS = append(newS, pair)
			newM[pair.Key()] = pair
			newKP[pair.Key()] = idx
			idx++
		}
	}

	c.s = newS
	c.m = newM
	c.kp = newKP
}

func (c *comfyMap[K, V]) Reverse() {
	newS := []Pair[K, V](nil)
	newKP := make(map[K]int)
	for i := len(c.s) - 1; i >= 0; i-- {
		newS = append(newS, c.s[i])
		newKP[c.s[i].Key()] = len(c.s) - i - 1
	}
	c.s = newS
	c.kp = newKP
}

func (c *comfyMap[K, V]) Set(k K, v V) {
	c.set(NewPair(k, v))
}

func (c *comfyMap[K, V]) SetMany(s []Pair[K, V]) {
	for _, pair := range s {
		c.set(pair)
	}
}

func (c *comfyMap[K, V]) Sort(compare PairComparator[K, V]) {
	c.s, c.kp = comfySortSliceAndKP(c.s, compare)
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

func (c *comfyMap[K, V]) ValuesRev() iter.Seq[Pair[K, V]] {
	return func(yield func(Pair[K, V]) bool) {
		for i := len(c.s) - 1; i >= 0; i-- {
			if !yield(c.s[i]) {
				break
			}
		}
	}
}

// Private functions:

//nolint:unused
func (c *comfyMap[K, V]) copy() baseInternal[Pair[K, V]] {
	newCm := &comfyMap[K, V]{
		s:  []Pair[K, V](nil),
		m:  make(map[K]Pair[K, V]),
		kp: make(map[K]int),
	}
	for i, pair := range c.s {
		p := pair.copy()
		newCm.s = append(newCm.s, p)
		newCm.m[pair.Key()] = p
		newCm.kp[pair.Key()] = i
	}

	return newCm
}

func (c *comfyMap[K, V]) set(pair Pair[K, V]) {
	pos, exists := c.kp[pair.Key()]
	if exists {
		c.s[pos] = pair
		c.m[pair.Key()] = pair
	} else {
		pos = len(c.s)
		c.s = append(c.s, pair)
		c.m[pair.Key()] = pair
		c.kp[pair.Key()] = pos
		return
	}
}

func (c *comfyMap[K, V]) prependAll(pairs []Pair[K, V]) {
	newS := []Pair[K, V](nil)
	newM := make(map[K]Pair[K, V])
	newKP := make(map[K]int)

	idx := 0
	for _, pair := range pairs {
		newS = append(newS, pair)
		newM[pair.Key()] = pair
		newKP[pair.Key()] = idx
		idx++
	}

	for _, pair := range c.s {
		if _, ok := newM[pair.Key()]; ok {
			continue
		}
		newS = append(newS, pair)
		newM[pair.Key()] = pair
		newKP[pair.Key()] = idx
		idx++
	}

	c.s = newS
	c.m = newM
	c.kp = newKP
}

func (c *comfyMap[K, V]) remove(k K) {
	pos, exists := c.kp[k]
	if !exists {
		return
	}

	removed, newSlice, _ := sliceRemoveAt(c.s, pos)

	newKP := make(map[K]int)
	for i, pair := range newSlice {
		newKP[pair.Key()] = i
	}

	c.s = newSlice
	delete(c.m, removed.Key())
	c.kp = newKP
}

func (c *comfyMap[K, V]) removeMany(keys []K) {
	if len(keys) == 0 {
		return
	}
	if len(keys) == 1 {
		c.remove(keys[0])
		return
	}

	newS := []Pair[K, V](nil)
	newM := make(map[K]Pair[K, V])
	newKP := make(map[K]int)

	keysToRemove := comfyMakeKeyPosMap(keys)

	idx := 0
	for _, pair := range c.s {
		if _, ok := keysToRemove[pair.Key()]; ok {
			continue
		}
		newS = append(newS, pair)
		newM[pair.Key()] = pair
		newKP[pair.Key()] = idx
		idx++
	}

	c.s = newS
	c.m = newM
	c.kp = newKP
}
