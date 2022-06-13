package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TUnique struct {
	name     string
	arg      []interface{}
	expected []interface{}
}

var TUniqueBenchs = []TUnique{
	{
		name: "10",
		arg:  []interface{}{},
	},
	{
		name: "100",
		arg:  []interface{}{},
	},
	{
		name: "1000",
		arg:  []interface{}{},
	},
	{
		name: "10000",
		arg:  []interface{}{},
	},
	{
		name: "100000",
		arg:  []interface{}{},
	},
}

func init() {
	for i := 0; i < len(TUniqueBenchs); i++ {
		k, _ := strconv.Atoi(TUniqueBenchs[i].name)
		for j := 0; j < k/10; j++ {
			TUniqueBenchs[i].arg = append(TUniqueBenchs[i].arg, []interface{}{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}...)
		}
	}
}

func TestUnique(t *testing.T) {
	var tests = []TUnique{
		{
			name:     "nil",
			arg:      nil,
			expected: nil,
		},
		{
			name:     "empty",
			arg:      []interface{}{},
			expected: []interface{}{},
		},
		{
			name:     "normal",
			arg:      []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "duplicate",
			arg:      []interface{}{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9},
			expected: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, sample := range tests {
		t.Run(sample.name, func(t *testing.T) {
			got, err := Unique(sample.arg)
			if err != nil {
				if sample.expected != nil {
					t.Errorf("got : %v but expected : %v", got, sample.expected)
				}
				return
			}
			if ok, _ := internal.Same(got, sample.expected); !ok {
				t.Errorf("got : %v but expected : %v", got, sample.expected)
				return
			}
		})
	}
}

func BenchmarkUnique(b *testing.B) {
	for _, sample := range TUniqueBenchs {
		b.Run(fmt.Sprintf("input_size_unique_%s", sample.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Unique(sample.arg)
			}
		})
	}
}
