package coll

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
