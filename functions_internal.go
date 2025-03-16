package coll

//lint:file-ignore U1000

import (
	"cmp"
	"slices"
)

func comfyAppendMap[K comparable, V any](c mapInternal[K, V], p ...Pair[K, V]) {
	keys := []K(nil)
	for _, pair := range p {
		keys = append(keys, pair.Key())
	}
	c.removeMany(keys)
	for _, pair := range p {
		c.set(pair)
	}
}

func comfyContains[C Base[V], V any](coll C, predicate Predicate[V]) bool {
	found := false
	coll.EachUntil(func(i int, v V) bool {
		if predicate(i, v) {
			found = true
			return false
		}

		return true
	})

	return found
}

func comfyContainsValue[C Base[V], V cmp.Ordered](coll C, search V) bool {
	return comfyContains(coll, func(_ int, v V) bool {
		return v == search
	})
}

func comfyContainsKV[M Map[K, V], K comparable, V any](m M, predicate KVPredicate[K, V]) bool {
	found := false
	m.EachUntil(func(i int, p Pair[K, V]) bool {
		if predicate(i, p.Key(), p.Val()) {
			found = true
			return false
		}

		return true
	})

	return found
}

func comfyCount[C Base[V], V any](coll C, predicate Predicate[V]) int {
	count := 0
	coll.Each(func(i int, v V) {
		if predicate(i, v) {
			count++
		}
	})

	return count
}

func comfyFind[C Base[V], V any](coll C, predicate Predicate[V], defaultValue V) V {
	found := false
	var foundValue V
	coll.EachUntil(func(i int, v V) bool {
		if predicate(i, v) {
			found = true
			foundValue = v
			return false
		}

		return true
	})

	if found {
		return foundValue
	}

	return defaultValue
}

func comfyFindLast[C Linear[V], V any](coll C, predicate Predicate[V], defaultValue V) V {
	found := false
	var foundValue V
	coll.EachRevUntil(func(i int, v V) bool {
		if predicate(i, v) {
			found = true
			foundValue = v
			return false
		}

		return true
	})

	if found {
		return foundValue
	}

	return defaultValue
}

func comfyMakeKeyPosMap[K comparable](s []K) map[K]int {
	kp := make(map[K]int, len(s))
	for i, k := range s {
		kp[k] = i
	}

	return kp
}

func comfyMax[C Base[V], V cmp.Ordered](c C) (V, error) {
	return c.Reduce(func(acc V, _ int, current V) V {
		if current > acc {
			return current
		}
		return acc
	})
}

func comfyMaxOfPairs[C BasePairs[K, V], K comparable, V cmp.Ordered](c C) (V, error) {
	first := true
	var foundVal V
	c.Each(func(_ int, p Pair[K, V]) {
		if first {
			foundVal = p.Val()
			first = false
		} else if p.Val() > foundVal {
			foundVal = p.Val()
		}
	})

	if first {
		return foundVal, ErrEmptyCollection
	}

	return foundVal, nil
}

func comfyMin[C Base[V], V cmp.Ordered](coll C) (V, error) {
	return coll.Reduce(func(acc V, _ int, current V) V {
		if current < acc {
			return current
		}
		return acc
	})
}

func comfyMinOfPairs[C BasePairs[K, V], K comparable, V cmp.Ordered](c C) (V, error) {
	first := true
	var foundVal V
	c.Each(func(_ int, p Pair[K, V]) {
		if first {
			foundVal = p.Val()
			first = false
		} else if p.Val() < foundVal {
			foundVal = p.Val()
		}
	})

	if first {
		return foundVal, ErrEmptyCollection
	}

	return foundVal, nil
}

func comfyFoldSlice[V any](s []V, reducer Reducer[V], initial V) V {
	acc := initial
	for i, v := range s {
		acc = reducer(acc, i, v)
	}

	return acc
}

func comfyFoldSliceRev[V any](s []V, reducer Reducer[V], initial V) V {
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = reducer(acc, i, s[i])
	}

	return acc
}

func comfyReduceSlice[V any](s []V, reducer Reducer[V]) (V, error) {
	var acc V
	if len(s) == 0 {
		return acc, ErrEmptyCollection
	}

	first := true
	for i, v := range s {
		if first {
			acc = v
			first = false
		} else {
			acc = reducer(acc, i, v)
		}
	}

	return acc, nil
}

func comfyReduceSliceRev[V any](s []V, reducer Reducer[V]) (V, error) {
	var acc V
	if len(s) == 0 {
		return acc, ErrEmptyCollection
	}

	first := true
	for i := len(s) - 1; i >= 0; i-- {
		if first {
			acc = s[i]
			first = false
		} else {
			acc = reducer(acc, i, s[i])
		}
	}

	return acc, nil
}

func comfySortSliceAndKP[K comparable, V any](s []Pair[K, V], compare PairComparator[K, V]) ([]Pair[K, V], map[K]int) {
	kp := make(map[K]int)
	if s == nil {
		return s, kp
	}

	slices.SortFunc(s, func(a, b Pair[K, V]) int {
		return compare(a, b)
	})
	for i, pair := range s {
		kp[pair.Key()] = i
	}

	return s, kp
}

func comfySum[C Base[V], V cmp.Ordered](coll C) V {
	var sum V
	for v := range coll.Values() {
		sum += v
	}

	return sum
}

func comfySumOfPairs[C BasePairs[K, V], K comparable, V cmp.Ordered](c C) V {
	var initial V
	for p := range c.Values() {
		initial += p.Val()
	}

	return initial
}

func sliceRemoveAt[V any](s []V, i int) (removed V, newSLice []V, err error) {
	if i < 0 || i >= len(s) {
		var v V
		return v, s, ErrOutOfBounds
	}

	removed = s[i]

	if len(s) == 1 {
		return removed, []V(nil), nil
	}

	return removed, append(s[:i], s[i+1:]...), nil
}

func sliceRemoveMatching[V any](s []V, predicate Predicate[V]) []V {
	newS := []V(nil)
	for i, v := range s {
		if !predicate(i, v) {
			newS = append(newS, v)
		}
	}

	return newS
}
