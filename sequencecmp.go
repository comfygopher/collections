package coll

import (
	"cmp"
	"iter"
	"slices"
)

// NewCmpSequence creates a new CmpSequence instance.
func NewCmpSequence[V cmp.Ordered]() CmpSequence[V] {
	return &comfyCmpSeq[V]{
		s:  []V(nil),
		vc: newValuesCounter[V](),
	}
}

// NewCmpSequenceFrom creates a new CmpSequence instance from a slice.
func NewCmpSequenceFrom[V cmp.Ordered](s []V) CmpSequence[V] {
	sq := NewCmpSequence[V]()
	sq.Append(s...)
	return sq
}

type comfyCmpSeq[V cmp.Ordered] struct {
	s  []V
	vc *valuesCounter[V]
}

func (c *comfyCmpSeq[V]) Apply(f Mapper[V]) {
	for i, v := range c.s {
		c.vc.Decrement(v)
		c.s[i] = f(i, v)
		c.vc.Increment(c.s[i])
	}
}

func (c *comfyCmpSeq[V]) Append(v ...V) {
	for _, v := range v {
		c.s = append(c.s, v)
		c.vc.Increment(v)
	}
}

func (c *comfyCmpSeq[V]) AppendColl(coll Linear[V]) {
	for v := range coll.Values() {
		c.s = append(c.s, v)
		c.vc.Increment(v)
	}
}

func (c *comfyCmpSeq[V]) At(i int) (V, bool) {
	if i < 0 || i >= len(c.s) {
		var v V
		return v, false
	}
	return c.s[i], true
}

func (c *comfyCmpSeq[V]) AtOrDefault(i int, defaultValue V) V {
	if i < 0 || i >= len(c.s) {
		return defaultValue
	}
	return c.s[i]
}

func (c *comfyCmpSeq[V]) Clear() {
	c.s = []V(nil)
	c.vc = newValuesCounter[V]()
}

func (c *comfyCmpSeq[V]) Contains(predicate Predicate[V]) bool {
	return comfyContains(c, predicate)
}

func (c *comfyCmpSeq[V]) ContainsValue(v V) bool {
	return c.vc.Count(v) > 0
}

func (c *comfyCmpSeq[V]) Count(predicate Predicate[V]) int {
	var count int
	for i, v := range c.s {
		if predicate(i, v) {
			count++
		}
	}
	return count
}

func (c *comfyCmpSeq[V]) CountValues(v V) int {
	return c.vc.Count(v)
}

func (c *comfyCmpSeq[V]) Each(f Visitor[V]) {
	for i, v := range c.s {
		f(i, v)
	}
}

func (c *comfyCmpSeq[V]) EachRev(f Visitor[V]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		f(i, c.s[i])
	}
}

func (c *comfyCmpSeq[V]) EachRevUntil(f Predicate[V]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if !f(i, c.s[i]) {
			return
		}
	}
}

func (c *comfyCmpSeq[V]) EachUntil(f Predicate[V]) {
	for i, v := range c.s {
		if !f(i, v) {
			return
		}
	}
}

func (c *comfyCmpSeq[V]) Find(predicate Predicate[V], defaultValue V) V {
	for i, v := range c.s {
		if predicate(i, v) {
			return v
		}
	}
	return defaultValue
}

func (c *comfyCmpSeq[V]) FindLast(predicate Predicate[V], defaultValue V) V {
	for i := len(c.s) - 1; i >= 0; i-- {
		if predicate(i, c.s[i]) {
			return c.s[i]
		}
	}
	return defaultValue
}

func (c *comfyCmpSeq[V]) Fold(reducer Reducer[V], initial V) V {
	return comfyFoldSlice(c.s, reducer, initial)
}

func (c *comfyCmpSeq[V]) Head() (V, bool) {
	if len(c.s) == 0 {
		var v V
		return v, false
	}
	return c.s[0], true
}

func (c *comfyCmpSeq[V]) HeadOrDefault(defaultValue V) V {
	if len(c.s) == 0 {
		return defaultValue
	}
	return c.s[0]
}

func (c *comfyCmpSeq[V]) IndexOf(v V) (i int, found bool) {
	if c.vc.Count(v) == 0 {
		return -1, false
	}
	for i, current := range c.s {
		if current == v {
			return i, true
		}
	}
	panic("invalid internal state of comfyCmpSeq")
}

func (c *comfyCmpSeq[V]) InsertAt(i int, v V) error {
	if i < 0 || i > len(c.s) {
		return ErrOutOfBounds
	}
	c.s = slices.Insert(c.s, i, v)
	c.vc.Increment(v)
	return nil
}

func (c *comfyCmpSeq[V]) IsEmpty() bool {
	return len(c.s) == 0
}

func (c *comfyCmpSeq[V]) LastIndexOf(v V) (i int, found bool) {
	if c.vc.Count(v) == 0 {
		return -1, false
	}
	for i := len(c.s) - 1; i >= 0; i-- {
		if c.s[i] == v {
			return i, true
		}
	}
	panic("invalid internal state of comfyCmpSeq")
}

func (c *comfyCmpSeq[V]) Len() int {
	return len(c.s)
}

func (c *comfyCmpSeq[V]) Max() (V, error) {
	return comfyMax[CmpSequence[V], V](c)
}

func (c *comfyCmpSeq[V]) Min() (V, error) {
	return comfyMin[CmpSequence[V], V](c)
}

func (c *comfyCmpSeq[V]) Prepend(v ...V) {
	if len(v) == 0 {
		return
	}
	c.s = append(v, c.s...)
	for _, v := range v {
		c.vc.Increment(v)
	}
}

func (c *comfyCmpSeq[V]) Reduce(reducer Reducer[V]) (V, error) {
	return comfyReduceSlice(c.s, reducer)
}

func (c *comfyCmpSeq[V]) RemoveAt(i int) (removed V, err error) {
	if removed, c.s, err = sliceRemoveAt(c.s, i); err != nil {
		return removed, err
	}
	c.vc.Decrement(removed)

	return removed, nil
}

func (c *comfyCmpSeq[V]) RemoveMatching(predicate Predicate[V]) {
	newS := []V(nil)
	newVC := newValuesCounter[V]()
	for i, v := range c.s {
		if !predicate(i, v) {
			newS = append(newS, v)
			newVC.Increment(v)
		}
	}
	c.s = newS
	c.vc = newVC
}

func (c *comfyCmpSeq[V]) RemoveValues(v V) {
	newS := []V(nil)
	newVC := newValuesCounter[V]()
	for _, current := range c.s {
		if current != v {
			newS = append(newS, current)
			newVC.Increment(current)
		}
	}
	c.s = newS
	c.vc = newVC
}

func (c *comfyCmpSeq[V]) Reverse() {
	slices.Reverse(c.s)
}

func (c *comfyCmpSeq[V]) Search(predicate Predicate[V]) (V, bool) {
	for i, v := range c.s {
		if predicate(i, v) {
			return v, true
		}
	}
	var v V
	return v, false
}

func (c *comfyCmpSeq[V]) SearchRev(predicate Predicate[V]) (V, bool) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if predicate(i, c.s[i]) {
			return c.s[i], true
		}
	}
	var v V
	return v, false
}

func (c *comfyCmpSeq[V]) Sort(cmp func(a, b V) int) {
	slices.SortFunc(c.s, cmp)
}

func (c *comfyCmpSeq[V]) SortAsc() {
	c.Sort(func(a, b V) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})
}

func (c *comfyCmpSeq[V]) SortDesc() {
	c.Sort(func(a, b V) int {
		if a < b {
			return 1
		} else if a > b {
			return -1
		}
		return 0
	})
}

func (c *comfyCmpSeq[V]) Sum() V {
	return comfySum[CmpSequence[V], V](c)
}

func (c *comfyCmpSeq[V]) Tail() (V, bool) {
	if len(c.s) == 0 {
		var v V
		return v, false
	}
	return c.s[len(c.s)-1], true
}

func (c *comfyCmpSeq[V]) TailOrDefault(defaultValue V) V {
	if len(c.s) == 0 {
		return defaultValue
	}
	return c.s[len(c.s)-1]
}

func (c *comfyCmpSeq[V]) ToSlice() []V {
	return append([]V{}, c.s...)
}

func (c *comfyCmpSeq[V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range c.s {
			if !yield(v) {
				break
			}
		}
	}
}

//nolint:unused
func (c *comfyCmpSeq[V]) copy() Base[V] {
	ccl := &comfyCmpSeq[V]{
		s:  []V(nil),
		vc: newValuesCounter[V](),
	}
	ccl.Append(c.s...)
	return ccl
}
