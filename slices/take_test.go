package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TTake struct {
	name     string
	arg1     []int
	arg2     int
	expected []int
}

var TTakeBenchs = []TTake{
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
	for i := 0; i < len(TTakeBenchs); i++ {
		k, _ := strconv.Atoi(TTakeBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TTakeBenchs[i].arg1 = append(TTakeBenchs[i].arg1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}...)
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
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := Take(sample.arg1, sample.arg2)
			if err != nil {
				if sample.expected != nil {
					t.Errorf("got : %v but expected : %v", got, sample.expected)
				}
				return
			}
			if ok, _ := same(got, sample.expected); !ok {
				t.Errorf("got : %v,%v but expected : %v", got, err, sample.expected)
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
