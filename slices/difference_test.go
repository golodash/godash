package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TDifference struct {
	name  string
	arr   []int
	notIn []int
	want  []int
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
	{
		name:  "1000000",
		arr:   []int{},
		notIn: []int{},
	},
}

func init() {
	for j := 0; j < len(tDifferenceBenchs); j++ {
		length, _ := strconv.Atoi(tDifferenceBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tDifferenceBenchs[j].arr = append(tDifferenceBenchs[j].arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
		tDifferenceBenchs[j].notIn = append(tDifferenceBenchs[j].notIn, 0, 1, 2, 3, 4, 5)
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
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := Difference(subject.arr, subject.notIn)
			if err != nil {
				if subject.want != nil {
					t.Errorf("Difference() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("Difference() got = %v, wanted = %v", got, subject.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if got[i] != subject.want[i] {
					t.Errorf("Difference() got = %v, wanted = %v", got, subject.want)
					return
				}
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
