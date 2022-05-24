package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TZip struct {
	name     string
	arg      [][]interface{}
	expected []interface{}
}

var TZipBenchs = []TZip{
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
	for i := 0; i < len(TZipBenchs); i++ {
		k, _ := strconv.Atoi(TZipBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TZipBenchs[i].arg = append(TZipBenchs[i].arg, []interface{}{"0", 1, true, 'e', false, 5, "name", 7, false, 9})
		}
	}
}

func TestZip(t *testing.T) {
	var tests = []TZip{
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
			name:     "error",
			arg:      [][]interface{}{{"a", 1, false, 15}, {"e"}},
			expected: nil,
		},
		{
			name:     "normal",
			arg:      [][]interface{}{{"a", "e"}, {1, 2}, {true, false}, {'b', 't'}},
			expected: []interface{}{[]interface{}{"a", 1, true, 'b'}, []interface{}{"e", 2, false, 't'}},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := Zip(sample.arg)
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

func BenchmarkZip(b *testing.B) {
	for _, sample := range TZipBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Zip(sample.arg)
			}
		})
	}
}
