package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TFindIndexBy struct {
	name string
	arr  interface{}
	want interface{}
}

var tFindIndexByBenchs = []TFindIndexBy{
	{
		name: "10",
		arr:  []int{},
	},
	{
		name: "100",
		arr:  []int{},
	},
	{
		name: "1000",
		arr:  []int{},
	},
	{
		name: "10000",
		arr:  []int{},
	},
	{
		name: "100000",
		arr:  []int{},
	},
}

func init() {
	for j := 0; j < len(tFindIndexByBenchs); j++ {
		length, _ := strconv.Atoi(tFindIndexByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tFindIndexByBenchs[j].arr = append(tFindIndexByBenchs[j].arr.([]int), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func compareFindIndexByTest(value interface{}) bool {
	return value.(int) == 5
}

func TestFindIndexBy(t *testing.T) {
	var tests = []TFindIndexBy{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  []int{},
			want: -1,
		},
		{
			name: "normal",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 5,
		},
		{
			name: "can't be found",
			arr:  []int{0, 1, 2, 3, 4, 6, 7, 8, 9},
			want: -1,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := FindIndexBy(subject.arr, compareFindIndexByTest)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkFindIndexBy(b *testing.B) {
	for j := 0; j < len(tFindIndexByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tFindIndexByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindIndexBy(tFindIndexByBenchs[j].arr, compareFindIndexByTest)
			}
		})
	}
}
