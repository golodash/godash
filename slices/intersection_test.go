package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TIntersection struct {
	name string
	arr  interface{}
	want interface{}
}

var tIntersectionBenchs = []TIntersection{
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
	for j := 0; j < len(tIntersectionBenchs); j++ {
		length, _ := strconv.Atoi(tIntersectionBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tIntersectionBenchs[j].arr = append(tIntersectionBenchs[j].arr.([]interface{}), []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		}
	}
}

func TestIntersection(t *testing.T) {
	var tests = []TIntersection{
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
			arr:  []interface{}{[]interface{}{}, []interface{}{}, 77},
			want: []interface{}{},
		},
		{
			name: "normal",
			arr:  []interface{}{[]interface{}{0, 1, 2, 3, 4}, []interface{}{3, 4}, []interface{}{5, 6, 7, 8, 9}, []interface{}{9}},
			want: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "more complex",
			arr:  []interface{}{[]interface{}{0, 1, 2, 3, 4}, []interface{}{0, 1, 2}, []interface{}{3, 4, 5, 6, 7, 8, 9}, []interface{}{5, 6, 7, 8, 9}, 55, 66},
			want: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "type based",
			arr:  [][]int{{0, 1, 2, 3}, {1, 2, 3}, {4, 5, 6}, {7}, {8, 9}},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Intersection(subject.arr)

			if !generals.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkIntersection(b *testing.B) {
	for j := 0; j < len(tIntersectionBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tIntersectionBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Intersection(tIntersectionBenchs[j].arr)
			}
		})
	}
}
