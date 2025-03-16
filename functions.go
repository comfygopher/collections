package coll

// Public:

// Copy creates a copy of the given collection.
func Copy[C Base[V], V any](coll C) C {
	// check if c is of type baseInternal[T]:
	if c, ok := any(coll).(baseInternal[V]); ok {
		return c.copy().(C)
	}
	panic("Copy() requires a collection that implements the baseInternal interface")
}

//// Filter creates a new, filtered collection from the given collection.
//func Filter[C Indexed[V], V any](c C, predicate func(int, V) bool) C {
//	panic("not implemented")
//}
//
//// MapTo creates a new, mapped collection from the given collection.
// Maybe this should be called "Transform"? Maybe "MapTo" should be an alias for "Transform"?
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
