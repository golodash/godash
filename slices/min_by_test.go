package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TMinBy struct {
	name string
	arr  interface{}
	want interface{}
}

var tMinByBenchs = []TMinBy{
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
	for j := 0; j < len(tMinByBenchs); j++ {
		length, _ := strconv.Atoi(tMinByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tMinByBenchs[j].arr = append(tMinByBenchs[j].arr.([]interface{}), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
	}
}

func returnSameMinByTest(value interface{}) interface{} {
	return value
}

func TestMinBy(t *testing.T) {
	var tests = []TMinBy{
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
			want: 0,
		},
		{
			name: "it is middle",
			arr:  []int{1, 2, 3, 0, 4, 5, 6, 7, 1, 2},
			want: 0,
		},
		{
			name: "interface",
			arr:  []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 0,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := MinBy(subject.arr, returnSameMinByTest)

			if !generals.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkMinBy(b *testing.B) {
	for j := 0; j < len(tMinByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tMinByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinBy(tMinByBenchs[j].arr, returnSameMinByTest)
			}
		})
	}
}
