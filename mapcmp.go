package coll

//lint:file-ignore U1000

import (
	"cmp"
	"iter"
	"slices"
)

// NewCmpMap creates a new CmpMap instance.
func NewCmpMap[K comparable, V cmp.Ordered]() CmpMap[K, V] {
	return &comfyCmpMap[K, V]{
		s:  []Pair[K, V](nil),
		m:  make(map[K]Pair[K, V]),
		kp: make(map[K]int),
		vc: newValuesCounter[V](),
	}
}

// NewCmpMapFrom creates a new CmpMap instance from a slice of pairs.
func NewCmpMapFrom[K comparable, V cmp.Ordered](s []Pair[K, V]) CmpMap[K, V] {
	cm := &comfyCmpMap[K, V]{
		s:  []Pair[K, V](nil),
		m:  make(map[K]Pair[K, V]),
		kp: make(map[K]int),
		vc: newValuesCounter[V](),
	}
	cm.setMany(s)
	return cm
}

type comfyCmpMap[K comparable, V cmp.Ordered] struct {
	s  []Pair[K, V]
	m  map[K]Pair[K, V]
	kp map[K]int
	vc *valuesCounter[V]
}

func (c *comfyCmpMap[K, V]) Append(p ...Pair[K, V]) {
	comfyAppendMap(c, p...)
}

func (c *comfyCmpMap[K, V]) AppendColl(coll Ordered[Pair[K, V]]) {
	c.Append(slices.Collect(coll.Values())...)
}

func (c *comfyCmpMap[K, V]) Apply(f Mapper[Pair[K, V]]) {
	newS := []Pair[K, V](nil)
	newM := make(map[K]Pair[K, V])
	newKP := make(map[K]int)
	newVC := newValuesCounter[V]()

	idx := 0
	for _, pair := range c.s {
		mapped := f(pair)
		newS = append(newS, mapped)
		newM[mapped.Key()] = mapped
		newKP[mapped.Key()] = idx
		newVC.Increment(mapped.Val())
		idx++
	}

	c.s = newS
	c.m = newM
	c.kp = newKP
	c.vc = newVC
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
	c.vc = newValuesCounter[V]()
}

func (c *comfyCmpMap[K, V]) ContainsValue(v V) bool {
	return c.vc.Count(v) > 0
}

func (c *comfyCmpMap[K, V]) CountValues(v V) int {
	return c.vc.Count(v)
}

func (c *comfyCmpMap[K, V]) Get(k K) (V, bool) {
	pair, ok := c.m[k]
	if !ok {
		var v V
		return v, false
	}
	return pair.Val(), true
}

func (c *comfyCmpMap[K, V]) GetOrDefault(k K, defaultValue V) V {
	pair, ok := c.m[k]
	if !ok {
		return defaultValue
	}
	return pair.Val()
}

func (c *comfyCmpMap[K, V]) Has(k K) bool {
	_, ok := c.m[k]
	return ok
}

func (c *comfyCmpMap[K, V]) HasValue(v V) bool {
	return c.ContainsValue(v)
}

func (c *comfyCmpMap[K, V]) IndexOf(v V) (pos int, found bool) {
	for i, current := range c.s {
		if current.Val() == v {
			return i, true
		}
	}
	return -1, false
}

func (c *comfyCmpMap[K, V]) IsEmpty() bool {
	return len(c.s) == 0
}

func (c *comfyCmpMap[K, V]) KeyValues() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, pair := range c.s {
			if !yield(pair.Key(), pair.Val()) {
				break
			}
		}
	}
}

func (c *comfyCmpMap[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for _, pair := range c.s {
			if !yield(pair.Key()) {
				break
			}
		}
	}
}

func (c *comfyCmpMap[K, V]) LastIndexOf(v V) (pos int, found bool) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if c.s[i].Val() == v {
			return i, true
		}
	}
	return -1, false
}

func (c *comfyCmpMap[K, V]) Len() int {
	return len(c.s)
}

func (c *comfyCmpMap[K, V]) Prepend(p ...Pair[K, V]) {
	c.prependAll(p)
}

func (c *comfyCmpMap[K, V]) Remove(k K) {
	c.remove(k)
}

func (c *comfyCmpMap[K, V]) RemoveAt(idx int) (removed Pair[K, V], err error) {
	if removed, c.s, err = sliceRemoveAt(c.s, idx); err != nil {
		return removed, err
	}
	delete(c.m, removed.Key())
	delete(c.kp, removed.Key())
	c.vc.Decrement(removed.Val())
	return removed, nil
}

func (c *comfyCmpMap[K, V]) RemoveMany(keys []K) {
	c.removeMany(keys)
}

func (c *comfyCmpMap[K, V]) RemoveMatching(predicate Predicate[Pair[K, V]]) (count int) {
	newS := []Pair[K, V](nil)
	newM := make(map[K]Pair[K, V])
	newKP := make(map[K]int)
	newVC := newValuesCounter[V]()

	idx := 0
	for _, pair := range c.s {
		if !predicate(pair) {
			newS = append(newS, pair)
			newM[pair.Key()] = pair
			newKP[pair.Key()] = idx
			newVC.Increment(pair.Val())
			idx++
		} else {
			count++
		}
	}

	c.s = newS
	c.m = newM
	c.kp = newKP
	c.vc = newVC
	return count
}

func (c *comfyCmpMap[K, V]) RemoveValues(v ...V) (count int) {
	toRemove := newValuesCounter[V]()
	for _, v := range v {
		if c.vc.Count(v) > 0 {
			toRemove.Set(v, c.vc.Count(v))
		}
	}

	if toRemove.IsEmpty() {
		return 0
	}

	return c.RemoveMatching(func(pair Pair[K, V]) bool {
		doRemove := toRemove.Count(pair.Val()) > 0
		if doRemove {
			toRemove.Decrement(pair.Val())
		}
		return doRemove
	})
}

func (c *comfyCmpMap[K, V]) Reverse() {
	newS := []Pair[K, V](nil)
	newKP := make(map[K]int)
	for i := len(c.s) - 1; i >= 0; i-- {
		newS = append(newS, c.s[i])
		newKP[c.s[i].Key()] = len(c.s) - i - 1
	}
	c.s = newS
	c.kp = newKP
}

func (c *comfyCmpMap[K, V]) Set(k K, v V) {
	c.set(NewPair(k, v))
}

func (c *comfyCmpMap[K, V]) SetMany(s []Pair[K, V]) {
	c.setMany(s)
}

func (c *comfyCmpMap[K, V]) Sort(compare PairComparator[K, V]) {
	c.s, c.kp = comfySortSliceAndKP(c.s, compare)
}

func (c *comfyCmpMap[K, V]) SortAsc() {
	c.Sort(func(a, b Pair[K, V]) int {
		if a.Val() < b.Val() {
			return -1
		} else if a.Val() > b.Val() {
			return 1
		}
		return 0
	})
}

func (c *comfyCmpMap[K, V]) SortDesc() {
	c.Sort(func(a, b Pair[K, V]) int {
		if a.Val() < b.Val() {
			return 1
		} else if a.Val() > b.Val() {
			return -1
		}
		return 0
	})
}

func (c *comfyCmpMap[K, V]) Values() iter.Seq[Pair[K, V]] {
	return func(yield func(Pair[K, V]) bool) {
		for _, pair := range c.s {
			if !yield(pair) {
				break
			}
		}
	}
}

func (c *comfyCmpMap[K, V]) ValuesRev() iter.Seq[Pair[K, V]] {
	return func(yield func(Pair[K, V]) bool) {
		for i := len(c.s) - 1; i >= 0; i-- {
			if !yield(c.s[i]) {
				break
			}
		}
	}
}

// Private:

func (c *comfyCmpMap[K, V]) copy() baseInternal[Pair[K, V]] {
	newCm := NewCmpMap[K, V]().(*comfyCmpMap[K, V])
	for _, pair := range c.s {
		newCm.set(pair.copy())
	}

	return newCm
}

func (c *comfyCmpMap[K, V]) set(pair Pair[K, V]) {
	pos, exists := c.kp[pair.Key()]
	if exists {
		c.vc.Decrement(c.s[pos].Val())
		c.s[pos] = pair
		c.m[pair.Key()] = pair
	} else {
		pos = len(c.s)
		c.s = append(c.s, pair)
		c.m[pair.Key()] = pair
		c.kp[pair.Key()] = pos
	}
	c.vc.Increment(pair.Val())
}

func (c *comfyCmpMap[K, V]) prependAll(pairs []Pair[K, V]) {
	newS := []Pair[K, V](nil)
	newM := make(map[K]Pair[K, V])
	newKP := make(map[K]int)
	newVC := newValuesCounter[V]()

	idx := 0
	for _, pair := range pairs {
		newS = append(newS, pair)
		newM[pair.Key()] = pair
		newKP[pair.Key()] = idx
		newVC.Increment(pair.Val())
		idx++
	}

	for _, pair := range c.s {
		if _, ok := newM[pair.Key()]; ok {
			continue
		}
		newS = append(newS, pair)
		newM[pair.Key()] = pair
		newKP[pair.Key()] = idx
		newVC.Increment(pair.Val())
		idx++
	}

	c.s = newS
	c.m = newM
	c.kp = newKP
	c.vc = newVC
}

func (c *comfyCmpMap[K, V]) remove(k K) {
	pos, exists := c.kp[k]
	if !exists {
		return
	}

	removed, newSlice, _ := sliceRemoveAt(c.s, pos)

	newKp := make(map[K]int)
	for i, pair := range newSlice {
		newKp[pair.Key()] = i
	}

	c.s = newSlice
	delete(c.m, removed.Key())
	c.kp = newKp
	c.vc.Decrement(removed.Val())
}

func (c *comfyCmpMap[K, V]) removeMany(keys []K) {
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
	newVC := newValuesCounter[V]()

	keysToRemove := comfyMakeKeyPosMap(keys)

	idx := 0
	for _, pair := range c.s {
		if _, ok := keysToRemove[pair.Key()]; ok {
			continue
		}
		newS = append(newS, pair)
		newM[pair.Key()] = pair
		newKP[pair.Key()] = idx
		newVC.Increment(pair.Val())
		idx++
	}

	c.s = newS
	c.m = newM
	c.kp = newKP
	c.vc = newVC
}

func (c *comfyCmpMap[K, V]) setMany(pairs []Pair[K, V]) {
	for _, pair := range pairs {
		c.set(pair)
	}
}
