package coll

import "testing"

type baseMapIntArgs = testArgs[Map[int, int], Pair[int, int]]
type baseMapTestCase = testCase[Map[int, int], Pair[int, int]]
type baseMapCollIntBuilder = testCollectionBuilder[Map[int, int], Pair[int, int]]

func getMapContainsCases(builder baseMapCollIntBuilder) []baseMapTestCase {
	return []baseMapTestCase{
		{
			name:  "Contains() on empty collection",
			coll:  builder.Empty(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return p.Val() == 1 }},
			want1: false,
		},
		{
			name:  "Contains() on one-item collection",
			coll:  builder.One(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return p.Val() == 111 }},
			want1: true,
		},
		{
			name:  "Contains() on three-item collection",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return p.Val() == 222 }},
			want1: true,
		},
		{
			name:  "Contains() on three-item collection, all false",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return false }},
			want1: false,
		},
		{
			name:  "Contains() on three-item collection, all true",
			coll:  builder.Three(),
			args:  baseMapIntArgs{predicate: func(i int, p Pair[int, int]) bool { return true }},
			want1: true,
		},
	}
}

func testMapContains(t *testing.T, builder baseMapCollIntBuilder) {
	cases := getMapContainsCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.coll.Contains(tt.args.predicate); got != tt.want1 {
				t.Errorf("Contains() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}
