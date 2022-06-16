package slices

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golodash/godash/internal"
)

type TIndexOf struct {
	name  string
	arr   interface{}
	value interface{}
	want  interface{}
}

var tIndexOfBenchs = []TIndexOf{
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
	for j := 0; j < len(tIndexOfBenchs); j++ {
		length, _ := strconv.Atoi(tIndexOfBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tIndexOfBenchs[j].arr = append(tIndexOfBenchs[j].arr.([]int), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
		tIndexOfBenchs[j].arr = append(tIndexOfBenchs[j].arr.([]int), 10)
		tIndexOfBenchs[j].value = 10
	}
}

func TestIndexOf(t *testing.T) {
	var tests = []TIndexOf{
		{
			name:  "nil",
			arr:   nil,
			value: -1,
			want:  -1,
		},
		{
			name:  "empty",
			arr:   []int{},
			value: -1,
			want:  -1,
		},
		{
			name:  "normal",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 4,
			want:  4,
		},
		{
			name:  "does not exist",
			arr:   []int{0, 1, 2, 3, 4, 6, 7, 8, 9},
			value: 10,
			want:  -1,
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			got, err := IndexOf(subject.arr, subject.value, 0)
			if err != nil {
				if subject.want != -1 {
					t.Errorf("got = %v, wanted = %v, err = %v", got, subject.want, err)
				}
				return
			}

			if ok, _ := internal.Same(got, subject.want); !ok {
				t.Errorf("got = %v, wanted = %v, err = %v", got, subject.want, err)
				return
			}
		})
	}
}

func BenchmarkIndexOf(b *testing.B) {
	for j := 0; j < len(tIndexOfBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tIndexOfBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IndexOf(tIndexOfBenchs[j].arr, tIndexOfBenchs[j].value, 0)
			}
		})
	}
}
