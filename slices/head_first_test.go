package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type THead struct {
	name string
	arr  []interface{}
	want interface{}
}

var tHeadBenchs = []THead{
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
	for j := 0; j < len(tHeadBenchs); j++ {
		length, _ := strconv.Atoi(tHeadBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tHeadBenchs[j].arr = append(tHeadBenchs[j].arr, []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestHead(t *testing.T) {
	tests := []THead{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  []interface{}{},
			want: nil,
		},
		{
			name: "normal",
			arr:  []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 0,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := Head(subject.arr)
			if err != nil {
				if subject.want != nil {
					t.Errorf("Head() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if got != subject.want {
				t.Errorf("Head() got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkHead(b *testing.B) {
	for j := 0; j < len(tHeadBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tHeadBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Head(tHeadBenchs[j].arr)
			}
		})
	}
}
