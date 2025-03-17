package coll

import (
	"iter"
	"slices"
)

type comfySeq[V any] struct {
	s []V
}

// NewSequence creates a new Sequence instance.
func NewSequence[V any]() Sequence[V] {
	return &comfySeq[V]{
		s: []V(nil),
	}
}

// NewSequenceFrom creates a new Sequence instance from a slice.
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

func (c *comfySeq[V]) AppendColl(coll Ordered[V]) {
	c.s = append(c.s, slices.Collect(coll.Values())...)
}

func (c *comfySeq[V]) Apply(f Mapper[V]) {
	for i, v := range c.s {
		c.s[i] = f(v)
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

func (c *comfySeq[V]) RemoveAt(i int) (removed V, err error) {
	if removed, c.s, err = sliceRemoveAt(c.s, i); err != nil {
		return removed, err
	}

	return removed, nil
}

func (c *comfySeq[V]) RemoveMatching(predicate Predicate[V]) (count int) {
	c.s, count = sliceRemoveMatching(c.s, predicate)
	return count
}

func (c *comfySeq[V]) Reverse() {
	slices.Reverse(c.s)
}

func (c *comfySeq[V]) Sort(cmp func(a, b V) int) {
	slices.SortFunc(c.s, cmp)
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

func (c *comfySeq[V]) ValuesRev() iter.Seq[V] {
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
func (c *comfySeq[V]) copy() baseInternal[V] {
	newCl := &comfySeq[V]{
		s: []V(nil),
	}
	newCl.s = append(newCl.s, c.s...)

	return newCl
}
