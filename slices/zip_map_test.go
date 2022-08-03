package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TZipMap struct {
	name     string
	keys     interface{}
	values   interface{}
	expected interface{}
}

var TZipMapBenchs = []TZipMap{
	{
		name:   "10",
		keys:   []interface{}{},
		values: []interface{}{},
	},
	{
		name:   "100",
		keys:   []interface{}{},
		values: []interface{}{},
	},
	{
		name:   "1000",
		keys:   []interface{}{},
		values: []interface{}{},
	},
	{
		name:   "10000",
		keys:   []interface{}{},
		values: []interface{}{},
	},
	{
		name:   "100000",
		keys:   []interface{}{},
		values: []interface{}{},
	},
}

func init() {
	for i := 0; i < len(TZipMapBenchs); i++ {
		k, _ := strconv.Atoi(TZipMapBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TZipMapBenchs[i].keys = append(TZipMapBenchs[i].keys.([]interface{}), 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
			TZipMapBenchs[i].values = append(TZipMapBenchs[i].values.([]interface{}), 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
		}
	}
}

func TestZipMap(t *testing.T) {
	var tests = []TZipMap{
		{
			name:     "nil",
			keys:     nil,
			expected: nil,
		},
		{
			name:     "empty",
			keys:     []interface{}{},
			values:   []interface{}{},
			expected: map[interface{}]interface{}{},
		},
		{
			name:     "error",
			keys:     []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			values:   []interface{}{0},
			expected: nil,
		},
		{
			name:     "normal",
			keys:     []int{0, 1, 2, 3},
			values:   []int{0, 11, 22, 33},
			expected: map[int]int{0: 0, 1: 11, 2: 22, 3: 33},
		},
		{
			name:     "type based",
			keys:     []int{0, 1, 2, 3},
			values:   []string{"A", "B", "C", "D"},
			expected: map[int]string{0: "A", 1: "B", 2: "C", 3: "D"},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, sample.expected)
			got := ZipMap(sample.keys, sample.values)

			if !internal.Same(got, sample.expected) {
				t.Errorf("got = %v, wanted = %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkZipMap(b *testing.B) {
	for _, sample := range TZipMapBenchs {
		b.Run(fmt.Sprintf("input_size_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ZipMap(sample.keys, sample.values)
			}
		})
	}
}
