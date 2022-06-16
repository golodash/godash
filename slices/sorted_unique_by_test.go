package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TSortedUniqueBy struct {
	name     string
	arg1     interface{}
	expected interface{}
}

var TSortedUniqueByBenchs = []TSortedUniqueBy{
	{
		name: "10",
		arg1: []interface{}{},
	},
	{
		name: "100",
		arg1: []interface{}{},
	},
	{
		name: "1000",
		arg1: []interface{}{},
	},
	{
		name: "10000",
		arg1: []interface{}{},
	},
	{
		name: "100000",
		arg1: []interface{}{},
	},
}

func init() {
	for i := 0; i < len(TSortedUniqueByBenchs); i++ {
		k, _ := strconv.Atoi(TSortedUniqueByBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TSortedUniqueByBenchs[i].arg1 = append(TSortedUniqueByBenchs[i].arg1.([]interface{}), []interface{}{1 * j, 2 * j, 3 * j, 4 * j, 5 * j, 6 * j, 7 * j, 8 * j, 9 * j, 10 * j}...)
		}
	}
}

func compareValueForSortedUniqueBy(value interface{}) interface{} {
	return value
}

func TestSortedUniqueBy(t *testing.T) {
	var tests = []TSortedUniqueBy{
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
			name:     "no need to change",
			arg1:     []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "normal",
			arg1:     []interface{}{0, 0, 1, 2, 2, 3, 4, 4, 4, 4, 4, 5, 6, 7, 7, 8, 9},
			expected: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "type based",
			arg1:     []int{0, 0, 1, 2, 2, 3, 4, 4, 4, 4, 4, 5, 6, 7, 7, 8, 9},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := SortedUniqueBy(sample.arg1, compareValueForSortedUniqueBy)
			if err != nil {
				if sample.expected != nil {
					t.Errorf("got : %v but expected : %v", got, sample.expected)
				}
				return
			}

			if ok, _ := internal.Same(got, sample.expected); !ok {
				t.Errorf("got : %v but expected : %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkSortedUniqueBy(b *testing.B) {
	for _, sample := range TSortedUniqueByBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortedUniqueBy(sample.arg1, compareValueForSortedUniqueBy)
			}
		})
	}
}
