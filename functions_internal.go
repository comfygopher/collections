package coll

//lint:file-ignore U1000

import (
	"slices"
)

func comfyAppendMap[K comparable, V any](c mapInternal[K, V], p ...Pair[K, V]) {
	keys := make([]K, 0, len(p))
	for _, pair := range p {
		keys = append(keys, pair.Key())
	}
	c.removeMany(keys)
	for _, pair := range p {
		c.set(pair)
	}
}

func comfyMakeKeyPosMap[K comparable](s []K) map[K]int {
	kp := make(map[K]int, len(s))
	for i, k := range s {
		kp[k] = i
	}

	return kp
}

func comfySortSliceAndKP[K comparable, V any](s []Pair[K, V], compare PairComparator[K, V]) ([]Pair[K, V], map[K]int) {
	kp := make(map[K]int, len(s))
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

// sliceRemoveMatching removes all elements for which the predicate returns true.
func sliceRemoveMatching[V any](s []V, predicate Predicate[V]) []V {
	newS := make([]V, 0, len(s))
	for _, v := range s {
		if !predicate(v) {
			newS = append(newS, v)
		}
	}

	if len(newS) == 0 {
		return []V(nil)
	}

	return newS
}
