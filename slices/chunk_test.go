package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TChunk struct {
	name string
	arr  interface{}
	size int
	want interface{}
}

var tChunkBenchs = []TChunk{
	{
		name: "10",
		arr:  []int{},
		size: 1,
	},
	{
		name: "100",
		arr:  []int{},
		size: 1,
	},
	{
		name: "1000",
		arr:  []int{},
		size: 1,
	},
	{
		name: "10000",
		arr:  []int{},
		size: 1,
	},
	{
		name: "100000",
		arr:  []int{},
		size: 1,
	},
}

func init() {
	for j := 0; j < len(tChunkBenchs); j++ {
		length, _ := strconv.Atoi(tChunkBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tChunkBenchs[j].arr = append(tChunkBenchs[j].arr.([]int), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestChunk(t *testing.T) {
	var tests = []TChunk{
		{
			name: "nil",
			arr:  nil,
			size: 5,
			want: nil,
		},
		{
			name: "empty",
			arr:  []int{},
			size: 5,
			want: []int{},
		},
		{
			name: "negative",
			arr: []int{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: -8,
			want: nil,
		},
		{
			name: "zero",
			arr: []int{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 0,
			want: nil,
		},
		{
			name: "lower size",
			arr: []int{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 2,
			want: [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}},
		},
		{
			name: "half size",
			arr: []int{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 5,
			want: [][]int{{0, 1, 2, 3, 4}, {5, 6, 7, 8, 9}},
		},
		{
			name: "more than half",
			arr: []int{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 8,
			want: [][]int{{0, 1, 2, 3, 4, 5, 6, 7}, {8, 9}},
		},
		{
			name: "equal to size",
			arr: []int{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 10,
			want: [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		{
			name: "more than size",
			arr: []int{
				0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			},
			size: 15,
			want: [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got, err := Chunk(subject.arr, subject.size)

			if ok := internal.Same(got, subject.want); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, subject.want, err)
				return
			}
		})
	}
}

func BenchmarkChunk(b *testing.B) {
	for j := 0; j < len(tChunkBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tChunkBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Chunk(tChunkBenchs[j].arr, tChunkBenchs[j].size)
			}
		})
	}
}
