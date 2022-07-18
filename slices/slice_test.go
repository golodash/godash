package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TSlice struct {
	name     string
	arg1     interface{}
	arg2     int
	arg3     int
	expected interface{}
}

var TSliceBenchs = []TSlice{
	{
		name: "10",
		arg1: []interface{}{},
		arg2: 5,
		arg3: 10,
	},
	{
		name: "100",
		arg1: []interface{}{},
		arg2: 50,
		arg3: 100,
	},
	{
		name: "1000",
		arg1: []interface{}{},
		arg2: 500,
		arg3: 1000,
	},
	{
		name: "10000",
		arg1: []interface{}{},
		arg2: 5000,
		arg3: 10000,
	},
	{
		name: "100000",
		arg1: []interface{}{},
		arg2: 50000,
		arg3: 100000,
	},
}

func init() {
	for i := 0; i < len(TSliceBenchs); i++ {
		k, _ := strconv.Atoi(TSliceBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TSliceBenchs[i].arg1 = append(TSliceBenchs[i].arg1.([]interface{}), 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
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
			name:     "empty-error",
			arg1:     []interface{}{},
			arg2:     2,
			arg3:     6,
			expected: nil,
		},
		{
			name:     "empty",
			arg1:     []interface{}{},
			arg2:     0,
			arg3:     0,
			expected: []interface{}{},
		},
		{
			name:     "normal",
			arg1:     []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			arg2:     2,
			arg3:     5,
			expected: []interface{}{3, 4, 5},
		},
		{
			name:     "error-1",
			arg1:     []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			arg2:     2,
			arg3:     100,
			expected: nil,
		},
		{
			name:     "error-2",
			arg1:     []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			arg2:     -5,
			arg3:     3,
			expected: nil,
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := Slice(sample.arg1, sample.arg2, sample.arg3)

			if ok := internal.Same(got, sample.expected); !ok {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
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
