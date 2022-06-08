package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TZipBy struct {
	name string
	arr  [][]int
	want []int
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
	{
		name: "1000000",
		arr:  [][]int{},
	},
}

func init() {
	for j := 0; j < len(tZipByBenchs); j++ {
		length, _ := strconv.Atoi(tZipByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tZipByBenchs[j].arr = append(tZipByBenchs[j].arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		}
	}
}

func returnZipByTest(inputs ...interface{}) int {
	output := 0
	for i := 0; i < len(inputs); i++ {
		output += inputs[i].(int)
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
			inputs := []interface{}{}
			for i := 0; i < len(subject.arr); i++ {
				var temp interface{} = subject.arr[i]
				inputs = append(inputs, temp)
			}
			got, err := ZipBy(returnZipByTest, inputs...)

			if err != nil {
				if subject.want != nil {
					t.Errorf("ZipBy() got = %v, wanted = %v, error = %v", got, subject.want, err)
				}
				return
			}

			if _, err := same(got, subject.want); err != nil {
				t.Errorf("ZipBy() got = %v, wanted = %v, error = %v", got, subject.want, err)

			}
		})
	}
}

func BenchmarkZipBy(b *testing.B) {
	for j := 0; j < len(tZipByBenchs); j++ {
		inputs := []interface{}{}
		for i := 0; i < len(tZipByBenchs[j].arr); i++ {
			var temp interface{} = tZipByBenchs[j].arr[i]
			inputs = append(inputs, temp)
		}
		b.Run(fmt.Sprintf("slice_size_%s", tZipByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ZipBy(returnZipByTest, inputs...)
			}
		})
	}
}
