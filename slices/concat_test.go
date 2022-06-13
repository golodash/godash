package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type (
	TConcat struct {
		name   string
		arr    []interface{}
		values []interface{}
		want   []interface{}
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
			tConcatBenchs[j].values = append(tConcatBenchs[j].values, []interface{}{0, nil, 2, false, 4, 5, "", nil, 8, 9}...)
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
			name:   "a bit weird again",
			arr:    []interface{}{0, nil, 0, false, nil, "", nil, 0, false, ""},
			values: []interface{}{[]interface{}{0}, ""},
			want:   []interface{}{0, nil, 0, false, nil, "", nil, 0, false, "", 0, ""},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := Concat(subject.arr, subject.values...)
			if err != nil {
				if subject.want != nil {
					t.Errorf("Concat() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("Concat() got = %v, wanted = %v", got, subject.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if got[i] != subject.want[i] {
					t.Errorf("Concat() got = %v, wanted = %v", got, subject.want)
					return
				}
			}
		})
	}
}

func BenchmarkConcat(b *testing.B) {
	for j := 0; j < len(tConcatBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tConcatBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Concat(tConcatBenchs[j].arr, tConcatBenchs[j].values...)
			}
		})
	}
}
