package coll

type comfyCmpMapIntBuilder[C mapInternal[int, int]] struct {
}

func (lcb *comfyCmpMapIntBuilder[C]) Empty() C {
	return lcb.make([]Pair[int, int]{}).(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) One() C {

	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
	}).(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) Two() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
	}).(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) Three() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
		NewPair(3, 333),
	}).(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) SixWithDuplicates() C {
	return lcb.make([]Pair[int, int]{
		NewPair(1, 111),
		NewPair(2, 222),
		NewPair(3, 333),
		NewPair(4, 111),
		NewPair(5, 222),
		NewPair(6, 333),
	}).(C)
}

func (lcb *comfyCmpMapIntBuilder[C]) extractUnderlyingSlice(c C) any {
	return (any(c)).(*comfyCmpMap[int, int]).s
}

func (lcb *comfyCmpMapIntBuilder[C]) extractUnderlyingMap(c C) any {
	return (any(c)).(*comfyCmpMap[int, int]).m
}

func (lcb *comfyCmpMapIntBuilder[C]) extractUnderlyingKp(c C) any {
	return (any(c)).(*comfyCmpMap[int, int]).kp
}

func (lcb *comfyCmpMapIntBuilder[C]) extractUnderlyingValsCount(c C) any {
	return (any(c)).(*comfyCmpMap[int, int]).valsCount
}

func (lcb *comfyCmpMapIntBuilder[C]) make(items []Pair[int, int]) mapInternal[int, int] {
	coll := &comfyCmpMap[int, int]{
		s:         items,
		m:         make(map[int]Pair[int, int]),
		kp:        make(map[int]int),
		valsCount: make(map[int]int),
	}

	for i, pair := range items {

		coll.m[pair.Key()] = pair
		coll.kp[pair.Key()] = i
	}

	return coll
}
