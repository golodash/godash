package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TZipBy struct {
	name string
	arr  interface{}
	want interface{}
}

var tZipByBenchs = []TZipBy{
	{
		name: "10",
		arr:  [][]int{},
	},
	{
		name: "100",
		arr:  [][]int{},
	},
	{
		name: "1000",
		arr:  [][]int{},
	},
	{
		name: "10000",
		arr:  [][]int{},
	},
	{
		name: "100000",
		arr:  [][]int{},
	},
}

func init() {
	for j := 0; j < len(tZipByBenchs); j++ {
		length, _ := strconv.Atoi(tZipByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tZipByBenchs[j].arr = append(tZipByBenchs[j].arr.([][]int), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		}
	}
}

func returnZipByTest(inputs interface{}) interface{} {
	output := 0
	literalInputs := inputs.([]int)
	for i := 0; i < len(literalInputs); i++ {
		output += literalInputs[i]
	}

	return output
}

func TestZipBy(t *testing.T) {
	var tests = []TZipBy{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "error",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  [][]int{},
			want: []int{},
		},
		{
			name: "normal",
			arr:  [][]int{{0, 1}, {2, 3, 4}, {5, 6, 7, 8, 9}},
			want: []int{7, 10, 11, 8, 9},
		},
		{
			name: "no change",
			arr:  [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := ZipBy(subject.arr, returnZipByTest)

			if ok := internal.Same(got, subject.want); !ok {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkZipBy(b *testing.B) {
	for j := 0; j < len(tZipByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tZipByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ZipBy(tZipByBenchs[j].arr, returnZipByTest)
			}
		})
	}
}
