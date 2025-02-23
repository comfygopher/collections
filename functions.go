package coll

import "cmp"

// public API:

// Copy creates a copy of the given collection.
func Copy[C Base[V], V any](c C) C {
	//var it V
	//
	//if _, ok := interface{}(it).(cmp.Cmp); ok {
	//	// c is of type CmpSequence[V]
	//}

	return c.copy().(C)

	//switch v := any(c).(type) {
	//case *comfySeq[V]:
	//	return c.copy().(C)
	//case *comfyCmpSeq[any]:
	//	return c.copy().(C)
	//}

	// check if c is of type Sequence[C]:
	//if cl, ok := any(c).(*comfySeq[V]); ok {
	//	s := make([]V, len(cl.s))
	//	for _, v := range cl.s {
	//		s = append(s, v)
	//	}
	//	c := &comfySeq[V]{
	//		s: s,
	//	}
	//	return any(c).(C)
	//}
}

//// Filter creates a new, filtered collection from the given collection.
//func Filter[C Indexed[V], V any](c C, predicate func(int, V) bool) C {
//	panic("not implemented")
//}
//
//// MapTo creates a new, mapped collection from the given collection.
//func MapTo[OUT Indexed[N], IN Indexed[V], V, N any](coll IN, transformer func(int, V) N) OUT {
//	panic("not implemented")
//}
//
//// Remap creates a new, mapped collection from the given collection.
//func Remap[C Indexed[V], V any](coll C) C {
//	panic("not implemented")
//}
//
//// Reverse creates a new, reversed collection from the given collection.
//func Reverse[C Indexed[V], V any](coll C) C {
//	panic("not implemented")
//}
//
//// Sort creates a new, sorted collection from the given collection.
//func Sort[C Indexed[V], V any](coll C, cmp func(a, b V) int) C {
//	panic("not implemented")
//}

// private API:

func comfyContains[C Indexed[V], V any](coll C, predicate Predicate[V]) bool {
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

func comfyContainsKV[M Map[*comfyPair[K, V], K, V], K comparable, V any](m M, predicate KVPredicate[K, V]) bool {
	found := false
	m.EachUntil(func(i int, p *comfyPair[K, V]) bool {
		if predicate(i, p.Key(), p.Value()) {
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

func comfyFindLast[C Base[V], V any](coll C, predicate Predicate[V], defaultValue V) V {
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

//func comfyIndexOf[C Base[V], V cmp.Cmp](coll C, value V) (int, error) {
//	if coll.IsEmpty() {
//		return -1, ErrEmptyCollection
//	}
//
//	foundIdx := -1
//	coll.EachUntil(func(i int, v V) bool {
//		if v == value {
//			foundIdx = i
//			return false
//		}
//
//		return true
//	})
//
//	if foundIdx == -1 {
//		return -1, ErrValueNotFound
//	}
//
//	return foundIdx, nil
//}

func comfyMax[C Base[V], V cmp.Ordered](coll C) (V, error) {
	var first = true
	var foundVal V
	coll.Each(func(i int, v V) {
		if first {
			foundVal = v
			first = false
		} else if v > foundVal {
			foundVal = v
		}
	})

	if first {
		return foundVal, ErrEmptyCollection
	}

	return foundVal, nil
}

func comfyMin[C Base[V], V cmp.Ordered](coll C) (V, error) {
	var first = true
	var foundVal V
	coll.Each(func(i int, v V) {
		if first {
			foundVal = v
			first = false
		} else if v < foundVal {
			foundVal = v
		}
	})

	if first {
		return foundVal, ErrEmptyCollection
	}

	return foundVal, nil
}

func comfyFold[C Base[V], V any](coll C, reducer Reducer[V], initial V) V {
	acc := initial
	coll.Each(func(i int, v V) {
		acc = reducer(acc, i, v)
	})

	return acc
}

func comfyFoldRev[C Base[V], V any](coll C, reducer Reducer[V], initial V) V {
	acc := initial
	coll.EachRev(func(i int, v V) {
		acc = reducer(acc, i, v)
	})

	return acc
}

func comfyReduce[C Base[V], V any](coll C, reducer Reducer[V]) (V, error) {
	var acc V
	if coll.IsEmpty() {
		return acc, ErrEmptyCollection
	}

	var first = true
	coll.Each(func(i int, v V) {
		if first {
			acc = v
			first = false
		} else {
			acc = reducer(acc, i, v)
		}
	})

	return acc, nil
}

func comfyReduceRev[C Base[V], V any](coll C, reducer Reducer[V]) (V, error) {
	var acc V
	if coll.IsEmpty() {
		return acc, ErrEmptyCollection
	}

	var first = true
	coll.EachRev(func(i int, v V) {
		if first {
			acc = v
			first = false
		} else {
			acc = reducer(acc, i, v)
		}
	})

	return acc, nil
}

func comfyReduceKV[M Map[*comfyPair[K, V], K, V], K comparable, V any](coll M, reducer KVReducer[K, V], initialKey K, initialValue V) (K, V) {
	kAcc := initialKey
	vAcc := initialValue
	coll.Each(func(i int, pair *comfyPair[K, V]) {
		kAcc, vAcc = reducer(kAcc, vAcc, pair.Key(), pair.Value())
	})

	return kAcc, vAcc
}

func comfySum[C Base[V], V cmp.Ordered](coll C) V {
	if coll.IsEmpty() {
		var v V
		return v
	}

	var initial V

	return coll.Fold(func(sum V, _ int, current V) V {
		return sum + current
	}, initial)
}

func sliceRemoveAt[V any](s []V, i int) ([]V, error) {
	if i < 0 || i >= len(s) {
		return nil, ErrOutOfBounds
	}

	return append(s[:i], s[i+1:]...), nil
}

func sliceRemoveMatching[V any](s []V, predicate Predicate[V]) []V {
	newS := make([]V, 0)
	for i, v := range s {
		if !predicate(i, v) {
			newS = append(newS, v)
		}
	}

	return newS
}

//func MapToSameColl[N any, V any, OUT Indexed[N], IN Indexed[V]](coll IN, transformer func(int, V) V) OUT {
//	panic("not implemented")
//}

//func Transform[IN Indexed[V], V, N any](coll C, transformer func(int, V) C) C {
//	panic("not implemented")
//}

//func Filter[T Indexed[V], V any](c T, predicate func(V) bool) T {
//	// check if c is of type Sequence[T]:
//	if _, ok := any(c).(Sequence[V]); ok {
//		return c.FindAll(predicate).(*comfySeq[V])
//	}
//
//	// check if c is of type CmpSequence[T]:
//	if _, ok := any(c).(CmpSequence[V]); ok {
//		return c.FindAll(predicate).(*comfyCmpSeq[V])
//	}
//
//	// check if c is of type Map[K, T]:
//}
//
//func Sort[T Indexed[V], V any](c T, cmp func(a, b V) int) T {
//
//}

//func test() {
//	l := NewSequence[int]()
//	l.Append(1, 2, 3)
//	l.AppendColl(l)
//	l.Prepend(0)
//
//	l2 := Copy(l)
//	l2.Append(4, 5, 6)
//
//	l3a := MapTo[Sequence[float64], Sequence[int]](l, func(coll int, v int) float64 {
//		return float64(v)
//	})
//
//	l3b := MapTo[Sequence[float64]](l, func(coll int, v int) float64 {
//		return float64(v)
//	})
//
//	//l3c := MapToSameColl[float64](l, func(coll int, v int) float64 {
//	//	return float64(v)
//	//})
//
//	//l4 := Transform(l, func(coll int, v int) float64 {
//	//	return float64(v)
//	//})
//
//	l4l := l4.(Sequence[float64])
//
//	l3.Append(4, 5, 6)
//	l4.Append(4, 5, 6)
//}
