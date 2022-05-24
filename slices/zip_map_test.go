package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TZipMap struct {
	name     string
	keys     []interface{}
	values   []interface{}
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
			TZipMapBenchs[i].keys = append(TZipMapBenchs[i].keys, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}...)
			TZipMapBenchs[i].values = append(TZipMapBenchs[i].values, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}...)
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
			name:     "default1",
			keys:     []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			values:   []interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"},
			expected: map[interface{}]interface{}{0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F", 6: "G", 7: "H", 8: "I", 9: "J"},
		},
	}

	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := ZipMap(sample.keys, sample.values)
			if err != nil {
				if sample.expected != nil {
					t.Errorf("got : %v but expected : %v, %v", got, sample.expected, err)
				}
				return
			}
			if ok, _ := same(got, sample.expected); !ok {
				t.Errorf("got : %v but expected : %v", got, sample.expected)
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
