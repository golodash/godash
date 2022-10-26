package generals

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TSame struct {
	name   string
	value1 interface{}
	value2 interface{}
	want   bool
}

var tSameBenchs = []TSame{
	{
		name:   "10",
		value1: []int{},
		value2: []int{},
	},
	{
		name:   "100",
		value1: []int{},
		value2: []int{},
	},
	{
		name:   "1000",
		value1: []int{},
		value2: []int{},
	},
	{
		name:   "10000",
		value1: []int{},
		value2: []int{},
	},
	{
		name:   "100000",
		value1: []int{},
		value2: []int{},
	},
}

func init() {
	for j := 0; j < len(tSameBenchs); j++ {
		length, _ := strconv.Atoi(tSameBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tSameBenchs[j].value1 = append(tSameBenchs[j].value1.([]int), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
			tSameBenchs[j].value2 = append(tSameBenchs[j].value2.([]int), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
	}
}

func TestSame(t *testing.T) {
	var tests = []TSame{
		{
			name:   "nil",
			value1: nil,
			value2: nil,
			want:   true,
		},
		{
			name:   "empty",
			value1: 0,
			value2: 0,
			want:   true,
		},
		{
			name:   "int int",
			value1: 15,
			value2: 3,
			want:   false,
		},
		{
			name:   "int float64",
			value1: 15,
			value2: 15.0,
			want:   false,
		},
		{
			name:   "float64 float64",
			value1: 15.1,
			value2: 15.1,
			want:   true,
		},
		{
			name:   "uint8 uint16",
			value1: uint8(3),
			value2: uint16(3),
			want:   false,
		},
		{
			name:   "uint16 uint16",
			value1: uint16(3),
			value2: uint16(3),
			want:   true,
		},
		{
			name:   "value1ay slice",
			value1: [3]int{1, 2, 3},
			value2: []int{1, 2, 3},
			want:   false,
		},
		{
			name:   "slice slice",
			value1: []int{1, 2, 3},
			value2: []int{1, 2, 3},
			want:   true,
		},
		{
			name:   "slice slice",
			value1: []int{1, 2, 3},
			value2: []int{1, 3, 2},
			want:   false,
		},
		{
			name:   "struct struct",
			value1: map[string]string{"1": "s", "2": "x"},
			value2: map[string]string{"1": "s", "2": "y"},
			want:   false,
		},
		{
			name:   "struct struct",
			value1: map[string]string{"1": "s", "2": "x"},
			value2: map[string]string{"1": "s", "2": "x"},
			want:   true,
		},
		{
			name:   "complex complex",
			value1: map[string]interface{}{"1": map[string]string{"1": "1"}, "2": []int{1}},
			value2: map[string]interface{}{"1": map[string]string{"1": "1"}, "2": []int{1}},
			want:   true,
		},
		{
			name:   "complex complex",
			value1: map[string]interface{}{"1": map[string]string{"1": "1"}, "2": []int{1}},
			value2: map[string]interface{}{"1": map[string]string{"1": "1"}, "2": []int{1, 1}},
			want:   false,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Same(subject.value1, subject.value2)

			if got != subject.want {
				t.Errorf("got = %v, wanted = %v", !subject.want, subject.want)
			}
		})
	}
}

func BenchmarkSame(b *testing.B) {
	for j := 0; j < len(tSameBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tSameBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Same(tSameBenchs[j].value1, tSameBenchs[j].value2)
			}
		})
	}
}
