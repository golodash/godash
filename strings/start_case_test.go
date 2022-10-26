package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TStartCase struct {
	name     string
	input    string
	expected string
}

var TStartCaseBenchs = []TStartCase{
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
	for i := 0; i < len(TStartCaseBenchs); i++ {
		k, _ := strconv.Atoi(TStartCaseBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TStartCaseBenchs[i].input += "kmsa daewq"
		}
	}
}

func TestStartCase(t *testing.T) {
	var tests = []TStartCase{
		{
			name:     "0",
			input:    "testCase",
			expected: "Test Case",
		},
		{
			name:     "1",
			input:    "TestCase",
			expected: "Test Case",
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
			expected: "Test Case",
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
			expected: "Many Many Words",
		},
		{
			name:     "11",
			input:    "manyManyWords",
			expected: "Many Many Words",
		},
		{
			name:     "12",
			input:    "AnyKind of-string",
			expected: "Any Kind Of String",
		},
		{
			name:     "13",
			input:    "numbers2and55with000",
			expected: "Numbers 2 And 55 With 000",
		},
		{
			name:     "14",
			input:    "JSONData",
			expected: "Json Data",
		},
		{
			name:     "15",
			input:    "userID",
			expected: "User Id",
		},
		{
			name:     "16",
			input:    "AAAbbb",
			expected: "Aa Abbb",
		},
		{
			name:     "17",
			input:    "1A2",
			expected: "1 A 2",
		},
		{
			name:     "18",
			input:    "A1B",
			expected: "A 1 B",
		},
		{
			name:     "19",
			input:    "A1A2A3",
			expected: "A 1 A 2 A 3",
		},
		{
			name:     "20",
			input:    "A1 A2 A3",
			expected: "A 1 A 2 A 3",
		},
		{
			name:     "21",
			input:    "AB1AB2AB3",
			expected: "Ab 1 Ab 2 Ab 3",
		},
		{
			name:     "22",
			input:    "AB1 AB2 AB3",
			expected: "Ab 1 Ab 2 Ab 3",
		},
		{
			name:     "23",
			input:    "some string",
			expected: "Some String",
		},
		{
			name:     "24",
			input:    " some string",
			expected: "Some String",
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := StartCase(sample.input)

			if !generals.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkStartCase(b *testing.B) {
	for _, sample := range TStartCaseBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StartCase(sample.input)
			}
		})
	}
}
