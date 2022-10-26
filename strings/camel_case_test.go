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

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TCamelCase struct {
	name     string
	input    string
	expected string
}

var TCamelCaseBenchs = []TCamelCase{
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
	for i := 0; i < len(TCamelCaseBenchs); i++ {
		k, _ := strconv.Atoi(TCamelCaseBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TCamelCaseBenchs[i].input += "0msa daewq"
		}
	}
}

func TestCamelCase(t *testing.T) {
	var tests = []TCamelCase{
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
			expected: "testCase",
		},
		{
			name:     "3",
			input:    " Test Case",
			expected: "testCase",
		},
		{
			name:     "4",
			input:    "Test Case ",
			expected: "testCase",
		},
		{
			name:     "5",
			input:    " Test Case ",
			expected: "testCase",
		},
		{
			name:     "6",
			input:    "test",
			expected: "test",
		},
		{
			name:     "7",
			input:    "test-case",
			expected: "testCase",
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
			expected: "anyKindOfString",
		},
		{
			name:     "13",
			input:    "numbers2and55with000",
			expected: "numbers2And55With000",
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
			input:    "AAA_bbb",
			expected: "aAABbb",
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
			expected: "a1A2A3",
		},
		{
			name:     "21",
			input:    "AB1AB2AB3",
			expected: "aB1AB2AB3",
		},
		{
			name:     "22",
			input:    "AB1 AB2 AB3",
			expected: "aB1AB2AB3",
		},
		{
			name:     "23",
			input:    "some string",
			expected: "someString",
		},
		{
			name:     "24",
			input:    " some string",
			expected: "someString",
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := CamelCase(sample.input)

			if !generals.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkCamelCase(b *testing.B) {
	for _, sample := range TCamelCaseBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CamelCase(sample.input)
			}
		})
	}
}
