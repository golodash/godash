package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TMax struct {
	name string
	arr  interface{}
	want interface{}
}

var tMaxBenchs = []TMax{
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
	for j := 0; j < len(tMaxBenchs); j++ {
		length, _ := strconv.Atoi(tMaxBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tMaxBenchs[j].arr = append(tMaxBenchs[j].arr.([]interface{}), []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestMax(t *testing.T) {
	var tests = []TMax{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  []interface{}{},
			want: nil,
		},
		{
			name: "just one item",
			arr:  []interface{}{0},
			want: 0,
		},
		{
			name: "normal",
			arr:  []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 9,
		},
		{
			name: "wrong data incompatible data types",
			arr:  []interface{}{0, []interface{}{1, 2}, []interface{}{3, 4, 5}, []interface{}{6, 7}, 8, []interface{}{9}},
			want: nil,
		},
		{
			name: "type based",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 9,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Max(subject.arr)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkMax(b *testing.B) {
	for j := 0; j < len(tMaxBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tMaxBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Max(tMaxBenchs[j].arr)
			}
		})
	}
}
