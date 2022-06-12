package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TXor struct {
	name     string
	arg      []interface{}
	expected interface{}
}

var TXorBenchs = []TXor{
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
	{
		name: "1000000",
		arg:  []interface{}{},
	},
}

func init() {
	for i := 0; i < len(TXorBenchs); i++ {
		k, _ := strconv.Atoi(TXorBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TXorBenchs[i].arg = append(TXorBenchs[i].arg, []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		}
	}
}

func TestXor(t *testing.T) {
	var tests = []TXor{
		{
			name:     "nil",
			arg:      nil,
			expected: []interface{}{},
		},
		{
			name:     "empty",
			arg:      []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "normal",
			arg:      []interface{}{[]int{1, 2}, []int{3, 4}, []int{5, 1, 2, 0}},
			expected: []int{3, 4, 5, 0},
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := Xor(sample.arg...)
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

func BenchmarkXor(b *testing.B) {
	for _, sample := range TXorBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Xor(sample.arg...)
			}
		})
	}
}
