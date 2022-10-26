package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TMaxBy struct {
	name string
	arr  interface{}
	want interface{}
}

var tMaxByBenchs = []TMaxBy{
	{
		name: "10",
		arr:  []interface{}{},
	},
	{
		name: "100",
		arr:  []interface{}{},
	},
	{
		name: "1000",
		arr:  []interface{}{},
	},
	{
		name: "10000",
		arr:  []interface{}{},
	},
	{
		name: "100000",
		arr:  []interface{}{},
	},
}

func init() {
	for j := 0; j < len(tMaxByBenchs); j++ {
		length, _ := strconv.Atoi(tMaxByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tMaxByBenchs[j].arr = append(tMaxByBenchs[j].arr.([]interface{}), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
	}
}

func returnSameMaxByTest(value interface{}) interface{} {
	return value
}

func TestMaxBy(t *testing.T) {
	var tests = []TMaxBy{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  []int{},
			want: nil,
		},
		{
			name: "normal",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 9,
		},
		{
			name: "it is middle",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 1, 2},
			want: 7,
		},
		{
			name: "interface",
			arr:  []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 9,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := MaxBy(subject.arr, returnSameMaxByTest)

			if !generals.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkMaxBy(b *testing.B) {
	for j := 0; j < len(tMaxByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tMaxByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxBy(tMaxByBenchs[j].arr, returnSameMaxByTest)
			}
		})
	}
}
