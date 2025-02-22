package coll

import (
	"reflect"
	"testing"
)

type linearIntArgs = testArgs[Linear[int], int]
type linearTestCase = testCase[Linear[int], int]
type linearCollIntBuilder = testCollectionBuilder[Linear[int], int]

func getHeadCases(builder linearCollIntBuilder) []linearTestCase {
	return []linearTestCase{
		{
			name:  "Head() on empty collection",
			coll:  builder.Empty(),
			want1: 0,
			want2: false,
		},
		{
			name:  "Head() on one-item collection",
			coll:  builder.One(),
			want1: 111,
			want2: true,
		},
		{
			name:  "Head() on three-item collection",
			coll:  builder.Three(),
			want1: 111,
			want2: true,
		},
	}
}

func testHead(t *testing.T, builder linearCollIntBuilder) {
	cases := getHeadCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.Head()
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Head() got1 = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Head() got2 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func getHeadOrDefaultCases(builder linearCollIntBuilder) []linearTestCase {
	return []linearTestCase{
		{
			name:  "HeadOrDefault() on empty collection",
			coll:  builder.Empty(),
			args:  linearIntArgs{defaultValue: -1},
			want1: -1,
		},
		{
			name:  "HeadOrDefault() on one-item collection",
			coll:  builder.One(),
			args:  linearIntArgs{defaultValue: -1},
			want1: 111,
		},
		{
			name:  "HeadOrDefault() on three-item collection",
			coll:  builder.Three(),
			args:  linearIntArgs{defaultValue: -1},
			want1: 111,
		},
	}
}

func testHeadOrDefault(t *testing.T, builder linearCollIntBuilder) {
	cases := getHeadOrDefaultCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.HeadOrDefault(tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("HeadOrDefault() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}

func getTailCases(builder linearCollIntBuilder) []linearTestCase {
	return []linearTestCase{
		{
			name:  "Tail() on empty collection",
			coll:  builder.Empty(),
			want1: 0,
			want2: false,
		},
		{
			name:  "Tail() on one-item collection",
			coll:  builder.One(),
			want1: 111,
			want2: true,
		},
		{
			name:  "Tail() on three-item collection",
			coll:  builder.Three(),
			want1: 333,
			want2: true,
		},
	}
}

func testTail(t *testing.T, builder linearCollIntBuilder) {
	cases := getTailCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.coll.Tail()
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Tail() got1 = %v, want1 %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Tail() got2 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func getTailOrDefaultCases(builder linearCollIntBuilder) []linearTestCase {
	return []linearTestCase{
		{
			name:  "TailOrDefault() on empty collection",
			coll:  builder.Empty(),
			args:  linearIntArgs{defaultValue: -1},
			want1: -1,
		},
		{
			name:  "TailOrDefault() on one-item collection",
			coll:  builder.One(),
			args:  linearIntArgs{defaultValue: -1},
			want1: 111,
		},
		{
			name:  "TailOrDefault() on three-item collection",
			coll:  builder.Three(),
			args:  linearIntArgs{defaultValue: -1},
			want1: 333,
		},
	}
}

func testTailOrDefault(t *testing.T, builder linearCollIntBuilder) {
	cases := getTailOrDefaultCases(builder)
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coll.TailOrDefault(tt.args.defaultValue)
			if !reflect.DeepEqual(got, tt.want1) {
				t.Errorf("TailOrDefault() = %v, want1 %v", got, tt.want1)
			}
		})
	}
}
