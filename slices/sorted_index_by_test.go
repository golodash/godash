package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TSortedIndexBy struct {
	name  string
	arr   interface{}
	value interface{}
	want  interface{}
}

var tSortedIndexByBenchs = []TSortedIndexBy{
	{
		name:  "10",
		arr:   []int{},
		value: 100,
	},
	{
		name:  "100",
		arr:   []int{},
		value: 1000,
	},
	{
		name:  "1000",
		arr:   []int{},
		value: 10000,
	},
	{
		name:  "10000",
		arr:   []int{},
		value: 100000,
	},
	{
		name:  "100000",
		arr:   []int{},
		value: 1000000,
	},
}

func init() {
	for j := 0; j < len(tSortedIndexByBenchs); j++ {
		length, _ := strconv.Atoi(tSortedIndexByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tSortedIndexByBenchs[j].arr = append(tSortedIndexByBenchs[j].arr.([]int), 0*i, 1*i, 2*i, 3*i, 4*i, 5*i, 6*i, 7*i, 8*i, 9*i)
		}
	}
}

func compareSortedIndexByTest(input interface{}) interface{} {
	return input
}

func TestSortedIndexBy(t *testing.T) {
	tests := []TSortedIndexBy{
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
			want:  0,
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
			name:  "more more sequence",
			arr:   []int{0, 1, 2, 3, 3, 3, 3, 3, 3, 4, 5, 6, 7, 8, 9},
			value: 3,
			want:  3,
		},
		{
			name:  "at the end",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 10,
			want:  10,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := SortedIndexBy(subject.arr, subject.value, compareSortedIndexByTest)

			if ok := internal.Same(got, subject.want); !ok {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkSortedIndexBy(b *testing.B) {
	for j := 0; j < len(tSortedIndexByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tSortedIndexByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortedIndexBy(tSortedIndexByBenchs[j].arr, tSortedIndexByBenchs[j].value, compareSortedIndexByTest)
			}
		})
	}
}
