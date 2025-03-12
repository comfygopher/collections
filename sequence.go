package coll

import (
	"iter"
	"slices"
)

type comfySeq[V any] struct {
	s []V
}

// NewSequence creates a new LinearMutable instance.
func NewSequence[V any]() Sequence[V] {
	return &comfySeq[V]{
		s: make([]V, 0),
	}
}

// NewSequenceFrom creates a new LinearMutable instance from a slice.
func NewSequenceFrom[V any](l []V) Sequence[V] {
	return &comfySeq[V]{
		s: l,
	}
}

func (c *comfySeq[V]) Append(v ...V) {
	if len(v) == 0 {
		return
	}
	c.s = append(c.s, v...)
}

func (c *comfySeq[V]) AppendColl(coll Linear[V]) {
	c.s = append(c.s, coll.ToSlice()...)
}

func (c *comfySeq[V]) Apply(f Mapper[V]) {
	for i, v := range c.s {
		c.s[i] = f(i, v)
	}
}

func (c *comfySeq[V]) At(i int) (V, bool) {
	if i < 0 || i >= len(c.s) {
		var v V
		return v, false
	}

	return c.s[i], true
}

func (c *comfySeq[V]) AtOrDefault(i int, defaultValue V) V {
	if i < 0 || i >= len(c.s) {
		return defaultValue
	}

	return c.s[i]
}

func (c *comfySeq[V]) Clear() {
	c.s = []V(nil)
}

func (c *comfySeq[V]) Contains(predicate Predicate[V]) bool {
	return comfyContains[Indexed[V], V](c, predicate)
}

func (c *comfySeq[V]) Count(predicate Predicate[V]) int {
	var count int
	for i, v := range c.s {
		if predicate(i, v) {
			count++
		}
	}

	return count
}

func (c *comfySeq[V]) Each(f Visitor[V]) {
	for i, v := range c.s {
		f(i, v)
	}
}

func (c *comfySeq[V]) EachRev(f Visitor[V]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		f(i, c.s[i])
	}
}

func (c *comfySeq[V]) EachRevUntil(f Predicate[V]) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if !f(i, c.s[i]) {
			return
		}
	}
}

func (c *comfySeq[V]) EachUntil(f Predicate[V]) {
	for i, v := range c.s {
		if !f(i, v) {
			return
		}
	}
}

func (c *comfySeq[V]) Find(predicate Predicate[V], defaultValue V) V {
	for i, v := range c.s {
		if predicate(i, v) {
			return v
		}
	}

	return defaultValue
}

func (c *comfySeq[V]) FindLast(predicate Predicate[V], defaultValue V) V {
	for i := len(c.s) - 1; i >= 0; i-- {
		if predicate(i, c.s[i]) {
			return c.s[i]
		}
	}

	return defaultValue
}

func (c *comfySeq[V]) Fold(reducer Reducer[V], initial V) V {
	return comfyFoldSlice(c.s, reducer, initial)
}

func (c *comfySeq[V]) FoldRev(reducer Reducer[V], initial V) V {
	return comfyFoldSliceRev(c.s, reducer, initial)
}

func (c *comfySeq[V]) Head() (V, bool) {
	if len(c.s) == 0 {
		var v V
		return v, false
	}

	return c.s[0], true
}

func (c *comfySeq[V]) HeadOrDefault(defaultValue V) V {
	if len(c.s) == 0 {
		return defaultValue
	}

	return c.s[0]
}

func (c *comfySeq[V]) InsertAt(i int, v V) error {
	if i < 0 || i > len(c.s) {
		return ErrOutOfBounds
	}

	c.s = slices.Insert(c.s, i, v)
	return nil
}

func (c *comfySeq[V]) IsEmpty() bool {
	return len(c.s) == 0
}

func (c *comfySeq[V]) Len() int {
	return len(c.s)
}

func (c *comfySeq[V]) Prepend(v ...V) {
	if len(v) == 0 {
		return
	}
	c.s = append(v, c.s...)
}

func (c *comfySeq[V]) Reduce(reducer Reducer[V]) (V, error) {
	return comfyReduceSlice(c.s, reducer)
}

func (c *comfySeq[V]) RemoveAt(i int) (removed V, err error) {
	if removed, c.s, err = sliceRemoveAt(c.s, i); err != nil {
		return removed, err
	}

	return removed, nil
}

func (c *comfySeq[V]) RemoveMatching(predicate Predicate[V]) {
	c.s = sliceRemoveMatching(c.s, predicate)
}

func (c *comfySeq[V]) Reverse() {
	slices.Reverse(c.s)
}

func (c *comfySeq[V]) Search(predicate Predicate[V]) (V, bool) {
	for i, v := range c.s {
		if predicate(i, v) {
			return v, true
		}
	}

	var v V
	return v, false
}

func (c *comfySeq[V]) SearchRev(predicate Predicate[V]) (V, bool) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if predicate(i, c.s[i]) {
			return c.s[i], true
		}
	}

	var v V
	return v, false
}

func (c *comfySeq[V]) Sort(cmp func(a, b V) int) {
	slices.SortFunc(c.s, cmp)
}

func (c *comfySeq[V]) Tail() (V, bool) {
	if len(c.s) == 0 {
		var v V
		return v, false
	}

	return c.s[len(c.s)-1], true
}

func (c *comfySeq[V]) TailOrDefault(defaultValue V) V {
	if len(c.s) == 0 {
		return defaultValue
	}

	return c.s[len(c.s)-1]
}

func (c *comfySeq[V]) ToSlice() []V {
	return append([]V{}, c.s...)
}

func (c *comfySeq[V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range c.s {
			if !yield(v) {
				break
			}
		}
	}
}

// Private:

//nolint:unused
func (c *comfySeq[V]) copy() Base[V] {
	newCl := &comfySeq[V]{
		s: []V(nil),
	}
	newCl.s = append(newCl.s, c.s...)

	return newCl
}
