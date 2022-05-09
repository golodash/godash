package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TSUnique struct {
	name     string
	arg1     []interface{}
	expected interface{}
}

var TSUniqueBenchs = []TSUnique{
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
	for i := 0; i < len(TSUniqueBenchs); i++ {
		k, _ := strconv.Atoi(TSUniqueBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TSUniqueBenchs[i].arg1 = append(TSUniqueBenchs[i].arg1, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}...)
		}

	}
}

func TestSUnique(t *testing.T) {
	var tests = []TSUnique{
		{
			name:     "nil",
			arg1:     nil,
			expected: nil,
		},
		{
			name:     "empty",
			arg1:     []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "default",
			arg1:     []interface{}{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 0},
			expected: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		},
		{
			name:     "default1",
			arg1:     []interface{}{"A", "B", "C", "D", "E", "F", "U", "U", "U"},
			expected: []interface{}{"A", "B", "C", "D", "E", "F", "U"},
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := SortedUnique(sample.arg1)
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

func BenchmarkSUnique(b *testing.B) {
	for _, sample := range TSUniqueBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortedUnique(sample.arg1)
			}
		})
	}

}
