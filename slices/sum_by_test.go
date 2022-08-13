package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TSumBy struct {
	name string
	arr  interface{}
	want interface{}
}

var tSumByBenchs = []TSumBy{
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
	for j := 0; j < len(tSumByBenchs); j++ {
		length, _ := strconv.Atoi(tSumByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tSumByBenchs[j].arr = append(tSumByBenchs[j].arr.([]interface{}), []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func returnSameSumByTest(value interface{}) interface{} {
	return value
}

func TestSumBy(t *testing.T) {
	var tests = []TSumBy{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  []interface{}{},
			want: 0.0,
		},
		{
			name: "just one item",
			arr:  []int{0},
			want: 0.0,
		},
		{
			name: "normal",
			arr:  []uint{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: 10.0,
		},
		{
			name: "wrong data incompatible data types",
			arr:  []float64{0, 0.2, 0.1, 0.2, 0.1, 0.2, 0.2},
			want: 1.0,
		},
		{
			name: "type based",
			arr:  []interface{}{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: 10.0,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := SumBy(subject.arr, returnSameSumByTest)

			if !internal.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkSumBy(b *testing.B) {
	for j := 0; j < len(tSumByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tSumByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SumBy(tSumByBenchs[j].arr, returnSameSumByTest)
			}
		})
	}
}
