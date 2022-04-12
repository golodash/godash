package slices

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

type TFlatten struct {
	name string
	arr  []interface{}
	want []interface{}
}

var tFlattenBenchs = []TFlatten{
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
	{
		name: "1000000",
		arr:  []interface{}{},
	},
}

func init() {
	for j := 0; j < len(tFlattenBenchs); j++ {
		length, _ := strconv.Atoi(tFlattenBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tFlattenBenchs[j].arr = append(tFlattenBenchs[j].arr, []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestFlatten(t *testing.T) {
	var tests = []TFlatten{
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
			name: "none",
			arr:  []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "normal",
			arr:  []interface{}{0, []interface{}{1, 2}, []interface{}{3, 4, 5}, []interface{}{6, 7}, 8, 9},
			want: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "more layer",
			arr:  []interface{}{[]interface{}{0, []interface{}{1, 2}}, 3, []interface{}{4, 5, []interface{}{6, 7}, 8}, []interface{}{9}},
			want: []interface{}{0, []interface{}{1, 2}, 3, 4, 5, []interface{}{6, 7}, 8, 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := Flatten(subject.arr)
			if err != nil {
				if subject.want != nil {
					t.Errorf("Flatten() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("Flatten() got = %v, wanted = %v", got, subject.want)
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
							t.Errorf("Flatten() got = %v, wanted = %v", got, subject.want)
							return
						}

					} else {
						if subgot[i] != want[i] {
							t.Errorf("Flatten() got = %v, wanted = %v", got, subject.want)
							return
						}
					}
				}
			}

			check(got, subject.want, check)
		})
	}
}

func BenchmarkFlatten(b *testing.B) {
	for j := 0; j < len(tFlattenBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tFlattenBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Flatten(tFlattenBenchs[j].arr)
			}
		})
	}
}
