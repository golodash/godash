package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TChunk struct {
	name string
	arr  []int
	size int
	want [][]int
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
	{
		name: "1000000",
		arr:  []int{},
		size: 1,
	},
}

func init() {
	for j := 0; j < len(tChunkBenchs); j++ {
		length, _ := strconv.Atoi(tChunkBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tChunkBenchs[j].arr = append(tChunkBenchs[j].arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
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
			want: nil,
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
			got, err := Chunk(subject.arr, subject.size)
			if err != nil {
				if subject.want != nil {
					t.Errorf("Compact() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("%v() got = %v, wanted = %v", "Chunk", got, subject.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if len(got[i]) != len(subject.want[i]) {
					t.Errorf("%v() got = %v, wanted = %v", "Chunk", got, subject.want)
					return
				}
			}

			for i := 0; i < len(got); i++ {
				for j := 0; j < len(got[i]); j++ {
					if got[i][j] != subject.want[i][j] {
						t.Errorf("%v() got = %v, wanted = %v", "Chunk", got, subject.want)
						return
					}
				}
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
