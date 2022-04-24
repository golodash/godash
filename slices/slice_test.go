package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TSlice struct {
	name     string
	arg1     []interface{}
	arg2     int
	arg3     int
	expected interface{}
}

var TSliceBenchs = []TSlice{
	{
		name: "10",
		arg1: []interface{}{},
	},
	{
		name: "100",
		arg1: []interface{}{},
	},
	{
		name: "1000",
		arg1: []interface{}{},
	},
	{
		name: "10000",
		arg1: []interface{}{},
	},
	{
		name: "100000",
		arg1: []interface{}{},
	},
}

func init() {
	for i := 0; i < len(TSliceBenchs); i++ {
		k, _ := strconv.Atoi(TSliceBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TSliceBenchs[i].arg1 = append(TSliceBenchs[i].arg1, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
		}

	}
}

func TestSlice(t *testing.T) {
	var tests = []TSlice{
		{
			name:     "nil",
			arg1:     nil,
			arg2:     3,
			arg3:     6,
			expected: nil,
		},
		{
			name:     "empty",
			arg1:     []interface{}{},
			arg2:     2,
			arg3:     6,
			expected: []interface{}{},
		},
		{
			name:     "default",
			arg1:     []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			arg2:     2,
			arg3:     5,
			expected: []interface{}{3, 4, 5},
		},
		{
			name:     "default1",
			arg1:     []interface{}{"A", "B", "C", "D", "E", "F", "U"},
			arg2:     1,
			arg3:     6,
			expected: []interface{}{"B", "C", "D", "E", "F"},
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := Slice(sample.arg1, sample.arg2, sample.arg3)
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

func BenchmarkSlice(b *testing.B) {
	for _, sample := range TSliceBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Slice(sample.arg1, sample.arg2, sample.arg3)
			}
		})
	}

}
