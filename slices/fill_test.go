package slices

import (
	"fmt"
	"strconv"
	"testing"
)

type TFill struct {
	name  string
	arr   []int
	value int
	start int
	end   int
	want  []int
}

var tFillBenchs = []TFill{
	{
		name:  "10",
		arr:   []int{},
		value: 0,
		start: 0,
		end:   -1,
	},
	{
		name:  "100",
		arr:   []int{},
		value: 0,
		start: 0,
		end:   -1,
	},
	{
		name:  "1000",
		arr:   []int{},
		value: 0,
		start: 0,
		end:   -1,
	},
	{
		name:  "10000",
		arr:   []int{},
		value: 0,
		start: 0,
		end:   -1,
	},
	{
		name:  "100000",
		arr:   []int{},
		value: 0,
		start: 0,
		end:   -1,
	},
	{
		name:  "1000000",
		arr:   []int{},
		value: 0,
		start: 0,
		end:   -1,
	},
}

func init() {
	for j := 0; j < len(tFillBenchs); j++ {
		length, _ := strconv.Atoi(tFillBenchs[j].name)
		for i := 0; i < length/10; i++ {
			tFillBenchs[j].arr = append(tFillBenchs[j].arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}...)
		}
	}
}

func TestFill(t *testing.T) {
	var tests = []TFill{
		{
			name:  "nil",
			arr:   nil,
			value: 0,
			start: 0,
			end:   -1,
			want:  nil,
		},
		{
			name:  "error",
			arr:   nil,
			value: 0,
			start: -1,
			end:   -1,
			want:  nil,
		},
		{
			name:  "empty",
			arr:   []int{},
			value: 0,
			start: 0,
			end:   -1,
			want:  []int{},
		},
		{
			name:  "normal",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 0,
			start: 0,
			end:   3,
			want:  []int{0, 0, 0, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "all remove",
			arr:   []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			value: 0,
			start: 0,
			end:   -1,
			want:  []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, subject := range tests {
		t.Run(subject.name, func(t *testing.T) {
			err := Fill(subject.arr, subject.value, subject.start, subject.end)

			if err != nil {
				if subject.want != nil && subject.arr != nil {
					t.Errorf("Fill() got = %v, wanted = %v", subject.arr, subject.want)
				}
				return
			}

			if len(subject.arr) != len(subject.want) {
				t.Errorf("Fill() got = %v, wanted = %v", subject.arr, subject.want)
				return
			}

			for i := 0; i < len(subject.arr); i++ {
				if subject.arr[i] != subject.want[i] {
					t.Errorf("Fill() got = %v, wanted = %v", subject.arr, subject.want)
					return
				}
			}
		})
	}
}

func BenchmarkFill(b *testing.B) {
	for j := 0; j < len(tFillBenchs); j++ {
		b.Run(fmt.Sprintf("slice_size_%s", tFillBenchs[j].name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fill(tFillBenchs[j].arr, tFillBenchs[j].value, tFillBenchs[j].start, tFillBenchs[j].end)
			}
		})
	}
}
