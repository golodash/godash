package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TXor struct {
	name     string
	arg1     interface{}
	arg2     interface{}
	expected interface{}
}

var TXorBenchs = []TXor{
	{
		name: "10",
		arg1: []interface{}{},
		arg2: []interface{}{},
	},
	{
		name: "100",
		arg1: []interface{}{},
		arg2: []interface{}{},
	},
	{
		name: "1000",
		arg1: []interface{}{},
		arg2: []interface{}{},
	},
	{
		name: "10000",
		arg1: []interface{}{},
		arg2: []interface{}{},
	},
	{
		name: "100000",
		arg1: []interface{}{},
		arg2: []interface{}{},
	},
}

func init() {
	for i := 0; i < len(TXorBenchs); i++ {
		k, _ := strconv.Atoi(TXorBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TXorBenchs[i].arg1 = append(TXorBenchs[i].arg1.([]interface{}), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
			TXorBenchs[i].arg2 = append(TXorBenchs[i].arg2.([]interface{}), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
	}
}

func TestXor(t *testing.T) {
	var tests = []TXor{
		{
			name:     "nil",
			arg1:     nil,
			arg2:     nil,
			expected: nil,
		},
		{
			name:     "empty",
			arg1:     []interface{}{},
			arg2:     []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "normal",
			arg1:     []int{1, 2},
			arg2:     []int{1, 4},
			expected: []int{2, 4},
		},
		{
			name:     "type based",
			arg1:     []string{"1", "2"},
			arg2:     []int{1, 4},
			expected: []interface{}{"1", "2", 1, 4},
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := Xor(sample.arg1, sample.arg2)
			if err != nil {
				if sample.expected != nil {
					t.Errorf("got : %v but expected : %v", got, sample.expected)
				}
				return
			}

			if ok, _ := internal.Same(got, sample.expected); !ok {
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
				Xor(sample.arg1, sample.arg2)
			}
		})
	}
}
