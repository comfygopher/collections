package coll

import "cmp"

// LinkedListCmp is a linked list of cmp.Ordered values
type LinkedListCmp[V cmp.Ordered] interface {
	LinkedList[V]
	Ordered[V]
}

//type linkedListCmp[V cmp.Ordered] struct {
//}
