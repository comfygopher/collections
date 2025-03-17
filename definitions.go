package coll

import (
	"cmp"
	"errors"
	"fmt"
	"iter"
)

var (
	// ErrOutOfBounds is returned when an operation is performed on an index that is out of bounds.
	ErrOutOfBounds = errors.New("index out of bounds")

	// ErrEmptyCollection is returned when an operation is performed on an empty collection.
	ErrEmptyCollection = fmt.Errorf("%w: collection is empty", ErrOutOfBounds)

	// ErrValueNotFound is returned when a value is not found in a collection.
	ErrValueNotFound = errors.New("value not found")
)

// Predicate is used to verify that collection element meets certain conditions.
type Predicate[V any] = func(val V) (valid bool)

// Mapper is used to map an element of a collection to a new value.
type Mapper[V any] = func(val V) V

// IndexedPredicate is used to verify that collection element meets certain conditions.
type IndexedPredicate[V any] = func(i int, val V) (valid bool)

// IndexedVisitor is used to visit each element of a collection.
type IndexedVisitor[V any] = func(i int, val V)

// IndexedMapper is used to map an element of a collection to a new value.
type IndexedMapper[V any] = func(i int, val V) V

// Comparator is a comparator function.
type Comparator[V any] = func(a, b V) int

// PairComparator is a comparator function for key-value pairs.
type PairComparator[K comparable, V any] = Comparator[Pair[K, V]]

// Base is the base interface for all collections.
type Base[V any] interface {
	// IsEmpty returns true if the collection is empty.
	IsEmpty() bool

	// Len returns the number of elements in the collection.
	Len() int

	// Values returns an iterator over all elements in the collection.
	// This method is implicitly ordered on all collections that also implement the Ordered interface.
	// There is no guarantee that the order of elements will be preserved if the collection does not implement
	// the Ordered interface.
	Values() iter.Seq[V]
}

// BasePairs is the base interface for all collections of key-value pairs.
type BasePairs[K comparable, V any] interface {
	Base[Pair[K, V]]
}

// Ordered interface indicates that given collection preserves the order of elements.
type Ordered[V any] interface {
	Base[V]

	// ValuesRev returns an iterator over all elements in the collection in reverse order.
	ValuesRev() iter.Seq[V]
}

// Indexed interface indicates that given collection can be accessed by index.
type Indexed[V any] interface {
	Ordered[V]

	// At returns the element at the given index.
	At(idx int) (V, bool)

	// AtOrDefault returns the element at the given index or the default value if the index is out of bounds.
	AtOrDefault(idx int, defaultValue V) V
}

// Mutable is a collection with methods that modify its contents.
type Mutable[V any] interface {
	Base[V]

	// Apply applies the given function to each element of the collection.
	Apply(f Mapper[V])

	// Clear removes all elements from the collection.
	Clear()

	// RemoveMatching removes all elements that match the given predicate.
	// Returns the number of removed items.
	RemoveMatching(predicate Predicate[V]) (count int)
}

// IndexedMutable is a mutable collection that can be modified based on the indexes.
type IndexedMutable[V any] interface {
	Indexed[V]
	Mutable[V]

	// RemoveAt removes the element at the given index.
	RemoveAt(idx int) (removed V, err error)

	// Sort sorts the collection using the given comparator.
	Sort(cmp Comparator[V])
}

// Cmp is a collection of elements of type cmp.Ordered
// It is called `cmp` (from `comparable`) instead of `ordered` to avoid confusion with collections
// that are not preserving order of elements (like the Go's native map).
type Cmp[V cmp.Ordered] interface {
	// ContainsValue returns true if the collection contains the given value.
	// Implementations of this method should utilize ValuesCounter making this operation O(1)
	// Alias: HasValue
	ContainsValue(v V) bool

	// CountValues returns the number of times the given value appears in the collection.
	// Implementations of this method should utilize ValuesCounter making this operation O(1)
	CountValues(v V) int

	// HasValue is an alias for ContainsValue.
	// Deprecated: use ContainsValue instead.
	HasValue(v V) bool

	// IndexOf returns the index of the first occurrence of the given value.
	IndexOf(val V) (i int, found bool)

	// LastIndexOf returns the index of the last occurrence of the given value.
	LastIndexOf(val V) (i int, found bool)
}

// CmpMutable is a mutable collection of elements of type cmp.Ordered
// It is called `cmp` (from `comparable`) instead of `ordered` to avoid confusion with collections
// that are not preserving order of elements (like the Go's native map).
type CmpMutable[V cmp.Ordered] interface {
	Cmp[V]

	// RemoveValues removes all occurrences of the given value.
	// Returns the number of removed items.
	RemoveValues(v ...V) (count int)

	// SortAsc sorts the collection in ascending order.
	SortAsc()

	// SortDesc sorts the collection in descending order.
	SortDesc()
}

// OrderedMutable is a mutable collection that can be modified by appending, prepending and inserting elements.
type OrderedMutable[V any] interface {
	Ordered[V]
	Mutable[V]

	// Append appends the given values to the collection.
	Append(v ...V)

	// AppendColl appends values of the given collection to the current collection.
	AppendColl(c Ordered[V])

	// Prepend prepends the given values to the collection.
	Prepend(v ...V)

	// Reverse reverses the order of elements in the collection.
	Reverse()
}

// Sequence is a list-like collection that wraps an underlying Go slice.
//
// Compared to a List, a Sequence allows for efficient O(1) access to arbitrary elements
// but slower insertion and removal time, making it suitable for situations where fast random access is needed.
type Sequence[V any] interface {
	OrderedMutable[V]
}

// CmpSequence is a ordered collection of elements that can be compared.
type CmpSequence[V cmp.Ordered] interface {
	Sequence[V]
	CmpMutable[V]
}

// List is a mutable collection of elements.
type List[V any] interface {
	OrderedMutable[V]

	// InsertAt inserts the given value at the given index.
	InsertAt(i int, val V) error
}

// CmpOrdered is a list of elements of type cmp.Ordered
type CmpOrdered[V cmp.Ordered] interface {
	OrderedMutable[V]
	CmpMutable[V]
}

// Map is a collection of key-value pairs.
type Map[K comparable, V any] interface {
	BasePairs[K, V]
	IndexedMutable[Pair[K, V]]
	OrderedMutable[Pair[K, V]]

	// Get returns the value associated with the given key.
	Get(key K) (val V, ok bool)

	// GetOrDefault returns the value associated with the given key or the default value if the key is not found.
	GetOrDefault(k K, defaultValue V) V

	// Has returns true if the given key is present in the map.
	Has(key K) bool

	// Keys returns an iterator over all keys in the map.
	Keys() iter.Seq[K]

	// KeyValues returns an iterator over all key-value pairs in the map.
	KeyValues() iter.Seq2[K, V]

	// Remove removes the value associated with the given key.
	Remove(key K) // TODO: Replace with Remove(key ...V)

	// RemoveMany removes the values associated with the given keys.
	RemoveMany(keys []K) // TODO: remove after implementing the above Remove functionality

	// Set sets the value associated with the given key.
	Set(key K, val V)

	// SetMany sets the Pair items associated with the given keys.
	SetMany(s []Pair[K, V])

	// Sort sorts the map using the given comparator.
	Sort(compare PairComparator[K, V])

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
	// Key returns the key of the pair.
	Key() K

	// Val returns the value of the pair.
	Val() V

	// SetVal sets the value of the pair.
	SetVal(v V)

	// copy is a private method that creates a deep copy of the pair.
	copy() Pair[K, V]
}

// NewPair creates a new Pair instance.
func NewPair[K comparable, V any](key K, val V) Pair[K, V] {
	return &comfyPair[K, V]{
		k: key,
		v: val,
	}
}
