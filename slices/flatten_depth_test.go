package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TFlattenDepth struct {
	name  string
	arr   interface{}
	depth int
	want  interface{}
}

var tFlattenDepthBenchs = []TFlattenDepth{
	{
		name:  "10",
		arr:   []interface{}{},
		depth: 100,
	},
	{
		name:  "100",
		arr:   []interface{}{},
		depth: 100,
	},
	{
		name:  "1000",
		arr:   []interface{}{},
		depth: 100,
	},
	{
		name:  "10000",
		arr:   []interface{}{},
		depth: 100,
	},
	{
		name:  "100000",
		arr:   []interface{}{},
		depth: 100,
	},
}

func init() {
	for j := 0; j < len(tFlattenDepthBenchs); j++ {
		length, _ := strconv.Atoi(tFlattenDepthBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tFlattenDepthBenchs[j].arr = append(tFlattenDepthBenchs[j].arr.([]interface{}), []interface{}{[]interface{}{[]interface{}{[]interface{}{0, []interface{}{1}, 2}, 3, 4}, 5}, 6, []interface{}{7, 8}, 9}...)
		}
	}
}

func TestFlattenDepth(t *testing.T) {
	var tests = []TFlattenDepth{
		{
			name:  "nil",
			arr:   nil,
			depth: 0,
			want:  nil,
		},
		{
			name:  "empty",
			arr:   []interface{}{},
			depth: 20,
			want:  []interface{}{},
		},
		{
			name:  "-2 level",
			arr:   []interface{}{[]interface{}{0, []interface{}{1, 2}}, 3, []interface{}{4, 5, []interface{}{6, 7}, 8}, []interface{}{9}},
			depth: -2,
			want:  []interface{}{[]interface{}{[]interface{}{[]interface{}{0, []interface{}{1, 2}}, 3, []interface{}{4, 5, []interface{}{6, 7}, 8}, []interface{}{9}}}},
		},
		{
			name:  "0 level",
			arr:   []interface{}{[]interface{}{0, []interface{}{1, 2}}, 3, []interface{}{4, 5, []interface{}{6, 7}, 8}, []interface{}{9}},
			depth: 0,
			want:  []interface{}{[]interface{}{0, []interface{}{1, 2}}, 3, []interface{}{4, 5, []interface{}{6, 7}, 8}, []interface{}{9}},
		},
		{
			name:  "1 level",
			arr:   []interface{}{[]interface{}{0, []interface{}{1, 2}}, 3, []interface{}{4, 5, []interface{}{6, 7}, 8}, []interface{}{9}},
			depth: 1,
			want:  []interface{}{0, []interface{}{1, 2}, 3, 4, 5, []interface{}{6, 7}, 8, 9},
		},
		{
			name:  "flat",
			arr:   []interface{}{[]interface{}{0, []interface{}{1, 2}}, 3, []interface{}{4, 5, []interface{}{6, 7}, 8}, []interface{}{9}},
			depth: 2,
			want:  []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "more than depths",
			arr:   []interface{}{[]interface{}{0, []interface{}{1, 2}}, 3, []interface{}{4, 5, []interface{}{6, 7}, 8}, []interface{}{9}},
			depth: 50,
			want:  []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := FlattenDepth(subject.arr, subject.depth)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkFlattenDepth(b *testing.B) {
	for j := 0; j < len(tFlattenDepthBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tFlattenDepthBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FlattenDepth(tFlattenDepthBenchs[j].arr, tFlattenDepthBenchs[j].depth)
			}
		})
	}
}
