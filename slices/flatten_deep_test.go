package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TFlattenDeep struct {
	name string
	arr  interface{}
	want interface{}
}

var tFlattenDeepBenchs = []TFlattenDeep{
	{
		name: "10",
		arr:  []interface{}{},
	},
	{
		name: "100",
		arr:  []interface{}{},
	},
	{
		name: "1000",
		arr:  []interface{}{},
	},
	{
		name: "10000",
		arr:  []interface{}{},
	},
	{
		name: "100000",
		arr:  []interface{}{},
	},
}

func init() {
	for j := 0; j < len(tFlattenDeepBenchs); j++ {
		length, _ := strconv.Atoi(tFlattenDeepBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tFlattenDeepBenchs[j].arr = append(tFlattenDeepBenchs[j].arr.([]interface{}), [][][]interface{}{{{0, 1, 2}}, {{3, 4, 5, 6, 7, 8, 9}}})
		}
	}
}

func TestFlattenDeep(t *testing.T) {
	var tests = []TFlattenDeep{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  []interface{}{},
			want: []interface{}{},
		},
		{
			name: "none",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "normal",
			arr:  []interface{}{0, []interface{}{1, 2}, [][]interface{}{{3}, {4, 5}}, []interface{}{6, 7}, 8, 9},
			want: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "more layer",
			arr:  [][][]int{{{0, 1, 2}}, {{3}, {4, 5, 6}}, {{}}, {{7}}, {{8, 9}}},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := FlattenDeep(subject.arr)

			if ok := internal.Same(got, subject.want); !ok {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkFlattenDeep(b *testing.B) {
	for j := 0; j < len(tFlattenDeepBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tFlattenDeepBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FlattenDeep(tFlattenDeepBenchs[j].arr)
			}
		})
	}
}
