package coll

import "cmp"

type baseInternal[V any] interface {
	Base[V]
	copy() Base[V]
}

type linearInternal[V any] interface {
	Linear[V]
	baseInternal[V]
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

type linearMutableInternal[V any] interface {
	LinearMutable[V]
	baseInternal[V]
}

type listInternal[V any] interface {
	List[V]
	baseInternal[V]
}

type mapInternal[K comparable, V any] interface {
	Map[K, V]
	copy() mapInternal[K, V]
	prependAll(pairs []Pair[K, V])
	remove(k K)
	removeMany(keys []K)
	set(pair Pair[K, V])
}

type cmpMapInternal[K comparable, V cmp.Ordered] interface {
	CmpMap[K, V]
	mapInternal[K, V]
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
