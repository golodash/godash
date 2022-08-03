package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TIntersectionBy struct {
	name string
	arr  interface{}
	want interface{}
}

var tIntersectionByBenchs = []TIntersectionBy{
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

func sameIntersectionByTest(value1, value2 interface{}) bool {
	ok := internal.Same(value1, value2)
	return ok
}

func init() {
	for j := 0; j < len(tIntersectionByBenchs); j++ {
		length, _ := strconv.Atoi(tIntersectionByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tIntersectionByBenchs[j].arr = append(tIntersectionByBenchs[j].arr.([]interface{}), []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		}
	}
}

func TestIntersectionBy(t *testing.T) {
	var tests = []TIntersectionBy{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  []interface{}{},
			want: []interface{}{},
		},
		{
			name: "none",
			arr:  []interface{}{[]interface{}{0, 1, 2, 3, 4}, []interface{}{5, 6, 7, 8, 9}},
			want: []interface{}{},
		},
		{
			name: "normal",
			arr:  []interface{}{[]interface{}{0}, []interface{}{0}, []interface{}{5, 6, 7, 8, 9}, []interface{}{9}},
			want: []interface{}{0},
		},
		{
			name: "complex",
			arr:  []interface{}{[]interface{}{0, 1, 2}, []interface{}{0, 1, 2}, []interface{}{3, 4, 5, 6, 7, 8, 9}, []interface{}{5, 6, 7, 8, 9}, 5, 8},
			want: []interface{}{0, 1, 2},
		},
		{
			name: "type based",
			arr:  [][]int{{0, 1, 2}, {0, 1, 2}, {4}, {4}, {8, 9}},
			want: []int{0, 1, 2, 4},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := IntersectionBy(subject.arr, sameIntersectionByTest)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkIntersectionBy(b *testing.B) {
	for j := 0; j < len(tIntersectionByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tIntersectionByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IntersectionBy(tIntersectionByBenchs[j].arr, sameIntersectionByTest)
			}
		})
	}
}
