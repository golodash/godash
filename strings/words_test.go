package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TWords struct {
	name     string
	input    string
	expected []string
}

var TWordsBenchs = []TWords{
	{
		name:  "10",
		input: "",
	},
	{
		name:  "100",
		input: "",
	},
	{
		name:  "1000",
		input: "",
	},
	{
		name:  "10000",
		input: "",
	},
	{
		name:  "100000",
		input: "",
	},
}

func init() {
	for i := 0; i < len(TWordsBenchs); i++ {
		k, _ := strconv.Atoi(TWordsBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TWordsBenchs[i].input += "kmsa daewq"
		}
	}
}

func TestWords(t *testing.T) {
	var tests = []TWords{
		{
			name:     "0",
			input:    "test      Case",
			expected: []string{"test", "Case"},
		},
		{
			name:     "1",
			input:    "Test, Case",
			expected: []string{"Test", "Case"},
		},
		{
			name:     "2",
			input:    "[Test] Case",
			expected: []string{"Test", "Case"},
		},
		{
			name:     "3",
			input:    " Test Case?",
			expected: []string{"Test", "Case"},
		},
		{
			name:     "4",
			input:    "Test Case !",
			expected: []string{"Test", "Case"},
		},
		{
			name:     "5",
			input:    " Test Case ",
			expected: []string{"Test", "Case"},
		},
		{
			name:     "6",
			input:    "[Test, Case]",
			expected: []string{"Test", "Case"},
		},
		{
			name:     "7",
			input:    "Test-Case",
			expected: []string{"Test", "Case"},
		},
		{
			name:     "8",
			input:    "This is a test text.",
			expected: []string{"This", "is", "a", "test", "text"},
		},
		{
			name:     "9",
			input:    "",
			expected: []string{},
		},
		{
			name:     "10",
			input:    "!Important!",
			expected: []string{"Important"},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := Words(sample.input)

			if !generals.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkWords(b *testing.B) {
	for _, sample := range TWordsBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Words(sample.input)
			}
		})
	}
}
