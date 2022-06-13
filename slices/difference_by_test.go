package slices

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type TDifferenceBy struct {
	name  string
	arr   []int
	notIn []int
	want  []int
}

var tDifferenceByBenchs = []TDifferenceBy{
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
	for j := 0; j < len(tDifferenceByBenchs); j++ {
		length, _ := strconv.Atoi(tDifferenceByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tDifferenceByBenchs[j].arr = append(tDifferenceByBenchs[j].arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
		tDifferenceByBenchs[j].notIn = append(tDifferenceByBenchs[j].notIn, 0, 1, 2, 3, 4, 5)
	}
}

func compareDifferenceByTest(value1, value2 interface{}) bool {
	v1 := reflect.ValueOf(value1).Int()
	v2 := reflect.ValueOf(value2).Int()

	return v1 == v2
}

func TestDifferenceBy(t *testing.T) {
	var tests = []TDifferenceBy{
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
			got, err := DifferenceBy(subject.arr, subject.notIn, compareDifferenceByTest)

			if err != nil {
				if subject.want != nil {
					t.Errorf("DifferenceBy() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("DifferenceBy() got = %v, wanted = %v", got, subject.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if got[i] != subject.want[i] {
					t.Errorf("DifferenceBy() got = %v, wanted = %v", got, subject.want)
					return
				}
			}
		})
	}
}

func BenchmarkDifferenceBy(b *testing.B) {
	for j := 0; j < len(tDifferenceByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tDifferenceByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DifferenceBy(tDifferenceByBenchs[j].arr, tDifferenceByBenchs[j].notIn, compareDifferenceByTest)
			}
		})
	}
}
