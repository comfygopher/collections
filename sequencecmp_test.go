package coll

import (
	"cmp"
	"reflect"
	"testing"
)

func TestNewCmpSequence(t *testing.T) {
	intSeq := NewCmpSequence[int]()
	if intSeq == nil {
		t.Error("NewCmpSequence[int]() returned nil")
	}
	if !reflect.DeepEqual(intSeq, &comfyCmpSeq[int]{s: make([]int, 0)}) {
		t.Error("NewCmpSequence[int]() did not return a comfyCmpSeq[int]")
	}

	stringSeq := NewCmpSequence[string]()
	if stringSeq == nil {
		t.Error("NewCmpSequence[string]() returned nil")
	}
	if !reflect.DeepEqual(stringSeq, &comfyCmpSeq[string]{s: make([]string, 0)}) {
		t.Error("NewCmpSequence[int]() did not return a comfyCmpSeq[int]")
	}
}

func TestNewCmpSequenceFrom(t *testing.T) {
	intSlice := []int{1, 2, 3}
	intSeq := NewCmpSequenceFrom[int](intSlice)
	if intSeq == nil {
		t.Error("NewSequence[int]() returned nil")
	}
	if !reflect.DeepEqual(intSeq, &comfyCmpSeq[int]{s: intSlice}) {
		t.Error("NewSequence[int]() did not return a comfyCmpSeq[int]")
	}

	stringSlice := []string{"a", "b", "c"}
	stringSeq := NewCmpSequenceFrom[string](stringSlice)
	if stringSeq == nil {
		t.Error("NewSequence[string]() returned nil")
	}
	if !reflect.DeepEqual(stringSeq, &comfyCmpSeq[string]{s: stringSlice}) {
		t.Error("NewSequence[int]() did not return a comfyCmpSeq[int]")
	}
}

func Test_comfyCmpSeq_Append_one(t *testing.T) {
	type args[V cmp.Ordered] struct {
		i V
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want []V
	}
	tests := []testCase[int]{
		{
			name: "Append() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{i: 123},
			want: []int{123},
		},
		{
			name: "Append() on one-item seq",
			c:    comfyCmpSeq[int]{s: []int{1}},
			args: args[int]{i: 123},
			want: []int{1, 123},
		},
		{
			name: "Append() on three-item seq",
			c:    comfyCmpSeq[int]{s: []int{1, 2, 3}},
			args: args[int]{i: 4},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Append(tt.args.i)
			if !reflect.DeepEqual(tt.c.s, tt.want) {
				t.Errorf("Each() called with wrong indices: %v, want1 %v", tt.c.s, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_Append_many(t *testing.T) {
	type args[V cmp.Ordered] struct {
		i []V
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want []V
	}
	tests := []testCase[int]{
		{
			name: "Append() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{i: []int{123}},
			want: []int{123},
		},
		{
			name: "Append() on one-item seq",
			c:    comfyCmpSeq[int]{s: []int{1}},
			args: args[int]{i: []int{123}},
			want: []int{1, 123},
		},
		{
			name: "Append() on three-item seq",
			c:    comfyCmpSeq[int]{s: []int{1, 2, 3}},
			args: args[int]{i: []int{4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Append() on none",
			c:    comfyCmpSeq[int]{s: []int{1, 2, 3}},
			args: args[int]{i: []int{}},
			want: []int{1, 2, 3},
		},
		{
			name: "Append() on none to empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{i: []int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Append(tt.args.i...)
			if !reflect.DeepEqual(tt.c.s, tt.want) {
				t.Errorf("Each() called with wrong indices: %v, want1 %v", tt.c.s, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_AppendColl(t *testing.T) {
	type args[V cmp.Ordered] struct {
		coll Indexed[V]
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want []V
	}
	tests := []testCase[int]{
		{
			name: "Append() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{coll: &comfyCmpSeq[int]{s: []int{123}}},
			want: []int{123},
		},
		{
			name: "Append() on one-item seq",
			c:    comfyCmpSeq[int]{s: []int{1}},
			args: args[int]{coll: &comfyCmpSeq[int]{s: []int{123}}},
			want: []int{1, 123},
		},
		{
			name: "Append() on three-item seq",
			c:    comfyCmpSeq[int]{s: []int{1, 2, 3}},
			args: args[int]{coll: &comfyCmpSeq[int]{s: []int{4, 5}}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Append() on none",
			c:    comfyCmpSeq[int]{s: []int{1, 2, 3}},
			args: args[int]{coll: &comfyCmpSeq[int]{s: []int{}}},
			want: []int{1, 2, 3},
		},
		{
			name: "Append() on none to empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{coll: &comfyCmpSeq[int]{s: []int{}}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.AppendColl(tt.args.coll)
			if !reflect.DeepEqual(tt.c.s, tt.want) {
				t.Errorf("Each() called with wrong indices: %v, want1 %v", tt.c.s, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_Apply(t *testing.T) {
	type args[V cmp.Ordered] struct {
		f Mapper[V]
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want []V
	}
	tests := []testCase[int]{
		{
			name: "Apply() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{f: func(i int, v int) int { return i + v }},
			want: []int{},
		},
		{
			name: "Apply() on one-item seq",
			c:    comfyCmpSeq[int]{s: []int{1}},
			args: args[int]{f: func(i int, v int) int { return i * v }},
			want: []int{0},
		},
		{
			name: "Apply() on three-item seq",
			c:    comfyCmpSeq[int]{s: []int{1, 2, 3}},
			args: args[int]{f: func(i int, v int) int { return i + v }},
			want: []int{1, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Apply(tt.args.f)
			if !reflect.DeepEqual(tt.c.s, tt.want) {
				t.Errorf("Each() called with wrong indices: %v, want1 %v", tt.c.s, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_At(t *testing.T) {
	type args struct {
		i int
	}
	type testCase[V any, C any] struct {
		name  string
		c     C
		args  args
		want1 V
		want2 bool
	}
	tests := []testCase[int, comfyCmpSeq[int]]{
		{
			name:  "At(0) on empty seq",
			c:     comfyCmpSeq[int]{s: []int{}},
			args:  args{i: 0},
			want1: 0,
			want2: false,
		},
		{
			name:  "At(0) on one-item",
			c:     comfyCmpSeq[int]{s: []int{123}},
			args:  args{i: 0},
			want1: 123,
			want2: true,
		},
		{
			name:  "At(0) on three item",
			c:     comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args:  args{i: 0},
			want1: 123,
			want2: true,
		},
		{
			name:  "At(1) on three item",
			c:     comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args:  args{i: 1},
			want1: 234,
			want2: true,
		},
		{
			name:  "At(2) on three item",
			c:     comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args:  args{i: 2},
			want1: 345,
			want2: true,
		},
		{
			name:  "At(2) on one-item",
			c:     comfyCmpSeq[int]{s: []int{123}},
			args:  args{i: 1},
			want1: 0,
			want2: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.c.At(tt.args.i)
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("At() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("At() got2 = %v, want1 %v", got2, tt.want2)
			}
		})
	}
}

func Test_comfyCmpSeq_AtOrDefault(t *testing.T) {
	type args struct {
		i            int
		defaultValue int
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args
		want V
	}
	tests := []testCase[int]{
		{
			name: "At(0) on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args{i: 0, defaultValue: -1},
			want: -1,
		},
		{
			name: "At(0) on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args{i: 0, defaultValue: -1},
			want: 123,
		},
		{
			name: "At(0) on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args{i: 0, defaultValue: -1},
			want: 123,
		},
		{
			name: "At(1) on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args{i: 1, defaultValue: -1},
			want: 234,
		},
		{
			name: "At(2) on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args{i: 2, defaultValue: -1},
			want: 345,
		},
		{
			name: "At(2) on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args{i: 1, defaultValue: -1},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.AtOrDefault(tt.args.i, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AtOrDefault() = %v, want1 %v", got, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_Clear(t *testing.T) {
	c1 := &comfyCmpSeq[int]{s: []int{1, 2, 3}}
	c1.Clear()

	if !reflect.DeepEqual(c1.s, []int{}) {
		t.Error("c.Clear() did not clear correctly")
	}

	c2 := &comfyCmpSeq[int]{s: []int{}}
	c2.Clear()

	if !reflect.DeepEqual(c1.s, []int{}) {
		t.Error("c.Clear() did not clear correctly")
	}
}

func Test_comfyCmpSeq_Contains(t *testing.T) {
	type args[V cmp.Ordered] struct {
		predicate Predicate[V]
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Contains() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 1
			}},
			want: false,
		},
		{
			name: "Contains() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 123
			}},
			want: true,
		},
		{
			name: "Contains() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 234
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Contains(tt.args.predicate); got != tt.want {
				t.Errorf("Contains() = %v, want1 %v", got, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_ContainsValue(t *testing.T) {
	type args[V cmp.Ordered] struct {
		val V
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Contains() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{val: 1},
			want: false,
		},
		{
			name: "Contains() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{val: 123},
			want: true,
		},
		{
			name: "Contains() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{val: 234},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ContainsValue(tt.args.val); got != tt.want {
				t.Errorf("ContainsValue() = %v, want1 %v", got, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_Count(t *testing.T) {
	type args[V cmp.Ordered] struct {
		predicate Predicate[V]
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want int
	}
	tests := []testCase[int]{
		{
			name: "Count() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 1
			}},
			want: 0,
		},
		{
			name: "Count() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 123
			}},
			want: 1,
		},
		{
			name: "Count() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 234
			}},
			want: 1,
		},
		{
			name: "Count() on three item, not found",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 1
			}},
			want: 0,
		},
		{
			name: "Count() on three item, all found",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return true
			}},
			want: 3,
		},
		{
			name: "Count() on three item, none found",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return false
			}},
			want: 0,
		},
		{
			name: "Count() on three item, some mod 2 found",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v%2 == 0
			}},
			want: 1,
		},
		{
			name: "Count() on three item, some not mod 2 found",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v%2 != 0
			}},
			want: 2,
		},
		{
			name: "Count() on three item, all found explicit compare",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 123 || v == 234 || v == 345
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Count(tt.args.predicate); got != tt.want {
				t.Errorf("Count() = %v, want1 %v", got, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_CountValues(t *testing.T) {
	type args[V cmp.Ordered] struct {
		val V
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want int
	}
	tests := []testCase[int]{
		{
			name: "Count() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{val: 1},
			want: 0,
		},
		{
			name: "Count() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{val: 123},
			want: 1,
		},
		{
			name: "Count() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{val: 234},
			want: 1,
		},
		{
			name: "Count() on three item, not found",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{val: 1},
			want: 0,
		},
		{
			name: "Count() on three item, all found",
			c:    comfyCmpSeq[int]{s: []int{123, 123, 123}},
			args: args[int]{val: 123},
			want: 3,
		},
		{
			name: "Count() on three item, none found",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{val: -1},
			want: 0,
		},
		{
			name: "Count() on three item, 2 found",
			c:    comfyCmpSeq[int]{s: []int{123, 123, 345}},
			args: args[int]{val: 123},
			want: 2,
		},
		{
			name: "Count() on three item, some not mod 2 found",
			c:    comfyCmpSeq[int]{s: []int{1, 123, 123}},
			args: args[int]{val: 123},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.CountValues(tt.args.val); got != tt.want {
				t.Errorf("CountValues() = %v, want1 %v", got, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_Each(t *testing.T) {
	type args[V cmp.Ordered] struct {
		f Visitor[V]
	}
	type testCase[V cmp.Ordered] struct {
		name    string
		c       comfyCmpSeq[V]
		args    args[V]
		wantIdx []int
		wantVal []V
	}

	var actualIdx []int
	var actualVal []int

	tests := []testCase[int]{
		{
			name: "Each() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{f: func(i int, v int) {
				t.Error("Each() called on empty seq")
			}},
		},
		{
			name: "Each() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{f: func(i int, v int) {
				if i != 0 || v != 123 {
					t.Error("Each() called with wrong values")
				}
			}},
		},
		{
			name: "Each() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{f: func(i int, v int) {
				if i < 0 || i > 2 || v < 123 || v > 345 {
					t.Error("Each() called with wrong values")
				}
				actualIdx = append(actualIdx, i)
				actualVal = append(actualVal, v)
			}},
			wantIdx: []int{0, 1, 2},
			wantVal: []int{123, 234, 345},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualIdx = []int{}
			actualVal = []int{}
			tt.c.Each(tt.args.f)
			if len(tt.wantIdx) == 0 {
				return
			}
			if !reflect.DeepEqual(actualIdx, tt.wantIdx) {
				t.Errorf("Each() called with wrong indices: %v, want1 %v", actualIdx, tt.wantIdx)
			}
			if !reflect.DeepEqual(actualVal, tt.wantVal) {
				t.Errorf("Each() called with wrong values: %v, want1 %v", actualVal, tt.wantVal)
			}
		})
	}
}

func Test_comfyCmpSeq_EachRev(t *testing.T) {
	type args[V cmp.Ordered] struct {
		f Visitor[V]
	}
	type testCase[V cmp.Ordered] struct {
		name    string
		c       comfyCmpSeq[V]
		args    args[V]
		wantIdx []int
		wantVal []V
	}

	var actualIdx []int
	var actualVal []int

	tests := []testCase[int]{
		{
			name: "Each() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{f: func(i int, v int) {
				t.Error("Each() called on empty seq")
			}},
		},
		{
			name: "Each() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{f: func(i int, v int) {
				if i != 0 || v != 123 {
					t.Error("Each() called with wrong values")
				}
			}},
		},
		{
			name: "Each() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{f: func(i int, v int) {
				if i < 0 || i > 2 || v < 123 || v > 345 {
					t.Error("Each() called with wrong values")
				}
				actualIdx = append(actualIdx, i)
				actualVal = append(actualVal, v)
			}},
			wantIdx: []int{2, 1, 0},
			wantVal: []int{345, 234, 123},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualIdx = []int{}
			actualVal = []int{}
			tt.c.EachRev(tt.args.f)
			if len(tt.wantIdx) == 0 {
				return
			}
			if !reflect.DeepEqual(actualIdx, tt.wantIdx) {
				t.Errorf("Each() called with wrong indices: %v, want1 %v", actualIdx, tt.wantIdx)
			}
			if !reflect.DeepEqual(actualVal, tt.wantVal) {
				t.Errorf("Each() called with wrong values: %v, want1 %v", actualVal, tt.wantVal)
			}
		})
	}
}

func Test_comfyCmpSeq_EachRevUntil(t *testing.T) {
	type args[V cmp.Ordered] struct {
		f Predicate[V]
	}
	type testCase[V cmp.Ordered] struct {
		name    string
		c       comfyCmpSeq[V]
		args    args[V]
		wantIdx []int
		wantVal []V
	}

	var actualIdx []int
	var actualVal []int

	tests := []testCase[int]{
		{
			name: "Each() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{f: func(i int, v int) bool {
				t.Error("Each() called on empty seq")
				return true
			}},
		},
		{
			name: "Each() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{f: func(i int, v int) bool {
				if i != 0 || v != 123 {
					t.Error("Each() called with wrong values")
				}
				return true
			}},
		},
		{
			name: "Each() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{f: func(i int, v int) bool {
				if i < 0 || i > 2 || v < 123 || v > 345 {
					t.Error("Each() called with wrong values")
				}
				actualIdx = append(actualIdx, i)
				actualVal = append(actualVal, v)
				return true
			}},
			wantIdx: []int{2, 1, 0},
			wantVal: []int{345, 234, 123},
		},
		{
			name: "Each() on three item, stop at 234",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{f: func(i int, v int) bool {
				if i < 0 || i > 2 || v < 123 || v > 345 {
					t.Error("Each() called with wrong values")
				}
				actualIdx = append(actualIdx, i)
				actualVal = append(actualVal, v)
				return v != 234
			}},
			wantIdx: []int{2, 1},
			wantVal: []int{345, 234},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualIdx = []int{}
			actualVal = []int{}
			tt.c.EachRevUntil(tt.args.f)
			if len(tt.wantIdx) == 0 {
				return
			}
			if !reflect.DeepEqual(actualIdx, tt.wantIdx) {
				t.Errorf("Each() called with wrong indices: %v, want1 %v", actualIdx, tt.wantIdx)
			}
			if !reflect.DeepEqual(actualVal, tt.wantVal) {
				t.Errorf("Each() called with wrong values: %v, want1 %v", actualVal, tt.wantVal)
			}
		})
	}
}

func Test_comfyCmpSeq_EachUntil(t *testing.T) {
	type args[V cmp.Ordered] struct {
		f Predicate[V]
	}
	type testCase[V cmp.Ordered] struct {
		name    string
		c       comfyCmpSeq[V]
		args    args[V]
		wantIdx []int
		wantVal []V
	}

	var actualIdx []int
	var actualVal []int

	tests := []testCase[int]{
		{
			name: "Each() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{f: func(i int, v int) bool {
				t.Error("Each() called on empty seq")
				return true
			}},
		},
		{
			name: "Each() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{f: func(i int, v int) bool {
				if i != 0 || v != 123 {
					t.Error("Each() called with wrong values")
				}
				return true
			}},
		},
		{
			name: "Each() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{f: func(i int, v int) bool {
				if i < 0 || i > 2 || v < 123 || v > 345 {
					t.Error("Each() called with wrong values")
				}
				actualIdx = append(actualIdx, i)
				actualVal = append(actualVal, v)
				return true
			}},
			wantIdx: []int{0, 1, 2},
			wantVal: []int{123, 234, 345},
		},
		{
			name: "Each() on three item, stop at 234",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{f: func(i int, v int) bool {
				if i < 0 || i > 2 || v < 123 || v > 345 {
					t.Error("Each() called with wrong values")
				}
				actualIdx = append(actualIdx, i)
				actualVal = append(actualVal, v)
				return v != 234
			}},
			wantIdx: []int{0, 1},
			wantVal: []int{123, 234},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualIdx = []int{}
			actualVal = []int{}
			tt.c.EachUntil(tt.args.f)
			if len(tt.wantIdx) == 0 {
				return
			}
			if !reflect.DeepEqual(actualIdx, tt.wantIdx) {
				t.Errorf("Each() called with wrong indices: %v, want1 %v", actualIdx, tt.wantIdx)
			}
			if !reflect.DeepEqual(actualVal, tt.wantVal) {
				t.Errorf("Each() called with wrong values: %v, want1 %v", actualVal, tt.wantVal)
			}
		})
	}
}

func Test_comfyCmpSeq_Find(t *testing.T) {
	type args[V cmp.Ordered] struct {
		predicate    Predicate[V]
		defaultValue V
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want V
	}
	tests := []testCase[int]{
		{
			name: "Find() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 1
			}, defaultValue: -1},
			want: -1,
		},
		{
			name: "Find() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 123
			}, defaultValue: -1},
			want: 123,
		},
		{
			name: "Find() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 234
			}, defaultValue: -1},
			want: 234,
		},
		{
			name: "Find() on three item, not found",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 1
			}, defaultValue: -1},
			want: -1,
		},
		{
			name: "Find() first one",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return true
			}, defaultValue: 0},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Find(tt.args.predicate, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want1 %v", got, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_FindLast(t *testing.T) {
	type args[V cmp.Ordered] struct {
		predicate    Predicate[V]
		defaultValue V
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want V
	}
	tests := []testCase[int]{
		{
			name: "FindLast() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 1
			}, defaultValue: -1},
			want: -1,
		},
		{
			name: "FindLast() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 123
			}, defaultValue: -1},
			want: 123,
		},
		{
			name: "FindLast() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 234
			}, defaultValue: -1},
			want: 234,
		},
		{
			name: "FindLast() on three item, not found",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return v == 1
			}, defaultValue: -1},
			want: -1,
		},
		{
			name: "FindLast() last one",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{predicate: func(i int, v int) bool {
				return true
			}, defaultValue: 0},
			want: 345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.FindLast(tt.args.predicate, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindLast() = %v, want1 %v", got, tt.want)
			}
		})
	}
}

func Test_comfyCmpSeq_Head(t *testing.T) {
	type testCase[V cmp.Ordered] struct {
		name  string
		c     comfyCmpSeq[V]
		want1 V
		want2 bool
	}
	tests := []testCase[int]{
		{
			name:  "Head() on empty seq",
			c:     comfyCmpSeq[int]{s: []int{}},
			want1: 0,
			want2: false,
		},
		{
			name:  "Head() on one-item",
			c:     comfyCmpSeq[int]{s: []int{123}},
			want1: 123,
			want2: true,
		},
		{
			name:  "Head() on three item",
			c:     comfyCmpSeq[int]{s: []int{123, 234, 345}},
			want1: 123,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.c.Head()
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Head() got1 = %v, want1 %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Head() got2 = %v, want2 %v", got2, tt.want2)
			}
		})
	}
}

func Test_comfyCmpSeq_HeadOrDefault(t *testing.T) {
	type args[V cmp.Ordered] struct {
		defaultValue V
	}
	type testCase[V cmp.Ordered] struct {
		name string
		c    comfyCmpSeq[V]
		args args[V]
		want V
	}
	tests := []testCase[int]{
		{
			name: "HeadOrDefault() on empty seq",
			c:    comfyCmpSeq[int]{s: []int{}},
			args: args[int]{defaultValue: -1},
			want: -1,
		},
		{
			name: "HeadOrDefault() on one-item",
			c:    comfyCmpSeq[int]{s: []int{123}},
			args: args[int]{defaultValue: -1},
			want: 123,
		},
		{
			name: "HeadOrDefault() on three item",
			c:    comfyCmpSeq[int]{s: []int{123, 234, 345}},
			args: args[int]{defaultValue: -1},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.HeadOrDefault(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeadOrDefault() = %v, want1 %v", got, tt.want)
			}
		})
	}
}

//func Test_comfyCmpSeq_IndexOf(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		v V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		testArgs    testArgs[V]
//		want    int
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.c.IndexOf(tt.testArgs.v)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("IndexOf() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("IndexOf() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_InsertAt(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		coll int
//		v V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		testArgs    testArgs[V]
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := tt.c.InsertAt(tt.testArgs.coll, tt.testArgs.v); (err != nil) != tt.wantErr {
//				t.Errorf("InsertAt() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_IsEmpty(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.IsEmpty(); got != tt.want {
//				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_LastIndexOf(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		v V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		testArgs    testArgs[V]
//		want    int
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.c.LastIndexOf(tt.testArgs.v)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("LastIndexOf() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("LastIndexOf() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Len(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want int
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.Len(); got != tt.want {
//				t.Errorf("Len() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Max(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		want    V
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.c.Max()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Max() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Max() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Min(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		want    V
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.c.Min()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Min() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Min() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Prepend(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		v []V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.Prepend(tt.testArgs.v...)
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Reduce(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		reducer Reducer
//		initial V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//		want V
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.Reduce(tt.testArgs.reducer, tt.testArgs.initial); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Reduce() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_RemoveAt(t *testing.T) {
//	type testArgs struct {
//		coll int
//	}
//	type testCase[V cmp.Ordered] struct {
//		name    string
//		c       comfyCmpSeq[V]
//		testArgs    testArgs
//		wantErr bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := tt.c.RemoveAt(tt.testArgs.coll); (err != nil) != tt.wantErr {
//				t.Errorf("RemoveAt() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_RemoveMatching(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		predicate Predicate
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.RemoveMatching(tt.testArgs.predicate)
//		})
//	}
//}
//
//func Test_comfyCmpSeq_RemoveValues(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		v V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.RemoveValues(tt.testArgs.v)
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Reverse(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.Reverse()
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Search(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		predicate Predicate
//	}
//	type testCase[V cmp.Ordered] struct {
//		name  string
//		c     comfyCmpSeq[V]
//		testArgs  testArgs[V]
//		want  V
//		want1 bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := tt.c.Search(tt.testArgs.predicate)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Search() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("Search() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_SearchRev(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		predicate Predicate
//	}
//	type testCase[V cmp.Ordered] struct {
//		name  string
//		c     comfyCmpSeq[V]
//		testArgs  testArgs[V]
//		want  V
//		want1 bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := tt.c.SearchRev(tt.testArgs.predicate)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("SearchRev() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("SearchRev() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Sort(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		cmp func(a, b V) int
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.Sort(tt.testArgs.cmp)
//		})
//	}
//}
//
//func Test_comfyCmpSeq_SortAsc(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.SortAsc()
//		})
//	}
//}
//
//func Test_comfyCmpSeq_SortDesc(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.SortDesc()
//		})
//	}
//}
//
//func Test_comfyCmpSeq_SortMut(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		cmp func(a, b V) int
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.c.SortMut(tt.testArgs.cmp)
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Sum(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want V
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.Sum(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Sum() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Tail(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name  string
//		c     comfyCmpSeq[V]
//		want  V
//		want1 bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := tt.c.Tail()
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Tail() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("Tail() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_TailOrDefault(t *testing.T) {
//	type testArgs[V cmp.Ordered] struct {
//		defaultValue V
//	}
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		testArgs testArgs[V]
//		want V
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.TailOrDefault(tt.testArgs.defaultValue); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("TailOrDefault() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_ToSlice(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want []V
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.ToSlice(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_Values(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want iter.Seq
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.Values(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Values() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_comfyCmpSeq_copy(t *testing.T) {
//	type testCase[V cmp.Ordered] struct {
//		name string
//		c    comfyCmpSeq[V]
//		want Indexed
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.c.copy(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("copy() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
