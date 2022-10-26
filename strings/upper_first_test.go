package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TUpperFirst struct {
	name     string
	input    string
	expected string
}

var TUpperFirstBenchs = []TUpperFirst{
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
	for i := 0; i < len(TUpperFirstBenchs); i++ {
		k, _ := strconv.Atoi(TUpperFirstBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TUpperFirstBenchs[i].input += "kmsa daewq"
		}
	}
}

func TestUpperFirst(t *testing.T) {
	var tests = []TUpperFirst{
		{
			name:     "0",
			input:    "testCase",
			expected: "TestCase",
		},
		{
			name:     "1",
			input:    "TestCase",
			expected: "TestCase",
		},
		{
			name:     "2",
			input:    "Test Case",
			expected: "Test Case",
		},
		{
			name:     "3",
			input:    " Test Case",
			expected: "Test Case",
		},
		{
			name:     "4",
			input:    "Test Case ",
			expected: "Test Case",
		},
		{
			name:     "5",
			input:    " Test Case ",
			expected: "Test Case",
		},
		{
			name:     "6",
			input:    "test",
			expected: "Test",
		},
		{
			name:     "7",
			input:    "test-case",
			expected: "Test-case",
		},
		{
			name:     "8",
			input:    "Test",
			expected: "Test",
		},
		{
			name:     "9",
			input:    "",
			expected: "",
		},
		{
			name:     "10",
			input:    "ManyManyWords",
			expected: "ManyManyWords",
		},
		{
			name:     "11",
			input:    "manyManyWords",
			expected: "ManyManyWords",
		},
		{
			name:     "12",
			input:    "AnyKind of-string",
			expected: "AnyKind of-string",
		},
		{
			name:     "13",
			input:    "numbers2and55with000",
			expected: "Numbers2and55with000",
		},
		{
			name:     "14",
			input:    "JSONData",
			expected: "JSONData",
		},
		{
			name:     "15",
			input:    "userID",
			expected: "UserID",
		},
		{
			name:     "16",
			input:    "AAAbbb",
			expected: "AAAbbb",
		},
		{
			name:     "17",
			input:    "1A2",
			expected: "1A2",
		},
		{
			name:     "18",
			input:    "A1B",
			expected: "A1B",
		},
		{
			name:     "19",
			input:    "A1A2A3",
			expected: "A1A2A3",
		},
		{
			name:     "20",
			input:    "A1 A2 A3",
			expected: "A1 A2 A3",
		},
		{
			name:     "21",
			input:    "AB1AB2AB3",
			expected: "AB1AB2AB3",
		},
		{
			name:     "22",
			input:    "AB1 AB2 AB3",
			expected: "AB1 AB2 AB3",
		},
		{
			name:     "23",
			input:    "some string",
			expected: "Some string",
		},
		{
			name:     "24",
			input:    " some string",
			expected: "Some string",
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := UpperFirst(sample.input)

			if !generals.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkUpperFirst(b *testing.B) {
	for _, sample := range TUpperFirstBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UpperFirst(sample.input)
			}
		})
	}
}
