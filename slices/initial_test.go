package slices

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type TInitial struct {
	name string
	arr  []interface{}
	want []interface{}
}

var tInitialBenchs = []TInitial{
	{
		name: "10",
		arr:  []interface{}{},
	},
	{
		name: "100",
		arr:  []interface{}{},
	},
	{
		name: "1000",
		arr:  []interface{}{},
	},
	{
		name: "10000",
		arr:  []interface{}{},
	},
	{
		name: "100000",
		arr:  []interface{}{},
	},
}

func init() {
	for j := 0; j < len(tInitialBenchs); j++ {
		length, _ := strconv.Atoi(tInitialBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tInitialBenchs[j].arr = append(tInitialBenchs[j].arr, []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestInitial(t *testing.T) {
	var tests = []TInitial{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  []interface{}{},
			want: []interface{}{},
		},
		{
			name: "just one item",
			arr:  []interface{}{0},
			want: []interface{}{},
		},
		{
			name: "normal",
			arr:  []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "more complex",
			arr:  []interface{}{0, []interface{}{1, 2}, []interface{}{3, 4, 5}, []interface{}{6, 7}, 8, []interface{}{9}},
			want: []interface{}{0, []interface{}{1, 2}, []interface{}{3, 4, 5}, []interface{}{6, 7}, 8},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := Initial(subject.arr)
			if err != nil {
				if subject.want != nil {
					t.Errorf("Initial() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("Initial() got = %v, wanted = %v", got, subject.want)
				return
			}

			check := func(subgot, want []interface{}, function interface{}) {
				for i := 0; i < len(subgot); i++ {
					if gotVal, ok := subgot[i].([]interface{}); ok {
						if wantVal, ok := want[i].([]interface{}); ok {
							for j := 0; j < len(subgot); j++ {
								reflect.ValueOf(function).Call([]reflect.Value{reflect.ValueOf(gotVal), reflect.ValueOf(wantVal), reflect.ValueOf(function)})
							}
						} else {
							t.Errorf("Initial() got = %v, wanted = %v", got, subject.want)
							return
						}

					} else {
						if subgot[i] != want[i] {
							t.Errorf("Initial() got = %v, wanted = %v", got, subject.want)
							return
						}
					}
				}
			}

			check(got, subject.want, check)
		})
	}
}

func BenchmarkInitial(b *testing.B) {
	for j := 0; j < len(tInitialBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tInitialBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Initial(tInitialBenchs[j].arr)
			}
		})
	}
}
