package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TSortedLastIndexOf struct {
	name  string
	arr   interface{}
	value interface{}
	want  interface{}
}

var tSortedLastIndexOfBenchs = []TSortedLastIndexOf{
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
	for j := 0; j < len(tSortedLastIndexOfBenchs); j++ {
		length, _ := strconv.Atoi(tSortedLastIndexOfBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tSortedLastIndexOfBenchs[j].arr = append(tSortedLastIndexOfBenchs[j].arr.([]int), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
		tSortedLastIndexOfBenchs[j].arr = append(tSortedLastIndexOfBenchs[j].arr.([]int), 10)
	}
}

func TestSortedLastIndexOf(t *testing.T) {
	var tests = []TSortedLastIndexOf{
		{
			name:  "nil",
			arr:   nil,
			value: -1,
			want:  nil,
		},
		{
			name:  "empty",
			arr:   []int{},
			value: -1,
			want:  -1,
		},
		{
			name:  "normal",
			arr:   []int{0, 1, 2, 3, 4, 5, 4, 7, 8, 9},
			value: 4,
			want:  4,
		},
		{
			name:  "more sequence",
			arr:   []int{0, 1, 2, 3, 4, 4, 4, 4, 4, 4, 4, 5, 6, 7, 8, 9},
			value: 4,
			want:  10,
		},
		{
			name:  "more more sequence",
			arr:   []int{0, 1, 2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 7, 8, 9},
			value: 4,
			want:  13,
		},
		{
			name:  "not more sequence",
			arr:   []int{0, 1, 2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 7, 8, 9},
			value: 5,
			want:  14,
		},
		{
			name:  "does not exist",
			arr:   []int{0, 1, 2, 3, 4, 6, 7, 8, 9},
			value: 10,
			want:  -1,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := SortedLastIndexOf(subject.arr, subject.value)

			if got != subject.want {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkSortedLastIndexOf(b *testing.B) {
	for j := 0; j < len(tSortedLastIndexOfBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tSortedLastIndexOfBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortedLastIndexOf(tSortedLastIndexOfBenchs[j].arr, tSortedLastIndexOfBenchs[j].value)
			}
		})
	}
}
