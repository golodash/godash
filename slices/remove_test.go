package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TRemove struct {
	name           string
	arg1           []int
	expectedSlice1 []int
	expectedSlice2 []int
}

var TRemoveBenchs = []TRemove{
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
	for i := 0; i < len(TRemoveBenchs); i++ {
		k, _ := strconv.Atoi(TRemoveBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TRemoveBenchs[i].arg1 = append(TRemoveBenchs[i].arg1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}...)
		}
	}
}

func removeFunctionTest(n int) bool {
	return n%2 != 0
}

func TestRemove(t *testing.T) {
	var tests = []TRemove{
		{
			name:           "nil",
			arg1:           nil,
			expectedSlice1: nil,
			expectedSlice2: nil,
		},
		{
			name:           "empty",
			arg1:           []int{},
			expectedSlice1: []int{},
			expectedSlice2: []int{},
		},
		{
			name:           "default",
			arg1:           []int{1, 2, 3, 4, 5, 6, 7, 8, 66, 44, 5, 6, 7, 99},
			expectedSlice1: []int{2, 4, 6, 8, 66, 44, 6},
			expectedSlice2: []int{99, 7, 5, 7, 5, 3, 1},
		},
		{
			name:           "default1",
			arg1:           []int{3, 4, 5, 6, 23, 34, 56, 68, 98},
			expectedSlice1: []int{4, 6, 34, 56, 68, 98},
			expectedSlice2: []int{23, 5, 3},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, got2, err := Remove(sample.arg1, removeFunctionTest)

			if err != nil {
				if sample.expectedSlice1 != nil && sample.expectedSlice2 != nil {
					t.Errorf("got: %v and %v, expected: %v and %v", got, got2, sample.expectedSlice1, sample.expectedSlice2)
				}
				return
			}

			if ok, _ := same(got, sample.expectedSlice1); !ok {
				t.Errorf("got : %v but expected : %v", got, sample.expectedSlice1)
				return
			}
			if ok, _ := same(got2, sample.expectedSlice2); !ok {
				t.Errorf("got : %v but expected : %v", got2, sample.expectedSlice2)
				return
			}
		})
	}
}

func BenchmarkRemove(b *testing.B) {
	for _, sample := range TRemoveBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Remove(sample.arg1, removeFunctionTest)
			}
		})
	}
}
