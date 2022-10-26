package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TPad struct {
	name     string
	input    string
	pattern  string
	length   int
	expected string
}

var TPadBenchs = []TPad{
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
	for i := 0; i < len(TPadBenchs); i++ {
		k, _ := strconv.Atoi(TPadBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TPadBenchs[i].input += "kmsa daewq"
		}
	}
}

func TestPad(t *testing.T) {
	var tests = []TPad{
		{
			name:     "0",
			input:    "TestHere",
			length:   16,
			pattern:  "_-",
			expected: "_-_-TestHere_-_-",
		},
		{
			name:     "1",
			input:    "TestHere",
			length:   12,
			pattern:  "_-",
			expected: "_-TestHere_-",
		},
		{
			name:     "2",
			input:    "TestHere",
			length:   11,
			pattern:  "_-",
			expected: "_-TestHere_",
		},
		{
			name:     "3",
			input:    "TestHere",
			length:   10,
			pattern:  "_-",
			expected: "-TestHere_",
		},
		{
			name:     "4",
			input:    "TestHere",
			length:   7,
			pattern:  "_-",
			expected: "estHere",
		},
		{
			name:     "5",
			input:    "TestHere",
			length:   4,
			pattern:  "_-",
			expected: "stHe",
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := Pad(sample.input, sample.length, sample.pattern)

			if !generals.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkPad(b *testing.B) {
	for _, sample := range TPadBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Pad(sample.input, sample.length, sample.pattern)
			}
		})
	}
}
