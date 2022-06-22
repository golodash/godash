package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TSortedUnique struct {
	name     string
	arg1     interface{}
	expected interface{}
}

var TSortedUniqueBenchs = []TSortedUnique{
	{
		name: "10",
		arg1: []int{},
	},
	{
		name: "100",
		arg1: []int{},
	},
	{
		name: "1000",
		arg1: []int{},
	},
	{
		name: "10000",
		arg1: []int{},
	},
	{
		name: "100000",
		arg1: []int{},
	},
}

func init() {
	for i := 0; i < len(TSortedUniqueBenchs); i++ {
		k, _ := strconv.Atoi(TSortedUniqueBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TSortedUniqueBenchs[i].arg1 = append(TSortedUniqueBenchs[i].arg1.([]int), 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
		}
	}
}

func TestSortedUnique(t *testing.T) {
	var tests = []TSortedUnique{
		{
			name:     "nil",
			arg1:     nil,
			expected: nil,
		},
		{
			name:     "empty",
			arg1:     []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "default",
			arg1:     []interface{}{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 0},
			expected: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		},
		{
			name:     "type based",
			arg1:     []string{"A", "B", "C", "D", "E", "F", "U", "U", "U"},
			expected: []string{"A", "B", "C", "D", "E", "F", "U"},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got, err := SortedUnique(sample.arg1)

			if ok := internal.Same(got, sample.expected); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, sample.expected, err)
				return
			}
		})
	}
}

func BenchmarkSortedUnique(b *testing.B) {
	for _, sample := range TSortedUniqueBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortedUnique(sample.arg1)
			}
		})
	}
}
