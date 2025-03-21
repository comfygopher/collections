package coll

type testArgs[C any, V any] struct {
	index           int
	key             int
	keys            []int
	value           V
	values          []V
	defaultValue    V
	defaultRawValue any
	visit           func(v V)
	predicate       func(v V) bool
	intPredicate    func(i int) bool
	mapper          func(val V) V
	comparer        func(a, b V) int
	initial         V
	coll            C
}

func (ta *testArgs[C, V]) valuesAsAnySlice() []any {
	anySlice := make([]any, len(ta.values))
	for i, v := range ta.values {
		anySlice[i] = v
	}
	return anySlice
}

type testCase[C any, V any] struct {
	name        string
	coll        C
	collBuilder func() C
	args        testArgs[C, V]
	want1       any
	want2       any
	want3       any
	want4       any
	want5       any
	got1        any
	got2        any
	got3        any
	err         error
	wantErr     bool
	metaInt1    int
	modify      func()
}

type testCollectionBuilder[C any] interface {
	Empty() C
	One() C
	Two() C
	Three() C
	ThreeRev() C
	SixWithDuplicates() C
	FromValues(values []any) C

	extractRawValues(c C) any
	extractUnderlyingSlice(c C) any
	extractUnderlyingMap(c C) any
	extractUnderlyingKp(c C) any
	extractUnderlyingValsCount(c C) any
}

type testPairCollectionBuilder[C Base[Pair[int, int]]] interface {
	SixWithDuplicates() C
}
