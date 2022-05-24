package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TUnzip struct {
	name     string
	arg      [][]interface{}
	expected []interface{}
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
			TUnzipBenchs[i].arg = append(TUnzipBenchs[i].arg, []interface{}{"0", 1, true, 'e', false, 5, "name", 7, false, 9})
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
			expected: []interface{}{},
		},
		{
			name:     "error-1",
			arg:      [][]interface{}{{"a", 1, false, 15}, {"e"}},
			expected: nil,
		},
		{
			name:     "error-2",
			arg:      [][]interface{}{{"a", 1, false, 15}, {"e", 1, 5, 15}},
			expected: nil,
		},
		{
			name:     "normal",
			arg:      [][]interface{}{{"a", 1, false, 15}, {"e", 2, true, 6}},
			expected: []interface{}{[]string{"a", "e"}, []int{1, 2}, []bool{false, true}, []int{15, 6}},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := Unzip(sample.arg)
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

func BenchmarkUnzip(b *testing.B) {
	for _, sample := range TUnzipBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Unzip(sample.arg)
			}
		})
	}
}
