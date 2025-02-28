package coll

import (
	"cmp"
	"errors"
	"fmt"
	"iter"
)

var (
	// ErrKeyNotFound     = errors.New("key not found")
	// ErrKeyAlreadyExists
	ErrOutOfBounds = errors.New("index out of bounds")
	// ErrValueNotFound
	ErrEmptyCollection = fmt.Errorf("%w: collection is empty", ErrOutOfBounds)
	// ErrKeyAlreadyExists
	ErrValueNotFound = errors.New("value not found")
)

// Predicate is used to verify collection element against implemented conditions.
type Predicate[V any] = func(i int, val V) (valid bool)

// Visitor is used to visit each element of a collection.
type Visitor[V any] = func(i int, val V)

// Reducer is used to reducer fold a collection into a single value.
type Reducer[V any] = func(acc V, i int, current V) V

// Mapper is used to map an element of a collection to a new value.
type Mapper[V any] = func(i int, val V) V

// Comparator is a comparator function.
type Comparator[V any] = func(a, b V) int

// PairComparator is a comparator function for key-value pairs.
type PairComparator[K comparable, V any] = Comparator[Pair[K, V]]

// KKVistor is a visitor function for key-value pairs.
type KVVistor[K comparable, V any] = func(i int, k K, val V)

// KVPredicate is a predicate function for key-value pairs.
type KVPredicate[K comparable, V any] = func(i int, k K, val V) (valid bool)

// KVReducer is a reducer function for key-value pairs.
type KVReducer[K comparable, V any] = func(keyAcc K, valueAcc V, currentKey K, currentValue V) (K, V)

// Base is the base interface for all collections.
type Base[V any] interface {
	Contains(predicate Predicate[V]) bool
	Count(predicate Predicate[V]) int
	Each(visit Visitor[V])
	EachRev(visit Visitor[V])
	EachRevUntil(valid Predicate[V])
	EachUntil(valid Predicate[V])
	// Find returns the first element that matches the predicate, or the default value if no element matches.
	// See: Search
	Find(predicate Predicate[V], defaultValue V) V
	// FindLast returns the last element that matches the predicate, or the default value if no element matches.
	// See: SearchRev
	FindLast(predicate Predicate[V], defaultValue V) V
	Fold(reducer Reducer[V], initial V) (result V)
	// FoldRev(reducer Reducer[V], initial V) (result V) // TODO
	IsEmpty() bool
	Len() int
	Search(predicate Predicate[V]) (val V, found bool)
	// SearchLastPos(predicate Predicate[V]) (val V, found bool) // TODO
	// SearchPos(predicate Predicate[V]) (val V, found bool) // TODO
	SearchRev(predicate Predicate[V]) (val V, found bool)
	Reduce(reducer Reducer[V]) (result V, err error)
	// ReduceRev(reducer Reducer[V]) (result V, err error) // TODO
	ToSlice() []V
	Values() iter.Seq[V]
	copy() Base[V]
}

// Linear interface indicates that given collection preserves the order of elements.
type Linear[V any] interface {
	Base[V]
	Head() (head V, ok bool)
	HeadOrDefault(defaultValue V) (head V)
	Tail() (tail V, ok bool)
	TailOrDefault(defaultValue V) (tail V)
	//LinearValues() iter.Seq2[int, V]  // TODO
}

// Indexed interface indicates that given collection can be accessed by index.
// There is no need for separate OrderedCollection interface, as all Comfy collections are ordered.
type Indexed[V any] interface {
	Linear[V]
	At(idx int) (V, bool)
	AtOrDefault(idx int, defaultValue V) V
}

// Sync is a thread-safe collection.
// TODO: To implement in the future, or drop in favor of immutable collections
//type Sync[V any] interface {
//	Indexed[V]
//	Lock()
//	Unlock()
//}

// Mutable is a collection with methods that modify its contents.
type Mutable[V any] interface {
	Base[V]
	Apply(f Mapper[V])
	Clear()
	RemoveMatching(predicate Predicate[V]) // TODO: return count of removed items
}

// IndexedMutable is a mutable collection that can be modified based on the indexes.
type IndexedMutable[V any] interface {
	Indexed[V]
	Mutable[V]
	RemoveAt(idx int) (removed V, err error)
	Sort(cmp Comparator[V])
}

// Cmp is a colection of elements of type cmp.Ordered
// It is called `cmp` (from `comparable`) instead of `ordered` to avoid confusion with collections
// that are not preserving order of elements (like the Go's native map).
type Cmp[V cmp.Ordered] interface {
	// ContainsValue returns true if the collection contains the given value.
	// Alias: HasValue
	ContainsValue(v V) bool
	CountValues(v V) int
	// HasValue is an alias for ContainsValue.
	// Deprecated: use ContainsValue instead.
	// HasValue(v V) bool // TODO
	IndexOf(val V) (i int, found bool)
	LastIndexOf(val V) (i int, found bool)
	Max() (v V, err error)
	Min() (v V, err error)
	Sum() (v V)
}

// CmpMutable is a mutable collection of elements of type cmp.Ordered
// It is called `cmp` (from `comparable`) instead of `ordered` to avoid confusion with collections
// that are not preserving order of elements (like the Go's native map).
type CmpMutable[V cmp.Ordered] interface {
	Cmp[V]
	RemoveValues(v V) // TODO: needed????
	SortAsc()
	SortDesc()
}

// LinearMutable is a mutable collection that can be modified by appending, prepending and inserting elements.
type LinearMutable[V any] interface {
	Linear[V]
	Mutable[V]
	Append(v ...V)
	AppendColl(c Linear[V])
	Prepend(v ...V)
	Reverse()
}

// Sequence is a list-like collection that wraps an underlying Go slice.
//
// Compared to a List, a Sequence allows for efficient O(1) access to arbitrary elements
// but slower insertion and removal time, making it suitable for situations where fast random access is needed.
type Sequence[V any] interface {
	LinearMutable[V]
}

// CmpSequence is a ordered collection of elements that can be compared.
type CmpSequence[V cmp.Ordered] interface {
	Sequence[V]
	CmpMutable[V]
}

// List is a mutable collection of elements.
type List[V any] interface {
	LinearMutable[V]
	InsertAt(i int, val V) error
}

// CmpLinear is a list of elements of type cmp.Ordered
type CmpLinear[V cmp.Ordered] interface {
	LinearMutable[V]
	CmpMutable[V]
}

// Map is a collection of key-value pairs.
type Map[K comparable, V any] interface {
	IndexedMutable[Pair[K, V]]
	LinearMutable[Pair[K, V]]
	// GetE(k K) (P, error) // TODO?

	Get(key K) (val V, ok bool)
	GetOrDefault(k K, defaultValue V) V
	Has(key K) bool
	Keys() iter.Seq[K]
	KeysToSlice() []K
	KeyValues() iter.Seq2[K, V]
	Remove(key K)
	RemoveMany(keys []K)
	Set(key K, val V)
	SetMany(s []Pair[K, V])
	Sort(cmp PairComparator[K, V])
	ToMap() map[K]V
	// Values returns values iterator.
	// Use KeyValues for key-value iterator.
	Values() iter.Seq[Pair[K, V]]
}

// CmpMap is a map of key-value pairs where values implement the cmp.Ordered interface
type CmpMap[K comparable, V cmp.Ordered] interface {
	Map[K, V]
	CmpMutable[V]
}

// Pair holds a key-value set of elements. It is used as the underlying value type for Map and similar collections.
// It is sealed with unexported `copy` method to prevent implementations outside the package that may allow changing
// the key. Tampering with the keys would most likely result in breaking the internal consistency of the collection.
type Pair[K comparable, V any] interface {
	Key() K
	Val() V
	SetVal(v V)
	copy() Pair[K, V]
}

type comfyPair[K comparable, V any] struct {
	k K
	v V
}

// NewPair creates a new Pair instance.
func NewPair[K comparable, V any](key K, val V) Pair[K, V] {
	return &comfyPair[K, V]{
		k: key,
		v: val,
	}
}

func NilPair[K comparable, V any]() Pair[K, V] {
	return &comfyPair[K, V]{}
}

func (p *comfyPair[K, V]) Key() K {
	return p.k
}

func (p *comfyPair[K, V]) Val() V {
	return p.v
}

func (p *comfyPair[K, V]) SetVal(v V) {
	p.v = v
}

func (p *comfyPair[K, V]) copy() Pair[K, V] {
	return &comfyPair[K, V]{
		k: p.k,
		v: p.v,
	}
}
