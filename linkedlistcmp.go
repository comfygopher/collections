package coll

import "cmp"

// LinkedListCmp is a linked list of cmp.Ordered values
type LinkedListCmp[V cmp.Ordered] interface {
	LinkedList[V]
	Cmp[V]
}

//type linkedListCmp[V cmp.Cmp] struct {
//}
