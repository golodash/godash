package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TXorBy struct {
	name     string
	arg1     interface{}
	arg2     interface{}
	expected interface{}
}

var TXorByBenchs = []TXorBy{
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
	for i := 0; i < len(TXorByBenchs); i++ {
		k, _ := strconv.Atoi(TXorByBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TXorByBenchs[i].arg1 = append(TXorByBenchs[i].arg1.([]interface{}), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
			TXorByBenchs[i].arg2 = append(TXorByBenchs[i].arg2.([]interface{}), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
	}
}

func compareXorByTest(value1 interface{}) interface{} {
	return value1
}

func TestXorBy(t *testing.T) {
	var tests = []TXorBy{
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
			arg2:     []int{3, 4, 1},
			expected: []int{2, 3, 4},
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got, err := XorBy(sample.arg1, sample.arg2, compareXorByTest)

			if ok, _ := internal.Same(got, sample.expected); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, sample.expected, err)
				return
			}
		})
	}
}

func BenchmarkXorBy(b *testing.B) {
	for _, sample := range TXorByBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				XorBy(sample.arg1, sample.arg2, compareXorByTest)
			}
		})
	}
}
