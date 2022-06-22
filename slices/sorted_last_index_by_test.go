package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TSortedLastIndexBy struct {
	name  string
	arr   interface{}
	value interface{}
	want  interface{}
}

var tSortedLastIndexByBenchs = []TSortedLastIndexBy{
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
	for j := 0; j < len(tSortedLastIndexByBenchs); j++ {
		length, _ := strconv.Atoi(tSortedLastIndexByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tSortedLastIndexByBenchs[j].arr = append(tSortedLastIndexByBenchs[j].arr.([]int), 0+(i*10), 1+(i*10), 2+(i*10), 3+(i*10), 4+(i*10), 5+(i*10), 6+(i*10), 7+(i*10), 8+(i*10), 9+(i*10))
		}
	}
}

func compareSortedLastIndexByTest(input interface{}) interface{} {
	return input
}

func TestSortedLastIndexBy(t *testing.T) {
	tests := []TSortedLastIndexBy{
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
			want:  4,
		},
		{
			name:  "more sequence",
			arr:   []int{0, 1, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9},
			value: 3,
			want:  6,
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
			got, err := SortedLastIndexBy(subject.arr, subject.value, compareSortedLastIndexByTest)

			if ok := internal.Same(got, subject.want); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, subject.want, err)
				return
			}
		})
	}
}

func BenchmarkSortedLastIndexBy(b *testing.B) {
	for j := 0; j < len(tSortedLastIndexByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tSortedLastIndexByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortedLastIndexBy(tSortedLastIndexByBenchs[j].arr, tSortedLastIndexByBenchs[j].value, compareSortedLastIndexByTest)
			}
		})
	}
}
