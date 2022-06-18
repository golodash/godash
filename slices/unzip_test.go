package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TUnzip struct {
	name     string
	arg      interface{}
	expected interface{}
}

var TUnzipBenchs = []TUnzip{
	{
		name: "10",
		arg:  [][]interface{}{},
	},
	{
		name: "100",
		arg:  [][]interface{}{},
	},
	{
		name: "1000",
		arg:  [][]interface{}{},
	},
	{
		name: "10000",
		arg:  [][]interface{}{},
	},
	{
		name: "100000",
		arg:  [][]interface{}{},
	},
}

func init() {
	for i := 0; i < len(TUnzipBenchs); i++ {
		k, _ := strconv.Atoi(TUnzipBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TUnzipBenchs[i].arg = append(TUnzipBenchs[i].arg.([][]interface{}), []interface{}{"0", 1, true, 'e', false, 5, "name", 7, false, 9})
		}
	}
}

func TestUnzip(t *testing.T) {
	var tests = []TUnzip{
		{
			name:     "nil",
			arg:      nil,
			expected: nil,
		},
		{
			name:     "empty",
			arg:      [][]interface{}{},
			expected: [][]interface{}{},
		},
		{
			name:     "error",
			arg:      [][]interface{}{{"a", 1, false, 15}, {"e"}},
			expected: nil,
		},
		{
			name:     "one different slice",
			arg:      [][]interface{}{{"a", 1, false, 15}, {"e", 2, 5, 16}},
			expected: [][]interface{}{{"a", "e"}, {1, 2}, {false, 5}, {15, 16}},
		},
		{
			name:     "normal",
			arg:      [][]interface{}{{"a", 1, false, 15}, {"e", 2, true, 6}},
			expected: [][]interface{}{{"a", "e"}, {1, 2}, {false, true}, {15, 6}},
		},
		{
			name:     "type based",
			arg:      [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}},
			expected: [][]int{{0, 4}, {1, 5}, {2, 6}, {3, 7}},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got, err := Unzip(sample.arg)

			if ok, _ := internal.Same(got, sample.expected); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, sample.expected, err)
				return
			}
		})
	}
}

func BenchmarkUnzip(b *testing.B) {
	for _, sample := range TUnzipBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Unzip(sample.arg)
			}
		})
	}
}
