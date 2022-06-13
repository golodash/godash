package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TPull struct {
	name string
	arr  []int
	rems []int
	want []int
}

var tPullBenchs = []TPull{
	{
		name: "10",
		arr:  []int{},
		rems: []int{},
	},
	{
		name: "100",
		arr:  []int{},
		rems: []int{},
	},
	{
		name: "1000",
		arr:  []int{},
		rems: []int{},
	},
	{
		name: "10000",
		arr:  []int{},
		rems: []int{},
	},
	{
		name: "100000",
		arr:  []int{},
		rems: []int{},
	},
}

func init() {
	for j := 0; j < len(tPullBenchs); j++ {
		length, _ := strconv.Atoi(tPullBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tPullBenchs[j].arr = append(tPullBenchs[j].arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
			tPullBenchs[j].rems = append(tPullBenchs[j].rems, 0+(i*10), 1+(i*10), 2+(i*10), 3+(i*10), 4+(i*10), 5+(i*10))
		}
	}
}

func TestPull(t *testing.T) {
	var tests = []TPull{
		{
			name: "nil",
			arr:  nil,
			rems: nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  []int{},
			rems: []int{},
			want: []int{},
		},
		{
			name: "normal",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			rems: []int{0, 1, 2},
			want: []int{3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "shuffle",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			rems: []int{0, 4, 9, 4, 9, 0},
			want: []int{1, 2, 3, 5, 6, 7, 8},
		},
		{
			name: "all remove",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			rems: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []int{},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := Pull(subject.arr, subject.rems)

			if err != nil {
				if subject.want != nil {
					t.Errorf("Pull() got = %v, want = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("Pull() got = %v, want = %v", got, subject.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if got[i].(int) != subject.want[i] {
					t.Errorf("Pull() got = %v, want = %v", got, subject.want)
					return
				}
			}
		})
	}
}

func BenchmarkPull(b *testing.B) {
	for j := 0; j < len(tPullBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tPullBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Pull(tPullBenchs[j].arr, tPullBenchs[j].rems)
			}
		})
	}
}
