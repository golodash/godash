package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TDropRight struct {
	name string
	arr  []int
	num  int
	want []int
}

var tDropRightBenchs = []TDropRight{
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
	for j := 0; j < len(tDropRightBenchs); j++ {
		length, _ := strconv.Atoi(tDropRightBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tDropRightBenchs[j].arr = append(tDropRightBenchs[j].arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestDropRight(t *testing.T) {
	var tests = []TDropRight{
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
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "half",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			num:  5,
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "most",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			num:  8,
			want: []int{0, 1},
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
			got, err := DropRight(subject.arr, subject.num)
			if err != nil {
				if subject.want != nil {
					t.Errorf("DropRight() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("DropRight() got = %v, wanted = %v", got, subject.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if got[i] != subject.want[i] {
					t.Errorf("DropRight() got = %v, wanted = %v", got, subject.want)
					return
				}
			}
		})
	}
}

func BenchmarkDropRight(b *testing.B) {
	for j := 0; j < len(tDropRightBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tDropRightBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DropRight(tDropRightBenchs[j].arr, tDropRightBenchs[j].num)
			}
		})
	}
}
