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
		c.s[i] = f(v)
		c.vc.Increment(c.s[i])
	}
}

func (c *comfyCmpSeq[V]) Append(v ...V) {
	for _, v := range v {
		c.s = append(c.s, v)
		c.vc.Increment(v)
	}
}

func (c *comfyCmpSeq[V]) AppendColl(coll Ordered[V]) {
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

func (c *comfyCmpSeq[V]) ContainsValue(v V) bool {
	return c.vc.Count(v) > 0
}

func (c *comfyCmpSeq[V]) CountValues(v V) int {
	return c.vc.Count(v)
}

func (c *comfyCmpSeq[V]) HasValue(v V) bool {
	return c.ContainsValue(v)
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

func (c *comfyCmpSeq[V]) Prepend(v ...V) {
	if len(v) == 0 {
		return
	}
	c.s = append(v, c.s...)
	for _, v := range v {
		c.vc.Increment(v)
	}
}

func (c *comfyCmpSeq[V]) RemoveAt(i int) (removed V, err error) {
	if removed, c.s, err = sliceRemoveAt(c.s, i); err != nil {
		return removed, err
	}
	c.vc.Decrement(removed)

	return removed, nil
}

func (c *comfyCmpSeq[V]) RemoveMatching(predicate Predicate[V]) (count int) {
	newS := []V(nil)
	newVC := newValuesCounter[V]()
	for _, v := range c.s {
		if !predicate(v) {
			newS = append(newS, v)
			newVC.Increment(v)
		} else {
			count++
		}
	}
	c.s = newS
	c.vc = newVC
	return count
}

func (c *comfyCmpSeq[V]) RemoveValues(v ...V) (count int) {
	newS := []V(nil)
	newVC := newValuesCounter[V]()

	toRemove := newValuesCounter[V]()
	for _, v := range v {
		if c.vc.Count(v) > 0 {
			toRemove.Set(v, c.vc.Count(v))
		}
	}

	if toRemove.IsEmpty() {
		return 0
	}

	for _, current := range c.s {
		if toRemove.Count(current) == 0 {
			newS = append(newS, current)
			newVC.Increment(current)
		} else {
			count++
			toRemove.Decrement(current)
		}
	}

	c.s = newS
	c.vc = newVC

	return count
}

func (c *comfyCmpSeq[V]) Reverse() {
	slices.Reverse(c.s)
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

func (c *comfyCmpSeq[V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range c.s {
			if !yield(v) {
				break
			}
		}
	}
}

func (c *comfyCmpSeq[V]) ValuesRev() iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := len(c.s) - 1; i >= 0; i-- {
			if !yield(c.s[i]) {
				break
			}
		}
	}
}

// Private:

//nolint:unused
func (c *comfyCmpSeq[V]) copy() baseInternal[V] {
	ccl := &comfyCmpSeq[V]{
		s:  []V(nil),
		vc: newValuesCounter[V](),
	}
	ccl.Append(c.s...)
	return ccl
}
