package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TJoin struct {
	name     string
	arg1     interface{}
	arg2     string
	expected interface{}
}

var TJoinBenchs = []TJoin{
	{
		name: "10",
		arg1: []string{},
	},
	{
		name: "100",
		arg1: []string{},
	},
	{
		name: "1000",
		arg1: []string{},
	},
	{
		name: "10000",
		arg1: []string{},
	},
	{
		name: "100000",
		arg1: []string{},
	},
}

func init() {
	for i := 0; i < len(TJoinBenchs); i++ {
		k, _ := strconv.Atoi(TJoinBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TJoinBenchs[i].arg1 = append(TJoinBenchs[i].arg1.([]string), []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}...)
		}

	}
}

func TestJoin(t *testing.T) {
	var tests = []TJoin{
		{
			name:     "nil",
			arg1:     nil,
			arg2:     "-",
			expected: nil,
		},
		{
			name:     "empty",
			arg1:     []string{},
			arg2:     "<",
			expected: "",
		},
		{
			name:     "normal",
			arg1:     []string{"A", "B", "C", "D"},
			arg2:     "-",
			expected: "A-B-C-D",
		},
		{
			name:     "type based",
			arg1:     []int{1, 2, 3, 4, 5, 6},
			arg2:     ", ",
			expected: "1, 2, 3, 4, 5, 6",
		},
	}
	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.expected)
			got := Join(subject.arg1, subject.arg2)

			if !generals.Same(got, subject.expected) {
				t.Errorf("got = %v, wanted = %v", got, subject.expected)
				return
			}
		})
	}
}

func BenchmarkJoin(b *testing.B) {
	for _, sample := range TJoinBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Join(sample.arg1, sample.arg2)
			}
		})
	}
}
