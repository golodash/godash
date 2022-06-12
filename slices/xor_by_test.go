package slices

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type TXorBy struct {
	name     string
	arg      []interface{}
	expected interface{}
}

var TXorByBenchs = []TXorBy{
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
	for i := 0; i < len(TXorByBenchs); i++ {
		k, _ := strconv.Atoi(TXorByBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TXorByBenchs[i].arg = append(TXorByBenchs[i].arg, []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		}
	}
}

func compareXorByTest(value1, value2 interface{}) bool {
	v1 := reflect.ValueOf(value1).Int()
	v2 := reflect.ValueOf(value2).Int()

	return v1 == v2
}

func TestXorBy(t *testing.T) {
	var tests = []TXorBy{
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
			got, err := XorBy(compareXorByTest, sample.arg...)
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

func BenchmarkXorBy(b *testing.B) {
	for _, sample := range TXorByBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				XorBy(compareXorByTest, sample.arg...)
			}
		})
	}
}
