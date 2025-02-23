package coll

import "testing"

type comfyMapIntBuilder[C Base[Pair[int, int]]] struct {
}

func (lcb *comfyMapIntBuilder[C]) Empty() C {
	return lcb.make([]Pair[int, int]{}).(C)
}

func (lcb *comfyMapIntBuilder[C]) One() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
	}).(C)
}

func (lcb *comfyMapIntBuilder[C]) Two() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
	}).(C)
}

func (lcb *comfyMapIntBuilder[C]) Three() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
		NewPair(3, 333),
	}).(C)
}

func (lcb *comfyMapIntBuilder[C]) SixWithDuplicates() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
		NewPair(3, 333),
		NewPair(4, 111),
		NewPair(5, 222),
		NewPair(6, 333),
	}).(C)
}

func (lcb *comfyMapIntBuilder[C]) make(items []Pair[int, int]) Base[Pair[int, int]] {
	coll := &comfyMap[int, int]{
		s: items,
		m: make(map[int]Pair[int, int]),
	}

	return coll
}

func Test_comfyMap_Contains(t *testing.T) {
	testMapContains(t, &comfyMapIntBuilder[Map[int, int]]{})
}
