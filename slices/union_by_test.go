package slices

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TUnionBy struct {
	name     string
	arg1     interface{}
	arg2     interface{}
	expected interface{}
}

var TUnionByBenchs = []TUnionBy{
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
	for i := 0; i < len(TUnionByBenchs); i++ {
		k, _ := strconv.Atoi(TUnionByBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TUnionByBenchs[i].arg1 = append(TUnionByBenchs[i].arg1.([]int), 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
			TUnionByBenchs[i].arg2 = append(TUnionByBenchs[i].arg2.([]int), 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
		}
	}
}

func compareUnionByByTest(input interface{}) interface{} {
	if reflect.TypeOf(input).Kind() == reflect.String {
		s := input.(string)
		output := []rune(s)
		return int(output[0])
	}
	return input.(int)
}

func TestUnionBy(t *testing.T) {
	var tests = []TUnionBy{
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
			got, err := UnionBy(sample.arg1, sample.arg2, compareUnionByByTest)

			if ok := internal.Same(got, sample.expected); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, sample.expected, err)
				return
			}
		})
	}
}

func BenchmarkUnionBy(b *testing.B) {
	for _, sample := range TUnionByBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UnionBy(sample.arg1, sample.arg2, compareUnionByByTest)
			}
		})
	}
}
