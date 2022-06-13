package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TTakeRight struct {
	name     string
	arg1     []int
	arg2     int
	expected []int
}

var TTakeRightBenchs = []TTakeRight{
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
	for i := 0; i < len(TTakeRightBenchs); i++ {
		k, _ := strconv.Atoi(TTakeRightBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TTakeRightBenchs[i].arg1 = append(TTakeRightBenchs[i].arg1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}...)
		}
	}
}

func TestTakeRight(t *testing.T) {
	var tests = []TTakeRight{
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
			name:     "empty",
			arg1:     []int{1, 2, 3, 4},
			arg2:     0,
			expected: []int{},
		},
		{
			name:     "default",
			arg1:     []int{2, 4, 6, 8, 0, 1, 3, 5, 7, 9},
			arg2:     5,
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "default1",
			arg1:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			arg2:     7,
			expected: []int{4, 5, 6, 7, 8, 9, 0},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := TakeRight(sample.arg1, sample.arg2)
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

func BenchmarkTakeRight(b *testing.B) {
	for _, sample := range TTakeRightBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TakeRight(sample.arg1, sample.arg2)
			}
		})
	}
}
