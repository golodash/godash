package slices

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TDifferenceBy struct {
	name  string
	arr   interface{}
	notIn interface{}
	want  interface{}
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
			tDifferenceByBenchs[j].arr = append(tDifferenceByBenchs[j].arr.([]int), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
		tDifferenceByBenchs[j].notIn = append(tDifferenceByBenchs[j].notIn.([]int), 0, 1, 2, 3, 4, 5)
	}
}

func compareDifferenceByTest(value1, value2 interface{}) bool {
	if reflect.TypeOf(value1).Kind() != reflect.TypeOf(value2).Kind() {
		return false
	}
	v1 := value1.(int)
	v2 := value2.(int)

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
		{
			name:  "type based",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			notIn: []string{"0", "1", "2", "3", "4"},
			want:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := DifferenceBy(subject.arr, subject.notIn, compareDifferenceByTest)

			if !generals.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
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
