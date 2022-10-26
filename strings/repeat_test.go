package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TRepeat struct {
	name     string
	input    string
	count    int
	expected string
}

var TRepeatBenchs = []TRepeat{
	{
		name:  "10",
		input: "a",
		count: 10,
	},
	{
		name:  "100",
		input: "a",
		count: 100,
	},
	{
		name:  "1000",
		input: "a",
		count: 1000,
	},
	{
		name:  "10000",
		input: "a",
		count: 10000,
	},
	// {
	// 	name:  "100000",
	// 	input: "a",
	// 	count: 100000,
	// },
}

func init() {
	for i := 0; i < len(TRepeatBenchs); i++ {
		k, _ := strconv.Atoi(TRepeatBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TRepeatBenchs[i].input += "abcdefghij"
		}
	}
}

func TestRepeat(t *testing.T) {
	var tests = []TRepeat{
		{
			name:     "empty",
			input:    "",
			count:    0,
			expected: "",
		},
		{
			name:     "zero count",
			input:    "abc",
			count:    0,
			expected: "",
		},
		{
			name:     "input empty",
			input:    "",
			count:    2,
			expected: "",
		},
		{
			name:     "normal",
			input:    "abc",
			count:    3,
			expected: "abcabcabc",
		},
		{
			name:     "negative count",
			input:    "abc",
			count:    -1,
			expected: "",
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := Repeat(sample.input, sample.count)

			if !generals.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for _, sample := range TRepeatBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Repeat(sample.input, sample.count)
			}
		})
	}
}
