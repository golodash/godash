package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TJoin struct {
	name     string
	arg1     []string
	arg2     string
	expected string
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
			TJoinBenchs[i].arg1 = append(TJoinBenchs[i].arg1, []string{"Hello", "Are", "You", "A", "Fucking", "Mother", "Fucker", "?", "Huh", ":)"}...)
		}

	}
}

func TestJoin(t *testing.T) {
	var tests = []TJoin{
		{
			name:     "nil",
			arg1:     nil,
			arg2:     "-",
			expected: "",
		},
		{
			name:     "empty",
			arg1:     []string{},
			arg2:     "<",
			expected: "",
		},
		{
			name:     "default",
			arg1:     []string{"Hello", "Mother", "Fucker", ":)"},
			arg2:     "-",
			expected: "Hello-Mother-Fucker-:)",
		},
		{
			name:     "default1",
			arg1:     []string{"A", "B", "C", "D", "E", "F", "U"},
			arg2:     "<<",
			expected: "A<<B<<C<<D<<E<<F<<U",
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := Join(sample.arg1, sample.arg2)
			if err != nil {
				if sample.expected != "" {
					t.Errorf("got : %v but expected : %v", got, sample.expected)
				}
				return
			}
			if got != sample.expected {
				t.Errorf("got : %v but expected : %v", got, sample.expected)
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
