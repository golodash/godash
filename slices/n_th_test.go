package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TNth struct {
	name     string
	arg1     interface{}
	arg2     int
	expected interface{}
}

var TNthBenchs = []TNth{
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
	for i := 0; i < len(TNthBenchs); i++ {
		k, _ := strconv.Atoi(TNthBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TNthBenchs[i].arg1 = append(TNthBenchs[i].arg1.([]interface{}), []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestNth(t *testing.T) {
	tests := []TNth{
		{
			name:     "nil",
			arg1:     nil,
			arg2:     4,
			expected: nil,
		},
		{
			name:     "empty",
			arg1:     []interface{}{},
			arg2:     -1,
			expected: nil,
		},
		{
			name:     "default",
			arg1:     []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			arg2:     5,
			expected: 6,
		},
		{
			name:     "default1",
			arg1:     []interface{}{"a", "b", "c", "d", "e", "f", "u"},
			arg2:     -2,
			expected: "f",
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := Nth(sample.arg1, sample.arg2)

			if ok := internal.Same(got, sample.expected); !ok {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkNth(b *testing.B) {
	for _, sample := range TNthBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Nth(sample.arg1, sample.arg2)
			}
		})
	}
}
