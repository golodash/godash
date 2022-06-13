package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TIntersectionBy struct {
	name string
	arr  []interface{}
	want []interface{}
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

func removeIntersectionByTest(value1 interface{}, value2 interface{}) bool {
	return value1 == value2
}

func init() {
	for j := 0; j < len(tIntersectionByBenchs); j++ {
		length, _ := strconv.Atoi(tIntersectionByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tIntersectionByBenchs[j].arr = append(tIntersectionByBenchs[j].arr, []interface{}{[]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}...)
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
			arr:  []interface{}{[]interface{}{0, 1, 2, 3, 4}, []interface{}{3, 4}, []interface{}{5, 6, 7, 8, 9}, []interface{}{9}},
			want: []interface{}{3, 4, 9},
		},
		{
			name: "all",
			arr:  []interface{}{[]interface{}{0, 1, 2, 3, 4}, []interface{}{0, 1, 2}, []interface{}{3, 4, 5, 6, 7, 8, 9}, []interface{}{5, 6, 7, 8, 9}},
			want: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := IntersectionBy(subject.arr, removeIntersectionByTest)

			if err != nil {
				if subject.want != nil {
					t.Errorf("IntersectionBy() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("IntersectionBy() got = %v, wanted = %v", got, subject.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if got[i] != subject.want[i] {
					t.Errorf("IntersectionBy() got = %v, wanted = %v", got, subject.want)
					return
				}
			}
		})
	}
}

func BenchmarkIntersectionBy(b *testing.B) {
	for j := 0; j < len(tIntersectionByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tIntersectionByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IntersectionBy(tIntersectionByBenchs[j].arr, removeIntersectionByTest)
			}
		})
	}
}
