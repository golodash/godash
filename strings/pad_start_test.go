package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TPadStart struct {
	name     string
	input    string
	pattern  string
	length   int
	expected string
}

var TPadStartBenchs = []TPadStart{
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
	for i := 0; i < len(TPadStartBenchs); i++ {
		k, _ := strconv.Atoi(TPadStartBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TPadStartBenchs[i].input += "kmsa daewq"
		}
	}
}

func TestPadStart(t *testing.T) {
	var tests = []TPadStart{
		{
			name:     "0",
			input:    "This is a freaking test just for the fun of it!??",
			length:   53,
			pattern:  "_-",
			expected: "_-_-This is a freaking test just for the fun of it!??",
		},
		{
			name:     "1",
			input:    "This is a freaking test just for the fun of it!??",
			length:   54,
			pattern:  "_-",
			expected: "-_-_-This is a freaking test just for the fun of it!??",
		},
		{
			name:     "2",
			input:    "This is a freaking test just for the fun of it!??",
			length:   44,
			pattern:  "_-",
			expected: "is a freaking test just for the fun of it!??",
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := PadStart(sample.input, sample.length, sample.pattern)

			if !internal.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkPadStart(b *testing.B) {
	for _, sample := range TPadStartBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PadStart(sample.input, sample.length, sample.pattern)
			}
		})
	}
}
