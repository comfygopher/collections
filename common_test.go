package coll

type testArgs[C Base[V], V any] struct {
	index        int
	value        V
	values       []V
	defaultValue V
	visit        Visitor[V]
	predicate    Predicate[V]
	reducer      Reducer[V]
	mapper       Mapper[V]
	comparer     func(a, b V) int
	initial      V
	coll         C
}

type testCase[C Base[V], V any] struct {
	name    string
	coll    C
	args    testArgs[C, V]
	want1   any
	want2   any
	want3   any
	got1    any
	got2    any
	got3    any
	err     error
	wantErr bool
}

type testCollectionBuilder[C Base[V], V any] interface {
	Empty() C
	One() C
	Two() C
	Three() C
	SixWithDuplicates() C
}

type testPairCollectionBuilder[C Base[Pair[int, int]]] interface {
	SixWithDuplicates() C
}
