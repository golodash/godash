package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TTakeWhile struct {
	name     string
	arg1     interface{}
	expected interface{}
}

var TTakeWhileBenchs = []TTakeWhile{
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
	for i := 0; i < len(TTakeWhileBenchs); i++ {
		k, _ := strconv.Atoi(TTakeWhileBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TTakeWhileBenchs[i].arg1 = append(TTakeWhileBenchs[i].arg1.([]int), []int{0, -1, -2, -3, -4, -5, -6, -7, -8, -9}...)
		}
	}
}

func compareTakeWhileTest(input interface{}) bool {
	return input.(int) < 6
}

func TestTakeWhile(t *testing.T) {
	var tests = []TTakeWhile{
		{
			name:     "nil",
			arg1:     nil,
			expected: nil,
		},
		{
			name:     "empty",
			arg1:     []int{},
			expected: []int{},
		},
		{
			name:     "default",
			arg1:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{6, 7, 8, 9},
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := TakeWhile(sample.arg1, compareTakeWhileTest)

			if !generals.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkTakeWhile(b *testing.B) {
	for _, sample := range TTakeWhileBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TakeWhile(sample.arg1, compareTakeWhileTest)
			}
		})
	}
}
