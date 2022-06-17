package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TTakeRightWhile struct {
	name     string
	arg1     interface{}
	expected interface{}
}

var TTakeRightWhileBenchs = []TTakeRightWhile{
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
	for i := 0; i < len(TTakeRightWhileBenchs); i++ {
		k, _ := strconv.Atoi(TTakeRightWhileBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TTakeRightWhileBenchs[i].arg1 = append(TTakeRightWhileBenchs[i].arg1.([]int), 10, 10, 20, 30, 40, 50, 60, 70, 80, 90)
		}
	}
}

func compareTakeRightWhileTest(input interface{}) bool {
	return input.(int) > 6
}

func TestTakeRightWhile(t *testing.T) {
	var tests = []TTakeRightWhile{
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
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := TakeRightWhile(sample.arg1, compareTakeRightWhileTest)
			if err != nil {
				if sample.expected != nil {
					t.Errorf("got = %v, wanted = %v, err = %v", got, sample.expected, err)
				}
				return
			}

			if ok, _ := internal.Same(got, sample.expected); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, sample.expected, err)
				return
			}
		})
	}
}

func BenchmarkTakeRightWhile(b *testing.B) {
	for _, sample := range TTakeRightWhileBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TakeRightWhile(sample.arg1, compareTakeRightWhileTest)
			}
		})
	}
}
