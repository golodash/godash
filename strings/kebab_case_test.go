/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 * Copyright (c) 2018 Ma_124, <github.com/Ma124>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package strings

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TKebabCase struct {
	name     string
	input    string
	expected string
}

var TKebabCaseBenchs = []TKebabCase{
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
	for i := 0; i < len(TKebabCaseBenchs); i++ {
		k, _ := strconv.Atoi(TKebabCaseBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TKebabCaseBenchs[i].input += "0msa daewq"
		}
	}
}

func TestKebabCase(t *testing.T) {
	var tests = []TKebabCase{
		{
			name:     "0",
			input:    "testCase",
			expected: "test-case",
		},
		{
			name:     "1",
			input:    "TestCase",
			expected: "test-case",
		},
		{
			name:     "2",
			input:    "Test Case",
			expected: "test-case",
		},
		{
			name:     "3",
			input:    " Test Case",
			expected: "test-case",
		},
		{
			name:     "4",
			input:    "Test Case ",
			expected: "test-case",
		},
		{
			name:     "5",
			input:    " Test Case ",
			expected: "test-case",
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
			expected: "many-many-words",
		},
		{
			name:     "11",
			input:    "manyManyWords",
			expected: "many-many-words",
		},
		{
			name:     "12",
			input:    "AnyKind of-string",
			expected: "any-kind-of-string",
		},
		{
			name:     "13",
			input:    "numbers2and55with000",
			expected: "numbers-2-and-55-with-000",
		},
		{
			name:     "14",
			input:    "JSONData",
			expected: "json-data",
		},
		{
			name:     "15",
			input:    "userID",
			expected: "user-id",
		},
		{
			name:     "16",
			input:    "AAAbbb",
			expected: "aa-abbb",
		},
		{
			name:     "17",
			input:    "1A2",
			expected: "1-a-2",
		},
		{
			name:     "18",
			input:    "A1B",
			expected: "a-1-b",
		},
		{
			name:     "19",
			input:    "A1A2A3",
			expected: "a-1-a-2-a-3",
		},
		{
			name:     "20",
			input:    "A1 A2 A3",
			expected: "a-1-a-2-a-3",
		},
		{
			name:     "21",
			input:    "AB1AB2AB3",
			expected: "ab-1-ab-2-ab-3",
		},
		{
			name:     "22",
			input:    "AB1 AB2 AB3",
			expected: "ab-1-ab-2-ab-3",
		},
		{
			name:     "23",
			input:    "some string",
			expected: "some-string",
		},
		{
			name:     "24",
			input:    " some string",
			expected: "some-string",
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := KebabCase(sample.input)

			if !internal.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkKebabCase(b *testing.B) {
	for _, sample := range TKebabCaseBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				KebabCase(sample.input)
			}
		})
	}
}
