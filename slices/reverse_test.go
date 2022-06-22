package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TReverse struct {
	name     string
	arg      interface{}
	expected interface{}
}

var TReverseBenchs = []TReverse{
	{
		name: "10",
		arg:  []interface{}{},
	},
	{
		name: "100",
		arg:  []interface{}{},
	},
	{
		name: "1000",
		arg:  []interface{}{},
	},
	{
		name: "10000",
		arg:  []interface{}{},
	},
	{
		name: "100000",
		arg:  []interface{}{},
	},
}

func init() {
	for i := 0; i < len(TReverseBenchs); i++ {
		k, _ := strconv.Atoi(TReverseBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TReverseBenchs[i].arg = append(TReverseBenchs[i].arg.([]interface{}), []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestReverse(t *testing.T) {
	var tests = []TReverse{
		{
			name:     "nil",
			arg:      nil,
			expected: nil,
		},
		{
			name:     "empty",
			arg:      []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "default",
			arg:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			expected: []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:     "default1",
			arg:      []string{"a", "b", "c", "d", "e", "f", "u"},
			expected: []string{"u", "f", "e", "d", "c", "b", "a"},
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got, err := Reverse(sample.arg)

			if ok := internal.Same(got, sample.expected); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, sample.expected, err)
				return
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for _, sample := range TReverseBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Reverse(sample.arg)
			}
		})
	}
}
