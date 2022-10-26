package generals

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TDuplicate struct {
	name  string
	value interface{}
}

var tDuplicateBenchs = []TDuplicate{
	{
		name:  "10",
		value: []int{},
	},
	{
		name:  "100",
		value: []int{},
	},
	{
		name:  "1000",
		value: []int{},
	},
	{
		name:  "10000",
		value: []int{},
	},
	{
		name:  "100000",
		value: []int{},
	},
}

func init() {
	for j := 0; j < len(tDuplicateBenchs); j++ {
		length, _ := strconv.Atoi(tDuplicateBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tDuplicateBenchs[j].value = append(tDuplicateBenchs[j].value.([]int), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
	}
}

func TestDuplicate(t *testing.T) {
	var tests = []TDuplicate{
		{
			name:  "nil",
			value: nil,
		},
		{
			name:  "empty",
			value: []int{},
		},
		{
			name:  "int",
			value: 15,
		},
		{
			name:  "float64",
			value: 15.1,
		},
		{
			name:  "array",
			value: [3]int{1, 2, 3},
		},
		{
			name:  "slice",
			value: []int{1, 2, 3},
		},
		{
			name:  "struct",
			value: map[string]string{"1": "s", "2": "x"},
		},
		{
			name:  "complex",
			value: map[string]interface{}{"1": map[string]string{"1": "1"}, "2": []int{1}},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.value)
			got := Duplicate(subject.value)

			if !Same(got, subject.value) {
				t.Errorf("got = %v, wanted = %v", got, subject.value)
			}
		})
	}
}

func BenchmarkDuplicate(b *testing.B) {
	for j := 0; j < len(tDuplicateBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tDuplicateBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Duplicate(tDuplicateBenchs[j].value)
			}
		})
	}
}
