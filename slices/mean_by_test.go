package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TMeanBy struct {
	name string
	arr  interface{}
	want interface{}
}

var tMeanByBenchs = []TMeanBy{
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
	for j := 0; j < len(tMeanByBenchs); j++ {
		length, _ := strconv.Atoi(tMeanByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tMeanByBenchs[j].arr = append(tMeanByBenchs[j].arr.([]interface{}), []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func returnSameMeanByTest(value interface{}) interface{} {
	return value
}

func TestMeanBy(t *testing.T) {
	var tests = []TMeanBy{
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
			want: 1.0,
		},
		{
			name: "float",
			arr:  []float64{0, 0.2, 0.1, 0.2, 0.1, 0.2, 0.2},
			want: 1.0 / 7,
		},
		{
			name: "type based",
			arr:  []interface{}{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want: 1.0,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := MeanBy(subject.arr, returnSameMeanByTest)

			if !generals.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkMeanBy(b *testing.B) {
	for j := 0; j < len(tMeanByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tMeanByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MeanBy(tMeanByBenchs[j].arr, returnSameMeanByTest)
			}
		})
	}
}
