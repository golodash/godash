package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TPadEnd struct {
	name     string
	input    string
	pattern  string
	length   int
	expected string
}

var TPadEndBenchs = []TPadEnd{
	{
		name:    "10",
		input:   "",
		pattern: "|-",
		length:  20,
	},
	{
		name:    "100",
		input:   "",
		pattern: "|-",
		length:  200,
	},
	{
		name:    "1000",
		input:   "",
		pattern: "|-",
		length:  2000,
	},
	{
		name:    "10000",
		input:   "",
		pattern: "|-",
		length:  20000,
	},
	{
		name:    "100000",
		input:   "",
		pattern: "|-",
		length:  200000,
	},
}

func init() {
	for i := 0; i < len(TPadEndBenchs); i++ {
		k, _ := strconv.Atoi(TPadEndBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TPadEndBenchs[i].input += "kmsa daewq"
		}
	}
}

func TestPadEnd(t *testing.T) {
	var tests = []TPadEnd{
		{
			name:     "0",
			input:    "This is a freaking test just for the fun of it!??",
			length:   53,
			pattern:  "_-",
			expected: "This is a freaking test just for the fun of it!??_-_-",
		},
		{
			name:     "1",
			input:    "This is a freaking test just for the fun of it!??",
			length:   54,
			pattern:  "_-",
			expected: "This is a freaking test just for the fun of it!??_-_-_",
		},
		{
			name:     "2",
			input:    "This is a freaking test just for the fun of it!??",
			length:   43,
			pattern:  "_-",
			expected: "This is a freaking test just for the fun of",
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := PadEnd(sample.input, sample.length, sample.pattern)

			if !internal.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkPadEnd(b *testing.B) {
	for _, sample := range TPadEndBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PadEnd(sample.input, sample.length, sample.pattern)
			}
		})
	}
}
