package coll

import (
	"cmp"
)

// Set is a collection of unique elements.
type Set[K comparable, V cmp.Ordered] interface {
	CmpMap[*comfyPair[K, V], K, V]
}

//type comfySet[K comparable, V cmp.Ordered] struct {
//	m       map[K]*kvPair[K, V]
//	l       []*kvPair[K, V]
//	revVals map[V]bool
//	mutex   sync.Mutex
//}

//func NewSet[K comparable, V cmp.Ordered]() Set[K, V] {
//	return &Set[K, V]{
//		m:       make(map[K]*kvPair[K, V]),
//		l:       make([]*kvPair[K, V], 0),
//		revVals: make(map[V]bool),
//	}
//}
//
//func NewSetFrom[K comparable, V cmp.Ordered](m map[K]V) Set[K, V] {
//	s := NewSet[K, V]()
//	for k, v := range m {
//		s.set(k, v)
//	}
//
//	return s
//}
//
//func (om *Set[K, V]) set(it *kvPair[K, V]) {
//	if _, ok := om.revVals[it.v]; !ok {
//		om.revVals[it.v] = true
//		om.comfyMap.set(it)
//	}
//}
//
//func (cs *comfySet[K, V]) set(it *kvPair[K, V]) {
//	if _, ok := ccm.revVals[it.v]; !ok {
//		ccm.revVals[it.v] = true
//		ccm.cm.set(it)
//	}
//}
//
//func (cs *comfySet[K, V]) remove(k K) {
//	if _, ok := ccm.cm.m[k]; !ok {
//		return
//	}
//
//	ccm.cm.remove(k)
//
//	for _, current := range ccm.cm.l {
//		if current.k == k {
//			delete(ccm.revVals, current.v)
//			return
//		}
//	}
//}
