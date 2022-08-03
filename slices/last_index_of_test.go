package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TLastIndexOf struct {
	name  string
	arr   interface{}
	value interface{}
	want  interface{}
}

var tLastIndexOfBenchs = []TLastIndexOf{
	{
		name:  "10",
		arr:   []int{},
		value: 10,
	},
	{
		name:  "100",
		arr:   []int{},
		value: 10,
	},
	{
		name:  "1000",
		arr:   []int{},
		value: 10,
	},
	{
		name:  "10000",
		arr:   []int{},
		value: 10,
	},
	{
		name:  "100000",
		arr:   []int{},
		value: 10,
	},
}

func init() {
	for j := 0; j < len(tLastIndexOfBenchs); j++ {
		length, _ := strconv.Atoi(tLastIndexOfBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tLastIndexOfBenchs[j].arr = append(tLastIndexOfBenchs[j].arr.([]int), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
		tLastIndexOfBenchs[j].arr = append(tLastIndexOfBenchs[j].arr.([]int), 10)
	}
}

func TestLastIndexOf(t *testing.T) {
	var tests = []TLastIndexOf{
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
			want:  6,
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
			got := LastIndexOf(subject.arr, subject.value, -1)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkLastIndexOf(b *testing.B) {
	for j := 0; j < len(tLastIndexOfBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tLastIndexOfBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LastIndexOf(tLastIndexOfBenchs[j].arr, tLastIndexOfBenchs[j].value, -1)
			}
		})
	}
}
