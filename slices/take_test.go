package slices

import (
	"testing"
)

type TTake struct {
	name     string
	arg1     []int
	arg2     int
	expected []int
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
