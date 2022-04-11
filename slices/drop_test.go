package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TDrop struct {
	name string
	arr  []int
	num  int
	want []int
}

var tDropBenchs = []TDrop{
	{
		name: "10",
		arr:  []int{},
		num:  5,
	},
	{
		name: "100",
		arr:  []int{},
		num:  50,
	},
	{
		name: "1000",
		arr:  []int{},
		num:  500,
	},
	{
		name: "10000",
		arr:  []int{},
		num:  5000,
	},
	{
		name: "100000",
		arr:  []int{},
		num:  50000,
	},
	{
		name: "1000000",
		arr:  []int{},
		num:  500000,
	},
}

func init() {
	for j := 0; j < len(tDropBenchs); j++ {
		length, _ := strconv.Atoi(tDropBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tDropBenchs[j].arr = append(tDropBenchs[j].arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestDrop(t *testing.T) {
	var tests = []TDrop{
		{
			name: "nil",
			arr:  nil,
			num:  0,
			want: nil,
		},
		{
			name: "empty",
			arr:  []int{},
			num:  0,
			want: []int{},
		},
		{
			name: "normal",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			num:  1,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "half",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			num:  5,
			want: []int{5, 6, 7, 8, 9},
		},
		{
			name: "most",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			num:  8,
			want: []int{8, 9},
		},
		{
			name: "all",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			num:  10,
			want: []int{},
		},
		{
			name: "more that length",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			num:  150,
			want: nil,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := Drop(subject.arr, subject.num)
			if err != nil {
				if subject.want != nil {
					t.Errorf("Drop() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("Drop() got = %v, wanted = %v", got, subject.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if got[i] != subject.want[i] {
					t.Errorf("Drop() got = %v, wanted = %v", got, subject.want)
					return
				}
			}
		})
	}
}

func BenchmarkDrop(b *testing.B) {
	for j := 0; j < len(tDropBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tDropBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Drop(tDropBenchs[j].arr, tDropBenchs[j].num)
			}
		})
	}
}
