package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TSortedIndexOf struct {
	name  string
	arr   interface{}
	value interface{}
	want  interface{}
}

var tSortedIndexOfBenchs = []TSortedIndexOf{
	{
		name:  "10",
		arr:   []int{},
		value: 0,
	},
	{
		name:  "100",
		arr:   []int{},
		value: 0,
	},
	{
		name:  "1000",
		arr:   []int{},
		value: 0,
	},
	{
		name:  "10000",
		arr:   []int{},
		value: 0,
	},
	{
		name:  "100000",
		arr:   []int{},
		value: 0,
	},
}

func init() {
	for j := 0; j < len(tSortedIndexOfBenchs); j++ {
		length, _ := strconv.Atoi(tSortedIndexOfBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tSortedIndexOfBenchs[j].arr = append(tSortedIndexOfBenchs[j].arr.([]int), 0+(i*10), 1+(i*10), 2+(i*10), 3+(i*10), 4+(i*10), 5+(i*10), 6+(i*10), 7+(i*10), 8+(i*10), 9+(i*10))
		}
	}
}

func TestSortedIndexOf(t *testing.T) {
	tests := []TSortedIndexOf{
		{
			name:  "nil",
			arr:   nil,
			value: 0,
			want:  nil,
		},
		{
			name:  "empty",
			arr:   []int{},
			value: 5,
			want:  -1,
		},
		{
			name:  "normal",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 3,
			want:  3,
		},
		{
			name:  "more sequence",
			arr:   []int{0, 1, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9},
			value: 3,
			want:  3,
		},
		{
			name:  "fail",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 88,
			want:  -1,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := SortedIndexOf(subject.arr, subject.value)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkSortedIndexOf(b *testing.B) {
	for j := 0; j < len(tSortedIndexOfBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tSortedIndexOfBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortedIndexOf(tSortedIndexOfBenchs[j].arr, tSortedIndexOfBenchs[j].value)
			}
		})
	}
}
