package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TFindIndex struct {
	name string
	arr  interface{}
	want int
}

var tFindIndexBenchs = []TFindIndex{
	{
		name: "10",
		arr:  []int{},
	},
	{
		name: "100",
		arr:  []int{},
	},
	{
		name: "1000",
		arr:  []int{},
	},
	{
		name: "10000",
		arr:  []int{},
	},
	{
		name: "100000",
		arr:  []int{},
	},
}

func init() {
	for j := 0; j < len(tFindIndexBenchs); j++ {
		length, _ := strconv.Atoi(tFindIndexBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tFindIndexBenchs[j].arr = append(tFindIndexBenchs[j].arr.([]int), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func compareFindIndexTest(value interface{}) bool {
	return value.(int) == 5
}

func TestFindIndex(t *testing.T) {
	var tests = []TFindIndex{
		{
			name: "nil",
			arr:  nil,
			want: -1,
		},
		{
			name: "empty",
			arr:  []int{},
			want: -1,
		},
		{
			name: "normal",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 5,
		},
		{
			name: "can't be found",
			arr:  []int{0, 1, 2, 3, 4, 6, 7, 8, 9},
			want: -1,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := FindIndex(subject.arr, compareFindIndexTest)
			if err != nil {
				if subject.want != -1 {
					t.Errorf("got = %v, wanted = %v, err = %v", got, subject.want, err)
				}
				return
			}

			if ok, _ := internal.Same(got, subject.want); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, subject.want, err)
				return
			}
		})
	}
}

func BenchmarkFindIndex(b *testing.B) {
	for j := 0; j < len(tFindIndexBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tFindIndexBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindIndex(tFindIndexBenchs[j].arr, compareFindIndexTest)
			}
		})
	}
}
