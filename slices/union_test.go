package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TUnion struct {
	name     string
	arg1     interface{}
	arg2     interface{}
	expected interface{}
}

var TUnionBenchs = []TUnion{
	{
		name: "10",
		arg1: []int{},
		arg2: []int{},
	},
	{
		name: "100",
		arg1: []int{},
		arg2: []int{},
	},
	{
		name: "1000",
		arg1: []int{},
		arg2: []int{},
	},
	{
		name: "10000",
		arg1: []int{},
		arg2: []int{},
	},
	{
		name: "100000",
		arg1: []int{},
		arg2: []int{},
	},
}

func init() {
	for i := 0; i < len(TUnionBenchs); i++ {
		k, _ := strconv.Atoi(TUnionBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TUnionBenchs[i].arg1 = append(TUnionBenchs[i].arg1.([]int), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}...)
			TUnionBenchs[i].arg2 = append(TUnionBenchs[i].arg2.([]int), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}...)
		}
	}
}

func TestUnion(t *testing.T) {
	var tests = []TUnion{
		{
			name:     "nil",
			arg1:     nil,
			arg2:     nil,
			expected: nil,
		},
		{
			name:     "empty",
			arg1:     []int{},
			arg2:     []int{},
			expected: []int{},
		},
		{
			name:     "default",
			arg1:     []int{2, 4, 6, 8, 0, 1, 3, 5, 7, 9},
			arg2:     []int{2, 4, 6, 10, 11, 12},
			expected: []int{2, 4, 6, 8, 0, 1, 3, 5, 7, 9, 10, 11, 12},
		},
		{
			name:     "default1",
			arg1:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			arg2:     []int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		},
		{
			name:     "interface output",
			arg1:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			arg2:     []string{"a", "b", "c"},
			expected: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, "a", "b", "c"},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := Union(sample.arg1, sample.arg2)

			if !internal.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkUnion(b *testing.B) {
	for _, sample := range TUnionBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Union(sample.arg1, sample.arg2)
			}
		})
	}
}
