// Package coll provides a set of collection data structures.
//
//nolint:unused
package coll

import (
	"cmp"
)

type baseInternal[V any] interface {
	Base[V]
	copy() baseInternal[V]
	// values() iter.Seq[V] // TODO
}

type orderedInternal[V any] interface {
	Ordered[V]
	baseInternal[V]
	// valuesRev() iter.Seq[V] // TODO
}

type indexedInternal[V any] interface {
	Indexed[V]
	baseInternal[V]
}

type mutableInternal[V any] interface {
	Mutable[V]
	baseInternal[V]
}

type indexedMutableInternal[V any] interface {
	IndexedMutable[V]
	baseInternal[V]
}

type cmpInternal[V cmp.Ordered] interface {
	Cmp[V]
}

type cmpBaseInternal[B any, V cmp.Ordered] interface {
	Base[B]
	cmpInternal[V]
}

type cmpMutableInternal[V cmp.Ordered] interface {
	CmpMutable[V]
}

type cmpBaseMutableInternal[B any, V cmp.Ordered] interface {
	Base[B]
	cmpInternal[V]
	CmpMutable[V]
}

type orderedMutableInternal[V any] interface {
	OrderedMutable[V]
	baseInternal[V]
}

type listInternal[V any] interface {
	List[V]
	baseInternal[V]
}

type mapInternal[K comparable, V any] interface {
	Map[K, V]
	baseInternal[Pair[K, V]]
	// keyValues() iter.Seq2[K, V] // TODO
	prependAll(pairs []Pair[K, V])
	remove(k K)
	removeMany(keys []K)
	set(pair Pair[K, V])
}

type cmpMapInternal[K comparable, V cmp.Ordered] interface {
	CmpMap[K, V]
	mapInternal[K, V]
	cmpInternal[V]
}

type cmpMapBaseInternal[K comparable, V cmp.Ordered] interface {
	Base[Pair[K, V]]
	cmpMapInternal[K, V]
	cmpBaseInternal[Pair[K, V], V]
}

type comfyPair[K comparable, V any] struct {
	k K
	v V
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
