package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TFindIndex struct {
	name  string
	arr   interface{}
	value interface{}
	want  interface{}
}

var tFindIndexBenchs = []TFindIndex{
	{
		name:  "10",
		arr:   []int{},
		value: 100,
	},
	{
		name:  "100",
		arr:   []int{},
		value: 100,
	},
	{
		name:  "1000",
		arr:   []int{},
		value: 100,
	},
	{
		name:  "10000",
		arr:   []int{},
		value: 100,
	},
	{
		name:  "100000",
		arr:   []int{},
		value: 100,
	},
}

func init() {
	for j := 0; j < len(tFindIndexBenchs); j++ {
		length, _ := strconv.Atoi(tFindIndexBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tFindIndexBenchs[j].arr = append(tFindIndexBenchs[j].arr.([]int), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestFindIndex(t *testing.T) {
	var tests = []TFindIndex{
		{
			name:  "nil",
			arr:   nil,
			value: 100,
			want:  nil,
		},
		{
			name:  "empty",
			arr:   []int{},
			value: 100,
			want:  -1,
		},
		{
			name:  "normal",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 5,
			want:  5,
		},
		{
			name:  "can't be found",
			arr:   []int{0, 1, 2, 3, 4, 6, 7, 8, 9},
			value: 10,
			want:  -1,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := FindIndex(subject.arr, subject.value)

			if ok := internal.Same(got, subject.want); !ok {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkFindIndex(b *testing.B) {
	for j := 0; j < len(tFindIndexBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tFindIndexBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindIndex(tFindIndexBenchs[j].arr, tFindIndexBenchs[j].value)
			}
		})
	}
}
