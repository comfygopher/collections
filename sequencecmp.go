package coll

import (
	"cmp"
	"iter"
	"slices"
)

// NewCmpSequence creates a new CmpSequence instance.
func NewCmpSequence[V cmp.Ordered]() CmpSequence[V] {
	return &comfyCmpSeq[V]{
		s: make([]V, 0),
	}
}

// NewCmpSequenceFrom creates a new CmpSequence instance from a slice.
func NewCmpSequenceFrom[V cmp.Ordered](l []V) CmpSequence[V] {
	return &comfyCmpSeq[V]{
		s: l,
	}
}

type comfyCmpSeq[V cmp.Ordered] struct {
	s []V
}

func (c *comfyCmpSeq[V]) Apply(f Mapper[V]) {
	for i, v := range c.s {
		c.s[i] = f(i, v)
	}
}

func (c *comfyCmpSeq[V]) Append(v ...V) {
	c.s = append(c.s, v...)
}

func (c *comfyCmpSeq[V]) AppendColl(coll Linear[V]) {
	c.s = append(c.s, coll.ToSlice()...)
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
	c.s = make([]V, 0)
}

func (c *comfyCmpSeq[V]) Contains(predicate Predicate[V]) bool {
	for i, v := range c.s {
		if predicate(i, v) {
			return true
		}
	}
	return false
}

func (c *comfyCmpSeq[V]) ContainsValue(v V) bool {
	for _, current := range c.s {
		if current == v {
			return true
		}
	}
	return false
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
	count := 0
	for _, current := range c.s {
		if current == v {
			count++
		}
	}
	return count
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
	return comfyFold(c, reducer, initial)
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
	for i, current := range c.s {
		if current == v {
			return i, true
		}
	}
	return -1, false
}

func (c *comfyCmpSeq[V]) InsertAt(i int, v V) error {
	if i < 0 || i > len(c.s) {
		return ErrOutOfBounds
	}
	c.s = slices.Insert(c.s, i, v)
	return nil
}

func (c *comfyCmpSeq[V]) IsEmpty() bool {
	return len(c.s) == 0
}

func (c *comfyCmpSeq[V]) LastIndexOf(v V) (i int, found bool) {
	for i := len(c.s) - 1; i >= 0; i-- {
		if c.s[i] == v {
			return i, true
		}
	}
	return -1, false
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
	c.s = append(v, c.s...)
}

func (c *comfyCmpSeq[V]) Reduce(reducer Reducer[V]) (V, error) {
	return comfyReduce(c, reducer)
}

func (c *comfyCmpSeq[V]) RemoveAt(i int) (removed V, err error) {
	if removed, c.s, err = sliceRemoveAt(c.s, i); err != nil {
		return removed, err
	}

	return removed, nil
}

func (c *comfyCmpSeq[V]) RemoveMatching(predicate Predicate[V]) {
	newS := make([]V, 0)
	for i, v := range c.s {
		if !predicate(i, v) {
			newS = append(newS, v)
		}
	}
	c.s = newS
}

func (c *comfyCmpSeq[V]) RemoveValues(v V) {
	newS := make([]V, 0)
	for _, current := range c.s {
		if current != v {
			newS = append(newS, current)
		}
	}
	c.s = newS
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
	newCcl := &comfyCmpSeq[V]{
		s: make([]V, 0),
	}
	newCcl.s = append(newCcl.s, c.s...)

	return newCcl
}
