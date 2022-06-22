package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type (
	TConcat struct {
		name   string
		arr    interface{}
		values interface{}
		want   interface{}
	}
)

var tConcatBenchs = []TConcat{
	{
		name:   "10",
		arr:    []interface{}{},
		values: []interface{}{},
	},
	{
		name:   "100",
		arr:    []interface{}{},
		values: []interface{}{},
	},
	{
		name:   "1000",
		arr:    []interface{}{},
		values: []interface{}{},
	},
	{
		name:   "10000",
		arr:    []interface{}{},
		values: []interface{}{},
	},
	{
		name:   "100000",
		arr:    []interface{}{},
		values: []interface{}{},
	},
}

func init() {
	for j := 0; j < len(tConcatBenchs); j++ {
		length, _ := strconv.Atoi(tConcatBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tConcatBenchs[j].values = append(tConcatBenchs[j].values.([]interface{}), []interface{}{0, nil, 2, false, 4, 5, "", nil, 8, 9}...)
		}
	}
}

func TestConcat(t *testing.T) {
	var tests = []TConcat{
		{
			name:   "nil",
			arr:    nil,
			values: nil,
			want:   nil,
		},
		{
			name:   "empty",
			arr:    []interface{}{},
			values: []interface{}{},
			want:   []interface{}{},
		},
		{
			name:   "normal",
			arr:    []interface{}{0, 1, 2, 3, nil, 5, '6', 0, false, ""},
			values: []interface{}{"45", "e"},
			want:   []interface{}{0, 1, 2, 3, nil, 5, '6', 0, false, "", "45", "e"},
		},
		{
			name:   "add to empty",
			arr:    []interface{}{},
			values: []interface{}{1, 2, 15, "l"},
			want:   []interface{}{1, 2, 15, "l"},
		},
		{
			name:   "more complicated",
			arr:    []interface{}{0, nil, 0, false, nil, "", nil, 0, false, ""},
			values: []interface{}{[]interface{}{0}, ""},
			want:   []interface{}{0, nil, 0, false, nil, "", nil, 0, false, "", 0, ""},
		},
		{
			name:   "type based",
			arr:    []int{0, 1, 2, 3, 4, 5},
			values: []int{6, 7, 8, 9},
			want:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Concat(subject.arr, subject.values)

			if ok := internal.Same(got, subject.want); !ok {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkConcat(b *testing.B) {
	for j := 0; j < len(tConcatBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tConcatBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Concat(tConcatBenchs[j].arr, tConcatBenchs[j].values)
			}
		})
	}
}
