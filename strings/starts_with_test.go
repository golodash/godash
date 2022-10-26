package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TStartsWith struct {
	name     string
	input    string
	starts   string
	expected bool
}

var TStartsWithBenchs = []TStartsWith{
	{
		name:   "10",
		input:  "",
		starts: "j",
	},
	{
		name:   "100",
		input:  "",
		starts: "j",
	},
	{
		name:   "1000",
		input:  "",
		starts: "j",
	},
	{
		name:   "10000",
		input:  "",
		starts: "j",
	},
	{
		name:   "100000",
		input:  "",
		starts: "j",
	},
}

func init() {
	for i := 0; i < len(TStartsWithBenchs); i++ {
		k, _ := strconv.Atoi(TStartsWithBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TStartsWithBenchs[i].input += "abcdefghij"
		}
	}
}

func TestStartsWith(t *testing.T) {
	var tests = []TStartsWith{
		{
			name:     "empty",
			input:    "",
			starts:   "",
			expected: true,
		},
		{
			name:     "starts empty",
			input:    "test",
			starts:   "",
			expected: true,
		},
		{
			name:     "input empty",
			input:    "",
			starts:   "te",
			expected: false,
		},
		{
			name:     "true",
			input:    "test",
			starts:   "te",
			expected: true,
		},
		{
			name:     "false",
			input:    "another_test",
			starts:   "no",
			expected: false,
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := StartsWith(sample.input, sample.starts)

			if !generals.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkStartsWith(b *testing.B) {
	for _, sample := range TStartsWithBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StartsWith(sample.input, sample.starts)
			}
		})
	}
}
