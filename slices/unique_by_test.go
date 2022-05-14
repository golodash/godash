package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TUniqueBy struct {
	name     string
	arg      []int
	expected []int
}

var TUniqueByBenchs = []TUniqueBy{
	{
		name: "10",
		arg:  []int{},
	},
	{
		name: "100",
		arg:  []int{},
	},
	{
		name: "1000",
		arg:  []int{},
	},
	{
		name: "10000",
		arg:  []int{},
	},
	{
		name: "100000",
		arg:  []int{},
	},
}

func init() {
	for i := 0; i < len(TUniqueByBenchs); i++ {
		k, _ := strconv.Atoi(TUniqueByBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TUniqueByBenchs[i].arg = append(TUniqueByBenchs[i].arg, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}...)
		}
	}
}

func compareItemUniqueByTest(input int) int {
	return input
}

func TestUniqueBy(t *testing.T) {
	var tests = []TUniqueBy{
		{
			name:     "nil",
			arg:      nil,
			expected: nil,
		},
		{
			name:     "empty",
			arg:      []int{},
			expected: []int{},
		},
		{
			name:     "normal",
			arg:      []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "duplicate",
			arg:      []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := UniqueBy(sample.arg, compareItemUniqueByTest)
			if err != nil {
				if sample.expected != nil {
					t.Errorf("got : %v but expected : %v", got, sample.expected)
				}
				return
			}
			if ok, _ := same(got, sample.expected); !ok {
				t.Errorf("got : %v but expected : %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkUniqueBy(b *testing.B) {
	for _, sample := range TUniqueByBenchs {
		b.Run(fmt.Sprintf("input_size_UniqueBy_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UniqueBy(sample.arg, compareItemUniqueByTest)
			}
		})
	}
}
