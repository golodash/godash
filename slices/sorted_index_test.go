package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TSortedIndex struct {
	name  string
	arr   []int
	value int
	want  int
}

var tSortedIndexBenchs = []TSortedIndex{
	{
		name: "10",
		arr:  []int{},
		want: 100,
	},
	{
		name: "100",
		arr:  []int{},
		want: 1000,
	},
	{
		name: "1000",
		arr:  []int{},
		want: 10000,
	},
	{
		name: "10000",
		arr:  []int{},
		want: 100000,
	},
	{
		name: "100000",
		arr:  []int{},
		want: 1000000,
	},
	{
		name: "1000000",
		arr:  []int{},
		want: 10000000,
	},
}

func init() {
	for j := 0; j < len(tSortedIndexBenchs); j++ {
		length, _ := strconv.Atoi(tSortedIndexBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tSortedIndexBenchs[j].arr = append(tSortedIndexBenchs[j].arr, []int{0 + (i * 10), 1 + (i * 10), 2 + (i * 10), 3 + (i * 10), 4 + (i * 10), 5 + (i * 10), 6 + (i * 10), 7 + (i * 10), 8 + (i * 10), 9 + (i * 10)}...)
		}
	}
}

func TestSortedIndex(t *testing.T) {
	tests := []TSortedIndex{
		{
			name:  "nil",
			arr:   nil,
			value: 0,
			want:  -1,
		},
		{
			name:  "empty",
			arr:   []int{},
			value: 5,
			want:  0,
		},
		{
			name:  "normal",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 3,
			want:  4,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := SortedIndex(subject.arr, subject.value)
			if err != nil {
				if subject.want != -1 {
					t.Errorf("SortedIndex() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if got != subject.want {
				t.Errorf("SortedIndex() got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkSortedIndex(b *testing.B) {
	for j := 0; j < len(tSortedIndexBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tSortedIndexBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortedIndex(tSortedIndexBenchs[j].arr, tSortedIndexBenchs[j].value)
			}
		})
	}
}
