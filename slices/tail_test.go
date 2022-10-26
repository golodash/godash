package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

type TTail struct {
	name string
	arr  interface{}
	want interface{}
}

var tTailBenchs = []TTail{
	{
		name: "10",
		arr:  []int{},
	},
	{
		name: "100",
		arr:  []int{},
	},
	{
		name: "1000",
		arr:  []int{},
	},
	{
		name: "10000",
		arr:  []int{},
	},
	{
		name: "100000",
		arr:  []int{},
	},
}

func init() {
	for j := 0; j < len(tTailBenchs); j++ {
		length, _ := strconv.Atoi(tTailBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tTailBenchs[j].arr = append(tTailBenchs[j].arr.([]int), 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		}
	}
}

func TestTail(t *testing.T) {
	var tests = []TTail{
		{
			name: "nil",
			arr:  nil,
			want: nil,
		},
		{
			name: "empty",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "just one item",
			arr:  []int{0},
			want: []int{},
		},
		{
			name: "normal-1",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "normal-2",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			name: "new type",
			arr:  []string{"a", "b", "c"},
			want: []string{"b", "c"},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			defer internal.DeferTestCases(t, subject.want)
			got := Tail(subject.arr)

			if !generals.Same(got, subject.want) {
				t.Errorf("got = %v, wanted = %v", got, subject.want)
				return
			}
		})
	}
}

func BenchmarkTail(b *testing.B) {
	for j := 0; j < len(tTailBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tTailBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Tail(tTailBenchs[j].arr)
			}
		})
	}
}
