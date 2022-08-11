package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TEndsWith struct {
	name     string
	input    string
	ends     string
	expected bool
}

var TEndsWithBenchs = []TEndsWith{
	{
		name:  "10",
		input: "",
		ends:  "j",
	},
	{
		name:  "100",
		input: "",
		ends:  "j",
	},
	{
		name:  "1000",
		input: "",
		ends:  "j",
	},
	{
		name:  "10000",
		input: "",
		ends:  "j",
	},
	{
		name:  "100000",
		input: "",
		ends:  "j",
	},
}

func init() {
	for i := 0; i < len(TEndsWithBenchs); i++ {
		k, _ := strconv.Atoi(TEndsWithBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TEndsWithBenchs[i].input += "abcdefghij"
		}
	}
}

func TestEndsWith(t *testing.T) {
	var tests = []TEndsWith{
		{
			name:     "true",
			input:    "test",
			ends:     "st",
			expected: true,
		},
		{
			name:     "false",
			input:    "another_test",
			ends:     "es",
			expected: false,
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := EndsWith(sample.input, sample.ends)

			if !internal.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkEndsWith(b *testing.B) {
	for _, sample := range TEndsWithBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				EndsWith(sample.input, sample.ends)
			}
		})
	}
}
