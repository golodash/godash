package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TTruncate struct {
	name     string
	input    string
	length   int
	expected string
}

var TTruncateBenchs = []TTruncate{
	{
		name:   "10",
		input:  "",
		length: 9,
	},
	{
		name:   "100",
		input:  "",
		length: 99,
	},
	{
		name:   "1000",
		input:  "",
		length: 999,
	},
	{
		name:   "10000",
		input:  "",
		length: 9999,
	},
	{
		name:   "100000",
		input:  "",
		length: 99999,
	},
}

func init() {
	for i := 0; i < len(TTruncateBenchs); i++ {
		k, _ := strconv.Atoi(TTruncateBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TTruncateBenchs[i].input += "kmsa daewq"
		}
	}
}

var separators = []rune{' ', '.', '?', '!'}

func TestTruncate(t *testing.T) {
	var tests = []TTruncate{
		{
			name:     "0",
			input:    "  This is a freaking test just for the fun of it!??",
			length:   20,
			expected: "This is a freaking...",
		},
		{
			name:     "1",
			input:    "  This is a freaking test just for the fun of it!??",
			length:   100,
			expected: "This is a freaking test just for the fun of it!??",
		},
		{
			name:     "2",
			input:    "  This is a freaking test just for the fun of it!??",
			length:   6,
			expected: "This...",
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := Truncate(sample.input, sample.length, separators, "...")

			if !internal.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkTruncate(b *testing.B) {
	for _, sample := range TTruncateBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Truncate(sample.input, sample.length, separators, "...")
			}
		})
	}
}
