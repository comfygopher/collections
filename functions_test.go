package coll

import (
	"iter"
	"reflect"
	"testing"
)

type baseFakeWithoutInternal[V any] struct{}

func (*baseFakeWithoutInternal[V]) Contains(_ Predicate[V]) bool {
	return false
}

func (*baseFakeWithoutInternal[V]) Count(_ Predicate[V]) int {
	return 0
}

func (*baseFakeWithoutInternal[V]) Each(_ Visitor[V]) {
}

func (*baseFakeWithoutInternal[V]) EachUntil(_ Predicate[V]) {
}

func (*baseFakeWithoutInternal[V]) Find(_ Predicate[V], defaultValue V) V {
	return defaultValue
}

func (*baseFakeWithoutInternal[V]) Fold(_ Reducer[V], initial V) (result V) {
	return initial
}

func (*baseFakeWithoutInternal[V]) IsEmpty() bool {
	return true
}

func (*baseFakeWithoutInternal[V]) Len() int {
	return 0
}

func (*baseFakeWithoutInternal[V]) Search(_ Predicate[V]) (val V, found bool) {
	return val, false
}

func (*baseFakeWithoutInternal[V]) Reduce(_ Reducer[V]) (result V, err error) {
	return result, nil
}

func (*baseFakeWithoutInternal[V]) ToSlice() []V {
	return nil
}

func (*baseFakeWithoutInternal[V]) Values() iter.Seq[V] {
	return nil
}

func Test_Copy_forEachFlatCollection(t *testing.T) {
	cases := []struct {
		name string
		coll LinearMutable[int]
	}{
		{
			name: "Copy() on empty Sequence",
			coll: NewSequence[int](),
		},
		{
			name: "Copy() three-item sequence",
			coll: NewSequenceFrom([]int{111, 222, 333}),
		},
		{
			name: "Copy() on empty CmpSequence",
			coll: NewCmpSequence[int](),
		},
		{
			name: "Copy() three-item CmpSequence",
			coll: NewCmpSequenceFrom([]int{111, 222, 333}),
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := Copy(tt.coll)
			if !reflect.DeepEqual(got, tt.coll) {
				t.Errorf("Copy() = %v, want %v", got, tt.coll)
			}

			tt.coll.Append(999)
			if reflect.DeepEqual(got, tt.coll) {
				t.Errorf("Copy() did not create a deep copy")
			}
			if got.Len() == tt.coll.Len() {
				t.Errorf(
					"Copy length %d should be different from original length %d after modification",
					got.Len(),
					tt.coll.Len(),
				)
			}
		})
	}
}

func Test_Copy_forEachMap(t *testing.T) {
	cases := []struct {
		name string
		coll Map[int, int]
	}{
		{
			name: "Copy() on empty Map",
			coll: NewMap[int, int](),
		},
		{
			name: "Copy() three-item Map",
			coll: NewMapFrom([]Pair[int, int]{
				NewPair(1, 111),
				NewPair(2, 222),
				NewPair(3, 333),
			}),
		},
		{
			name: "Copy() on empty CmpMap",
			coll: NewCmpMap[int, int](),
		},
		{
			name: "Copy() three-item CmpMap",
			coll: NewCmpMapFrom([]Pair[int, int]{
				NewPair(1, 111),
				NewPair(2, 222),
				NewPair(3, 333),
			}),
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := Copy(tt.coll)
			if !reflect.DeepEqual(got, tt.coll) {
				t.Errorf("Copy() = %v, want %v", got, tt.coll)
			}

			tt.coll.Append(NewPair(4, 444))
			if reflect.DeepEqual(got, tt.coll) {
				t.Errorf("Copy() did not create a deep copy")
			}
			if got.Len() == tt.coll.Len() {
				t.Errorf(
					"Copy length %d should be different from original length %d after modification",
					got.Len(),
					tt.coll.Len(),
				)
			}
		})
	}
}

func Test_Copy_ofCollectionWithoutInternal(t *testing.T) {
	t.Run("Copy() on collection without internal", func(t *testing.T) {
		coll := &baseFakeWithoutInternal[int]{}
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("Copy() did not panic")
			}
			if r != "Copy() requires a collection that implements the baseInternal interface" {
				t.Errorf("Copy() panicked with wrong error: %v", r)
			}
		}()
		Copy(coll)
	})
}
