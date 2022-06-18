package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TTake struct {
	name     string
	arg1     interface{}
	arg2     int
	expected interface{}
}

var TTakeBenchs = []TTake{
	{
		name: "10",
		arg1: []int{},
		arg2: 10,
	},
	{
		name: "100",
		arg1: []int{},
		arg2: 100,
	},
	{
		name: "1000",
		arg1: []int{},
		arg2: 1000,
	},
	{
		name: "10000",
		arg1: []int{},
		arg2: 10000,
	},
	{
		name: "100000",
		arg1: []int{},
		arg2: 100000,
	},
}

func init() {
	for i := 0; i < len(TTakeBenchs); i++ {
		k, _ := strconv.Atoi(TTakeBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TTakeBenchs[i].arg1 = append(TTakeBenchs[i].arg1.([]int), 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
		}
	}
}

func TestTake(t *testing.T) {
	var tests = []TTake{
		{
			name:     "nil",
			arg1:     nil,
			arg2:     2,
			expected: nil,
		},
		{
			name:     "empty",
			arg1:     []int{},
			arg2:     4,
			expected: []int{},
		},
		{
			name:     "empty1",
			arg1:     []int{1, 2, 3, 4},
			arg2:     0,
			expected: []int{},
		},
		{
			name:     "default",
			arg1:     []int{2, 4, 6, 8, 0, 1, 3, 5, 7, 9},
			arg2:     5,
			expected: []int{2, 4, 6, 8, 0},
		},
		{
			name:     "default1",
			arg1:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			arg2:     7,
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "new type",
			arg1:     []string{"a", "b", "c"},
			arg2:     1,
			expected: []string{"a"},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got, err := Take(sample.arg1, sample.arg2)

			if ok, _ := internal.Same(got, sample.expected); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, sample.expected, err)
				return
			}
		})
	}
}

func BenchmarkTake(b *testing.B) {
	for _, sample := range TTakeBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Take(sample.arg1, sample.arg2)
			}
		})
	}
}
