package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TLowerFirst struct {
	name     string
	input    string
	expected string
}

var TLowerFirstBenchs = []TLowerFirst{
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
	for i := 0; i < len(TLowerFirstBenchs); i++ {
		k, _ := strconv.Atoi(TLowerFirstBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TLowerFirstBenchs[i].input += "kmsa daewq"
		}
	}
}

func TestLowerFirst(t *testing.T) {
	var tests = []TLowerFirst{
		{
			name:     "0",
			input:    "testCase",
			expected: "testCase",
		},
		{
			name:     "1",
			input:    "TestCase",
			expected: "testCase",
		},
		{
			name:     "2",
			input:    "Test Case",
			expected: "test Case",
		},
		{
			name:     "3",
			input:    " Test Case",
			expected: "test Case",
		},
		{
			name:     "4",
			input:    "Test Case ",
			expected: "test Case",
		},
		{
			name:     "5",
			input:    " Test Case ",
			expected: "test Case",
		},
		{
			name:     "6",
			input:    "test",
			expected: "test",
		},
		{
			name:     "7",
			input:    "test-case",
			expected: "test-case",
		},
		{
			name:     "8",
			input:    "Test",
			expected: "test",
		},
		{
			name:     "9",
			input:    "",
			expected: "",
		},
		{
			name:     "10",
			input:    "ManyManyWords",
			expected: "manyManyWords",
		},
		{
			name:     "11",
			input:    "manyManyWords",
			expected: "manyManyWords",
		},
		{
			name:     "12",
			input:    "AnyKind of-string",
			expected: "anyKind of-string",
		},
		{
			name:     "13",
			input:    "numbers2and55with000",
			expected: "numbers2and55with000",
		},
		{
			name:     "14",
			input:    "JSONData",
			expected: "jSONData",
		},
		{
			name:     "15",
			input:    "userID",
			expected: "userID",
		},
		{
			name:     "16",
			input:    "AAAbbb",
			expected: "aAAbbb",
		},
		{
			name:     "17",
			input:    "1A2",
			expected: "1A2",
		},
		{
			name:     "18",
			input:    "A1B",
			expected: "a1B",
		},
		{
			name:     "19",
			input:    "A1A2A3",
			expected: "a1A2A3",
		},
		{
			name:     "20",
			input:    "A1 A2 A3",
			expected: "a1 A2 A3",
		},
		{
			name:     "21",
			input:    "AB1AB2AB3",
			expected: "aB1AB2AB3",
		},
		{
			name:     "22",
			input:    "AB1 AB2 AB3",
			expected: "aB1 AB2 AB3",
		},
		{
			name:     "23",
			input:    "some string",
			expected: "some string",
		},
		{
			name:     "24",
			input:    " some string",
			expected: "some string",
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := LowerFirst(sample.input)

			if !internal.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkLowerFirst(b *testing.B) {
	for _, sample := range TLowerFirstBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LowerFirst(sample.input)
			}
		})
	}
}
