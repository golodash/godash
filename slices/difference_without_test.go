package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TDifference struct {
	name  string
	arr   interface{}
	notIn interface{}
	want  interface{}
}

var tDifferenceBenchs = []TDifference{
	{
		name:  "10",
		arr:   []int{},
		notIn: []int{},
	},
	{
		name:  "100",
		arr:   []int{},
		notIn: []int{},
	},
	{
		name:  "1000",
		arr:   []int{},
		notIn: []int{},
	},
	{
		name:  "10000",
		arr:   []int{},
		notIn: []int{},
	},
	{
		name:  "100000",
		arr:   []int{},
		notIn: []int{},
	},
}

func init() {
	for j := 0; j < len(tDifferenceBenchs); j++ {
		length, _ := strconv.Atoi(tDifferenceBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tDifferenceBenchs[j].arr = append(tDifferenceBenchs[j].arr.([]int), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
		tDifferenceBenchs[j].notIn = append(tDifferenceBenchs[j].notIn.([]int), 0, 1, 2, 3, 4, 5)
	}
}

func TestDifference(t *testing.T) {
	var tests = []TDifference{
		{
			name:  "nil",
			arr:   nil,
			notIn: nil,
			want:  nil,
		},
		{
			name:  "empty",
			arr:   []int{},
			notIn: []int{},
			want:  []int{},
		},
		{
			name:  "normal",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			notIn: []int{0, 1, 2},
			want:  []int{3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "all remove",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			notIn: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want:  []int{},
		},
		{
			name:  "type based",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			notIn: []string{"0", "1"},
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := Difference(subject.arr, subject.notIn)
			if err != nil {
				if subject.want != nil {
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

func BenchmarkDifference(b *testing.B) {
	for j := 0; j < len(tDifferenceBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tDifferenceBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Difference(tDifferenceBenchs[j].arr, tDifferenceBenchs[j].notIn)
			}
		})
	}
}
