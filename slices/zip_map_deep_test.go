package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TZipMapDeep struct {
	name     string
	keys     []string
	values   interface{}
	expected interface{}
}

var TZipMapDeepBenchs = []TZipMapDeep{
	{
		name:   "10",
		keys:   []string{},
		values: []string{},
	},
	{
		name:   "100",
		keys:   []string{},
		values: []string{},
	},
	{
		name:   "1000",
		keys:   []string{},
		values: []string{},
	},
	{
		name:   "10000",
		keys:   []string{},
		values: []string{},
	},
	{
		name:   "100000",
		keys:   []string{},
		values: []string{},
	},
}

func init() {
	for i := 0; i < len(TZipMapDeepBenchs); i++ {
		k, _ := strconv.Atoi(TZipMapDeepBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TZipMapDeepBenchs[i].keys = append(TZipMapDeepBenchs[i].keys, "1", "2", "3", "4", "5", "6", "7", "8", "9", "0")
			TZipMapDeepBenchs[i].values = append(TZipMapDeepBenchs[i].values.([]string), "1", "2", "3", "4", "5", "6", "7", "8", "9", "0")
		}
	}
}

func TestZipMapDeep(t *testing.T) {
	var tests = []TZipMapDeep{
		{
			name:     "nil",
			keys:     nil,
			values:   nil,
			expected: nil,
		},
		{
			name:     "empty",
			keys:     []string{},
			values:   []string{},
			expected: nil,
		},
		{
			name:     "normal",
			keys:     []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
			values:   []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
			expected: map[string]string{"0": "0", "1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7", "8": "8", "9": "9"},
		},
		{
			name:     "little complicated",
			keys:     []string{"a.bid", "a.b", "b.c"},
			values:   []string{"A", "B", "C"},
			expected: map[string]map[string]string{"a": {"bid": "A", "b": "B"}, "b": {"c": "C"}},
		},
		{
			name:     "little complicated with slices",
			keys:     []string{"a.bid[0]", "a.b[0]", "a.b[1]", "b.c[0]"},
			values:   []string{"A", "B", "C", "D"},
			expected: map[string]map[string]*[]string{"a": {"bid": {"A"}, "b": {"B", "C"}}, "b": {"c": {"D"}}},
		},
		{
			name:     "another with slices",
			keys:     []string{"[0].a", "[0].b", "[3].c"},
			values:   []string{"A", "B", "C"},
			expected: &[]map[string]string{{"a": "A", "b": "B"}, {}, {}, {"c": "C"}},
		},
		{
			name:     "complicated error",
			keys:     []string{"a.bid", "a.b.c[0][1]", "b.c[0]"},
			values:   []string{"A", "B", "C"},
			expected: nil,
		},
		{
			name:     "wrong format error",
			keys:     []string{"]a.bid", "[]]", "b.c[0]"},
			values:   []string{"A", "B", "C"},
			expected: nil,
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got, err := ZipMapDeep(sample.keys, sample.values)

			if ok, _ := internal.Same(got, sample.expected); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, sample.expected, err)
				return
			}
		})
	}
}

func BenchmarkZipMapDeep(b *testing.B) {
	for _, sample := range TZipMapDeepBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ZipMapDeep(sample.keys, sample.values)
			}
		})
	}
}
