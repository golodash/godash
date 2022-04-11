package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TDropBy struct {
	name string
	arr  []int
	want []int
}

var tDropByBenchs = []TDropBy{
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
	{
		name: "1000000",
		arr:  []int{},
	},
}

func removeDropByTest(intput int) bool {
	if intput%2 == 0 {
		return true
	}
	return false
}

func init() {
	for j := 0; j < len(tDropByBenchs); j++ {
		length, _ := strconv.Atoi(tDropByBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tDropByBenchs[j].arr = append(tDropByBenchs[j].arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestDropBy(t *testing.T) {
	var tests = []TDropBy{
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
			name: "normal",
			arr:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []int{1, 3, 5, 7, 9},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := DropBy(subject.arr, removeDropByTest)
			if err != nil {
				if subject.want != nil && got != nil {
					t.Errorf("DropBy() got = %v, wanted = %v", got, subject.want)
				}
				return
			}

			if len(got) != len(subject.want) {
				t.Errorf("DropBy() got = %v, wanted = %v", got, subject.want)
				return
			}

			for i := 0; i < len(got); i++ {
				if got[i] != subject.want[i] {
					t.Errorf("DropBy() got = %v, wanted = %v", got, subject.want)
					return
				}
			}
		})
	}
}

func BenchmarkDropBy(b *testing.B) {
	for j := 0; j < len(tDropByBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tDropByBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DropBy(tDropByBenchs[j].arr, removeDropByTest)
			}
		})
	}
}
